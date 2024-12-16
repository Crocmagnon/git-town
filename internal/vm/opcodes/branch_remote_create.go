package opcodes

import (
	"github.com/git-town/git-town/v17/internal/git/gitdomain"
	"github.com/git-town/git-town/v17/internal/vm/shared"
)

// BranchRemoteCreate pushes the given local branch up to origin.
type BranchRemoteCreate struct {
	Branch                  gitdomain.LocalBranchName
	SHA                     gitdomain.SHA
	undeclaredOpcodeMethods `exhaustruct:"optional"`
}

func (self *BranchRemoteCreate) Run(args shared.RunArgs) error {
	return args.Git.CreateRemoteBranch(args.Frontend, self.SHA, self.Branch, args.Config.Value.NormalConfig.DevRemote, args.Config.Value.NormalConfig.NoPushHook())
}
