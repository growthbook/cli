// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
)

func TestSegmentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"segments", "create",
			"--datasource-id", "datasourceId",
			"--identifier-type", "identifierType",
			"--name", "name",
			"--type", "SQL",
			"--description", "description",
			"--fact-table-id", "factTableId",
			"--filter", "string",
			"--managed-by", "",
			"--owner", "owner",
			"--project", "string",
			"--query", "query",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"datasourceId: datasourceId\n" +
			"identifierType: identifierType\n" +
			"name: name\n" +
			"type: SQL\n" +
			"description: description\n" +
			"factTableId: factTableId\n" +
			"filters:\n" +
			"  - string\n" +
			"managedBy: ''\n" +
			"owner: owner\n" +
			"projects:\n" +
			"  - string\n" +
			"query: query\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"segments", "create",
		)
	})
}

func TestSegmentsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"segments", "retrieve",
			"--id", "id",
		)
	})
}

func TestSegmentsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"segments", "update",
			"--id", "id",
			"--datasource-id", "datasourceId",
			"--description", "description",
			"--fact-table-id", "factTableId",
			"--filter", "string",
			"--identifier-type", "identifierType",
			"--managed-by", "",
			"--name", "name",
			"--owner", "owner",
			"--project", "string",
			"--query", "query",
			"--type", "SQL",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"datasourceId: datasourceId\n" +
			"description: description\n" +
			"factTableId: factTableId\n" +
			"filters:\n" +
			"  - string\n" +
			"identifierType: identifierType\n" +
			"managedBy: ''\n" +
			"name: name\n" +
			"owner: owner\n" +
			"projects:\n" +
			"  - string\n" +
			"query: query\n" +
			"type: SQL\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"segments", "update",
			"--id", "id",
		)
	})
}

func TestSegmentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"segments", "list",
			"--datasource-id", "datasourceId",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestSegmentsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"segments", "delete",
			"--id", "id",
		)
	})
}
