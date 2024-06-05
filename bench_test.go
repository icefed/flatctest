package flatctest

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"testing"
	"time"

	"github.com/icefed/flatctest/helloflatc"
	"github.com/icefed/flatctest/helloproto"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

var (
	_0k  = 0
	_1k  = 1024
	_4k  = _1k * 4
	_64k = _1k * 64
	_1m  = _1k * 1024
	_4m  = _1m * 4

	_testname      = "name"
	_testage       = 1
	_testphone     = "phone"
	_testaddress   = "address"
	_testemptydata = []byte{}
	_testdata      = []byte{}

	_testdatasize = 0
	_testdataname = "0k"
	_maxmsgsize   = 0

	_testsizenamearray = []string{
		// "0k",
		"4k",
		// "64k",
		// "1m",
		"4m",
	}
	_testsizearray = map[string]int{
		"0k":  _0k,
		"4k":  _4k,
		"64k": _64k,
		"1m":  _1m,
		"4m":  _4m,
	}
)

func initTestData(size int) {
	_testdata = _testdata[:0]
	for i := 0; i < 256; i++ {
		for j := 0; j < size/256; j++ {
			_testdata = append(_testdata, byte(i))
		}
	}
}

func BenchmarkHelloSerial(b *testing.B) {
	p := NewBuilderPool(0)
	for _, name := range _testsizenamearray {
		_testdataname = name
		_testdatasize = _testsizearray[name]
		_maxmsgsize = _testdatasize + 256
		initTestData(_testdatasize)

		b.Run(_testdataname, func(b *testing.B) {
			b.Run("flatc", func(b *testing.B) {
				b.Run("marshal", func(b *testing.B) {
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							bd := p.GetBuilder()

							// bd := flatbuffers.NewBuilder(_maxmsgsize)
							name := bd.CreateString(_testname)
							phone := bd.CreateString(_testphone)
							address := bd.CreateString(_testaddress)
							data := bd.CreateByteVector(_testdata)
							helloflatc.HelloRequestStart(bd)
							helloflatc.HelloRequestAddName(bd, name)
							helloflatc.HelloRequestAddAge(bd, 1)
							helloflatc.HelloRequestAddPhone(bd, phone)
							helloflatc.HelloRequestAddAddress(bd, address)
							helloflatc.HelloRequestAddData(bd, data)
							bd.Finish(helloflatc.HelloRequestEnd(bd))

							buf := bd.FinishedBytes()
							_ = buf

							p.PutBuilder(bd)
						}
					})
				})
				b.Run("unmarshal", func(b *testing.B) {
					b.RunParallel(func(pb *testing.PB) {
						bd := p.GetBuilder()
						defer p.PutBuilder(bd)

						// bd := flatbuffers.NewBuilder(_maxmsgsize)
						name := bd.CreateString(_testname)
						phone := bd.CreateString(_testphone)
						address := bd.CreateString(_testaddress)
						data := bd.CreateByteVector(_testdata)
						helloflatc.HelloRequestStart(bd)
						helloflatc.HelloRequestAddName(bd, name)
						helloflatc.HelloRequestAddAge(bd, 1)
						helloflatc.HelloRequestAddPhone(bd, phone)
						helloflatc.HelloRequestAddAddress(bd, address)
						helloflatc.HelloRequestAddData(bd, data)
						bd.Finish(helloflatc.HelloRequestEnd(bd))
						buf := bd.FinishedBytes()
						for pb.Next() {
							hellorequest := helloflatc.GetRootAsHelloRequest(buf, 0)
							_ = hellorequest
						}
					})
				})
			})

			b.Run("proto", func(b *testing.B) {
				b.Run("marshal", func(b *testing.B) {
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							h := &helloproto.HelloRequest{
								Name:    _testname,
								Age:     1,
								Phone:   _testphone,
								Address: _testaddress,
								Data:    _testdata,
							}

							buf, err := proto.Marshal(h)
							assert.NoError(b, err)
							_ = buf
						}
					})
				})
				b.Run("unmarshal", func(b *testing.B) {
					b.RunParallel(func(pb *testing.PB) {
						h := &helloproto.HelloRequest{
							Name:    _testname,
							Age:     1,
							Phone:   _testphone,
							Address: _testaddress,
							Data:    _testdata,
						}

						buf, err := proto.Marshal(h)
						assert.NoError(b, err)
						for pb.Next() {
							h2 := &helloproto.HelloRequest{}
							err = proto.Unmarshal(buf, h2)
							assert.NoError(b, err)
						}
					})
				})
			})
		})
	}
}

