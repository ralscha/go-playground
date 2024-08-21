package main

import (
	"fmt"
	"github.com/carlmjohnson/deque"
	"strings"
)

type History struct {
	data []string
	size int
	pos  int
}

func NewHistory(size int) *History {
	return &History{
		data: make([]string, size),
		size: size,
		pos:  0,
	}
}

func (h *History) Add(element string) {
	h.data[h.pos] = element
	h.pos = (h.pos + 1) % h.size
}

func (h *History) Get() string {
	return h.data[h.pos]
}

func (h *History) GetLast() string {
	return h.data[(h.pos-1+h.size)%h.size]
}

func (h *History) String() string {
	sb := strings.Builder{}
	sb.Grow(h.size * (len(h.data[0]) + 1))

	idx := h.pos
	for i := 0; i < h.size; i++ {
		sb.WriteString(h.data[idx])
		if i < h.size-1 {
			sb.WriteByte('\n')
		}
		idx = (idx + 1) % h.size
	}

	return sb.String()
}

func main() {
	h := NewHistory(3)
	for _, s := range []string{"foo", "bar", "car", "dar", "ear"} {
		h.Add(s)
	}
	fmt.Println(h)

	d := deque.Of(9, 8, 7, 6)
	fmt.Println(d)

	hd := deque.Make[int](3)
	for _, i := range []int{1, 2, 3, 4, 5} {
		hd.PushBack(i)
	}

	fmt.Println(hd)

}
