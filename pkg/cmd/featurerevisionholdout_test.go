// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
	"github.com/stainless-sdks/growthbook-cli/internal/requestflag"
)

func TestFeaturesRevisionsHoldoutCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"features:revisions:holdout", "create",
			"--id", "id",
			"--version", "new",
			"--holdout", "{id: id, value: value}",
			"--revision-comment", "revisionComment",
			"--revision-title", "revisionTitle",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(featuresRevisionsHoldoutCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"features:revisions:holdout", "create",
			"--id", "id",
			"--version", "new",
			"--holdout.id", "id",
			"--holdout.value", "value",
			"--revision-comment", "revisionComment",
			"--revision-title", "revisionTitle",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"holdout:\n" +
			"  id: id\n" +
			"  value: value\n" +
			"revisionComment: revisionComment\n" +
			"revisionTitle: revisionTitle\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"features:revisions:holdout", "create",
			"--id", "id",
			"--version", "new",
		)
	})
}
