// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
)

func TestFactTablesFiltersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"fact-tables:filters", "create",
			"--fact-table-id", "factTableId",
			"--name", "name",
			"--value", "country = 'US'",
			"--description", "description",
			"--managed-by", "",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"value: country = 'US'\n" +
			"description: description\n" +
			"managedBy: ''\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"fact-tables:filters", "create",
			"--fact-table-id", "factTableId",
		)
	})
}

func TestFactTablesFiltersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"fact-tables:filters", "retrieve",
			"--fact-table-id", "factTableId",
			"--id", "id",
		)
	})
}

func TestFactTablesFiltersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"fact-tables:filters", "update",
			"--fact-table-id", "factTableId",
			"--id", "id",
			"--description", "description",
			"--managed-by", "",
			"--name", "name",
			"--value", "country = 'US'",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"managedBy: ''\n" +
			"name: name\n" +
			"value: country = 'US'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"fact-tables:filters", "update",
			"--fact-table-id", "factTableId",
			"--id", "id",
		)
	})
}

func TestFactTablesFiltersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"fact-tables:filters", "list",
			"--fact-table-id", "factTableId",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestFactTablesFiltersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"fact-tables:filters", "delete",
			"--fact-table-id", "factTableId",
			"--id", "id",
		)
	})
}
