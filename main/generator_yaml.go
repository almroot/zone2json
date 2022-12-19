package main

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"

	"github.com/miekg/dns"
)

type GeneratorYaml struct {
	output io.Writer
	parser *dns.ZoneParser
}

func NewGeneratorYaml(output io.Writer, parser *dns.ZoneParser) *GeneratorYaml {
	return &GeneratorYaml{
		output: output,
		parser: parser,
	}
}

func (g *GeneratorYaml) Generate(rr dns.RR) (err error) {
	var data []byte
	var document = make(map[string]dns.RR)
	document["record"] = rr
	if data, err = yaml.Marshal(document); err == nil {
		_, err = fmt.Fprint(g.output, string(data))
	}
	return
}
