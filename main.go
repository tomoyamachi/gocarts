package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tomoyamachi/gocarts/cmd"
)

// Revision of Git
var revision string

func main() {
	var v = flag.Bool("v", false, "Show version")
	flag.Parse()
	if *v {
		fmt.Printf("gocarts %s\n", revision)
		os.Exit(0)
	}

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