func BenchmarkHelloGRPC(b *testing.B) {
	p := NewBuilderPool(0)
	for _, name := range _testsizenamearray {
		_testdataname = name
		_testdatasize = _testsizearray[name]
		_maxmsgsize = _testdatasize + 256
		initTestData(_testdatasize)

		b.Run(_testdataname, func(b *testing.B) {
			b.Run("flatc", func(b *testing.B) {
				b.Run("send_data", func(b *testing.B) {
					client, closer := flatcServe(b, false)
					defer closer()
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							bd := p.GetBuilder()
							// bd := flatbuffers.NewBuilder(_maxmsgsize)
							name := bd.CreateString("name")
							phone := bd.CreateString("phone")
							address := bd.CreateString("address")
							data := bd.CreateByteVector(_testdata)
							helloflatc.HelloRequestStart(bd)
							helloflatc.HelloRequestAddName(bd, name)
							helloflatc.HelloRequestAddAge(bd, 1)
							helloflatc.HelloRequestAddPhone(bd, phone)
							helloflatc.HelloRequestAddAddress(bd, address)
							helloflatc.HelloRequestAddData(bd, data)
							bd.Finish(helloflatc.HelloRequestEnd(bd))

							res, err := client.SayHello(context.Background(), bd)
							assert.NoError(b, err)
							_ = res
							p.PutBuilder(bd)
						}
					})
				})
				b.Run("recv_data", func(b *testing.B) {
					client, closer := flatcServe(b, true)
					defer closer()
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							bd := p.GetBuilder()
							// bd := flatbuffers.NewBuilder(_maxmsgsize)
							name := bd.CreateString("name")
							phone := bd.CreateString("phone")
							address := bd.CreateString("address")
							data := bd.CreateByteVector(_testemptydata)
							helloflatc.HelloRequestStart(bd)
							helloflatc.HelloRequestAddName(bd, name)
							helloflatc.HelloRequestAddAge(bd, 1)
							helloflatc.HelloRequestAddPhone(bd, phone)
							helloflatc.HelloRequestAddAddress(bd, address)
							helloflatc.HelloRequestAddData(bd, data)
							bd.Finish(helloflatc.HelloRequestEnd(bd))

							res, err := client.SayHello(context.Background(), bd)
							assert.NoError(b, err)
							_ = res
							p.PutBuilder(bd)
						}
					})
				})
				b.Run("sendrecv", func(b *testing.B) {
					client, closer := flatcServe(b, true)
					defer closer()
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							bd := p.GetBuilder()
							// bd := flatbuffers.NewBuilder(_maxmsgsize)
							name := bd.CreateString("name")
							phone := bd.CreateString("phone")
							address := bd.CreateString("address")
							data := bd.CreateByteVector(_testdata)
							helloflatc.HelloRequestStart(bd)
							helloflatc.HelloRequestAddName(bd, name)
							helloflatc.HelloRequestAddAge(bd, 1)
							helloflatc.HelloRequestAddPhone(bd, phone)
							helloflatc.HelloRequestAddAddress(bd, address)
							helloflatc.HelloRequestAddData(bd, data)
							bd.Finish(helloflatc.HelloRequestEnd(bd))

							res, err := client.SayHello(context.Background(), bd)
							assert.NoError(b, err)
							_ = res
							p.PutBuilder(bd)
						}
					})
				})
			})
			b.Run("proto", func(b *testing.B) {
				b.Run("send_data", func(b *testing.B) {
					client, closer := protoServe(b, false)
					defer closer()
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							res, err := client.SayHello(context.Background(), &helloproto.HelloRequest{
								Name:    "name",
								Age:     1,
								Phone:   "phone",
								Address: "address",
								Data:    _testdata,
							})

							assert.NoError(b, err)
							_ = res
						}
					})
				})
				b.Run("recv_data", func(b *testing.B) {
					client, closer := protoServe(b, true)
					defer closer()
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							res, err := client.SayHello(context.Background(), &helloproto.HelloRequest{
								Name:    "name",
								Age:     1,
								Phone:   "phone",
								Address: "address",
								Data:    _testemptydata,
							})

							assert.NoError(b, err)
							_ = res
						}
					})
				})
				b.Run("sendrecv", func(b *testing.B) {
					client, closer := protoServe(b, true)
					defer closer()
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							res, err := client.SayHello(context.Background(), &helloproto.HelloRequest{
								Name:    "name",
								Age:     1,
								Phone:   "phone",
								Address: "address",
								Data:    _testdata,
							})

							assert.NoError(b, err)
							_ = res
						}
					})
				})
			})
		})
	}
}

