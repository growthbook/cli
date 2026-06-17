package customcfg

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/growthbook/cli/internal/config"
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

// mockKeyring is an in-memory keyring backend so tests never touch the real OS keychain.
type mockKeyring struct{ store map[string]string }

func (m *mockKeyring) Get(service, key string) (string, error) {
	if v, ok := m.store[service+"/"+key]; ok {
		return v, nil
	}
	return "", keyring.ErrNotFound
}
func (m *mockKeyring) Set(service, key, value string) error {
	m.store[service+"/"+key] = value
	return nil
}
func (m *mockKeyring) Delete(service, key string) error {
	delete(m.store, service+"/"+key)
	return nil
}

// setup isolates HOME to a temp dir and installs the mock keyring + fresh config.
func setup(t *testing.T) (home string, kr *mockKeyring) {
	t.Helper()
	home = t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("GBCLI_PROFILE", "")
	t.Setenv("GBCLI_BEARER_AUTH", "")
	kr = &mockKeyring{store: map[string]string{}}
	config.SetKeyringBackend(kr)
	config.Reset()
	if err := config.Init("growthbook", "GBCLI"); err != nil {
		t.Fatalf("config.Init: %v", err)
	}
	t.Cleanup(func() { config.ResetKeyring(); config.Reset() })
	return home, kr
}

func writeLegacy(t *testing.T, home, content string) {
	t.Helper()
	dir := filepath.Join(home, ".growthbook")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "config.toml"), []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
}

func TestMigrate_AllProfiles(t *testing.T) {
	home, kr := setup(t)
	writeLegacy(t, home, `
[default]
growthbook_secret = "secret_abc123"
api_base_url = "https://gb.internal.example.com"

[staging]
growthbook_secret = "secret_staging"
`)
	cmd := &cobra.Command{}
	if err := maybeMigrateLegacy(cmd); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	if got := kr.store["growthbook/profile.default.bearer-auth"]; got != "secret_abc123" {
		t.Errorf("default secret = %q, want secret_abc123", got)
	}
	if got := kr.store["growthbook/profile.staging.bearer-auth"]; got != "secret_staging" {
		t.Errorf("staging secret = %q, want secret_staging", got)
	}
	f, _ := Load()
	if f.Profiles["default"].ServerURL != "https://gb.internal.example.com/api" {
		t.Errorf("default ServerURL = %q", f.Profiles["default"].ServerURL)
	}
	if f.Profiles["staging"].ServerURL != "" {
		t.Errorf("staging ServerURL should be empty, got %q", f.Profiles["staging"].ServerURL)
	}
	if f.Current != "default" {
		t.Errorf("Current = %q, want default", f.Current)
	}
	if !f.MigratedFromLegacy {
		t.Error("MigratedFromLegacy not set")
	}

	// Idempotent: a second run must not re-import or clobber.
	kr.store["growthbook/profile.default.bearer-auth"] = "user-changed"
	if err := maybeMigrateLegacy(cmd); err != nil {
		t.Fatalf("second migrate: %v", err)
	}
	if kr.store["growthbook/profile.default.bearer-auth"] != "user-changed" {
		t.Error("second run clobbered the credential")
	}
}

func TestMigrate_CloudDefaultURLNotPersisted(t *testing.T) {
	home, _ := setup(t)
	writeLegacy(t, home, "[default]\ngrowthbook_secret = \"secret_x\"\napi_base_url = \"https://api.growthbook.io\"\n")
	if err := maybeMigrateLegacy(&cobra.Command{}); err != nil {
		t.Fatal(err)
	}
	f, _ := Load()
	if _, ok := f.Profiles["default"]; !ok {
		t.Fatal("default profile not created")
	}
	if f.Profiles["default"].ServerURL != "" {
		t.Errorf("cloud default should not persist a server URL, got %q", f.Profiles["default"].ServerURL)
	}
}

func TestMigrate_NoLegacyFile(t *testing.T) {
	setup(t)
	if err := maybeMigrateLegacy(&cobra.Command{}); err != nil {
		t.Fatalf("migrate with no legacy file: %v", err)
	}
	if f, _ := Load(); f.MigratedFromLegacy {
		t.Error("marked migrated despite no legacy file")
	}
}

func TestApplyProfile_Injection(t *testing.T) {
	setup(t)
	f := &File{
		Current:  "staging",
		Profiles: map[string]Profile{"staging": {ServerURL: "https://s.example.com/api"}},
	}
	if err := f.Save(); err != nil {
		t.Fatal(err)
	}
	if err := SetProfileSecret("staging", "secret_staging"); err != nil {
		t.Fatal(err)
	}

	cmd := testCmd()
	if err := applyProfile(cmd); err != nil {
		t.Fatalf("applyProfile: %v", err)
	}
	if got, _ := cmd.Flags().GetString("server-url"); got != "https://s.example.com/api" {
		t.Errorf("server-url = %q, want injected profile URL", got)
	}
	if got := os.Getenv("GBCLI_BEARER_AUTH"); got != "secret_staging" {
		t.Errorf("GBCLI_BEARER_AUTH = %q, want injected secret", got)
	}
}

func TestApplyProfile_ExplicitFlagWins(t *testing.T) {
	setup(t)
	f := &File{Current: "staging", Profiles: map[string]Profile{"staging": {ServerURL: "https://s.example.com/api"}}}
	if err := f.Save(); err != nil {
		t.Fatal(err)
	}
	cmd := testCmd()
	if err := cmd.Flags().Set("server-url", "https://override.example.com/api"); err != nil {
		t.Fatal(err)
	}
	if err := applyProfile(cmd); err != nil {
		t.Fatalf("applyProfile: %v", err)
	}
	if got, _ := cmd.Flags().GetString("server-url"); got != "https://override.example.com/api" {
		t.Errorf("explicit --server-url overridden: got %q", got)
	}
}

func TestApplyProfile_UnknownExplicitProfileErrors(t *testing.T) {
	setup(t)
	(&File{Profiles: map[string]Profile{}}).Save()
	t.Setenv("GBCLI_PROFILE", "ghost")
	if err := applyProfile(testCmd()); err == nil {
		t.Error("expected error for unknown explicitly-requested profile")
	}
}

func TestToServerURL(t *testing.T) {
	cases := map[string]string{
		"https://h.com":      "https://h.com/api",
		"https://h.com/":     "https://h.com/api",
		"https://h.com/api":  "https://h.com/api",
		"https://h.com/api/": "https://h.com/api",
	}
	for in, want := range cases {
		if got := toServerURL(in); got != want {
			t.Errorf("toServerURL(%q) = %q, want %q", in, got, want)
		}
	}
}

// testCmd builds a command exposing the flags applyProfile inspects.
func testCmd() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Flags().String("server-url", "", "")
	cmd.Flags().String("server", "", "")
	cmd.Flags().String("domain", "", "")
	cmd.Flags().String("bearer-auth", "", "")
	cmd.Flags().String("profile", "", "")
	return cmd
}
