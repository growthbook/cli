// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
)

func TestSavedGroupsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"saved-groups", "create",
			"--name", "name",
			"--attribute-key", "attributeKey",
			"--condition", "condition",
			"--owner", "owner",
			"--project", "string",
			"--type", "condition",
			"--value", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"attributeKey: attributeKey\n" +
			"condition: condition\n" +
			"owner: owner\n" +
			"projects:\n" +
			"  - string\n" +
			"type: condition\n" +
			"values:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"saved-groups", "create",
		)
	})
}

func TestSavedGroupsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"saved-groups", "retrieve",
			"--id", "id",
		)
	})
}

func TestSavedGroupsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"saved-groups", "update",
			"--id", "id",
			"--condition", "condition",
			"--name", "name",
			"--owner", "owner",
			"--project", "string",
			"--value", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"condition: condition\n" +
			"name: name\n" +
			"owner: owner\n" +
			"projects:\n" +
			"  - string\n" +
			"values:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"saved-groups", "update",
			"--id", "id",
		)
	})
}

func TestSavedGroupsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"saved-groups", "list",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestSavedGroupsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"saved-groups", "delete",
			"--id", "id",
		)
	})
}

func TestSavedGroupsArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"saved-groups", "archive",
			"--id", "id",
		)
	})
}

func TestSavedGroupsUnarchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"saved-groups", "unarchive",
			"--id", "id",
		)
	})
}
