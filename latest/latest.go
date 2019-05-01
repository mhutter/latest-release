// Package latest contains helper methods of all kinds
package latest

import (
	"strings"

	"github.com/blang/semver"
)

// ExpandRepo checks whether "repo" contains a "/", and expands to
// "repo/repo" otherwise.
func ExpandRepo(repo string) string {
	if strings.ContainsRune(repo, '/') {
		return repo
	}
	return repo + "/" + repo
}

// All simply passes through all Releases
func All(in <-chan *Release) <-chan *Release {
	out := make(chan *Release)
	go func() {
		for r := range in {
			out <- r
		}
		close(out)
	}()
	return out
}

// NoPrereleases filters out all Release that are prereleases
func NoPrereleases(in <-chan *Release) <-chan *Release {
	out := make(chan *Release)
	go func() {
		for r := range in {
			if !r.Prerelease {
				out <- r
			}
		}
		close(out)
	}()
	return out
}

// OnlyPrerelease filters out all Release that are NOT prereleases
func OnlyPrerelease(in <-chan *Release) <-chan *Release {
	out := make(chan *Release)
	go func() {
		for r := range in {
			if r.Prerelease {
				out <- r
			}
		}
		close(out)
	}()
	return out
}

// Latest reads all releases from IN and returns the latest
func Latest(in <-chan *Release) *Release {
	var l *Release
	var lv semver.Version
	for r := range in {
		v := semver.MustParse(strings.TrimPrefix(r.TagName, "v"))
		if l == nil || v.GT(lv) {
			l = r
			lv = v
		}
	}

	return l
}
