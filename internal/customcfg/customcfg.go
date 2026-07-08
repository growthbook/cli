// Package customcfg holds hand-written (NOT generated) configuration that the
// generated config.Config struct can't carry — named profiles, each with a
// server URL and (keychain-stored) secret — plus one-time migration from the
// legacy TS CLI. It lives in its own package and a separate file
// (~/.config/growthbook/cli.yaml) so it survives regeneration; the generated
// config layer is never edited.
//
// Profile resolution sits in front of the generated credential chain
// (flag > env > keychain > config): the active profile injects its server URL
// into the --server-url flag and its secret into GBCLI_BEARER_AUTH, but only
// when the user hasn't supplied those explicitly — so explicit flags/env always
// win and the rest of the chain still applies when no profile is active.
package customcfg

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/growthbook/cli/internal/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// Profile is one named environment. The secret is stored in the OS keychain
// (key profileSecretKey(name)), not on disk.
type Profile struct {
	ServerURL string `yaml:"server_url,omitempty"`
}

// File is the on-disk custom config, kept beside the generated config.yaml.
type File struct {
	Current            string             `yaml:"current,omitempty"`
	Profiles           map[string]Profile `yaml:"profiles,omitempty"`
	MigratedFromLegacy bool               `yaml:"migrated_from_legacy,omitempty"`
}

// Path returns the on-disk location of the custom config (cli.yaml), where
// profiles are stored. Exported for `whoami` to surface alongside config.yaml.
func Path() string {
	return filePath()
}

func filePath() string {
	dir := filepath.Dir(config.GetConfigPath())
	if dir == "" || dir == "." {
		home, err := os.UserHomeDir()
		if err != nil {
			return ""
		}
		dir = filepath.Join(home, ".config", "growthbook")
	}
	return filepath.Join(dir, "cli.yaml")
}

// Load reads the custom config; a missing file yields an empty File (not an error).
func Load() (*File, error) {
	p := filePath()
	if p == "" {
		return &File{}, nil
	}
	data, err := os.ReadFile(p)
	if os.IsNotExist(err) {
		return &File{}, nil
	}
	if err != nil {
		return nil, err
	}
	var f File
	if err := yaml.Unmarshal(data, &f); err != nil {
		return nil, err
	}
	return &f, nil
}

// Save writes the custom config (0600, dir 0700).
func (f *File) Save() error {
	p := filePath()
	if p == "" {
		return fmt.Errorf("cannot determine config directory")
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o700); err != nil {
		return err
	}
	data, err := yaml.Marshal(f)
	if err != nil {
		return err
	}
	return os.WriteFile(p, data, 0o600)
}

// Names returns the profile names, sorted.
func (f *File) Names() []string {
	names := make([]string, 0, len(f.Profiles))
	for n := range f.Profiles {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}

// profileSecretKey namespaces a profile's secret in the keychain so it never
// collides with the generated CLI's own "bearer-auth" entry (auth login).
func profileSecretKey(name string) string { return "profile." + name + ".bearer-auth" }

// GetProfileSecret / SetProfileSecret / DeleteProfileSecret manage per-profile
// secrets in the OS keychain.
func GetProfileSecret(name string) string { return config.GetKeyringValue(profileSecretKey(name)) }
func SetProfileSecret(name, secret string) error {
	return config.SetKeyringValue(profileSecretKey(name), secret)
}
func DeleteProfileSecret(name string) error { return config.DeleteKeyringValue(profileSecretKey(name)) }

// activeProfileName resolves which profile applies: explicit --profile flag,
// then GBCLI_PROFILE env, then the configured current profile.
func activeProfileName(cmd *cobra.Command, f *File) (name string, explicit bool) {
	if cmd != nil {
		if fl := cmd.Flags().Lookup("profile"); fl != nil && fl.Changed && fl.Value.String() != "" {
			return fl.Value.String(), true
		}
	}
	if v := os.Getenv("GBCLI_PROFILE"); v != "" {
		return v, true
	}
	return f.Current, false
}

// OnStartup runs custom startup logic from the root PersistentPreRunE (added via
// persistentEdits): one-time legacy migration, then active-profile resolution.
func OnStartup(cmd *cobra.Command) error {
	if err := maybeMigrateLegacy(cmd); err != nil {
		// Best-effort: never block the command on a migration hiccup.
		fmt.Fprintf(cmd.ErrOrStderr(), "growthbook: legacy config migration skipped: %v\n", err)
	}
	warnIfDeprecated(cmd)
	if err := applyProfile(cmd); err != nil {
		return err
	}
	MaybeCheckForUpdate(cmd)
	return nil
}

// deprecatedGroups maps a deprecated command group to the group that replaces it.
var deprecatedGroups = map[string]string{
	"features-v1":          "features",
	"feature-revisions-v1": "feature-revisions",
}

// warnIfDeprecated prints a one-line notice to stderr (never stdout, so JSON
// output stays clean) when the invoked command lives under a deprecated group.
func warnIfDeprecated(cmd *cobra.Command) {
	for c := cmd; c != nil; c = c.Parent() {
		if replacement, ok := deprecatedGroups[c.Name()]; ok {
			fmt.Fprintf(cmd.ErrOrStderr(),
				"growthbook: `%s` targets the legacy v1 API and is deprecated; prefer `%s`. See MIGRATION.md.\n",
				c.Name(), replacement)
			return
		}
	}
}

// applyProfile injects the active profile's server URL and secret, preserving
// the precedence flag > env > profile.
func applyProfile(cmd *cobra.Command) error {
	f, err := Load()
	if err != nil {
		return nil
	}
	name, explicit := activeProfileName(cmd, f)
	if name == "" {
		return nil
	}
	prof, ok := f.Profiles[name]
	if !ok {
		if explicit {
			return fmt.Errorf("profile %q not found (see `growthbook profiles list`)", name)
		}
		return nil
	}

	if prof.ServerURL != "" && !serverSelectedByFlag(cmd) {
		if fl := cmd.Flags().Lookup("server-url"); fl != nil {
			if err := cmd.Flags().Set("server-url", prof.ServerURL); err != nil {
				return err
			}
		}
	}

	if !flagChanged(cmd, "bearer-auth") && os.Getenv("GBCLI_BEARER_AUTH") == "" {
		if secret := GetProfileSecret(name); secret != "" {
			_ = os.Setenv("GBCLI_BEARER_AUTH", secret)
		}
	}
	return nil
}

func serverSelectedByFlag(cmd *cobra.Command) bool {
	for _, name := range []string{"server-url", "server", "domain"} {
		if flagChanged(cmd, name) {
			return true
		}
	}
	return false
}

func flagChanged(cmd *cobra.Command, name string) bool {
	fl := cmd.Flags().Lookup(name)
	return fl != nil && fl.Changed
}
