package packet

import (
	"github.com/volf52/dnsgo/pkg/dns/answer"
	"github.com/volf52/dnsgo/pkg/dns/buffer"
	"github.com/volf52/dnsgo/pkg/dns/header"
	"github.com/volf52/dnsgo/pkg/dns/question"
)

type Packet struct {
	header    *header.Header
	questions []*question.Question
	answers   []*answer.Answer
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
