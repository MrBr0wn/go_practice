package randbyte

import (
	"encoding/binary"
	"io"
	"math/rand"
)

type generator struct {
	rnd rand.Source
}

func New(seed int64) io.Reader {
	return &generator{
		rnd: rand.NewSource(seed),
	}
}

func (g *generator) Read(bytes []byte) (n int, err error) {
	for i := 0; i+8 < len(bytes); i += 8 {
		randInt := g.rnd.Int63()
		binary.LittleEndian.PutUint64(bytes[i:i+8], uint64(randInt))
	}
	return len(bytes), nil
}
