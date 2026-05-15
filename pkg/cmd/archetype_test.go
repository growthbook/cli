// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
)

func TestArchetypesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"archetypes", "create",
			"--is-public=true",
			"--name", "name",
			"--attributes", "{foo: bar}",
			"--description", "description",
			"--project", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"isPublic: true\n" +
			"name: name\n" +
			"attributes:\n" +
			"  foo: bar\n" +
			"description: description\n" +
			"projects:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"archetypes", "create",
		)
	})
}

func TestArchetypesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"archetypes", "retrieve",
			"--id", "id",
		)
	})
}

func TestArchetypesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"archetypes", "update",
			"--id", "id",
			"--attributes", "{foo: bar}",
			"--description", "description",
			"--is-public=true",
			"--name", "name",
			"--project", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"attributes:\n" +
			"  foo: bar\n" +
			"description: description\n" +
			"isPublic: true\n" +
			"name: name\n" +
			"projects:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"archetypes", "update",
			"--id", "id",
		)
	})
}

func TestArchetypesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"archetypes", "list",
		)
	})
}

func TestArchetypesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"archetypes", "delete",
			"--id", "id",
		)
	})
}
