package main

import (
	"os"

	"github.com/jessevdk/go-flags"
)

type Arguments struct {
	Input         string `short:"i" long:"input" description:"The file to read from, use - for STDIN"`
	Output        string `short:"o" long:"output" description:"The file to write to, use - for STDOUT"`
	Format        Format `short:"f" long:"format" description:"Output format: bind/json"`
	DefaultOrigin string `long:"origin" description:"The default origin for relative domains"`
	DefaultTTL    uint32 `long:"ttl" description:"The default TTL to be used"`
	AllowIncludes bool   `long:"allow-includes" description:"Enables support for bind $INCLUDE directives"`
}

func NewArguments() *Arguments {
	return &Arguments{
		Input:         "-",
		Output:        "-",
		Format:        FormatBind,
		DefaultOrigin: ".",
		DefaultTTL:    86400,
		AllowIncludes: false,
	}
}

func (a *Arguments) Parse() (int, bool) {
	_, err := flags.ParseArgs(a, os.Args)
	if err == nil {
		err = a.Format.Valid()
	}
	if err == nil {
		return 0, false
	} else if flags.WroteHelp(err) {
		return 0, true
	} else {
		return 1, true
	}
}
