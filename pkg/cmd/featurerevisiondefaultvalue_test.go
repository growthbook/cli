// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
)

func TestFeaturesRevisionsDefaultValueUpdateDefaultValue(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"features:revisions:default-value", "update-default-value",
			"--id", "id",
			"--version", "new",
			"--default-value", "defaultValue",
			"--revision-comment", "revisionComment",
			"--revision-title", "revisionTitle",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"defaultValue: defaultValue\n" +
			"revisionComment: revisionComment\n" +
			"revisionTitle: revisionTitle\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"features:revisions:default-value", "update-default-value",
			"--id", "id",
			"--version", "new",
		)
	})
}
