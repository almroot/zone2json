package main

import (
	"fmt"
	"io"

	"github.com/miekg/dns"
)

type Format string

const (
	FormatBind Format = "bind"
	FormatJson Format = "json"
	FormatYaml Format = "yaml"
	FormatXml  Format = "xml"
)

func (f Format) Valid() error {
	switch f {
	case FormatBind:
		return nil
	case FormatJson:
		return nil
	case FormatYaml:
		return nil
	case FormatXml:
		return nil
	default:
		return fmt.Errorf("unrecognized format: %s", f)
	}
}

func (f Format) Generator(output io.Writer, parser *dns.ZoneParser) (func(rr dns.RR) error, error) {
	switch f {
	case FormatBind:
		return NewGeneratorBind(output, parser).Generate, nil
	case FormatJson:
		return NewGeneratorJson(output, parser).Generate, nil
	case FormatYaml:
		return NewGeneratorYaml(output, parser).Generate, nil
	case FormatXml:
		return NewGeneratorXml(output, parser).Generate, nil
	default:
		return nil, fmt.Errorf("not implemented: %s", f)
	}
}
