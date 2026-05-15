// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/growthbook/cli/internal/mocktest"
	"github.com/growthbook/cli/internal/requestflag"
)

func TestVisualChangesetsVisualChangeCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"visual-changesets:visual-change", "create",
			"--id", "id",
			"--variation", "variation",
			"--id", "id",
			"--css", "css",
			"--description", "description",
			"--dom-mutation", "{action: append, attribute: attribute, selector: selector, insertBeforeSelector: insertBeforeSelector, parentSelector: parentSelector, value: value}",
			"--js", "js",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(visualChangesetsVisualChangeCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"visual-changesets:visual-change", "create",
			"--id", "id",
			"--variation", "variation",
			"--id", "id",
			"--css", "css",
			"--description", "description",
			"--dom-mutation.action", "append",
			"--dom-mutation.attribute", "attribute",
			"--dom-mutation.selector", "selector",
			"--dom-mutation.insert-before-selector", "insertBeforeSelector",
			"--dom-mutation.parent-selector", "parentSelector",
			"--dom-mutation.value", "value",
			"--js", "js",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"variation: variation\n" +
			"id: id\n" +
			"css: css\n" +
			"description: description\n" +
			"domMutations:\n" +
			"  - action: append\n" +
			"    attribute: attribute\n" +
			"    selector: selector\n" +
			"    insertBeforeSelector: insertBeforeSelector\n" +
			"    parentSelector: parentSelector\n" +
			"    value: value\n" +
			"js: js\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"visual-changesets:visual-change", "create",
			"--id", "id",
		)
	})
}

func TestVisualChangesetsVisualChangeUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"visual-changesets:visual-change", "update",
			"--id", "id",
			"--visual-change-id", "visualChangeId",
			"--id", "id",
			"--css", "css",
			"--description", "description",
			"--dom-mutation", "{action: append, attribute: attribute, selector: selector, insertBeforeSelector: insertBeforeSelector, parentSelector: parentSelector, value: value}",
			"--js", "js",
			"--variation", "variation",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(visualChangesetsVisualChangeUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"visual-changesets:visual-change", "update",
			"--id", "id",
			"--visual-change-id", "visualChangeId",
			"--id", "id",
			"--css", "css",
			"--description", "description",
			"--dom-mutation.action", "append",
			"--dom-mutation.attribute", "attribute",
			"--dom-mutation.selector", "selector",
			"--dom-mutation.insert-before-selector", "insertBeforeSelector",
			"--dom-mutation.parent-selector", "parentSelector",
			"--dom-mutation.value", "value",
			"--js", "js",
			"--variation", "variation",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: id\n" +
			"css: css\n" +
			"description: description\n" +
			"domMutations:\n" +
			"  - action: append\n" +
			"    attribute: attribute\n" +
			"    selector: selector\n" +
			"    insertBeforeSelector: insertBeforeSelector\n" +
			"    parentSelector: parentSelector\n" +
			"    value: value\n" +
			"js: js\n" +
			"variation: variation\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"visual-changesets:visual-change", "update",
			"--id", "id",
			"--visual-change-id", "visualChangeId",
		)
	})
}
