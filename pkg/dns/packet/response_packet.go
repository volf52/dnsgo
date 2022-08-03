package packet

import (
	"strings"

	"github.com/volf52/dnsgo/pkg/dns/answer"
	"github.com/volf52/dnsgo/pkg/dns/buffer"
	"github.com/volf52/dnsgo/pkg/dns/header"
	"github.com/volf52/dnsgo/pkg/dns/question"
)

type ResponsePacket struct {
	*Packet
}

func FromResponse(b []byte) *ResponsePacket {
	buff := buffer.From(b)

	header := header.ParseFrom(buff)

	questions := make([]*question.Question, header.QdCount())
	answers := make([]*answer.Answer, header.AnsCount())

	for i := range questions {
		questions[i] = question.ParseFrom(buff)
	}

	for i := range answers {
		answers[i] = answer.ParseFrom(buff)
	}

	p := &Packet{header, questions, answers}

	return &ResponsePacket{p}
}

func (r *ResponsePacket) String() string {
	sb := strings.Builder{}

	sb.WriteString("--------- Header -------- \n")
	sb.WriteString(r.header.String())

	sb.WriteString("\n-------- Questions ---------\n")
	for _, q := range r.questions {
		sb.WriteString(q.String())
		sb.WriteRune('\n')
	}

	sb.WriteString("\n--------- Answers ----------\n")
	for _, a := range r.answers {
		sb.WriteString(a.String())
		sb.WriteRune('\n')
	}

	return sb.String()
}
