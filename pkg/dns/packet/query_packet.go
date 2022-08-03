package packet

import (
	"github.com/volf52/dnsgo/pkg/dns/header"
	"github.com/volf52/dnsgo/pkg/dns/question"
	"github.com/volf52/dnsgo/pkg/dns/record_type"
)

type QueryPacket struct {
	*Packet
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
