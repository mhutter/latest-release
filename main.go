package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/mhutter/latest-release/latest"
)

func main() {
	if versionFlag {
		printVersion()
		os.Exit(0)
	}

	if flag.NArg() != 1 {
		usage()
		os.Exit(1)
	}

	repo := latest.ExpandRepo(flag.Arg(0))
	releases, err := FetchReleases(repo)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	filter := latest.All
	if onlyPre {
		filter = latest.OnlyPrerelease
	} else if !includePre {
		filter = latest.NoPrereleases
	}

	latestRelease := latest.Latest(filter(releases.Emit()))
	fmt.Println(latestRelease.TagName)
}

// FetchReleases fetches the latest releases for the given repository
// from the GitHub API
func FetchReleases(repo string) (latest.Releases, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases", repo)
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("Repository '%s' does not exist on GitHub", repo)
	}

	var releases latest.Releases
	err = json.NewDecoder(res.Body).Decode(&releases)
	return releases, err
}
