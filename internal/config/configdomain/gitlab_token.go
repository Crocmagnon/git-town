package configdomain

import (
	"strings"

	. "github.com/git-town/git-town/v17/pkg/prelude"
)

// GitLabToken is a bearer token to use with the GitLab API.
type GitLabToken string

func (self GitLabToken) String() string {
	return string(self)
}

func ParseGitLabToken(value string) Option[GitLabToken] {
	value = strings.TrimSpace(value)
	if value == "" {
		return None[GitLabToken]()
	}
	return Some(GitLabToken(value))
}
