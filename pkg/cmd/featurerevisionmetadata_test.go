// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
	"github.com/stainless-sdks/growthbook-cli/internal/requestflag"
)

func TestFeaturesRevisionsMetadataCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"features:revisions:metadata", "create",
			"--id", "id",
			"--version", "new",
			"--comment", "comment",
			"--custom-fields", "{foo: bar}",
			"--description", "description",
			"--json-schema", "{date: {}, enabled: true, schema: schema, schemaType: schema, simple: {fields: [{default: default, description: description, enum: [string], key: key, max: 0, min: 0, required: true, type: integer}], type: object}}",
			"--never-stale=true",
			"--owner", "owner",
			"--project", "project",
			"--tag", "string",
			"--title", "title",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(featuresRevisionsMetadataCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"features:revisions:metadata", "create",
			"--id", "id",
			"--version", "new",
			"--comment", "comment",
			"--custom-fields", "{foo: bar}",
			"--description", "description",
			"--json-schema.date", "{}",
			"--json-schema.enabled=true",
			"--json-schema.schema", "schema",
			"--json-schema.schema-type", "schema",
			"--json-schema.simple", "{fields: [{default: default, description: description, enum: [string], key: key, max: 0, min: 0, required: true, type: integer}], type: object}",
			"--never-stale=true",
			"--owner", "owner",
			"--project", "project",
			"--tag", "string",
			"--title", "title",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"comment: comment\n" +
			"customFields:\n" +
			"  foo: bar\n" +
			"description: description\n" +
			"jsonSchema:\n" +
			"  date: {}\n" +
			"  enabled: true\n" +
			"  schema: schema\n" +
			"  schemaType: schema\n" +
			"  simple:\n" +
			"    fields:\n" +
			"      - default: default\n" +
			"        description: description\n" +
			"        enum:\n" +
			"          - string\n" +
			"        key: key\n" +
			"        max: 0\n" +
			"        min: 0\n" +
			"        required: true\n" +
			"        type: integer\n" +
			"    type: object\n" +
			"neverStale: true\n" +
			"owner: owner\n" +
			"project: project\n" +
			"tags:\n" +
			"  - string\n" +
			"title: title\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"features:revisions:metadata", "create",
			"--id", "id",
			"--version", "new",
		)
	})
}
