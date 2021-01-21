package main

import (
	"fmt"
	"os"

	"github.com/TiraTom/cobraStudy/cmd"
)

func main() {
	if err := cmd.CaesarCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
