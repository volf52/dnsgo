package packet

import (
	"github.com/volf52/dnsgo/pkg/dns/buffer"
	"github.com/volf52/dnsgo/pkg/dns/header"
	"github.com/volf52/dnsgo/pkg/dns/question"
	"github.com/volf52/dnsgo/pkg/dns/record_type"
)

type Packet struct {
	header    *header.Header
	questions []*question.Question
}

type QueryPacket struct {
	*Packet
}

func (p *Packet) SetHeader(h *header.Header) {
	p.header = h
}

func (p *Packet) AddQuestion(q *question.Question) {
	p.header.IncQuestionCount()
	p.questions = append(p.questions, q)
}

func (p *Packet) Bytes() []byte {
	buff := buffer.WithCap(512)

	buff.WriteSlice(p.header.Bytes())

	for _, q := range p.questions {
		buff.WriteSlice(q.Bytes())
	}

	b := make([]byte, buff.Pos())
	copy(b, buff.Slice(0, buff.Pos()))

	return b
}

func NewQuery(domain string, r record_type.RecordType) *QueryPacket {
	h := header.ForQuery()
	q := question.New(domain, r)

	p := &Packet{
		header:    h,
		questions: nil,
	}

	p.AddQuestion(q)

	return &QueryPacket{p}
}

func IpQuery(domain string) *QueryPacket {
	return NewQuery(domain, record_type.A)
}

func CNameQuery(domain string) *QueryPacket {
	return NewQuery(domain, record_type.CNAME)
}

func Ipv6Query(domain string) *QueryPacket {
	return NewQuery(domain, record_type.AAAA)
}
