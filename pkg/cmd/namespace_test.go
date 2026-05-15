// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/growthbook/cli/internal/mocktest"
)

func TestNamespacesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"namespaces", "create",
			"--display-name", "displayName",
			"--description", "description",
			"--format", "legacy",
			"--hash-attribute", "hashAttribute",
			"--status", "active",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"displayName: displayName\n" +
			"description: description\n" +
			"format: legacy\n" +
			"hashAttribute: hashAttribute\n" +
			"status: active\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"namespaces", "create",
		)
	})
}

func TestNamespacesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"namespaces", "retrieve",
			"--id", "id",
		)
	})
}

func TestNamespacesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"namespaces", "update",
			"--id", "id",
			"--description", "description",
			"--display-name", "displayName",
			"--hash-attribute", "hashAttribute",
			"--status", "active",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"displayName: displayName\n" +
			"hashAttribute: hashAttribute\n" +
			"status: active\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"namespaces", "update",
			"--id", "id",
		)
	})
}

func TestNamespacesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"namespaces", "list",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestNamespacesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"namespaces", "delete",
			"--id", "id",
		)
	})
}

func TestNamespacesGetMemberships(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"namespaces", "get-memberships",
			"--id", "id",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestNamespacesRotateSeed(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"namespaces", "rotate-seed",
			"--id", "id",
			"--seed", "seed",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("seed: seed")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"namespaces", "rotate-seed",
			"--id", "id",
		)
	})
}
