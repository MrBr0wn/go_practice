package hashbyte

import "io"

type Hasher interface {
	io.Writer
	Hash() byte
}

type hash struct {
	result byte
}

func New(_init byte) Hasher {
	return &hash{
		result: _init,
	}
}

func (h *hash) Write(bytes []byte) (n int, err error) {
	for _, b := range bytes {
		h.result = (h.result^b)<<1 + b%2
	}

	return len(bytes), nil
}

func (h hash) Hash() byte {
	return h.result
}
