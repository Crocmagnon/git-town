package configdomain

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/git-town/git-town/v14/src/git/gitdomain"
	. "github.com/git-town/git-town/v14/src/gohacks/prelude"
)

// Key contains all the keys used in Git Town's Git metadata configuration.
type Key string

func (self Key) IsAliasKey() bool {
	return strings.HasPrefix(self.String(), "alias.")
}

// IsLineage indicates using the returned option whether this key is a lineage key.
// The option contains the name of the child branch of the lineage entry.
// This method returns a string instead of a gitdomain.LocalBranchName to indicate that
// the returned child name isn't verified and might be empty.
func (self Key) IsLineage() bool {
	return isLineageKey(self.String())
}

// MarshalJSON is used when serializing this LocalBranchName to JSON.
func (self Key) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.String())
}

func (self Key) String() string { return string(self) }

// UnmarshalJSON is used when de-serializing JSON into a Location.
func (self *Key) UnmarshalJSON(b []byte) error {
	value := ""
	err := json.Unmarshal(b, &value)
	*self = Key(value)
	return err
}

const (
	KeyAliasAppend                         = Key("alias.append")
	KeyAliasCompress                       = Key("alias.compress")
	KeyAliasContribute                     = Key("alias.contribute")
	KeyAliasDiffParent                     = Key("alias.diff-parent")
	KeyAliasHack                           = Key("alias.hack")
	KeyAliasKill                           = Key("alias.kill")
	KeyAliasObserve                        = Key("alias.observe")
	KeyAliasPark                           = Key("alias.park")
	KeyAliasPrepend                        = Key("alias.prepend")
	KeyAliasPropose                        = Key("alias.propose")
	KeyAliasRenameBranch                   = Key("alias.rename-branch")
	KeyAliasRepo                           = Key("alias.repo")
	KeyAliasSetParent                      = Key("alias.set-parent")
	KeyAliasShip                           = Key("alias.ship")
	KeyAliasSync                           = Key("alias.sync")
	KeyContributionBranches                = Key("git-town.contribution-branches")
	KeyCreatePrototypeBranches             = Key("git-town.create-prototype-branches")
	KeyDeprecatedCodeHostingDriver         = Key("git-town.code-hosting-driver")
	KeyDeprecatedCodeHostingOriginHostname = Key("git-town.code-hosting-origin-hostname")
	KeyDeprecatedCodeHostingPlatform       = Key("git-town.code-hosting-platform")
	KeyDeprecatedMainBranchName            = Key("git-town.main-branch-name")
	KeyDeprecatedNewBranchPushFlag         = Key("git-town.new-branch-push-flag")
	KeyDeprecatedPerennialBranchNames      = Key("git-town.perennial-branch-names")
	KeyDeprecatedPullBranchStrategy        = Key("git-town.pull-branch-strategy")
	KeyDeprecatedPushVerify                = Key("git-town.push-verify")
	KeyDeprecatedShipDeleteRemoteBranch    = Key("git-town.ship-delete-remote-branch")
	KeyDeprecatedSyncStrategy              = Key("git-town.sync-strategy")
	KeyGiteaToken                          = Key("git-town.gitea-token")
	KeyGithubToken                         = Key("git-town.github-token")
	KeyGitlabToken                         = Key("git-town.gitlab-token")
	KeyHostingOriginHostname               = Key("git-town.hosting-origin-hostname")
	KeyHostingPlatform                     = Key("git-town.hosting-platform")
	KeyMainBranch                          = Key("git-town.main-branch")
	KeyObservedBranches                    = Key("git-town.observed-branches")
	KeyOffline                             = Key("git-town.offline")
	KeyParkedBranches                      = Key("git-town.parked-branches")
	KeyPerennialBranches                   = Key("git-town.perennial-branches")
	KeyPerennialRegex                      = Key("git-town.perennial-regex")
	KeyPrototypeBranches                   = Key("git-town.prototype-branches")
	KeyPushHook                            = Key("git-town.push-hook")
	KeyPushNewBranches                     = Key("git-town.push-new-branches")
	KeyShipDeleteTrackingBranch            = Key("git-town.ship-delete-tracking-branch")
	KeySyncBeforeShip                      = Key("git-town.sync-before-ship")
	KeySyncFeatureStrategy                 = Key("git-town.sync-feature-strategy")
	KeySyncPerennialStrategy               = Key("git-town.sync-perennial-strategy")
	KeySyncPrototypeStrategy               = Key("git-town.sync-prototype-strategy")
	KeySyncStrategy                        = Key("git-town.sync-strategy")
	KeySyncUpstream                        = Key("git-town.sync-upstream")
	KeyGitUserEmail                        = Key("user.email")
	KeyGitUserName                         = Key("user.name")
)

