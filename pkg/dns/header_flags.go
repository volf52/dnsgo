package dns

import (
	"fmt"
	"github.com/volf52/dnsgo/internal/utils"
)

const (
	QrMask uint16 = 0x8000
	AaMask uint16 = 0x0400
	TcMask uint16 = 0x0200
	RdMask uint16 = 0x0100
	RaMask uint16 = 0x0080
)

type HeaderFlags struct {
	qr bool
	aa bool
	tc bool
	rd bool
	ra bool

	val uint16
}

func QueryFlags() *HeaderFlags {
	return &HeaderFlags{
		qr: false,
		aa: false,
		tc: false,
		rd: true,
		ra: false,
	}
}

func ParseFlags(container uint16) *HeaderFlags {
	return &HeaderFlags{
		qr:  utils.IsSet(container, QrMask),
		aa:  utils.IsSet(container, AaMask),
		tc:  utils.IsSet(container, TcMask),
		rd:  utils.IsSet(container, RdMask),
		ra:  utils.IsSet(container, RaMask),
		val: container,
	}
}

func (f *HeaderFlags) Pack() uint16 {
	b := uint16(0)

	if f.qr {
		b |= QrMask
	}

	if f.aa {
		b |= AaMask
	}

	if f.tc {
		b |= TcMask
	}

	if f.rd {
		b |= RdMask
	}

	if f.ra {
		b |= RaMask
	}

	return b
}

func (f *HeaderFlags) String() string {
	return fmt.Sprintf("QR=%t AA=%t TC=%t RD=%t RA=%t",
		f.qr,
		f.aa,
		f.tc,
		f.rd,
		f.ra,
	)
}

func (f *HeaderFlags) Val() uint16 {
	return f.val
}
