package main

import (
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/fraenky8/urldecoder-go/internal"
)

func main() {
	in, out, err := internal.NewInOut()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = urlencode(in, out)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func urlencode(r io.ReadCloser, w io.WriteCloser) error {
	defer r.Close()
	defer w.Close()

	b, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("could not read: %w", err)
	}

	_, err = fmt.Fprint(w, url.QueryEscape(string(b)))
	return err
}
