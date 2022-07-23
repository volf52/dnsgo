package dns

import "fmt"

const (
	ZMask     uint16 = 0x70
	RCodeMask uint16 = 0x0F
)

type Header struct {
	id    uint16
	flags *HeaderFlags
	z     uint16
	rcode uint16

	qdCount uint16
	anCount uint16
	nsCount uint16
	arCount uint16
}

func ParseHeader(b []byte) *Header {
	return ParseHeaderFrom(BufferFrom(b))
}

func ParseHeaderFrom(b *Buffer) *Header {
	if b.Remaining() < 12 {
		panic("must have >= 12 bytes to parse header")
	}

	id := b.ReadUint16()
	flagsVal := b.ReadUint16()
	flags := ParseFlags(flagsVal)

	z := (flagsVal & ZMask) >> 4
	rcode := flagsVal & RCodeMask

	qdCount := b.ReadUint16()
	anCount := b.ReadUint16()
	nsCount := b.ReadUint16()
	arCount := b.ReadUint16()

	return &Header{
		id,
		flags,
		z,
		rcode,
		qdCount,
		anCount,
		nsCount,
		arCount,
	}
}

func (h *Header) String() string {
	return fmt.Sprintf("ID %d\nFlags: %s Z=%d RCODE=%d\n"+
		"Questions: %d\tAnswers: %d",
		h.id, h.flags, h.z, h.rcode,
		h.qdCount, h.anCount,
	)
}
