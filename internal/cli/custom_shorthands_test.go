package cli

import "testing"

// Fails if a regen renames or drops a flag whose shorthand is pinned.
func TestPinnedShorthands(t *testing.T) {
	root, err := NewRootCommand()
	if err != nil {
		t.Fatal(err)
	}
	for group, leaves := range pinnedShorthands {
		for leaf, pins := range leaves {
			cmd, _, err := root.Find([]string{group, leaf})
			if err != nil || cmd.Name() != leaf {
				t.Errorf("`%s %s` not found", group, leaf)
				continue
			}
			for name, want := range pins {
				f := cmd.Flags().Lookup(name)
				if f == nil || f.Shorthand != want {
					t.Errorf("`%s %s --%s`: want shorthand %q, got %v", group, leaf, name, want, f)
				}
				if want != "" && cmd.Flags().ShorthandLookup(want) != f {
					t.Errorf("`%s %s -%s` does not resolve to --%s", group, leaf, want, name)
				}
			}
		}
	}
}
