// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
)

func TestDimensionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"dimensions", "create",
			"--datasource-id", "datasourceId",
			"--identifier-type", "identifierType",
			"--name", "name",
			"--query", "query",
			"--description", "description",
			"--managed-by", "",
			"--owner", "owner",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"datasourceId: datasourceId\n" +
			"identifierType: identifierType\n" +
			"name: name\n" +
			"query: query\n" +
			"description: description\n" +
			"managedBy: ''\n" +
			"owner: owner\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"dimensions", "create",
		)
	})
}

func TestDimensionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"dimensions", "retrieve",
			"--id", "id",
		)
	})
}

func TestDimensionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"dimensions", "update",
			"--id", "id",
			"--datasource-id", "datasourceId",
			"--description", "description",
			"--identifier-type", "identifierType",
			"--managed-by", "",
			"--name", "name",
			"--owner", "owner",
			"--query", "query",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"datasourceId: datasourceId\n" +
			"description: description\n" +
			"identifierType: identifierType\n" +
			"managedBy: ''\n" +
			"name: name\n" +
			"owner: owner\n" +
			"query: query\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"dimensions", "update",
			"--id", "id",
		)
	})
}

func TestDimensionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"dimensions", "list",
			"--datasource-id", "datasourceId",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestDimensionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"dimensions", "delete",
			"--id", "id",
		)
	})
}
