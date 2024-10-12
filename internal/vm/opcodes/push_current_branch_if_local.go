package opcodes

import (
	"github.com/git-town/git-town/v16/internal/git/gitdomain"
	"github.com/git-town/git-town/v16/internal/vm/shared"
)

// PushCurrentBranch pushes the current branch to its existing tracking branch.
type PushCurrentBranchIfLocal struct {
	CurrentBranch           gitdomain.LocalBranchName
	undeclaredOpcodeMethods `exhaustruct:"optional"`
}

func (self *PushCurrentBranchIfLocal) Run(args shared.RunArgs) error {
	hasTrackingBranch := args.Git.CurrentBranchHasTrackingBranch(args.Backend)
	if !hasTrackingBranch {
		args.PrependOpcodes(&CreateTrackingBranch{
			Branch: self.CurrentBranch,
		})
	}
	return nil
}