type protoServer struct {
	writeData bool
}

func (s *protoServer) SayHello(ctx context.Context, in *helloproto.HelloRequest) (*helloproto.HelloReply, error) {
	writedata := _testemptydata
	if s.writeData {
		writedata = _testdata
	}
	return &helloproto.HelloReply{
		Name:    in.Name,
		Age:     in.Age,
		Phone:   in.Phone,
		Address: in.Address,
		Data:    writedata,
	}, nil
}

func protoServe(t assert.TestingT, writeData bool) (client helloproto.GreeterClient, closer func()) {
	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(_maxmsgsize),
		grpc.MaxSendMsgSize(_maxmsgsize),
	)
	helloproto.RegisterGreeterServer(grpcServer, &protoServer{writeData: writeData})

	port := randPort()
	grpcEndpoint := fmt.Sprintf("%s:%d", "127.0.0.1", port)
	grpcL, err := net.Listen("tcp", grpcEndpoint)
	assert.NoError(t, err)
	go grpcServer.Serve(grpcL)
	waitForServerRunning("127.0.0.1", port)

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", "127.0.0.1", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(_maxmsgsize),
			grpc.MaxCallSendMsgSize(_maxmsgsize),
		),
	)
	assert.NoError(t, err)
	client = helloproto.NewGreeterClient(conn)
	assert.NoError(t, err)
	return client, func() {
		conn.Close()
		grpcServer.GracefulStop()
	}
}

type flatcServer struct {
	helloflatc.UnimplementedGreeterServer
	writeData bool
}

func (s *flatcServer) SayHello(ctx context.Context, req *helloflatc.HelloRequest) (*flatbuffers.Builder, error) {
	out := flatbuffers.NewBuilder(_maxmsgsize)
	name := out.CreateString(string(req.Name()))
	phone := out.CreateString(string(req.Phone()))
	address := out.CreateString(string(req.Address()))
	writedata := _testemptydata
	if s.writeData {
		writedata = _testdata
	}
	data := out.CreateByteVector(writedata)
	helloflatc.HelloRequestStart(out)
	helloflatc.HelloRequestAddName(out, name)
	helloflatc.HelloRequestAddAge(out, req.Age())
	helloflatc.HelloRequestAddPhone(out, phone)
	helloflatc.HelloRequestAddAddress(out, address)
	helloflatc.HelloRequestAddData(out, data)
	out.Finish(helloflatc.HelloRequestEnd(out))

	return out, nil
}

func flatcServe(t assert.TestingT, writeData bool) (client helloflatc.GreeterClient, closer func()) {
	grpcServer := grpc.NewServer(
		grpc.ForceServerCodec(&flatbuffers.FlatbuffersCodec{}),
		grpc.MaxRecvMsgSize(_maxmsgsize),
		grpc.MaxSendMsgSize(_maxmsgsize),
	)
	helloflatc.RegisterGreeterServer(grpcServer, &flatcServer{writeData: writeData})

	port := randPort()
	grpcEndpoint := fmt.Sprintf("%s:%d", "127.0.0.1", port)
	grpcL, err := net.Listen("tcp", grpcEndpoint)
	assert.NoError(t, err)
	go grpcServer.Serve(grpcL)
	waitForServerRunning("127.0.0.1", port)

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", "127.0.0.1", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.ForceCodec(&flatbuffers.FlatbuffersCodec{}),
			grpc.MaxCallRecvMsgSize(_maxmsgsize),
			grpc.MaxCallSendMsgSize(_maxmsgsize),
		),
	)
	assert.NoError(t, err)
	client = helloflatc.NewGreeterClient(conn)
	assert.NoError(t, err)
	return client, func() {
		conn.Close()
		grpcServer.GracefulStop()
	}
}

func randPort() int {
	return int(rand.Int63n(55535) + 10000)
}

func waitForServerRunning(host string, port int) {
	for {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
		if err == nil {
			_ = conn.Close()
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}
