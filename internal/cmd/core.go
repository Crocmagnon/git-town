// Package cmd defines the Git Town commands.
package cmd

import (
	"github.com/git-town/git-town/v15/internal/cmd/config"
	"github.com/git-town/git-town/v15/internal/cmd/debug"
	"github.com/git-town/git-town/v15/internal/cmd/status"
)

// Execute runs the Cobra stack.
func Execute() error {
	rootCmd := rootCmd()
	rootCmd.AddCommand(appendCmd())
	rootCmd.AddCommand(completionsCmd(&rootCmd))
	rootCmd.AddCommand(compressCmd())
	rootCmd.AddCommand(config.RootCmd())
	rootCmd.AddCommand(continueCmd())
	rootCmd.AddCommand(contributeCmd())
	rootCmd.AddCommand(debug.RootCmd())
	rootCmd.AddCommand(diffParentCommand())
	rootCmd.AddCommand(hackCmd())
	rootCmd.AddCommand(killCommand())
	rootCmd.AddCommand(newPullRequestCommand())
	rootCmd.AddCommand(observeCmd())
	rootCmd.AddCommand(offlineCmd())
	rootCmd.AddCommand(parkCmd())
	rootCmd.AddCommand(proposeCommand())
	rootCmd.AddCommand(prependCommand())
	rootCmd.AddCommand(prototypeCmd())
	rootCmd.AddCommand(renameBranchCommand())
	rootCmd.AddCommand(repoCommand())
	rootCmd.AddCommand(status.RootCommand())
	rootCmd.AddCommand(setParentCommand())
	rootCmd.AddCommand(shipCmd())
	rootCmd.AddCommand(skipCmd())
	rootCmd.AddCommand(switchCmd())
	rootCmd.AddCommand(syncCmd())
	rootCmd.AddCommand(undoCmd())
	return rootCmd.Execute()
}
