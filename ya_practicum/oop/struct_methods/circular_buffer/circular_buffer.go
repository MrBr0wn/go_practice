package main

import "fmt"

type CircularBuffer struct {
	values  []float64
	headIdx int
	tailIdx int
}

func (b CircularBuffer) GetCurrentSize() int {
	if b.tailIdx < b.headIdx {
		return b.tailIdx + cap(b.values) - b.headIdx
	}

	return b.tailIdx - b.headIdx
}

func (b CircularBuffer) GetValues() (retValues []float64) {
	for i := b.headIdx; i != b.tailIdx; i = (i + 1) % cap(b.values) {
		retValues = append(retValues, b.values[i])
	}

	return
}

func (b *CircularBuffer) AddValue(v float64) {
	b.values[b.tailIdx] = v
	b.tailIdx = (b.tailIdx + 1) % cap(b.values)
	if b.tailIdx == b.headIdx {
		b.headIdx = (b.headIdx + 1) % cap(b.values)
	}
}

func NewCircularBuffer(size int) CircularBuffer {
	return CircularBuffer{values: make([]float64, size+1)}
}

func (b CircularBuffer) ForceSetValueByIdx(idx int, v float64) {
	b.values[idx] = v
}

func (b *CircularBuffer) TestFunc(idx int, v float64) {
	b.values[idx] = v
}

func main() {
	buf := NewCircularBuffer(4)
	for i := 0; i < 6; i++ {
		if i > 0 {
			buf.AddValue(float64(i))
		}

		fmt.Printf("[%d]: %v\n", buf.GetCurrentSize(), buf.GetValues())
	}
	fmt.Println()

	buf1 := NewCircularBuffer(5)
	buf1.ForceSetValueByIdx(0, -1.0)
	buf1.ForceSetValueByIdx(1, -2.0)
	fmt.Println(buf1.values)

	buf2 := NewCircularBuffer(4)
	buf2.TestFunc(0, 3)
	buf2.TestFunc(1, 2)
	buf2.ForceSetValueByIdx(0, 1)
	fmt.Println(buf2.values)
}
