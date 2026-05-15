// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/growthbook/cli/internal/mocktest"
)

func TestCustomFieldsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"custom-fields", "create",
			"--id", "x",
			"--name", "name",
			"--required=true",
			"--section", "feature",
			"--type", "text",
			"--default-value", "string",
			"--description", "description",
			"--placeholder", "placeholder",
			"--project", "string",
			"--values", "values",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: x\n" +
			"name: name\n" +
			"required: true\n" +
			"sections:\n" +
			"  - feature\n" +
			"type: text\n" +
			"defaultValue: string\n" +
			"description: description\n" +
			"placeholder: placeholder\n" +
			"projects:\n" +
			"  - string\n" +
			"values: values\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"custom-fields", "create",
		)
	})
}

func TestCustomFieldsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"custom-fields", "retrieve",
			"--id", "id",
		)
	})
}

func TestCustomFieldsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"custom-fields", "update",
			"--id", "id",
			"--active=true",
			"--default-value", "string",
			"--description", "description",
			"--name", "name",
			"--placeholder", "placeholder",
			"--project", "string",
			"--required=true",
			"--section", "feature",
			"--values", "values",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"active: true\n" +
			"defaultValue: string\n" +
			"description: description\n" +
			"name: name\n" +
			"placeholder: placeholder\n" +
			"projects:\n" +
			"  - string\n" +
			"required: true\n" +
			"sections:\n" +
			"  - feature\n" +
			"values: values\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"custom-fields", "update",
			"--id", "id",
		)
	})
}

func TestCustomFieldsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"custom-fields", "list",
			"--project-id", "projectId",
		)
	})
}

func TestCustomFieldsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"custom-fields", "delete",
			"--id", "id",
			"--index", "index",
		)
	})
}
