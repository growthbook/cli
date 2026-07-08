// Hand-written (NOT generated). All customization of the generated root command
// is funneled through customRegister, so the ONLY edit to the generated root.go
// is a single `customRegister(rootCmd)` line (see .speakeasy/patches/). Keeping
// the patch to one line makes it robust across regenerations — there are no
// multi-line hunks whose context can drift.
package cli

import (
	"github.com/growthbook/cli/internal/customcfg"
	"github.com/growthbook/cli/internal/usage"
	"github.com/spf13/cobra"
)

// customRegister wires the hand-written layer onto the generated root command:
// the custom persistent flags, the extra top-level commands, and the startup
// hook (chained after the generated PersistentPreRunE).
func customRegister(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().String("profile", "", "Use a named credential/server profile (manage with `growthbook profiles`)")
	rootCmd.PersistentFlags().Bool("no-update-check", false, "Disable the once-a-day check for a newer CLI version")

	// Add a top-level `--version` flag alongside the existing `version` subcommand.
	// Setting Version makes Cobra register the flag and short-circuit on it; the
	// template matches the subcommand's output ("growthbook <version>").
	rootCmd.Version = Version
	rootCmd.SetVersionTemplate("growthbook {{.Version}}\n")

	rootCmd.AddCommand(newGenerateTypesCmd())
	rootCmd.AddCommand(newProfilesCmd())

	// Run custom startup after the generated PersistentPreRunE (which initializes
	// config), but skip it for --usage/schema introspection, exactly as the
	// generated code does for its own setup.
	genPreRun := rootCmd.PersistentPreRunE
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if genPreRun != nil {
			if err := genPreRun(cmd, args); err != nil {
				return err
			}
		}
		if usage.UsageRequested(cmd) {
			return nil
		}
		return customcfg.OnStartup(cmd)
	}
}
