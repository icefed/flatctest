package flatctest

import (
	"testing"

	"github.com/icefed/flatctest/helloflatc"
	"github.com/icefed/flatctest/helloproto"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestFlatc(t *testing.T) {
	bd := flatbuffers.NewBuilder(0)
	name := bd.CreateString("name")
	phone := bd.CreateString("phone")
	address := bd.CreateString("address")
	data := bd.CreateByteVector([]byte("data"))
	helloflatc.UserStart(bd)
	helloflatc.UserAddName(bd, name)
	helloflatc.UserAddAge(bd, 1)
	helloflatc.UserAddPhone(bd, phone)
	helloflatc.UserAddAddress(bd, address)
	helloflatc.UserAddData(bd, data)
	bd.Finish(helloflatc.UserEnd(bd))

	buf := bd.FinishedBytes()

	user := helloflatc.GetRootAsUser(buf, 0)
	assert.EqualValues(t, "name", user.Name())
	assert.EqualValues(t, 1, user.Age())
	assert.EqualValues(t, "phone", user.Phone())
	assert.EqualValues(t, "address", user.Address())
	assert.Equal(t, []byte("data"), user.DataBytes())
}

func TestProto(t *testing.T) {
	h := &helloproto.User{
		Name:    "name",
		Age:     1,
		Phone:   "phone",
		Address: "address",
		Data:    []byte("data"),
	}

	buf, err := proto.Marshal(h)
	assert.NoError(t, err)

	h2 := &helloproto.User{}
	err = proto.Unmarshal(buf, h2)
	assert.NoError(t, err)
	assert.EqualExportedValues(t, h, h2)
}
