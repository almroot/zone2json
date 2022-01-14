package main

import (
	"fmt"
	"io"
	"os"

	"github.com/miekg/dns"
)

func main() {

	// Parse the arguments...
	args := NewArguments()
	if code, terminate := args.Parse(); terminate {
		os.Exit(code)
	}

	// Figure out how to treat I/O.
	var filename = ""
	var input io.Reader
	if args.Input == "" || args.Input == "-" {
		input = os.Stdin
	}
	var output io.Writer
	if args.Output == "" || args.Output == "-" {
		output = os.Stdout
	}

	// Setup the zone parser...
	parser := dns.NewZoneParser(input, args.DefaultOrigin, filename)
	parser.SetIncludeAllowed(args.AllowIncludes)
	parser.SetDefaultTTL(args.DefaultTTL)

	// Setup the DNS RR interpreters (and output format generator).
	interpreter, err := args.Format.Generator(output, parser)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Exhaust the zone contents...
	for rr, ok := parser.Next(); ok; rr, ok = parser.Next() {
		err = interpreter(rr)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}
}
