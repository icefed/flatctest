package flatctest

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"sync"
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

	_testdatamu  = sync.Mutex{}
	_testdatamap = make(map[int][]byte)

	_testsizenamearray = []string{
		"0k",
		"4k",
		"64k",
		"1m",
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

func initTestData(size int) []byte {
	testdata := make([]byte, 0, size)
	for i := 0; i < 256; i++ {
		for j := 0; j < size/256; j++ {
			testdata = append(testdata, byte(i))
		}
	}
	return testdata
}

func getTestData(size int) []byte {
	_testdatamu.Lock()
	defer _testdatamu.Unlock()

	if size == 0 {
		return _testemptydata
	}
	if data, ok := _testdatamap[size]; ok {
		return data
	}

	testdata := initTestData(size)
	_testdatamap[size] = testdata

	return testdata
}

func getMaxMsgSize(size int) int {
	return size + 256
}

func TestMain(m *testing.M) {
	for _, size := range _testsizearray {
		testdata := initTestData(size)
		_testdatamap[size] = testdata
	}
	m.Run()
}

func BenchmarkSerial(b *testing.B) {
	for _, sizename := range _testsizenamearray {
		b.Run(sizename, func(b *testing.B) {
			benchSerialFlatc(b, sizename)
			benchSerialProto(b, sizename)
		})
	}
}

func benchSerialFlatc(b *testing.B, sizename string) {
	testdatasize := _testsizearray[sizename]
	maxmsgsize := getMaxMsgSize(testdatasize)

	// p := NewBuilderPool(0)
	// bd := p.GetBuilder()
	// defer p.PutBuilder(bd)
	bd := flatbuffers.NewBuilder(maxmsgsize)
	name := bd.CreateString(_testname)
	phone := bd.CreateString(_testphone)
	address := bd.CreateString(_testaddress)
	data := bd.CreateByteVector(getTestData(testdatasize))
	helloflatc.UserStart(bd)
	helloflatc.UserAddName(bd, name)
	helloflatc.UserAddAge(bd, 1)
	helloflatc.UserAddPhone(bd, phone)
	helloflatc.UserAddAddress(bd, address)
	helloflatc.UserAddData(bd, data)
	bd.Finish(helloflatc.UserEnd(bd))
	buf := bd.FinishedBytes()

	b.Run("flatc", func(b *testing.B) {
		b.Run("marshal", func(b *testing.B) {
			b.SetBytes(int64(testdatasize))
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					// bd := p.GetBuilder()
					bd := flatbuffers.NewBuilder(maxmsgsize)
					name := bd.CreateString(_testname)
					phone := bd.CreateString(_testphone)
					address := bd.CreateString(_testaddress)
					data := bd.CreateByteVector(getTestData(testdatasize))
					helloflatc.UserStart(bd)
					helloflatc.UserAddName(bd, name)
					helloflatc.UserAddAge(bd, 1)
					helloflatc.UserAddPhone(bd, phone)
					helloflatc.UserAddAddress(bd, address)
					helloflatc.UserAddData(bd, data)
					bd.Finish(helloflatc.UserEnd(bd))

					buf := bd.FinishedBytes()
					_ = buf

					// p.PutBuilder(bd)
				}
			})
		})

		b.Run("unmarshal", func(b *testing.B) {
			b.SetBytes(int64(testdatasize))

			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					user := helloflatc.GetRootAsUser(buf, 0)
					_ = user
				}
			})
		})
	})
}

func benchSerialProto(b *testing.B, sizename string) {
	testdatasize := _testsizearray[sizename]

	h := &helloproto.User{
		Name:    _testname,
		Age:     1,
		Phone:   _testphone,
		Address: _testaddress,
		Data:    getTestData(testdatasize),
	}
	buf, _ := proto.Marshal(h)

	b.Run("proto", func(b *testing.B) {
		b.Run("marshal", func(b *testing.B) {
			b.SetBytes(int64(testdatasize))
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					h := &helloproto.User{
						Name:    _testname,
						Age:     1,
						Phone:   _testphone,
						Address: _testaddress,
						Data:    getTestData(testdatasize),
					}

					buf, err := proto.Marshal(h)
					_ = err
					_ = buf
				}
			})
		})

		b.Run("unmarshal", func(b *testing.B) {
			b.SetBytes(int64(testdatasize))
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					h2 := &helloproto.User{}
					err := proto.Unmarshal(buf, h2)
					_ = err
				}
			})
		})
	})
}

func BenchmarkGRPC(b *testing.B) {
	for _, sizename := range _testsizenamearray {
		b.Run(sizename, func(b *testing.B) {
			benchGRPCFlatc(b, sizename)
			benchGRPCProto(b, sizename)
		})
	}
}

