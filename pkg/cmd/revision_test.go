// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
)

func TestRevisionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"revisions", "list",
			"--author", "author",
			"--feature-id", "featureId",
			"--limit", "1",
			"--mine", "'true'",
			"--offset", "0",
			"--skip-pagination", "'true'",
			"--status", "draft",
		)
	})
}
