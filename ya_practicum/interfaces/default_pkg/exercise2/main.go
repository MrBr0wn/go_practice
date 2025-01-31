package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type limitedReader struct {
	r    io.Reader
	tail int
}

func (r *limitedReader) Read(p []byte) (int, error) {
	if r.tail <= 0 {
		return 0, io.EOF
	}

	if len(p) > r.tail {
		p = p[:r.tail]
	}

	n, err := r.r.Read(p)
	r.tail -= n

	return n, err
}

func LimitReader(r io.Reader, n int) io.Reader {
	return &limitedReader{r: r, tail: n}
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := LimitReader(r, 4)

	_, err := io.Copy(os.Stdout, lr)
	if err != nil {
		log.Fatal(err)
	}
}
