// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/growthbook/cli/internal/mocktest"
)

func TestMetricGroupsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"metric-groups", "create",
			"--datasource", "datasource",
			"--description", "description",
			"--metric", "string",
			"--name", "name",
			"--project", "string",
			"--archived=true",
			"--owner", "owner",
			"--tag", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"datasource: datasource\n" +
			"description: description\n" +
			"metrics:\n" +
			"  - string\n" +
			"name: name\n" +
			"projects:\n" +
			"  - string\n" +
			"archived: true\n" +
			"owner: owner\n" +
			"tags:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"metric-groups", "create",
		)
	})
}

func TestMetricGroupsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"metric-groups", "retrieve",
			"--id", "id",
		)
	})
}

func TestMetricGroupsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"metric-groups", "update",
			"--id", "id",
			"--archived=true",
			"--datasource", "datasource",
			"--description", "description",
			"--metric", "string",
			"--name", "name",
			"--owner", "owner",
			"--project", "string",
			"--tag", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"archived: true\n" +
			"datasource: datasource\n" +
			"description: description\n" +
			"metrics:\n" +
			"  - string\n" +
			"name: name\n" +
			"owner: owner\n" +
			"projects:\n" +
			"  - string\n" +
			"tags:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"metric-groups", "update",
			"--id", "id",
		)
	})
}

func TestMetricGroupsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"metric-groups", "list",
		)
	})
}

func TestMetricGroupsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"metric-groups", "delete",
			"--id", "id",
		)
	})
}
