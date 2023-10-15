package opcode

import (
	"github.com/git-town/git-town/v9/src/domain"
	"github.com/git-town/git-town/v9/src/vm/shared"
)

// CreateTrackingBranch pushes the given local branch up to origin
// and marks it as tracking the current branch.
type CreateTrackingBranch struct {
	Branch     domain.LocalBranchName
	NoPushHook bool
	undeclaredOpcodeMethods
}

func (op *CreateTrackingBranch) Run(args shared.RunArgs) error {
	return args.Runner.Frontend.CreateTrackingBranch(op.Branch, domain.OriginRemote, op.NoPushHook)
}