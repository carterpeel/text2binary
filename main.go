package main

import (
	"flag"
	"fmt"
	"os"
	"text2bin/core"
)

var (
	text string
	help bool
	delim string
	bufLen int64
)

func main() {
	// Parse the command line arguments
	flag.StringVar(&text, "text", "", "The string to be converted to its' binary format")
	flag.BoolVar(&help, "help", false, "Displays this help")
	flag.StringVar(&delim, "delim", "", "Set the delimiter between binary values")
	flag.Int64Var(&bufLen, "buffersize", 0, "the size of the buffer to be used (only applies to data piped through stdin)")
	flag.Parse()

	// Initialize the new converter
	conv := core.NewEncoder()

	// Get stdin info
	inStat, err := os.Stdin.Stat()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// Is data coming via Stdin?
	// If yes, use streaming IO to write it to Stdout.
	// This is faster because it doesn't require reading large chunks of data into RAM.
	// (This pisses off the Java fanboys. Get away from my RAM.)
	//
	// ₜₕᵢₛ ᵢₛ ₐ ⱼₒₖₑ, ⱼₐᵥₐ ᵢₛ ₒₖ
	if (inStat.Mode() & os.ModeCharDevice) == 0 {
		if err := conv.ConvertAndWrite(os.Stdin, os.Stdout, bufLen, []byte(delim)); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		return
	}

	// If the text flag was empty or the help flag
	// was invoked, print the defaults for all flags & their
	// descriptions.
	if len(text) <= 0  || help {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Print the output of the binary conversion directly to stdout.
	fmt.Println(core.NewEncoder().Encode([]byte(text)).Delim([]byte(delim)).String())

}