func benchGRPCFlatc(b *testing.B, sizename string) {
	testdatasize := _testsizearray[sizename]
	maxmsgsize := getMaxMsgSize(testdatasize)

	// p := NewBuilderPool(0)
	b.Run("flatc", func(b *testing.B) {
		client, closer := flatcServe(b, maxmsgsize)
		defer closer()

		b.Run("write", func(b *testing.B) {
			b.SetBytes(int64(testdatasize))
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					// bd := p.GetBuilder()
					bd := flatbuffers.NewBuilder(maxmsgsize)
					data := bd.CreateByteVector(getTestData(testdatasize))
					helloflatc.WriteRequestStart(bd)
					helloflatc.WriteRequestAddData(bd, data)
					bd.Finish(helloflatc.WriteRequestEnd(bd))

					res, err := client.Write(context.Background(), bd)
					_ = res
					_ = err
					// p.PutBuilder(bd)
				}
			})
		})
		b.Run("read", func(b *testing.B) {
			b.SetBytes(int64(testdatasize))
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					// bd := p.GetBuilder()
					bd := flatbuffers.NewBuilder(maxmsgsize)
					helloflatc.ReadRequestStart(bd)
					helloflatc.ReadRequestAddReadBytes(bd, int32(testdatasize))
					bd.Finish(helloflatc.ReadRequestEnd(bd))

					res, err := client.Read(context.Background(), bd)
					_ = res
					_ = err
					// p.PutBuilder(bd)
				}
			})
		})
		b.Run("writestream", func(b *testing.B) {
			b.SetBytes(int64(testdatasize))
			b.RunParallel(func(pb *testing.PB) {
				writer, err := client.WriteStream(context.Background())
				_ = err
				defer writer.CloseSend()

				for pb.Next() {
					// bd := p.GetBuilder()
					bd := flatbuffers.NewBuilder(maxmsgsize)
					data := bd.CreateByteVector(getTestData(testdatasize))
					helloflatc.WriteRequestStart(bd)
					helloflatc.WriteRequestAddData(bd, data)
					bd.Finish(helloflatc.WriteRequestEnd(bd))

					err := writer.Send(bd)
					_ = err
					resp, err := writer.Recv()
					_ = resp
					_ = err

					// p.PutBuilder(bd)
				}
			})
		})
		b.Run("readstream", func(b *testing.B) {
			b.SetBytes(int64(testdatasize))
			b.RunParallel(func(pb *testing.PB) {
				reader, err := client.ReadStream(context.Background())
				_ = err
				defer reader.CloseSend()

				for pb.Next() {
					// bd := p.GetBuilder()
					bd := flatbuffers.NewBuilder(maxmsgsize)
					helloflatc.ReadRequestStart(bd)
					helloflatc.ReadRequestAddReadBytes(bd, int32(testdatasize))
					bd.Finish(helloflatc.ReadRequestEnd(bd))

					err := reader.Send(bd)
					_ = err
					resp, err := reader.Recv()
					_ = resp
					_ = err

					// p.PutBuilder(bd)
				}
			})
		})
	})
}

func benchGRPCProto(b *testing.B, sizename string) {
	testdatasize := _testsizearray[sizename]
	maxmsgsize := getMaxMsgSize(testdatasize)

	b.Run("proto", func(b *testing.B) {
		client, closer := protoServe(b, maxmsgsize)
		defer closer()

		b.Run("write", func(b *testing.B) {
			b.SetBytes(int64(testdatasize))
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					res, err := client.Write(context.Background(), &helloproto.WriteRequest{
						Data: getTestData(testdatasize),
					})

					_ = res
					_ = err
				}
			})
		})
		b.Run("read", func(b *testing.B) {
			b.SetBytes(int64(testdatasize))
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					res, err := client.Read(context.Background(), &helloproto.ReadRequest{
						ReadBytes: int32(testdatasize),
					})

					_ = res
					_ = err
				}
			})
		})

		b.Run("writestream", func(b *testing.B) {
			b.SetBytes(int64(testdatasize))
			b.RunParallel(func(pb *testing.PB) {
				writer, err := client.WriteStream(context.Background())
				_ = err
				defer writer.CloseSend()

				for pb.Next() {
					err := writer.Send(&helloproto.WriteRequest{
						Data: getTestData(testdatasize),
					})
					_ = err
					resp, err := writer.Recv()
					_ = resp
					_ = err
				}
			})
		})
		b.Run("readstream", func(b *testing.B) {
			b.SetBytes(int64(testdatasize))
			b.RunParallel(func(pb *testing.PB) {
				reader, err := client.ReadStream(context.Background())
				_ = err
				defer reader.CloseSend()

				for pb.Next() {
					err := reader.Send(&helloproto.ReadRequest{
						ReadBytes: int32(testdatasize),
					})
					_ = err
					resp, err := reader.Recv()
					_ = resp
					_ = err
				}
			})
		})
	})
}

type protoServer struct {
	helloproto.UnimplementedHelloServiceServer
}

func (s *protoServer) SayHello(ctx context.Context, in *helloproto.HelloRequest) (*helloproto.HelloReply, error) {
	return &helloproto.HelloReply{
		Name: in.Name,
	}, nil
}

func (s *protoServer) Read(ctx context.Context, in *helloproto.ReadRequest) (*helloproto.ReadReply, error) {
	return &helloproto.ReadReply{
		Data: getTestData(int(in.ReadBytes)),
		Eof:  false,
	}, nil
}

