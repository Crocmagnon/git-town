Feature: make another remote branch a contribution branch

  Background:
    Given a Git repo with origin
    And the branches
      | NAME           | TYPE    | PARENT | LOCATIONS |
      | remote-feature | feature | main   | origin    |
    And I run "git fetch"
    When I run "git-town contribute remote-feature"

  Scenario: result
    Then it runs the commands
      | BRANCH | COMMAND                     |
      |        | git checkout remote-feature |
    And it prints:
      """
      branch "remote-feature" is now a contribution branch
      """
    And branch "remote-feature" is now a contribution branch
    And the current branch is now "remote-feature"

  Scenario: undo
    When I run "git-town undo"
    Then it runs the commands
      | BRANCH         | COMMAND                      |
      | remote-feature | git checkout main            |
      | main           | git branch -D remote-feature |
    And the current branch is now "main"
    And there are now no contribution branches
