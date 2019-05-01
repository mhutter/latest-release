package main

import (
	"flag"
	"fmt"
)

func usage() {
	fmt.Fprintln(flag.CommandLine.Output(), `
Check the latest release of a GitHub repository

USAGE:
  $ latest-release [<flags>] repository

ARGUMENTS:
  repository
        Repository to check, either in the form of ORG/REPO or just REPO,
        in which case it will be expanded to REPO/REPO

FLAGS:`)
	flag.PrintDefaults()
}

var (
	includePre  = false
	onlyPre     = false
	versionFlag = false
)

func init() {
	flag.CommandLine.Usage = usage

	flag.BoolVar(&includePre, "p", includePre, "Include prereleases")
	flag.BoolVar(&onlyPre, "P", onlyPre, "ONLY include prereleases")
	flag.BoolVar(&versionFlag, "v", versionFlag, "Print version and exit")

	flag.Parse()
}
