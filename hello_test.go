package flatctest

import (
	"testing"

	"github.com/icefed/flatctest/helloflatc"
	"github.com/icefed/flatctest/helloproto"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestHello(t *testing.T) {
	t.Run("flatc", func(t *testing.T) {
		bd := flatbuffers.NewBuilder(0)
		name := bd.CreateString("name")
		phone := bd.CreateString("phone")
		address := bd.CreateString("address")
		data := bd.CreateByteVector([]byte("data"))
		helloflatc.HelloRequestStart(bd)
		helloflatc.HelloRequestAddName(bd, name)
		helloflatc.HelloRequestAddAge(bd, 1)
		helloflatc.HelloRequestAddPhone(bd, phone)
		helloflatc.HelloRequestAddAddress(bd, address)
		helloflatc.HelloRequestAddData(bd, data)
		bd.Finish(helloflatc.HelloRequestEnd(bd))

		buf := bd.FinishedBytes()

		hellorequest := helloflatc.GetRootAsHelloRequest(buf, 0)
		assert.EqualValues(t, "name", hellorequest.Name())
		assert.EqualValues(t, 1, hellorequest.Age())
		assert.EqualValues(t, "phone", hellorequest.Phone())
		assert.EqualValues(t, "address", hellorequest.Address())
		assert.Equal(t, []byte("data"), hellorequest.DataBytes())
	})
	t.Run("proto", func(t *testing.T) {
		h := &helloproto.HelloRequest{
			Name:    "name",
			Age:     1,
			Phone:   "phone",
			Address: "address",
			Data:    []byte("data"),
		}

		buf, err := proto.Marshal(h)
		assert.NoError(t, err)

		h2 := &helloproto.HelloRequest{}
		err = proto.Unmarshal(buf, h2)
		assert.NoError(t, err)
		assert.EqualExportedValues(t, h, h2)
	})
}