var keys = []Key{ //nolint:gochecknoglobals
	KeyHostingOriginHostname,
	KeyHostingPlatform,
	KeyContributionBranches,
	KeyCreatePrototypeBranches,
	KeyDeprecatedCodeHostingDriver,
	KeyDeprecatedCodeHostingOriginHostname,
	KeyDeprecatedCodeHostingPlatform,
	KeyDeprecatedMainBranchName,
	KeyDeprecatedNewBranchPushFlag,
	KeyDeprecatedPerennialBranchNames,
	KeyDeprecatedPullBranchStrategy,
	KeyDeprecatedPushVerify,
	KeyDeprecatedShipDeleteRemoteBranch,
	KeyDeprecatedSyncStrategy,
	KeyGiteaToken,
	KeyGithubToken,
	KeyGitlabToken,
	KeyGitUserEmail,
	KeyGitUserName,
	KeyMainBranch,
	KeyObservedBranches,
	KeyOffline,
	KeyParkedBranches,
	KeyPerennialBranches,
	KeyPerennialRegex,
	KeyPrototypeBranches,
	KeyPushHook,
	KeyPushNewBranches,
	KeyShipDeleteTrackingBranch,
	KeySyncBeforeShip,
	KeySyncFeatureStrategy,
	KeySyncPerennialStrategy,
	KeySyncPrototypeStrategy,
	KeySyncStrategy,
	KeySyncUpstream,
}

func AliasableCommandForKey(key Key) Option[AliasableCommand] {
	for _, aliasableCommand := range AllAliasableCommands() {
		if KeyForAliasableCommand(aliasableCommand) == key {
			return Some(aliasableCommand)
		}
	}
	return None[AliasableCommand]()
}

func KeyForAliasableCommand(aliasableCommand AliasableCommand) Key {
	switch aliasableCommand {
	case AliasableCommandAppend:
		return KeyAliasAppend
	case AliasableCommandCompress:
		return KeyAliasCompress
	case AliasableCommandContribute:
		return KeyAliasContribute
	case AliasableCommandDiffParent:
		return KeyAliasDiffParent
	case AliasableCommandHack:
		return KeyAliasHack
	case AliasableCommandKill:
		return KeyAliasKill
	case AliasableCommandObserve:
		return KeyAliasObserve
	case AliasableCommandPark:
		return KeyAliasPark
	case AliasableCommandPrepend:
		return KeyAliasPrepend
	case AliasableCommandPropose:
		return KeyAliasPropose
	case AliasableCommandRenameBranch:
		return KeyAliasRenameBranch
	case AliasableCommandRepo:
		return KeyAliasRepo
	case AliasableCommandSetParent:
		return KeyAliasSetParent
	case AliasableCommandShip:
		return KeyAliasShip
	case AliasableCommandSync:
		return KeyAliasSync
	}
	panic(fmt.Sprintf("don't know how to convert alias type %q into a config key", &aliasableCommand))
}

func NewParentKey(branch gitdomain.LocalBranchName) Key {
	return Key(LineageKeyPrefix + branch + LineageKeySuffix)
}

func ParseKey(name string) Option[Key] {
	for _, configKey := range keys {
		if configKey.String() == name {
			return Some(configKey)
		}
	}
	if isLineageKey(name) {
		return Some(Key(name))
	}
	for _, aliasableCommand := range AllAliasableCommands() {
		key := KeyForAliasableCommand(aliasableCommand)
		if key.String() == name {
			return Some(key)
		}
	}
	return None[Key]()
}

const (
	LineageKeyPrefix = "git-town-branch."
	LineageKeySuffix = ".parent"
)

func isLineageKey(key string) bool {
	return strings.HasPrefix(key, LineageKeyPrefix) && strings.HasSuffix(key, LineageKeySuffix)
}

// DeprecatedKeys defines the up-to-date counterparts to deprecated configuration settings.
var DeprecatedKeys = map[Key]Key{ //nolint:gochecknoglobals
	KeyDeprecatedCodeHostingDriver:         KeyHostingPlatform,
	KeyDeprecatedCodeHostingOriginHostname: KeyHostingOriginHostname,
	KeyDeprecatedCodeHostingPlatform:       KeyHostingPlatform,
	KeyDeprecatedMainBranchName:            KeyMainBranch,
	KeyDeprecatedNewBranchPushFlag:         KeyPushNewBranches,
	KeyDeprecatedPerennialBranchNames:      KeyPerennialBranches,
	KeyDeprecatedPullBranchStrategy:        KeySyncPerennialStrategy,
	KeyDeprecatedPushVerify:                KeyPushHook,
	KeyDeprecatedShipDeleteRemoteBranch:    KeyShipDeleteTrackingBranch,
	KeyDeprecatedSyncStrategy:              KeySyncFeatureStrategy,
}