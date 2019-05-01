package latest

// Release contains relevant fields of a GitHub release
type Release struct {
	TagName    string `json:"tag_name"`
	Prerelease bool   `json:"prerelease"`
}

// Releases represents a list of releases :tada:
type Releases []*Release

// Filter represents a function that filters releases for certain criteria
type Filter func(<-chan *Release) <-chan *Release

// Emit spews out all releases in a channel
func (r Releases) Emit() <-chan *Release {
	out := make(chan *Release)

	go func() {
		for _, rel := range r {
			out <- rel
		}
		close(out)
	}()

	return out
}
