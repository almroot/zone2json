package main

import (
	"fmt"
	"io"

	"github.com/miekg/dns"
)

type GeneratorBind struct {
	output io.Writer
	parser *dns.ZoneParser
}

func NewGeneratorBind(output io.Writer, parser *dns.ZoneParser) *GeneratorBind {
	return &GeneratorBind{
		output: output,
		parser: parser,
	}
}

func (g *GeneratorBind) Generate(rr dns.RR) (err error) {
	_, err = fmt.Fprintln(g.output, rr.String())
	return
}
