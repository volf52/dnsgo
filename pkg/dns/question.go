package dns

import (
	"encoding/binary"
	"fmt"
)

type Question struct {
	lbl   *LabelSequence
	rType RecordType

	packed []byte
}

func NewQuestion(domain string, rType RecordType) *Question {
	lbl := NewLabelSequence(domain)

	lblLen := len(lbl.packed)

	packed := make([]byte, lblLen+4)
	copy(packed[:lblLen], lbl.packed)
	binary.BigEndian.PutUint16(packed[lblLen:], rType.Value())
	binary.BigEndian.PutUint16(packed[lblLen+2:], 1) // class

	return &Question{
		lbl,
		rType,
		packed,
	}
}

func ParseQuestion(b []byte) *Question {
	return ParseQuestionFrom(BufferFrom(b))
}

func ParseQuestionFrom(buff *Buffer) *Question {
	lbl := ParseLabelSequenceFrom(buff)
	if buff.Remaining() < 4 {
		panic("not enough bytes left to parse question")
	}

	rType := NewRecordType(buff.ReadUint16())
	_ = buff.ReadUint16()

	lblBytesLen := len(lbl.packed)

	packed := make([]byte, lblBytesLen+4)
	copy(packed[:lblBytesLen], lbl.packed)
	copy(packed[lblBytesLen:], buff.Slice(buff.pos-4, buff.pos))

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
