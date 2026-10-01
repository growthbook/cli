// Hand-written (NOT generated). Keeps the command names renamed in 3.0.0 working
// as deprecated aliases until 4.0.0.
package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// renamedCommands maps group -> new leaf name -> pre-3.0.0 leaf name.
var renamedCommands = map[string]map[string]string{
	"SDK-connections": {"lookup": "lookup-SDK-connection-by-key"},
	"features":        {"list-keys": "get-feature-keys", "get-stale": "get-feature-stale"},
	"features-v1":     {"list-keys": "get-feature-keys", "get-stale": "get-feature-stale"},
	"members":         {"update-role": "update-member-role"},
	"ramp-schedules": {
		"get-status":        "get-ramp-schedule-status",
		"update-lockdown":   "update-ramp-schedule-lockdown",
		"update-monitoring": "update-ramp-schedule-monitoring",
		"update-steps":      "update-ramp-schedule-steps",
	},
	"saved-groups": {"get-references": "get-saved-group-references"},
	"teams":        {"add-members": "add-team-members", "remove-members": "remove-team-member"},
}

func registerRenamedAliases(rootCmd *cobra.Command) {
	for _, group := range rootCmd.Commands() {
		for _, leaf := range group.Commands() {
			if old, ok := renamedCommands[group.Name()][leaf.Name()]; ok {
				leaf.Aliases = append(leaf.Aliases, old)
			}
		}
	}
}

func warnIfRenamedAlias(cmd *cobra.Command) {
	if cmd.Parent() == nil || cmd.CalledAs() == cmd.Name() {
		return
	}
	if old, ok := renamedCommands[cmd.Parent().Name()][cmd.Name()]; ok && cmd.CalledAs() == old {
		fmt.Fprintf(cmd.ErrOrStderr(),
			"growthbook: `%s %s` was renamed to `%s %s` in 3.0.0; the old name will be removed in 4.0.0.\n",
			cmd.Parent().Name(), old, cmd.Parent().Name(), cmd.Name())
	}
}
