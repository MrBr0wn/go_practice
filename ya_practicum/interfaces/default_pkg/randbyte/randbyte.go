package randbyte

import (
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
	for i := range bytes {
		randInt := g.rnd.Int63()
		randByte := byte(randInt)
		bytes[i] = randByte
	}
	return len(bytes), nil
}
