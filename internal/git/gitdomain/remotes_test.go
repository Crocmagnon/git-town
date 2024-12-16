package gitdomain_test

import (
	"testing"

	"github.com/git-town/git-town/v17/internal/git/gitdomain"
	"github.com/git-town/git-town/v17/test/git"
	"github.com/shoenig/test/must"
)

func TestRemotes(t *testing.T) {
	t.Parallel()

	t.Run("HasOrigin", func(t *testing.T) {
		t.Parallel()
		t.Run("origin remote exists", func(t *testing.T) {
			t.Parallel()
			remotes := gitdomain.Remotes{git.RemoteOrigin}
			must.True(t, remotes.Contains(git.RemoteOrigin))
		})
		t.Run("origin remote does not exist", func(t *testing.T) {
			t.Parallel()
			remotes := gitdomain.Remotes{gitdomain.RemoteUpstream}
			must.False(t, remotes.Contains(git.RemoteOrigin))
		})
	})

	t.Run("HasUpstream", func(t *testing.T) {
		t.Parallel()
		t.Run("upstream remote exists", func(t *testing.T) {
			t.Parallel()
			remotes := gitdomain.Remotes{gitdomain.RemoteUpstream}
			must.True(t, remotes.HasUpstream())
		})
		t.Run("upstream remote does not exist", func(t *testing.T) {
			t.Parallel()
			remotes := gitdomain.Remotes{git.RemoteOrigin}
			must.False(t, remotes.HasUpstream())
		})
	})
}
