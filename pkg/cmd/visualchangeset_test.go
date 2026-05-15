// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
	"github.com/stainless-sdks/growthbook-cli/internal/requestflag"
)

func TestVisualChangesetsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"visual-changesets", "retrieve",
			"--id", "id",
			"--include-experiment", "0",
		)
	})
}

func TestVisualChangesetsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"visual-changesets", "update",
			"--id", "id",
			"--editor-url", "editorUrl",
			"--url-pattern", "{pattern: pattern, type: simple, include: true}",
			"--visual-change", "{variation: variation, id: id, css: css, description: description, domMutations: [{action: append, attribute: attribute, selector: selector, insertBeforeSelector: insertBeforeSelector, parentSelector: parentSelector, value: value}], js: js}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(visualChangesetsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"visual-changesets", "update",
			"--id", "id",
			"--editor-url", "editorUrl",
			"--url-pattern.pattern", "pattern",
			"--url-pattern.type", "simple",
			"--url-pattern.include=true",
			"--visual-change.variation", "variation",
			"--visual-change.id", "id",
			"--visual-change.css", "css",
			"--visual-change.description", "description",
			"--visual-change.dom-mutations", "[{action: append, attribute: attribute, selector: selector, insertBeforeSelector: insertBeforeSelector, parentSelector: parentSelector, value: value}]",
			"--visual-change.js", "js",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"editorUrl: editorUrl\n" +
			"urlPatterns:\n" +
			"  - pattern: pattern\n" +
			"    type: simple\n" +
			"    include: true\n" +
			"visualChanges:\n" +
			"  - variation: variation\n" +
			"    id: id\n" +
			"    css: css\n" +
			"    description: description\n" +
			"    domMutations:\n" +
			"      - action: append\n" +
			"        attribute: attribute\n" +
			"        selector: selector\n" +
			"        insertBeforeSelector: insertBeforeSelector\n" +
			"        parentSelector: parentSelector\n" +
			"        value: value\n" +
			"    js: js\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"visual-changesets", "update",
			"--id", "id",
		)
	})
}
