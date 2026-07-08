// Hand-written (NOT generated). A golden-file snapshot of the CLI's command
// surface. Because the command tree is generated from the OpenAPI spec, a spec
// or overlay change can silently add, drop, or rename a command or flag — the
// kind of regression a reviewer would otherwise have to eyeball. This test
// turns that into a reviewable diff.
//
// The golden file is a generated artifact: `go generate ./internal/cli/...`
// refreshes it (see the go:generate directive below), and CI fails if the
// committed copy is stale — so a spec-driven surface change lands as a
// reviewable diff instead of drifting silently.
package cli

//go:generate go test -run ^TestCommandSurface$ -update

import (
	"flag"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var update = flag.Bool("update", false, "update golden files")

// TestCommandSurface snapshots every command path and its local flags.
func TestCommandSurface(t *testing.T) {
	root, err := NewRootCommand()
	if err != nil {
		t.Fatalf("NewRootCommand: %v", err)
	}

	var lines []string
	var walk func(cmd *cobra.Command)
	walk = func(cmd *cobra.Command) {
		// Skip cobra's built-in help/completion scaffolding — it's stable and noisy.
		if cmd.Name() == "help" || cmd.Name() == "completion" {
			return
		}
		var flags []string
		cmd.LocalFlags().VisitAll(func(f *pflag.Flag) {
			flags = append(flags, "--"+f.Name)
		})
		sort.Strings(flags)
		line := cmd.CommandPath()
		if len(flags) > 0 {
			line += "  [" + strings.Join(flags, " ") + "]"
		}
		lines = append(lines, line)
		for _, sub := range cmd.Commands() {
			walk(sub)
		}
	}
	walk(root)
	sort.Strings(lines)
	got := strings.Join(lines, "\n") + "\n"

	golden := filepath.Join("testdata", "command-surface.txt")
	if *update {
		if err := os.MkdirAll("testdata", 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(golden, []byte(got), 0o644); err != nil {
			t.Fatal(err)
		}
		t.Logf("wrote %s", golden)
		return
	}

	want, err := os.ReadFile(golden)
	if err != nil {
		t.Fatalf("read golden (run with -update to create it): %v", err)
	}
	if got != string(want) {
		t.Errorf("command surface changed. If this is intentional, run:\n"+
			"  go generate ./internal/cli/...\n"+
			"and commit internal/cli/testdata/command-surface.txt.\n\n"+
			"diff (-want +got):\n%s", diff(string(want), got))
	}
}

// diff produces a minimal line-by-line diff for the failure message.
func diff(want, got string) string {
	wl, gl := strings.Split(want, "\n"), strings.Split(got, "\n")
	seen := map[string]int{}
	for _, l := range wl {
		seen[l]++
	}
	for _, l := range gl {
		seen[l]--
	}
	var b strings.Builder
	for _, l := range wl {
		if seen[l] > 0 {
			b.WriteString("- " + l + "\n")
			seen[l]--
		}
	}
	seen = map[string]int{}
	for _, l := range wl {
		seen[l]++
	}
	for _, l := range gl {
		if seen[l] <= 0 {
			b.WriteString("+ " + l + "\n")
		} else {
			seen[l]--
		}
	}
	return b.String()
}
