package question

import (
	"encoding/binary"
	"fmt"
	"github.com/volf52/dnsgo/pkg/dns/buffer"
	"github.com/volf52/dnsgo/pkg/dns/label_sequence"
	"github.com/volf52/dnsgo/pkg/dns/record_type"
)

type Question struct {
	lbl   *label_sequence.LabelSequence
	rType record_type.RecordType

	packed []byte
}

func New(domain string, rType record_type.RecordType) *Question {
	lbl := label_sequence.New(domain)

	lblLen := lbl.Len()

	packed := make([]byte, lblLen+4)
	copy(packed[:lblLen], lbl.Data())
	binary.BigEndian.PutUint16(packed[lblLen:], rType.Value())
	binary.BigEndian.PutUint16(packed[lblLen+2:], 1) // class

	return &Question{
		lbl,
		rType,
		packed,
	}
}

func Parse(b []byte) *Question {
	return ParseFrom(buffer.From(b))
}

func ParseFrom(buff *buffer.Buffer) *Question {
	lbl := label_sequence.ParseFrom(buff)
	if buff.Remaining() < 4 {
		panic("not enough bytes left to parse question")
	}

	rType := record_type.New(buff.ReadUint16())
	_ = buff.ReadUint16()

	lblBytesLen := lbl.Len()

	packed := make([]byte, lblBytesLen+4)
	copy(packed[:lblBytesLen], lbl.Data())
	copy(packed[lblBytesLen:], buff.Slice(buff.Pos()-4, buff.Pos()))

	return &Question{
		lbl,
		rType,
		packed,
	}
}

func (q *Question) String() string {
	return fmt.Sprintf("%s\tIN\t%s", q.lbl, q.rType)
}

func (q *Question) Bytes() []byte {
	return q.packed
}
