// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
	"github.com/stainless-sdks/growthbook-cli/internal/requestflag"
)

func TestFactTablesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"fact-tables", "create",
			"--datasource", "datasource",
			"--name", "name",
			"--sql", "sql",
			"--user-id-type", "string",
			"--description", "description",
			"--event-name", "eventName",
			"--managed-by", "",
			"--owner", "owner",
			"--project", "string",
			"--tag", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"datasource: datasource\n" +
			"name: name\n" +
			"sql: sql\n" +
			"userIdTypes:\n" +
			"  - string\n" +
			"description: description\n" +
			"eventName: eventName\n" +
			"managedBy: ''\n" +
			"owner: owner\n" +
			"projects:\n" +
			"  - string\n" +
			"tags:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"fact-tables", "create",
		)
	})
}

func TestFactTablesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"fact-tables", "retrieve",
			"--id", "id",
		)
	})
}

func TestFactTablesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"fact-tables", "update",
			"--id", "id",
			"--archived=true",
			"--column", "{column: column, datatype: number, alwaysInlineFilter: true, autoSlices: [string], deleted: true, description: description, isAutoSliceColumn: true, jsonFields: {foo: {datatype: number}}, lockedAutoSlices: [string], name: name, numberFormat: ''}",
			"--columns-error", "columnsError",
			"--description", "description",
			"--event-name", "eventName",
			"--managed-by", "",
			"--name", "name",
			"--owner", "owner",
			"--project", "string",
			"--sql", "sql",
			"--tag", "string",
			"--user-id-type", "string",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(factTablesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"fact-tables", "update",
			"--id", "id",
			"--archived=true",
			"--column.column", "column",
			"--column.datatype", "number",
			"--column.always-inline-filter=true",
			"--column.auto-slices", "[string]",
			"--column.deleted=true",
			"--column.description", "description",
			"--column.is-auto-slice-column=true",
			"--column.json-fields", "{foo: {datatype: number}}",
			"--column.locked-auto-slices", "[string]",
			"--column.name", "name",
			"--column.number-format", "",
			"--columns-error", "columnsError",
			"--description", "description",
			"--event-name", "eventName",
			"--managed-by", "",
			"--name", "name",
			"--owner", "owner",
			"--project", "string",
			"--sql", "sql",
			"--tag", "string",
			"--user-id-type", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"archived: true\n" +
			"columns:\n" +
			"  - column: column\n" +
			"    datatype: number\n" +
			"    alwaysInlineFilter: true\n" +
			"    autoSlices:\n" +
			"      - string\n" +
			"    deleted: true\n" +
			"    description: description\n" +
			"    isAutoSliceColumn: true\n" +
			"    jsonFields:\n" +
			"      foo:\n" +
			"        datatype: number\n" +
			"    lockedAutoSlices:\n" +
			"      - string\n" +
			"    name: name\n" +
			"    numberFormat: ''\n" +
			"columnsError: columnsError\n" +
			"description: description\n" +
			"eventName: eventName\n" +
			"managedBy: ''\n" +
			"name: name\n" +
			"owner: owner\n" +
			"projects:\n" +
			"  - string\n" +
			"sql: sql\n" +
			"tags:\n" +
			"  - string\n" +
			"userIdTypes:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"fact-tables", "update",
			"--id", "id",
		)
	})
}

func TestFactTablesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"fact-tables", "list",
			"--datasource-id", "datasourceId",
			"--limit", "1",
			"--offset", "0",
			"--project-id", "projectId",
		)
	})
}

func TestFactTablesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"fact-tables", "delete",
			"--id", "id",
		)
	})
}
