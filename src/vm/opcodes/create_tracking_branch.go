package opcodes

import (
	"github.com/git-town/git-town/v14/src/git/gitdomain"
	"github.com/git-town/git-town/v14/src/vm/shared"
)

// CreateTrackingBranch pushes the given local branch up to origin
// and marks it as tracking the current branch.
type CreateTrackingBranch struct {
	Branch                  gitdomain.LocalBranchName
	undeclaredOpcodeMethods `exhaustruct:"optional"`
}

func (self *CreateTrackingBranch) CreateContinueProgram() []shared.Opcode {
	return []shared.Opcode{self}
}

func (self *CreateTrackingBranch) Run(args shared.RunArgs) error {
	return args.Git.CreateTrackingBranch(args.Frontend, self.Branch, gitdomain.RemoteOrigin, args.Config.Config.NoPushHook())
}
