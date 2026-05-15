// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
)

func TestEnvironmentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"environments", "create",
			"--id", "id",
			"--default-state=true",
			"--description", "description",
			"--parent", "parent",
			"--project", "string",
			"--toggle-on-list=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: id\n" +
			"defaultState: true\n" +
			"description: description\n" +
			"parent: parent\n" +
			"projects:\n" +
			"  - string\n" +
			"toggleOnList: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"environments", "create",
		)
	})
}

func TestEnvironmentsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"environments", "update",
			"--id", "id",
			"--default-state=true",
			"--description", "description",
			"--project", "string",
			"--toggle-on-list=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"defaultState: true\n" +
			"description: description\n" +
			"projects:\n" +
			"  - string\n" +
			"toggleOnList: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"environments", "update",
			"--id", "id",
		)
	})
}

func TestEnvironmentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"environments", "list",
		)
	})
}

func TestEnvironmentsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"environments", "delete",
			"--id", "id",
		)
	})
}
