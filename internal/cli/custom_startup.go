// Hand-written (NOT generated). Bridges the generated root command to custom
// startup logic without adding imports to the generated root.go: root.go only
// references the same-package identifier customOnStartup, so its generated
// import block is never edited (keeping the persistentEdits merge trivial).
package cli

import (
	"github.com/growthbook/cli/internal/customcfg"
	"github.com/spf13/cobra"
)

func customOnStartup(cmd *cobra.Command) error {
	return customcfg.OnStartup(cmd)
}