func (s *protoServer) Write(ctx context.Context, in *helloproto.WriteRequest) (*helloproto.WriteReply, error) {
	return &helloproto.WriteReply{
		WrittenBytes: int32(len(in.Data)),
	}, nil
}

func (s *protoServer) ReadStream(stream helloproto.HelloService_ReadStreamServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		if err := stream.Send(&helloproto.ReadReply{
			Data: getTestData(int(req.ReadBytes)),
		}); err != nil {
			return err
		}
	}
}

func (s *protoServer) WriteStream(stream helloproto.HelloService_WriteStreamServer) error {
	for {
		in, err := stream.Recv()
		if err != nil {
			return err
		}
		if err := stream.Send(&helloproto.WriteReply{
			WrittenBytes: int32(len(in.Data)),
		}); err != nil {
			return err
		}
	}
}

func protoServe(t assert.TestingT, maxmsgsize int) (client helloproto.HelloServiceClient, closer func()) {
	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(maxmsgsize),
		grpc.MaxSendMsgSize(maxmsgsize),
	)
	helloproto.RegisterHelloServiceServer(grpcServer, &protoServer{})

	port := randPort()
	grpcEndpoint := fmt.Sprintf("%s:%d", "127.0.0.1", port)
	grpcL, err := net.Listen("tcp", grpcEndpoint)
	assert.NoError(t, err)
	go grpcServer.Serve(grpcL)
	waitForServerRunning("127.0.0.1", port)

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", "127.0.0.1", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(maxmsgsize),
			grpc.MaxCallSendMsgSize(maxmsgsize),
		),
	)
	assert.NoError(t, err)
	client = helloproto.NewHelloServiceClient(conn)
	assert.NoError(t, err)
	return client, func() {
		conn.Close()
		grpcServer.GracefulStop()
	}
}

type flatcServer struct {
	helloflatc.UnimplementedHelloServiceServer
}

func (s *flatcServer) SayHello(ctx context.Context, req *helloflatc.HelloRequest) (*flatbuffers.Builder, error) {
	out := flatbuffers.NewBuilder(0)
	name := out.CreateString(string(req.Name()))
	helloflatc.HelloReplyStart(out)
	helloflatc.HelloReplyAddName(out, name)
	out.Finish(helloflatc.HelloReplyEnd(out))

	return out, nil
}

func (s *flatcServer) Read(ctx context.Context, req *helloflatc.ReadRequest) (*flatbuffers.Builder, error) {
	out := flatbuffers.NewBuilder(getMaxMsgSize(int(req.ReadBytes())))
	data := out.CreateByteVector(getTestData(int(req.ReadBytes())))
	helloflatc.ReadReplyStart(out)
	helloflatc.ReadReplyAddData(out, data)
	helloflatc.ReadReplyAddEof(out, false)
	out.Finish(helloflatc.ReadReplyEnd(out))

	return out, nil
}

func (s *flatcServer) Write(ctx context.Context, req *helloflatc.WriteRequest) (*flatbuffers.Builder, error) {
	out := flatbuffers.NewBuilder(0)
	helloflatc.WriteReplyStart(out)
	helloflatc.WriteReplyAddWrittenBytes(out, int32(len(req.DataBytes())))
	out.Finish(helloflatc.WriteReplyEnd(out))

	return out, nil
}

func (s *flatcServer) ReadStream(stream helloflatc.HelloService_ReadStreamServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		out := flatbuffers.NewBuilder(getMaxMsgSize(int(req.ReadBytes())))
		data := out.CreateByteVector(getTestData(int(req.ReadBytes())))
		helloflatc.ReadReplyStart(out)
		helloflatc.ReadReplyAddData(out, data)
		helloflatc.ReadReplyAddEof(out, false)
		out.Finish(helloflatc.ReadReplyEnd(out))

		if err := stream.Send(out); err != nil {
			return err
		}
	}
}

func (s *flatcServer) WriteStream(stream helloflatc.HelloService_WriteStreamServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		out := flatbuffers.NewBuilder(0)
		helloflatc.WriteReplyStart(out)
		helloflatc.WriteReplyAddWrittenBytes(out, int32(len(req.DataBytes())))
		out.Finish(helloflatc.WriteReplyEnd(out))

		if err := stream.Send(out); err != nil {
			return err
		}
	}
}

func flatcServe(t assert.TestingT, maxmsgsize int) (client helloflatc.HelloServiceClient, closer func()) {
	grpcServer := grpc.NewServer(
		grpc.ForceServerCodec(&flatbuffers.FlatbuffersCodec{}),
		grpc.MaxRecvMsgSize(maxmsgsize),
		grpc.MaxSendMsgSize(maxmsgsize),
	)
	helloflatc.RegisterHelloServiceServer(grpcServer, &flatcServer{})

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
			grpc.MaxCallRecvMsgSize(maxmsgsize),
			grpc.MaxCallSendMsgSize(maxmsgsize),
		),
	)
	assert.NoError(t, err)
	client = helloflatc.NewHelloServiceClient(conn)
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
