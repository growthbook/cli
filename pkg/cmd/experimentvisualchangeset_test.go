// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/growthbook/cli/internal/mocktest"
	"github.com/growthbook/cli/internal/requestflag"
)

func TestExperimentsVisualChangesetsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments:visual-changesets", "create",
			"--id", "id",
			"--editor-url", "editorUrl",
			"--url-pattern", "{pattern: pattern, type: simple, include: true}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(experimentsVisualChangesetsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments:visual-changesets", "create",
			"--id", "id",
			"--editor-url", "editorUrl",
			"--url-pattern.pattern", "pattern",
			"--url-pattern.type", "simple",
			"--url-pattern.include=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"editorUrl: editorUrl\n" +
			"urlPatterns:\n" +
			"  - pattern: pattern\n" +
			"    type: simple\n" +
			"    include: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiments:visual-changesets", "create",
			"--id", "id",
		)
	})
}

func TestExperimentsVisualChangesetsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments:visual-changesets", "list",
			"--id", "id",
		)
	})
}
