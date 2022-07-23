package dns

import "fmt"

type Question struct {
	lbl   *LabelSequence
	rType RecordType

	packed []byte
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

	lblBytes := lbl.Bytes()
	lblBytesLen := len(lblBytes)

	packed := make([]byte, lblBytesLen+4)
	copy(packed[:lblBytesLen], lblBytes)
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
