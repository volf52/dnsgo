package label_sequence

import (
	"fmt"
	"strings"

	"github.com/volf52/dnsgo/internal/utils"
	"github.com/volf52/dnsgo/pkg/dns/buffer"
)

const (
	JmpByte byte   = 0xc0
	JmpMask uint16 = 0xc000
)

type LabelSequence struct {
	domain string

	packed []byte
}

func New(domain string) *LabelSequence {
	parts := strings.Split(domain, ".")
	b := buffer.WithCap(512)

	for _, part := range parts {
		l := len(part)

		b.Write(uint8(l))
		b.WriteSlice([]byte(part))
	}

	packed := make([]byte, b.Pos()+1)
	copy(packed, b.Till(b.Pos()))

	return &LabelSequence{
		domain,
		packed,
	}
}

func Parse(b []byte) *LabelSequence {
	return ParseFrom(buffer.From(b))
}

func ParseFrom(buff *buffer.Buffer) *LabelSequence {
	initPos := buff.Pos()

	if utils.IsSet(buff.Peek(), JmpByte) {
		if buff.Remaining() < 2 {
			panic("not enough bytes remaining")
		}

		jmpIdx := buff.ReadUint16()
		jmpIdx ^= JmpMask

		if jmpIdx > buff.Pos() {
			panic("invalid jmp idx")
		}

		slice := buff.Slice(jmpIdx, buff.Len())

		return Parse(slice)
	}

	idx := initPos
	var parts []string
	for idx < buff.Len() && buff.Peek() != 0x0 {
		partLen := uint16(buff.Pop())

		if idx+partLen >= buff.Len() {
			panic("invalid part length")
		}

		b := buff.Get(partLen)
		part := string(b)
		parts = append(parts, part)

		idx = buff.Pos()
	}
	buff.Pop() // pop null byte
	domain := strings.Join(parts, ".")
	packed := make([]byte, buff.Pos()-initPos)
	copy(packed, buff.Slice(initPos, buff.Pos()))

	return &LabelSequence{
		domain,
		packed,
	}
}

func (lbl *LabelSequence) Domain() string {
	return lbl.domain
}

func (lbl *LabelSequence) Len() int {
	return len(lbl.packed)
}

func (lbl *LabelSequence) Data() []byte {
	return lbl.packed
}

func (lbl *LabelSequence) String() string {
	return fmt.Sprintf(";; %s ;;", lbl.domain)
}

func (lbl *LabelSequence) Bytes() []byte {
	return lbl.packed
}
