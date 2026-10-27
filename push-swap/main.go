package main

import (
	"fmt"
	"io"
	"os"

	"push-swap/internal/pushswap"
)

// main lit les nombres, calcule les instructions de tri et les affiche une par ligne.
func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		os.Exit(1)
	}
}

// run validates arguments and writes the generated instructions to output.
func run(arguments []string, output, errorOutput io.Writer) error {
	if len(arguments) == 0 {
		return nil
	}

	values, err := pushswap.ParseArguments(arguments)
	if err != nil {
		fmt.Fprintln(errorOutput, "Error")
		return err
	}

	for _, instruction := range pushswap.Generate(values) {
		if _, err := fmt.Fprintln(output, instruction); err != nil {
			return err
		}
	}
	return nil
}
