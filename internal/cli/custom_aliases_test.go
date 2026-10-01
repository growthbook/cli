package cli

import "testing"

// Fails if a regen renames a command that still carries a deprecated alias.
func TestRenamedAliasTargetsExist(t *testing.T) {
	root, err := NewRootCommand()
	if err != nil {
		t.Fatal(err)
	}
	for group, leaves := range renamedCommands {
		for name, old := range leaves {
			cmd, _, err := root.Find([]string{group, old})
			if err != nil || cmd.Name() != name || cmd.Parent().Name() != group {
				t.Errorf("`%s %s` does not resolve to `%s %s`", group, old, group, name)
			}
		}
	}
}
