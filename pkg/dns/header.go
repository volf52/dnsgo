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

	b []byte
}

func ParseHeader(b []byte) *Header {
	return ParseHeaderFrom(BufferFrom(b))
}

func ParseHeaderFrom(buff *Buffer) *Header {
	if buff.Remaining() < 12 {
		panic("must have >= 12 bytes to parse header")
	}
	b := buff.Get(12)
	innerBuff := BufferFrom(b)

	id := innerBuff.ReadUint16()
	flagsVal := innerBuff.ReadUint16()
	flags := ParseFlags(flagsVal)

	z := (flagsVal & ZMask) >> 4
	rcode := flagsVal & RCodeMask

	qdCount := innerBuff.ReadUint16()
	anCount := innerBuff.ReadUint16()
	nsCount := innerBuff.ReadUint16()
	arCount := innerBuff.ReadUint16()
	innerBuff = nil

	return &Header{
		id,
		flags,
		z,
		rcode,
		qdCount,
		anCount,
		nsCount,
		arCount,
		b,
	}
}

func (h *Header) String() string {
	return fmt.Sprintf("ID %d\nFlags: %s Z=%d RCODE=%d\n"+
		"Questions: %d\tAnswers: %d",
		h.id, h.flags, h.z, h.rcode,
		h.qdCount, h.anCount,
	)
}

func (h *Header) Bytes() []byte {
	return h.b
}
