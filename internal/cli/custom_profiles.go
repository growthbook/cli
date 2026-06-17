// Hand-written (NOT generated). `growthbook profiles` — manage named
// credential/server profiles (parity with the legacy CLI's profiles).
// Registered from root.go via a same-package reference, so the generated
// import block is never touched.
package cli

import (
	"fmt"

	"github.com/growthbook/cli/internal/customcfg"
	"github.com/spf13/cobra"
)

func newProfilesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "profiles",
		Short: "Manage named credential/server profiles",
		Long: "Profiles let you switch between GrowthBook environments (e.g. cloud, staging, " +
			"self-hosted), each with its own API key and server URL. Select one with " +
			"`profiles use`, override per-command with --profile, or via GBCLI_PROFILE.",
	}
	cmd.AddCommand(newProfilesListCmd(), newProfilesUseCmd(), newProfilesSetCmd(), newProfilesRemoveCmd())
	return cmd
}

func newProfilesListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List configured profiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			f, err := customcfg.Load()
			if err != nil {
				return err
			}
			names := f.Names()
			if len(names) == 0 {
				fmt.Fprintln(cmd.OutOrStdout(), "No profiles configured. Add one with `growthbook profiles set <name>`.")
				return nil
			}
			out := cmd.OutOrStdout()
			for _, n := range names {
				marker := "  "
				if n == f.Current {
					marker = "* "
				}
				url := f.Profiles[n].ServerURL
				if url == "" {
					url = "(default server)"
				}
				secret := "no key"
				if customcfg.GetProfileSecret(n) != "" {
					secret = "key set"
				}
				fmt.Fprintf(out, "%s%s\t%s\t%s\n", marker, n, url, secret)
			}
			return nil
		},
	}
}

func newProfilesUseCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "use <name>",
		Short: "Set the active profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			f, err := customcfg.Load()
			if err != nil {
				return err
			}
			if _, ok := f.Profiles[name]; !ok {
				return fmt.Errorf("profile %q not found (see `growthbook profiles list`)", name)
			}
			f.Current = name
			if err := f.Save(); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Active profile is now %q\n", name)
			return nil
		},
	}
}

func newProfilesSetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set <name>",
		Short: "Create or update a profile",
		Long: "Create or update a profile's server URL and/or API key. The key is read from " +
			"the global --bearer-auth flag and stored in the OS keychain.",
		Example: "  growthbook profiles set staging --server-url https://gb.staging.example.com/api --bearer-auth secret_xxx",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			f, err := customcfg.Load()
			if err != nil {
				return err
			}
			if f.Profiles == nil {
				f.Profiles = map[string]customcfg.Profile{}
			}
			prof := f.Profiles[name]
			if serverURL, _ := cmd.Flags().GetString("server-url"); serverURL != "" {
				prof.ServerURL = serverURL
			}
			f.Profiles[name] = prof
			if f.Current == "" {
				f.Current = name
			}
			if err := f.Save(); err != nil {
				return err
			}
			// The secret comes from the inherited global --bearer-auth flag.
			if secret, _ := cmd.Flags().GetString("bearer-auth"); secret != "" {
				if err := customcfg.SetProfileSecret(name, secret); err != nil {
					return fmt.Errorf("store profile secret: %w", err)
				}
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Saved profile %q\n", name)
			return nil
		},
	}
	return cmd
}

func newProfilesRemoveCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "remove <name>",
		Aliases: []string{"rm"},
		Short:   "Remove a profile",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			f, err := customcfg.Load()
			if err != nil {
				return err
			}
			if _, ok := f.Profiles[name]; !ok {
				return fmt.Errorf("profile %q not found", name)
			}
			delete(f.Profiles, name)
			if f.Current == name {
				f.Current = ""
			}
			_ = customcfg.DeleteProfileSecret(name)
			if err := f.Save(); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Removed profile %q\n", name)
			return nil
		},
	}
}
