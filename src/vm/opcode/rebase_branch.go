package opcode

import (
	"github.com/git-town/git-town/v9/src/domain"
	"github.com/git-town/git-town/v9/src/vm/shared"
)

// RebaseBranch rebases the current branch
// against the branch with the given name.
type RebaseBranch struct {
	Branch domain.BranchName
	undeclaredOpcodeMethods
}

func (op *RebaseBranch) CreateAbortProgram() []shared.Opcode {
	return []shared.Opcode{&AbortRebase{}}
}

func (op *RebaseBranch) CreateContinueProgram() []shared.Opcode {
	return []shared.Opcode{&ContinueRebase{}}
}

func (op *RebaseBranch) Run(args shared.RunArgs) error {
	return args.Runner.Frontend.Rebase(op.Branch)
}