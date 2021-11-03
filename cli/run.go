package cli

import (
	"os"

	"github.com/itchyny/gojq"
)

// Run gojq.
func Run(opts ...gojq.CompilerOption) int {
	return (&cli{
		inStream:  os.Stdin,
		outStream: os.Stdout,
		errStream: os.Stderr,
		options:   opts,
	}).run(os.Args[1:])
}
