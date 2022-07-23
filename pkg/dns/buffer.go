package dns

import (
	"encoding/binary"
	"fmt"
)

type Buffer struct {
	pos  uint16
	len  uint16
	data []byte
}

const BuffLen = 1024

func NewBuffer() *Buffer {
	buff := make([]byte, BuffLen)

	return &Buffer{
		pos:  0,
		data: buff,
		len:  BuffLen,
	}
}

func BufferFrom(d []byte) *Buffer {
	l := uint16(len(d))

	return &Buffer{
		pos:  0,
		len:  l,
		data: d,
	}
}

func (b *Buffer) Get(n uint16) []byte {
	d := b.data[b.pos : b.pos+n]
	b.pos += n

	return d
}

func (b *Buffer) Slice(start, end uint16) []byte {
	return b.data[start:end]
}

func (b *Buffer) Pop() byte {
	d := b.data[b.pos]
	b.pos += 1

	return d
}

func (b *Buffer) PeekUint16() uint16 {
	d := b.data[b.pos : b.pos+2]
	return binary.BigEndian.Uint16(d)
}

func (b *Buffer) ReadUint16() uint16 {
	return binary.BigEndian.Uint16(b.Get(2))
}

func (b *Buffer) Peek() byte {
	return b.data[b.pos]
}

func (b *Buffer) Remaining() uint16 {
	return b.len - b.pos
}

func (b *Buffer) Rest() []byte {
	return b.data[b.pos:]
}

func (b *Buffer) Clear() {
	b.pos = 0

	for idx := range b.data {
		b.data[idx] = 0
	}
}

func (b *Buffer) Write(d byte) {
	b.data[b.pos] = d
	b.pos += 1
}

func (b *Buffer) WriteSlice(d []byte) {
	l := uint16(len(d))
	copy(b.data[b.pos:b.pos+l], d)

	b.pos += l
}

func (b *Buffer) Bytes() []byte {
	d := make([]byte, b.len)
	copy(d, b.data)

	return d
}

func (b *Buffer) String() string {
	return fmt.Sprintf("Buffer: len<%d> pos<%d>", b.len, b.pos)
}
