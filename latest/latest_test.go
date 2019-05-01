package latest_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mhutter/latest-release/latest"
)

func TestExpandRepo(t *testing.T) {
	for in, expected := range map[string]string{
		"/":                      "/",
		"mhutter/latest-release": "mhutter/latest-release",
		"kubernetes":             "kubernetes/kubernetes",
	} {
		assert.Equal(t, expected, latest.ExpandRepo(in))
	}
}

var testData = latest.Releases{
	&latest.Release{TagName: "v1.15.0-alpha.2", Prerelease: true},
	&latest.Release{TagName: "v1.12.8"},
	&latest.Release{TagName: "v1.15.0-alpha.1", Prerelease: true},
	&latest.Release{TagName: "v1.14.1"},
	&latest.Release{TagName: "v1.9.0"},
}

func TestAll(t *testing.T) {
	i := 0
	for r := range latest.All(testData.Emit()) {
		assert.Equal(t, testData[i].TagName, r.TagName)
		i++
	}
}

func TestNoPrereleases(t *testing.T) {
	for r := range latest.NoPrereleases(testData.Emit()) {
		assert.False(t, r.Prerelease)
	}
}

func TestOnlyPrerelease(t *testing.T) {
	for r := range latest.OnlyPrerelease(testData.Emit()) {
		assert.True(t, r.Prerelease)
	}
}

func TestLatest(t *testing.T) {
	l := latest.Latest(testData.Emit())
	assert.Equal(t, "v1.15.0-alpha.2", l.TagName)
}
