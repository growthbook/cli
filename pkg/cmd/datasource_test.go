// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
)

func TestDataSourcesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"data-sources", "retrieve",
			"--id", "id",
		)
	})
}

func TestDataSourcesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"data-sources", "list",
			"--limit", "1",
			"--offset", "0",
			"--project-id", "projectId",
		)
	})
}

func TestDataSourcesGetInformationSchema(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"data-sources", "get-information-schema",
			"--data-source-id", "dataSourceId",
		)
	})
}
