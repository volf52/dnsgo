package dns

import (
	"fmt"
	"github.com/volf52/dnsgo/internal/utils"
	"strings"
)

const (
	JmpByte byte   = 0xc0
	JmpMask uint16 = 0xc000
)

type LabelSequence struct {
	domain string

	packed []byte
}

func NewLabelSequence(domain string) *LabelSequence {
	parts := strings.Split(domain, ".")
	b := BufferWithCap(512)

	for _, part := range parts {
		l := len(part)

		b.Write(uint8(l))
		b.WriteSlice([]byte(part))
	}

	packed := make([]byte, b.pos+1)
	copy(packed, b.data[:b.pos])

	return &LabelSequence{
		domain,
		packed,
	}
}

func ParseLabelSequence(b []byte) *LabelSequence {
	return ParseLabelSequenceFrom(BufferFrom(b))
}

func ParseLabelSequenceFrom(buff *Buffer) *LabelSequence {
	initPos := buff.pos

	if utils.IsSet(buff.Peek(), JmpByte) {
		if buff.Remaining() < 2 {
			panic("not enough bytes remaining")
		}

		jmpIdx := buff.ReadUint16()
		jmpIdx ^= JmpMask

		if jmpIdx > buff.pos {
			panic("invalid jmp idx")
		}

		slice := buff.Slice(jmpIdx, buff.len)

		return ParseLabelSequence(slice)
	}

	idx := buff.pos
	var parts []string
	for idx < buff.len && buff.Peek() != 0x0 {
		partLen := uint16(buff.Pop())

		if idx+partLen >= buff.len {
			panic("invalid part length")
		}

		b := buff.Get(partLen)
		part := string(b)
		parts = append(parts, part)

		idx = buff.pos
	}
	buff.Pop() // pop null byte
	domain := strings.Join(parts, ".")
	packed := make([]byte, buff.pos-initPos)
	copy(packed, buff.Slice(initPos, buff.pos))

	return &LabelSequence{
		domain,
		packed,
	}
}

func (lbl *LabelSequence) Domain() string {
	return lbl.domain
}

func (lbl *LabelSequence) String() string {
	return fmt.Sprintf(";; %s ;;", lbl.domain)
}

func (lbl *LabelSequence) Bytes() []byte {
	return lbl.packed
}
