package main

import (
	"fmt"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func printVersion() {
	fmt.Printf("%s, commit %s, built on %s\n", version, commit, date)
}
