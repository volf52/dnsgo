package record_type

import (
	"encoding/binary"
	"fmt"
)

type RecordType uint16

const (
	Unknown RecordType = 0
	A       RecordType = 1
	CNAME   RecordType = 5
	AAAA    RecordType = 28
)

func New(val uint16) RecordType {
	return RecordType(val)
}

func (r RecordType) String() string {
	switch r {
	case A:
		return "A"
	case CNAME:
		return "CNAME"
	case AAAA:
		return "AAAA"
	default:
		return fmt.Sprintf("Unknown %d", r)
	}
}

func (r RecordType) Value() uint16 {
	return uint16(r)
}

func (r RecordType) Bytes() []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, r.Value())

	return b
}
