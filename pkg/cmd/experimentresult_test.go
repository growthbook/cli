// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
)

func TestExperimentsResultsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments:results", "retrieve",
			"--id", "id",
			"--dimension", "dimension",
			"--phase", "phase",
		)
	})
}

func TestExperimentsResultsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments:results", "list",
			"--datasource-id", "datasourceId",
			"--limit", "1",
			"--offset", "0",
			"--project-id", "projectId",
			"--status", "draft",
			"--tracking-key", "trackingKey",
		)
	})
}
