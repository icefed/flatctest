package flatctest

import (
	"sync"

	flatbuffers "github.com/google/flatbuffers/go"
)

const (
	// default use 4M + 256 bytes, 256 for header and additional values
	defaultBuilderSize = 1024*1024*4 + 256
)

type BuilderPool interface {
	GetBuilder() *flatbuffers.Builder
	PutBuilder(*flatbuffers.Builder)
}

type builderPool struct {
	pool sync.Pool
}

func NewBuilderPool(builderSize int) BuilderPool {
	return &builderPool{
		pool: sync.Pool{
			New: func() interface{} {
				if builderSize <= 0 {
					builderSize = defaultBuilderSize
				}
				return flatbuffers.NewBuilder(builderSize)
			},
		},
	}
}

func (p *builderPool) GetBuilder() *flatbuffers.Builder {
	return p.pool.Get().(*flatbuffers.Builder)
}

func (p *builderPool) PutBuilder(b *flatbuffers.Builder) {
	b.Reset()
	p.pool.Put(b)
}
