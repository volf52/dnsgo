package answer

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/volf52/dnsgo/pkg/dns/buffer"
	"github.com/volf52/dnsgo/pkg/dns/label_sequence"
	"github.com/volf52/dnsgo/pkg/dns/record_type"
)

type Answer struct {
	lbl *label_sequence.LabelSequence

	rType    record_type.RecordType
	ttl      uint32
	rDataLen uint16
	rData    string

	packed []byte
}

func Parse(b []byte) *Answer {
	return ParseFrom(buffer.From(b))
}

func ParseFrom(buff *buffer.Buffer) *Answer {
	initPos := buff.Pos()
	lbl := label_sequence.ParseFrom(buff)
	if buff.Remaining() < 10 {
		panic("not enough bytes left to parse question")
	}

	rTypeVal := buff.ReadUint16()
	rType := record_type.New(rTypeVal)

	_ = buff.ReadUint16() // class
	ttl := buff.ReadUint32()
	rDataLen := buff.ReadUint16()

	if buff.Remaining() < rDataLen {
		panic("not enough bytes left to parse question")
	}

	var rData string

	switch rType {
	case record_type.A:
		rData = parseIp(buff)
		break
	case record_type.CNAME:
		rData = parseCname(buff)
		break
	case record_type.AAAA:
		rData = parseIp(buff)
		break
	default:
		panic("unsupported recordType")
	}

	s := buff.Slice(initPos, buff.Pos())
	packed := make([]byte, len(s))
	copy(packed, s)

	return &Answer{
		lbl,
		rType,
		ttl,
		rDataLen,
		rData,
		packed,
	}
}

func parseIp(buff *buffer.Buffer) string {
	parts := make([]string, 0, 4)

	for i := 0; i < 4; i++ {
		b := buff.Pop()
		bStr := strconv.Itoa(int(b))

		parts = append(parts, bStr)
	}

	return strings.Join(parts, ".")
}

func parseCname(buff *buffer.Buffer) string {
	lbl := label_sequence.ParseFrom(buff)

	return lbl.String()
}

func parseIpv6(buff *buffer.Buffer) string {
	sb := strings.Builder{}

	for i := 0; i < 4; i++ {
		b := buff.Pop()
		sb.WriteByte(b)
	}

	return sb.String()
}

func (r *Answer) String() string {
	return fmt.Sprintf("%s\tIN\t%s\t%s", r.lbl, r.rType, r.rData)
}

func (r *Answer) Bytes() []byte {
	return r.packed
}
