package github

import "github.com/git-town/git-town/v15/internal/git/giturl"

// Detect indicates whether the current repository is hosted on a GitHub server.
func Detect(originURL giturl.Parts) bool {
	return originURL.Host == "github.com"
}
