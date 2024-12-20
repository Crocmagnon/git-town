package fixture

import (
	"log"
	"os"
	"path/filepath"

	"github.com/git-town/git-town/v17/test/filesystem"
	"github.com/git-town/git-town/v17/test/helpers"
)

// Factory manages the Git setup for the entire test suite.
// For each scenario, it provides a standardized, empty Fixture consisting of a local and remote Git repository.
//
// Setting up a Fixture is an expensive operation and has to be done for every scenario.
// As a performance optimization, Factory creates a fully set up Fixture (including the main branch and configuration)
// (the "memoized" environment) at the beginning of the test suite and makes copies of it for each scenario.
// Making copies of a fully set up Git repo is much faster than creating it from scratch.
// End-to-end tests run multi-threaded, all threads share a global Factory instance.
type Factory struct {
	Counter helpers.AtomicCounter

	// path of the folder that this class operates in
	Dir string

	// the memoized environment
	memoized Memoized
}

// creates a new FixtureFactory instance
func CreateFactory() Factory {
	baseDir, err := os.MkdirTemp("", "")
	if err != nil {
		log.Fatalf("cannot create base directory for feature specs: %s", err)
	}
	// Evaluate symlinks as Mac temp dir is symlinked
	evalBaseDir, err := filepath.EvalSymlinks(baseDir)
	if err != nil {
		log.Fatalf("cannot evaluate symlinks of base directory for feature specs: %s", err)
	}
	return Factory{
		Counter:  helpers.AtomicCounter{},
		Dir:      evalBaseDir,
		memoized: NewMemoized(filepath.Join(evalBaseDir, "memoized")),
	}
}

// CreateFixture provides a new Fixture for the scenario with the given name.
func (self *Factory) CreateFixture(scenarioName string) Fixture {
	envDirName := filesystem.FolderName(scenarioName) + "_" + self.Counter.NextAsString()
	envPath := filepath.Join(self.Dir, envDirName)
	return self.memoized.CloneInto(envPath)
}

func (self *Factory) Remove() {
	os.RemoveAll(self.Dir)
}
