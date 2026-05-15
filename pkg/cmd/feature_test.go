// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
	"github.com/stainless-sdks/growthbook-cli/internal/requestflag"
)

func TestFeaturesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"features", "create",
			"--id", "x",
			"--default-value", "defaultValue",
			"--owner", "owner",
			"--value-type", "boolean",
			"--archived=true",
			"--custom-fields", "{foo: string}",
			"--description", "description",
			"--environments", "{foo: {enabled: true}}",
			"--json-schema", "jsonSchema",
			"--object-schema", "{fields: [{default: default, description: description, enum: [string], key: key, max: 0, min: 0, required: true, type: integer}], type: object}",
			"--prerequisite", "string",
			"--project", "project",
			"--rule", "{type: force, value: value, id: id, allEnvironments: true, condition: condition, description: description, enabled: true, environments: [string], prerequisites: [{id: id, condition: condition}], savedGroupTargeting: [{matchType: all, savedGroups: [string]}], scheduleRules: [{enabled: true, timestamp: timestamp}]}",
			"--tag", "string",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(featuresCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"features", "create",
			"--id", "x",
			"--default-value", "defaultValue",
			"--owner", "owner",
			"--value-type", "boolean",
			"--archived=true",
			"--custom-fields", "{foo: string}",
			"--description", "description",
			"--environments", "{foo: {enabled: true}}",
			"--json-schema", "jsonSchema",
			"--object-schema.fields", "[{default: default, description: description, enum: [string], key: key, max: 0, min: 0, required: true, type: integer}]",
			"--object-schema.type", "object",
			"--prerequisite", "string",
			"--project", "project",
			"--rule", "{type: force, value: value, id: id, allEnvironments: true, condition: condition, description: description, enabled: true, environments: [string], prerequisites: [{id: id, condition: condition}], savedGroupTargeting: [{matchType: all, savedGroups: [string]}], scheduleRules: [{enabled: true, timestamp: timestamp}]}",
			"--tag", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: x\n" +
			"defaultValue: defaultValue\n" +
			"owner: owner\n" +
			"valueType: boolean\n" +
			"archived: true\n" +
			"customFields:\n" +
			"  foo: string\n" +
			"description: description\n" +
			"environments:\n" +
			"  foo:\n" +
			"    enabled: true\n" +
			"jsonSchema: jsonSchema\n" +
			"objectSchema:\n" +
			"  fields:\n" +
			"    - default: default\n" +
			"      description: description\n" +
			"      enum:\n" +
			"        - string\n" +
			"      key: key\n" +
			"      max: 0\n" +
			"      min: 0\n" +
			"      required: true\n" +
			"      type: integer\n" +
			"  type: object\n" +
			"prerequisites:\n" +
			"  - string\n" +
			"project: project\n" +
			"rules:\n" +
			"  - type: force\n" +
			"    value: value\n" +
			"    id: id\n" +
			"    allEnvironments: true\n" +
			"    condition: condition\n" +
			"    description: description\n" +
			"    enabled: true\n" +
			"    environments:\n" +
			"      - string\n" +
			"    prerequisites:\n" +
			"      - id: id\n" +
			"        condition: condition\n" +
			"    savedGroupTargeting:\n" +
			"      - matchType: all\n" +
			"        savedGroups:\n" +
			"          - string\n" +
			"    scheduleRules:\n" +
			"      - enabled: true\n" +
			"        timestamp: timestamp\n" +
			"tags:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"features", "create",
		)
	})
}

func TestFeaturesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"features", "retrieve",
			"--id", "id",
			"--with-revisions", "all",
		)
	})
}

func TestFeaturesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"features", "update",
			"--id", "id",
			"--archived=true",
			"--custom-fields", "{foo: string}",
			"--default-value", "defaultValue",
			"--description", "description",
			"--environments", "{foo: {enabled: true}}",
			"--holdout", "{id: id, value: value}",
			"--json-schema", "jsonSchema",
			"--object-schema", "{fields: [{default: default, description: description, enum: [string], key: key, max: 0, min: 0, required: true, type: integer}], type: object}",
			"--owner", "owner",
			"--prerequisite", "string",
			"--project", "project",
			"--rule", "{type: force, value: value, id: id, allEnvironments: true, condition: condition, description: description, enabled: true, environments: [string], prerequisites: [{id: id, condition: condition}], savedGroupTargeting: [{matchType: all, savedGroups: [string]}], scheduleRules: [{enabled: true, timestamp: timestamp}]}",
			"--tag", "string",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(featuresUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"features", "update",
			"--id", "id",
			"--archived=true",
			"--custom-fields", "{foo: string}",
			"--default-value", "defaultValue",
			"--description", "description",
			"--environments", "{foo: {enabled: true}}",
			"--holdout.id", "id",
			"--holdout.value", "value",
			"--json-schema", "jsonSchema",
			"--object-schema.fields", "[{default: default, description: description, enum: [string], key: key, max: 0, min: 0, required: true, type: integer}]",
			"--object-schema.type", "object",
			"--owner", "owner",
			"--prerequisite", "string",
			"--project", "project",
			"--rule", "{type: force, value: value, id: id, allEnvironments: true, condition: condition, description: description, enabled: true, environments: [string], prerequisites: [{id: id, condition: condition}], savedGroupTargeting: [{matchType: all, savedGroups: [string]}], scheduleRules: [{enabled: true, timestamp: timestamp}]}",
			"--tag", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"archived: true\n" +
			"customFields:\n" +
			"  foo: string\n" +
			"defaultValue: defaultValue\n" +
			"description: description\n" +
			"environments:\n" +
			"  foo:\n" +
			"    enabled: true\n" +
			"holdout:\n" +
			"  id: id\n" +
			"  value: value\n" +
			"jsonSchema: jsonSchema\n" +
			"objectSchema:\n" +
			"  fields:\n" +
			"    - default: default\n" +
			"      description: description\n" +
			"      enum:\n" +
			"        - string\n" +
			"      key: key\n" +
			"      max: 0\n" +
			"      min: 0\n" +
			"      required: true\n" +
			"      type: integer\n" +
			"  type: object\n" +
			"owner: owner\n" +
			"prerequisites:\n" +
			"  - string\n" +
			"project: project\n" +
			"rules:\n" +
			"  - type: force\n" +
			"    value: value\n" +
			"    id: id\n" +
			"    allEnvironments: true\n" +
			"    condition: condition\n" +
			"    description: description\n" +
			"    enabled: true\n" +
			"    environments:\n" +
			"      - string\n" +
			"    prerequisites:\n" +
			"      - id: id\n" +
			"        condition: condition\n" +
			"    savedGroupTargeting:\n" +
			"      - matchType: all\n" +
			"        savedGroups:\n" +
			"          - string\n" +
			"    scheduleRules:\n" +
			"      - enabled: true\n" +
			"        timestamp: timestamp\n" +
			"tags:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"features", "update",
			"--id", "id",
		)
	})
}

func TestFeaturesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"features", "list",
			"--client-key", "clientKey",
			"--limit", "1",
			"--offset", "0",
			"--project-id", "projectId",
			"--skip-pagination", "'true'",
		)
	})
}

func TestFeaturesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"features", "delete",
			"--id", "id",
		)
	})
}

func TestFeaturesRevert(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"features", "revert",
			"--id", "id",
			"--revision", "0",
			"--comment", "comment",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"revision: 0\n" +
			"comment: comment\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"features", "revert",
			"--id", "id",
		)
	})
}

func TestFeaturesToggle(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"features", "toggle",
			"--id", "id",
			"--environments", "{foo: 'true'}",
			"--reason", "reason",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"environments:\n" +
			"  foo: 'true'\n" +
			"reason: reason\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"features", "toggle",
			"--id", "id",
		)
	})
}
