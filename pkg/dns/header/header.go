package header

import (
	"encoding/binary"
	"fmt"
	"math/rand"

	"github.com/volf52/dnsgo/pkg/dns/buffer"
	dnsFlags "github.com/volf52/dnsgo/pkg/dns/header/flags"
)

const (
	ZMask     uint16 = 0x70
	RCodeMask uint16 = 0x0F
	MaxId     int    = 65535 // uint16
)

type Header struct {
	id    uint16
	flags *dnsFlags.Flags
	z     uint16
	rcode uint16

	qdCount uint16
	anCount uint16
	nsCount uint16
	arCount uint16

	b []byte
}

func ForQuery() *Header {
	id := uint16(rand.Intn(MaxId))
	b := make([]byte, 12)

	binary.BigEndian.PutUint16(b, id)
	binary.BigEndian.PutUint16(b[2:4], 288)

	return &Header{
		id:      id,
		flags:   dnsFlags.ForQuery(),
		z:       2,
		rcode:   0,
		qdCount: 0,
		anCount: 0,
		nsCount: 0,
		arCount: 0,
		b:       b,
	}
}

func Parse(b []byte) *Header {
	return ParseFrom(buffer.From(b))
}

func ParseFrom(buff *buffer.Buffer) *Header {
	if buff.Remaining() < 12 {
		panic("must have >= 12 bytes to parse header")
	}
	b := buff.Get(12)
	innerBuff := buffer.From(b)

	id := innerBuff.ReadUint16()
	flagsVal := innerBuff.ReadUint16()
	flags := dnsFlags.Parse(flagsVal)

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

func (h *Header) IncQuestionCount() {
	h.qdCount += 1

	binary.BigEndian.PutUint16(h.b[4:6], h.qdCount)
}

func (h *Header) QdCount() uint16  {
	return h.qdCount
}

func (h *Header) AnsCount() uint16  {
	return h.anCount
}

func (h *Header) IncAnswerCount() {
	h.anCount += 1

	binary.BigEndian.PutUint16(h.b[6:8], h.anCount)
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
