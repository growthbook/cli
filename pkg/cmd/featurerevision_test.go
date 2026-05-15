// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
	"github.com/stainless-sdks/growthbook-cli/internal/requestflag"
)

func TestFeaturesRevisionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "create",
			"--id", "id",
			"--comment", "comment",
			"--title", "title",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"comment: comment\n" +
			"title: title\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"features:revisions", "create",
			"--id", "id",
		)
	})
}

func TestFeaturesRevisionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "retrieve",
			"--id", "id",
			"--version", "0",
		)
	})
}

func TestFeaturesRevisionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "list",
			"--id", "id",
			"--author", "author",
			"--limit", "1",
			"--offset", "0",
			"--skip-pagination", "'true'",
			"--status", "draft",
		)
	})
}

func TestFeaturesRevisionsDiscard(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "discard",
			"--id", "id",
			"--version", "0",
			"--body", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("foo: bar")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"features:revisions", "discard",
			"--id", "id",
			"--version", "0",
		)
	})
}

func TestFeaturesRevisionsGetMergeStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "get-merge-status",
			"--id", "id",
			"--version", "0",
		)
	})
}

func TestFeaturesRevisionsPublish(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "publish",
			"--id", "id",
			"--version", "0",
			"--comment", "comment",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("comment: comment")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"features:revisions", "publish",
			"--id", "id",
			"--version", "0",
		)
	})
}

func TestFeaturesRevisionsRebase(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "rebase",
			"--id", "id",
			"--version", "0",
			"--conflict-resolutions", "{foo: overwrite}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"conflictResolutions:\n" +
			"  foo: overwrite\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"features:revisions", "rebase",
			"--id", "id",
			"--version", "0",
		)
	})
}

func TestFeaturesRevisionsRequestReview(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "request-review",
			"--id", "id",
			"--version", "0",
			"--comment", "comment",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("comment: comment")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"features:revisions", "request-review",
			"--id", "id",
			"--version", "0",
		)
	})
}

func TestFeaturesRevisionsRetrieveLatest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "retrieve-latest",
			"--id", "id",
			"--mine", "'true'",
		)
	})
}

func TestFeaturesRevisionsRevert(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "revert",
			"--id", "id",
			"--version", "0",
			"--comment", "comment",
			"--strategy", "draft",
			"--title", "title",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"comment: comment\n" +
			"strategy: draft\n" +
			"title: title\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"features:revisions", "revert",
			"--id", "id",
			"--version", "0",
		)
	})
}

func TestFeaturesRevisionsSubmitReview(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "submit-review",
			"--id", "id",
			"--version", "0",
			"--action", "approve",
			"--comment", "comment",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"action: approve\n" +
			"comment: comment\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"features:revisions", "submit-review",
			"--id", "id",
			"--version", "0",
		)
	})
}

func TestFeaturesRevisionsToggle(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "toggle",
			"--id", "id",
			"--version", "new",
			"--enabled=true",
			"--environment", "environment",
			"--revision-comment", "revisionComment",
			"--revision-title", "revisionTitle",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"enabled: true\n" +
			"environment: environment\n" +
			"revisionComment: revisionComment\n" +
			"revisionTitle: revisionTitle\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"features:revisions", "toggle",
			"--id", "id",
			"--version", "new",
		)
	})
}

func TestFeaturesRevisionsUpdateArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "update-archive",
			"--id", "id",
			"--version", "new",
			"--archived=true",
			"--revision-comment", "revisionComment",
			"--revision-title", "revisionTitle",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"archived: true\n" +
			"revisionComment: revisionComment\n" +
			"revisionTitle: revisionTitle\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"features:revisions", "update-archive",
			"--id", "id",
			"--version", "new",
		)
	})
}

func TestFeaturesRevisionsUpdateDefaultValue(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "update-default-value",
			"--id", "id",
			"--version", "new",
			"--default-value", "defaultValue",
			"--revision-comment", "revisionComment",
			"--revision-title", "revisionTitle",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"defaultValue: defaultValue\n" +
			"revisionComment: revisionComment\n" +
			"revisionTitle: revisionTitle\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"features:revisions", "update-default-value",
			"--id", "id",
			"--version", "new",
		)
	})
}

func TestFeaturesRevisionsUpdateHoldout(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "update-holdout",
			"--id", "id",
			"--version", "new",
			"--holdout", "{id: id, value: value}",
			"--revision-comment", "revisionComment",
			"--revision-title", "revisionTitle",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(featuresRevisionsUpdateHoldout)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "update-holdout",
			"--id", "id",
			"--version", "new",
			"--holdout.id", "id",
			"--holdout.value", "value",
			"--revision-comment", "revisionComment",
			"--revision-title", "revisionTitle",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"holdout:\n" +
			"  id: id\n" +
			"  value: value\n" +
			"revisionComment: revisionComment\n" +
			"revisionTitle: revisionTitle\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"features:revisions", "update-holdout",
			"--id", "id",
			"--version", "new",
		)
	})
}

func TestFeaturesRevisionsUpdateMetadata(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "update-metadata",
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
		requestflag.CheckInnerFlags(featuresRevisionsUpdateMetadata)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "update-metadata",
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
			"features:revisions", "update-metadata",
			"--id", "id",
			"--version", "new",
		)
	})
}

func TestFeaturesRevisionsUpdatePrerequisites(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "update-prerequisites",
			"--id", "id",
			"--version", "new",
			"--prerequisite", "{id: id, condition: condition}",
			"--revision-comment", "revisionComment",
			"--revision-title", "revisionTitle",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(featuresRevisionsUpdatePrerequisites)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features:revisions", "update-prerequisites",
			"--id", "id",
			"--version", "new",
			"--prerequisite.id", "id",
			"--prerequisite.condition", "condition",
			"--revision-comment", "revisionComment",
			"--revision-title", "revisionTitle",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"prerequisites:\n" +
			"  - id: id\n" +
			"    condition: condition\n" +
			"revisionComment: revisionComment\n" +
			"revisionTitle: revisionTitle\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"features:revisions", "update-prerequisites",
			"--id", "id",
			"--version", "new",
		)
	})
}
