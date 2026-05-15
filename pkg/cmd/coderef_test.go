// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
	"github.com/stainless-sdks/growthbook-cli/internal/requestflag"
)

func TestCodeRefsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"code-refs", "create",
			"--branch", "branch",
			"--ref", "{contentHash: contentHash, filePath: filePath, flagKey: flagKey, lines: lines, startingLineNumber: 0}",
			"--repo-name", "repoName",
			"--delete-missing", "true",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(codeRefsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"code-refs", "create",
			"--branch", "branch",
			"--ref.content-hash", "contentHash",
			"--ref.file-path", "filePath",
			"--ref.flag-key", "flagKey",
			"--ref.lines", "lines",
			"--ref.starting-line-number", "0",
			"--repo-name", "repoName",
			"--delete-missing", "true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"branch: branch\n" +
			"refs:\n" +
			"  - contentHash: contentHash\n" +
			"    filePath: filePath\n" +
			"    flagKey: flagKey\n" +
			"    lines: lines\n" +
			"    startingLineNumber: 0\n" +
			"repoName: repoName\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"code-refs", "create",
			"--delete-missing", "true",
		)
	})
}

func TestCodeRefsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"code-refs", "retrieve",
			"--id", "id",
		)
	})
}

func TestCodeRefsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"code-refs", "list",
			"--limit", "1",
			"--offset", "0",
		)
	})
}
