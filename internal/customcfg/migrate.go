package customcfg

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const (
	legacyDefaultProfile = "default"
	// cloudHost is the legacy default api_base_url; nothing to carry over for it.
	cloudHost = "https://api.growthbook.io"
)

// legacyConfigPath is the legacy TS CLI's config: ~/.growthbook/config.toml.
func legacyConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".growthbook", "config.toml")
}

// maybeMigrateLegacy imports the legacy CLI's profiles on first run. It runs at
// most once (guarded by the MigratedFromLegacy marker) and never overwrites a
// profile that already exists in the new config.
func maybeMigrateLegacy(cmd *cobra.Command) error {
	cur, err := Load()
	if err != nil {
		cur = &File{}
	}
	if cur.MigratedFromLegacy {
		return nil
	}

	profiles, err := parseLegacyTOML(legacyConfigPath())
	if os.IsNotExist(err) {
		return nil // no legacy config — nothing to do, don't even mark.
	}
	if err != nil {
		return err
	}

	cur.MigratedFromLegacy = true // mark regardless so we don't re-scan every run
	if cur.Profiles == nil {
		cur.Profiles = map[string]Profile{}
	}

	var imported []string
	for name, kv := range profiles {
		if _, exists := cur.Profiles[name]; exists {
			continue // never clobber an existing profile
		}
		prof := Profile{}
		if u := kv["api_base_url"]; u != "" && strings.TrimRight(u, "/") != cloudHost {
			prof.ServerURL = toServerURL(u)
		}
		cur.Profiles[name] = prof
		if secret := kv["growthbook_secret"]; secret != "" {
			_ = SetProfileSecret(name, secret) // best-effort; keychain may be unavailable
		}
		imported = append(imported, name)
	}

	// Default the active profile to the legacy default (or the only one).
	if cur.Current == "" {
		if _, ok := cur.Profiles[legacyDefaultProfile]; ok {
			cur.Current = legacyDefaultProfile
		} else if len(imported) == 1 {
			cur.Current = imported[0]
		}
	}

	if err := cur.Save(); err != nil {
		return err
	}

	if len(imported) > 0 {
		w := cmd.ErrOrStderr()
		fmt.Fprintf(w, "growthbook: migrated %d profile(s) from ~/.growthbook/config.toml: %s\n",
			len(imported), strings.Join(imported, ", "))
		if cur.Current != "" {
			fmt.Fprintf(w, "  active profile: %s (switch with `growthbook profiles use <name>`)\n", cur.Current)
		}
	}
	return nil
}

// toServerURL translates a legacy api_base_url (host root) to the new
// --server-url form (host + /api). The legacy CLI appended /api/v1 itself; the
// generated SDK appends /v1 or /v2, so the persisted base must end in /api.
func toServerURL(legacy string) string {
	u := strings.TrimRight(legacy, "/")
	if !strings.HasSuffix(u, "/api") {
		u += "/api"
	}
	return u
}

// parseLegacyTOML reads the legacy config.toml into profile -> key -> value.
// The legacy file is a flat set of `[profile]` sections with `key = "value"`
// string entries, so a minimal parser avoids pulling in a TOML dependency.
func parseLegacyTOML(path string) (map[string]map[string]string, error) {
	if path == "" {
		return nil, os.ErrNotExist
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	profiles := map[string]map[string]string{}
	current := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			current = strings.TrimSpace(line[1 : len(line)-1])
			profiles[current] = map[string]string{}
			continue
		}
		eq := strings.IndexByte(line, '=')
		if eq < 0 || current == "" {
			continue
		}
		key := strings.TrimSpace(line[:eq])
		val := strings.Trim(strings.TrimSpace(line[eq+1:]), `"'`)
		profiles[current][key] = val
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return profiles, nil
}
