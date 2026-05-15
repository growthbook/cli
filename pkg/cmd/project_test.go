// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
	"github.com/stainless-sdks/growthbook-cli/internal/requestflag"
)

func TestProjectsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"projects", "create",
			"--name", "name",
			"--description", "description",
			"--public-id", "publicId",
			"--settings", "{confidenceLevel: 0, pValueThreshold: 0, statsEngine: statsEngine}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(projectsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"projects", "create",
			"--name", "name",
			"--description", "description",
			"--public-id", "publicId",
			"--settings.confidence-level", "0",
			"--settings.p-value-threshold", "0",
			"--settings.stats-engine", "statsEngine",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"description: description\n" +
			"publicId: publicId\n" +
			"settings:\n" +
			"  confidenceLevel: 0\n" +
			"  pValueThreshold: 0\n" +
			"  statsEngine: statsEngine\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"projects", "create",
		)
	})
}

func TestProjectsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"projects", "retrieve",
			"--id", "id",
		)
	})
}

func TestProjectsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"projects", "update",
			"--id", "id",
			"--description", "description",
			"--name", "name",
			"--public-id", "publicId",
			"--settings", "{confidenceLevel: 0, pValueThreshold: 0, statsEngine: statsEngine}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(projectsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"projects", "update",
			"--id", "id",
			"--description", "description",
			"--name", "name",
			"--public-id", "publicId",
			"--settings.confidence-level", "0",
			"--settings.p-value-threshold", "0",
			"--settings.stats-engine", "statsEngine",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"name: name\n" +
			"publicId: publicId\n" +
			"settings:\n" +
			"  confidenceLevel: 0\n" +
			"  pValueThreshold: 0\n" +
			"  statsEngine: statsEngine\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"projects", "update",
			"--id", "id",
		)
	})
}

func TestProjectsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"projects", "list",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestProjectsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"projects", "delete",
			"--id", "id",
		)
	})
}
