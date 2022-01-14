package main

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/miekg/dns"
)

type GeneratorXml struct {
	output io.Writer
	parser *dns.ZoneParser
}

func NewGeneratorXml(output io.Writer, parser *dns.ZoneParser) *GeneratorXml {
	return &GeneratorXml{
		output: output,
		parser: parser,
	}
}

func (g *GeneratorXml) Generate(rr dns.RR) (err error) {
	var data []byte
	if data, err = xml.Marshal(rr); err == nil {
		_, err = fmt.Fprintln(g.output, string(data))
	}
	return
}
