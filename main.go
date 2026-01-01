package main

import (
	"fmt"
	"os"

	"github.com/ismailtsdln/socialrecon/cmd/socialrecon"
)

func main() {
	if err := socialrecon.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
