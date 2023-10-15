package opcode

import "github.com/git-town/git-town/v9/src/vm/shared"

// AbortMerge aborts the current merge conflict.
type AbortMerge struct {
	undeclaredOpcodeMethods
}

func (op *AbortMerge) Run(args shared.RunArgs) error {
	return args.Runner.Frontend.AbortMerge()
}