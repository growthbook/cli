// Hand-written (NOT generated). Pins flag shorthands the generator would otherwise move:
// it gives a flag its first letter only while no sibling flag shares it.
package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// pinnedShorthands maps group -> leaf -> flag -> shorthand; "" strips one.
var pinnedShorthands = map[string]map[string]map[string]string{
	"features": {
		"create": {"archived": "a"},
		"update": {"archived": "a"},
	},
	// A guardrail bypass shouldn't be one keystroke away.
	"feature-revisions": {
		"add-rule":    {"allow-unscoped-rules": ""},
		"update-rule": {"allow-unscoped-rules": ""},
	},
}

func pinShorthands(rootCmd *cobra.Command) {
	for _, group := range rootCmd.Commands() {
		for _, leaf := range group.Commands() {
			pins, ok := pinnedShorthands[group.Name()][leaf.Name()]
			if !ok {
				continue
			}
			// pflag indexes shorthands at registration, so re-register the local flags.
			var flags []*pflag.Flag
			leaf.LocalNonPersistentFlags().VisitAll(func(f *pflag.Flag) {
				if s, ok := pins[f.Name]; ok {
					f.Shorthand = s
				}
				flags = append(flags, f)
			})
			leaf.ResetFlags()
			for _, f := range flags {
				leaf.Flags().AddFlag(f)
			}
		}
	}
}
