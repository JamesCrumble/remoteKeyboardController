package server

type Buffer struct {
	arr []byte
}

func (buffer *Buffer) ResetBuffer() {
	for i := 0; i < len(buffer.arr); i++ {
		buffer.arr[i] = 0
	}
}

func NewBuffer(size int) Buffer {
	return Buffer{arr: make([]byte, size)}
}
