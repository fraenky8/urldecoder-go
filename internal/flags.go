package internal

import (
	"flag"
	"io"
	"os"
	"strings"
)

func NewInOut() (io.ReadCloser, io.WriteCloser, error) {
	fileFlag := flag.String("file", "", "Specify a file to read from.")
	outFlag := flag.String("output", "", "Write result to FILE instead of standard output.")
	flag.Parse()

	var in io.ReadCloser
	switch {
	case *fileFlag != "":
		f, err := os.Open(*fileFlag)
		if err != nil {
			return nil, nil, err
		}
		in = f
	case len(flag.Args()) > 0:
		in = io.NopCloser(strings.NewReader(strings.Join(flag.Args(), " ")))
	default:
		in = os.Stdin
	}

	var out io.WriteCloser = os.Stdout
	if *outFlag != "" {
		f, err := os.Create(*outFlag)
		if err != nil {
			return nil, nil, err
		}
		out = f
	}

	return in, out, nil
}
