package dns

type Packet struct {
	header    *Header
	questions []*Question
}

type QueryPacket struct {
	*Packet
}

func (p *Packet) SetHeader(h *Header) {
	p.header = h
}

func (p *Packet) AddQuestion(q *Question) {
	p.header.IncQuestionCount()
	p.questions = append(p.questions, q)
}

func (p *Packet) Bytes() []byte {
	buff := BufferWithCap(512)

	buff.WriteSlice(p.header.Bytes())

	for _, q := range p.questions {
		buff.WriteSlice(q.Bytes())
	}

	b := make([]byte, buff.pos)
	copy(b, buff.data[:buff.pos])

	return b
}

func NewQuery(domain string, r RecordType) *QueryPacket {
	h := QueryHeader()
	q := NewQuestion(domain, r)

	p := &Packet{
		header:    h,
		questions: nil,
	}

	p.AddQuestion(q)

	return &QueryPacket{p}
}

func IpQuery(domain string) *QueryPacket {
	return NewQuery(domain, A)
}

func CNameQuery(domain string) *QueryPacket {
	return NewQuery(domain, CNAME)
}

func Ipv6Query(domain string) *QueryPacket {
	return NewQuery(domain, AAAA)
}
