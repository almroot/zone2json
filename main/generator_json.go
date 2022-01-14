package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/miekg/dns"
)

type GeneratorJson struct {
	output io.Writer
	parser *dns.ZoneParser
}

func NewGeneratorJson(output io.Writer, parser *dns.ZoneParser) *GeneratorJson {
	return &GeneratorJson{
		output: output,
		parser: parser,
	}
}

func (g *GeneratorJson) Generate(rr dns.RR) (err error) {
	var data []byte
	if data, err = json.Marshal(rr); err == nil {
		_, err = fmt.Fprintln(g.output, string(data))
	}
	return
}
