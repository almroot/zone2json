package main

import (
	"fmt"
	"io"

	"github.com/miekg/dns"
)

type Format string

const (
	FormatBind Format = "bind"
	FormatJSON Format = "json"
)

func (f Format) Valid() error {
	switch f {
	case FormatBind:
		return nil
	case FormatJSON:
		return nil
	default:
		return fmt.Errorf("unrecognized format: %s", f)
	}
}

func (f Format) Generator(output io.Writer, parser *dns.ZoneParser) (func(rr dns.RR) error, error) {
	switch f {
	case FormatBind:
		return NewGeneratorBind(output, parser).Generate, nil
	case FormatJSON:
		return NewGeneratorJson(output, parser).Generate, nil
	default:
		return nil, fmt.Errorf("not implemented: %s", f)
	}
}
