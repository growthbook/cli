// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/growthbook/cli/internal/mocktest"
)

func TestExperimentsVariationScreenshotDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments:variation:screenshot", "delete",
			"--id", "id",
			"--variation-id", "variationId",
			"--path", "path",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("path: path")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiments:variation:screenshot", "delete",
			"--id", "id",
			"--variation-id", "variationId",
		)
	})
}

func TestExperimentsVariationScreenshotUpload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments:variation:screenshot", "upload",
			"--id", "id",
			"--variation-id", "variationId",
			"--content-type", "image/png",
			"--screenshot", "screenshot",
			"--description", "description",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"contentType: image/png\n" +
			"screenshot: screenshot\n" +
			"description: description\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiments:variation:screenshot", "upload",
			"--id", "id",
			"--variation-id", "variationId",
		)
	})
}
