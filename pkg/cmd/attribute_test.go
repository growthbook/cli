// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
)

func TestAttributesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"attributes", "create",
			"--datatype", "boolean",
			"--property", "property",
			"--archived=true",
			"--description", "description",
			"--enum", "enum",
			"--format", "",
			"--hash-attribute=true",
			"--project", "string",
			"--tag", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"datatype: boolean\n" +
			"property: property\n" +
			"archived: true\n" +
			"description: description\n" +
			"enum: enum\n" +
			"format: ''\n" +
			"hashAttribute: true\n" +
			"projects:\n" +
			"  - string\n" +
			"tags:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"attributes", "create",
		)
	})
}

func TestAttributesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"attributes", "update",
			"--property", "property",
			"--archived=true",
			"--datatype", "boolean",
			"--description", "description",
			"--enum", "enum",
			"--format", "",
			"--hash-attribute=true",
			"--project", "string",
			"--tag", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"archived: true\n" +
			"datatype: boolean\n" +
			"description: description\n" +
			"enum: enum\n" +
			"format: ''\n" +
			"hashAttribute: true\n" +
			"projects:\n" +
			"  - string\n" +
			"tags:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"attributes", "update",
			"--property", "property",
		)
	})
}

func TestAttributesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"attributes", "list",
		)
	})
}

func TestAttributesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"attributes", "delete",
			"--property", "property",
		)
	})
}
