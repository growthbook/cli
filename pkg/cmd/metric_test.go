// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
	"github.com/stainless-sdks/growthbook-cli/internal/requestflag"
)

func TestMetricsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"metrics", "create",
			"--datasource-id", "datasourceId",
			"--name", "name",
			"--type", "binomial",
			"--archived=true",
			"--behavior", "{cap: 0, capping: absolute, cappingSettings: {type: none, ignoreZeros: true, value: 0}, capValue: 0, conversionWindowEnd: 0, conversionWindowStart: 0, goal: increase, maxPercentChange: 0, minPercentChange: 0, minSampleSize: 0, priorSettings: {mean: 0, override: true, proper: true, stddev: 1}, riskThresholdDanger: 0, riskThresholdSuccess: 0, targetMDE: 0, windowSettings: {type: none, delayHours: 0, delayUnit: minutes, delayValue: 0, windowUnit: minutes, windowValue: 0}}",
			"--description", "description",
			"--managed-by", "",
			"--mixpanel", "{eventName: eventName, userAggregation: userAggregation, conditions: [{operator: operator, property: property, value: value}], eventValue: eventValue}",
			"--owner", "owner",
			"--project", "string",
			"--sql", "{conversionSQL: conversionSQL, identifierTypes: [string], denominatorMetricId: denominatorMetricId, userAggregationSQL: userAggregationSQL}",
			"--sql-builder", "{identifierTypeColumns: [{columnName: columnName, identifierType: identifierType}], tableName: tableName, timestampColumnName: timestampColumnName, conditions: [{column: column, operator: operator, value: value}], valueColumnName: valueColumnName}",
			"--tag", "string",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(metricsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"metrics", "create",
			"--datasource-id", "datasourceId",
			"--name", "name",
			"--type", "binomial",
			"--archived=true",
			"--behavior.cap", "0",
			"--behavior.capping", "absolute",
			"--behavior.capping-settings", "{type: none, ignoreZeros: true, value: 0}",
			"--behavior.cap-value", "0",
			"--behavior.conversion-window-end", "0",
			"--behavior.conversion-window-start", "0",
			"--behavior.goal", "increase",
			"--behavior.max-percent-change", "0",
			"--behavior.min-percent-change", "0",
			"--behavior.min-sample-size", "0",
			"--behavior.prior-settings", "{mean: 0, override: true, proper: true, stddev: 1}",
			"--behavior.risk-threshold-danger", "0",
			"--behavior.risk-threshold-success", "0",
			"--behavior.target-mde", "0",
			"--behavior.window-settings", "{type: none, delayHours: 0, delayUnit: minutes, delayValue: 0, windowUnit: minutes, windowValue: 0}",
			"--description", "description",
			"--managed-by", "",
			"--mixpanel.event-name", "eventName",
			"--mixpanel.user-aggregation", "userAggregation",
			"--mixpanel.conditions", "[{operator: operator, property: property, value: value}]",
			"--mixpanel.event-value", "eventValue",
			"--owner", "owner",
			"--project", "string",
			"--sql.conversion-sql", "conversionSQL",
			"--sql.identifier-types", "[string]",
			"--sql.denominator-metric-id", "denominatorMetricId",
			"--sql.user-aggregation-sql", "userAggregationSQL",
			"--sql-builder.identifier-type-columns", "[{columnName: columnName, identifierType: identifierType}]",
			"--sql-builder.table-name", "tableName",
			"--sql-builder.timestamp-column-name", "timestampColumnName",
			"--sql-builder.conditions", "[{column: column, operator: operator, value: value}]",
			"--sql-builder.value-column-name", "valueColumnName",
			"--tag", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"datasourceId: datasourceId\n" +
			"name: name\n" +
			"type: binomial\n" +
			"archived: true\n" +
			"behavior:\n" +
			"  cap: 0\n" +
			"  capping: absolute\n" +
			"  cappingSettings:\n" +
			"    type: none\n" +
			"    ignoreZeros: true\n" +
			"    value: 0\n" +
			"  capValue: 0\n" +
			"  conversionWindowEnd: 0\n" +
			"  conversionWindowStart: 0\n" +
			"  goal: increase\n" +
			"  maxPercentChange: 0\n" +
			"  minPercentChange: 0\n" +
			"  minSampleSize: 0\n" +
			"  priorSettings:\n" +
			"    mean: 0\n" +
			"    override: true\n" +
			"    proper: true\n" +
			"    stddev: 1\n" +
			"  riskThresholdDanger: 0\n" +
			"  riskThresholdSuccess: 0\n" +
			"  targetMDE: 0\n" +
			"  windowSettings:\n" +
			"    type: none\n" +
			"    delayHours: 0\n" +
			"    delayUnit: minutes\n" +
			"    delayValue: 0\n" +
			"    windowUnit: minutes\n" +
			"    windowValue: 0\n" +
			"description: description\n" +
			"managedBy: ''\n" +
			"mixpanel:\n" +
			"  eventName: eventName\n" +
			"  userAggregation: userAggregation\n" +
			"  conditions:\n" +
			"    - operator: operator\n" +
			"      property: property\n" +
			"      value: value\n" +
			"  eventValue: eventValue\n" +
			"owner: owner\n" +
			"projects:\n" +
			"  - string\n" +
			"sql:\n" +
			"  conversionSQL: conversionSQL\n" +
			"  identifierTypes:\n" +
			"    - string\n" +
			"  denominatorMetricId: denominatorMetricId\n" +
			"  userAggregationSQL: userAggregationSQL\n" +
			"sqlBuilder:\n" +
			"  identifierTypeColumns:\n" +
			"    - columnName: columnName\n" +
			"      identifierType: identifierType\n" +
			"  tableName: tableName\n" +
			"  timestampColumnName: timestampColumnName\n" +
			"  conditions:\n" +
			"    - column: column\n" +
			"      operator: operator\n" +
			"      value: value\n" +
			"  valueColumnName: valueColumnName\n" +
			"tags:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"metrics", "create",
		)
	})
}

func TestMetricsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"metrics", "retrieve",
			"--id", "id",
		)
	})
}

func TestMetricsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"metrics", "update",
			"--id", "id",
			"--archived=true",
			"--behavior", "{cap: 0, capping: absolute, cappingSettings: {type: none, ignoreZeros: true, value: 0}, capValue: 0, conversionWindowEnd: 0, conversionWindowStart: 0, goal: increase, maxPercentChange: 0, minPercentChange: 0, minSampleSize: 0, priorSettings: {mean: 0, override: true, proper: true, stddev: 1}, riskThresholdDanger: 0, riskThresholdSuccess: 0, targetMDE: 0, windowSettings: {type: none, delayHours: 0, delayUnit: minutes, delayValue: 0, windowUnit: minutes, windowValue: 0}}",
			"--description", "description",
			"--managed-by", "",
			"--mixpanel", "{conditions: [{operator: operator, property: property, value: value}], eventName: eventName, eventValue: eventValue, userAggregation: userAggregation}",
			"--name", "name",
			"--owner", "owner",
			"--project", "string",
			"--sql", "{conversionSQL: conversionSQL, denominatorMetricId: denominatorMetricId, identifierTypes: [string], userAggregationSQL: userAggregationSQL}",
			"--sql-builder", "{conditions: [{column: column, operator: operator, value: value}], identifierTypeColumns: [{columnName: columnName, identifierType: identifierType}], tableName: tableName, timestampColumnName: timestampColumnName, valueColumnName: valueColumnName}",
			"--tag", "string",
			"--type", "binomial",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(metricsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"metrics", "update",
			"--id", "id",
			"--archived=true",
			"--behavior.cap", "0",
			"--behavior.capping", "absolute",
			"--behavior.capping-settings", "{type: none, ignoreZeros: true, value: 0}",
			"--behavior.cap-value", "0",
			"--behavior.conversion-window-end", "0",
			"--behavior.conversion-window-start", "0",
			"--behavior.goal", "increase",
			"--behavior.max-percent-change", "0",
			"--behavior.min-percent-change", "0",
			"--behavior.min-sample-size", "0",
			"--behavior.prior-settings", "{mean: 0, override: true, proper: true, stddev: 1}",
			"--behavior.risk-threshold-danger", "0",
			"--behavior.risk-threshold-success", "0",
			"--behavior.target-mde", "0",
			"--behavior.window-settings", "{type: none, delayHours: 0, delayUnit: minutes, delayValue: 0, windowUnit: minutes, windowValue: 0}",
			"--description", "description",
			"--managed-by", "",
			"--mixpanel.conditions", "[{operator: operator, property: property, value: value}]",
			"--mixpanel.event-name", "eventName",
			"--mixpanel.event-value", "eventValue",
			"--mixpanel.user-aggregation", "userAggregation",
			"--name", "name",
			"--owner", "owner",
			"--project", "string",
			"--sql.conversion-sql", "conversionSQL",
			"--sql.denominator-metric-id", "denominatorMetricId",
			"--sql.identifier-types", "[string]",
			"--sql.user-aggregation-sql", "userAggregationSQL",
			"--sql-builder.conditions", "[{column: column, operator: operator, value: value}]",
			"--sql-builder.identifier-type-columns", "[{columnName: columnName, identifierType: identifierType}]",
			"--sql-builder.table-name", "tableName",
			"--sql-builder.timestamp-column-name", "timestampColumnName",
			"--sql-builder.value-column-name", "valueColumnName",
			"--tag", "string",
			"--type", "binomial",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"archived: true\n" +
			"behavior:\n" +
			"  cap: 0\n" +
			"  capping: absolute\n" +
			"  cappingSettings:\n" +
			"    type: none\n" +
			"    ignoreZeros: true\n" +
			"    value: 0\n" +
			"  capValue: 0\n" +
			"  conversionWindowEnd: 0\n" +
			"  conversionWindowStart: 0\n" +
			"  goal: increase\n" +
			"  maxPercentChange: 0\n" +
			"  minPercentChange: 0\n" +
			"  minSampleSize: 0\n" +
			"  priorSettings:\n" +
			"    mean: 0\n" +
			"    override: true\n" +
			"    proper: true\n" +
			"    stddev: 1\n" +
			"  riskThresholdDanger: 0\n" +
			"  riskThresholdSuccess: 0\n" +
			"  targetMDE: 0\n" +
			"  windowSettings:\n" +
			"    type: none\n" +
			"    delayHours: 0\n" +
			"    delayUnit: minutes\n" +
			"    delayValue: 0\n" +
			"    windowUnit: minutes\n" +
			"    windowValue: 0\n" +
			"description: description\n" +
			"managedBy: ''\n" +
			"mixpanel:\n" +
			"  conditions:\n" +
			"    - operator: operator\n" +
			"      property: property\n" +
			"      value: value\n" +
			"  eventName: eventName\n" +
			"  eventValue: eventValue\n" +
			"  userAggregation: userAggregation\n" +
			"name: name\n" +
			"owner: owner\n" +
			"projects:\n" +
			"  - string\n" +
			"sql:\n" +
			"  conversionSQL: conversionSQL\n" +
			"  denominatorMetricId: denominatorMetricId\n" +
			"  identifierTypes:\n" +
			"    - string\n" +
			"  userAggregationSQL: userAggregationSQL\n" +
			"sqlBuilder:\n" +
			"  conditions:\n" +
			"    - column: column\n" +
			"      operator: operator\n" +
			"      value: value\n" +
			"  identifierTypeColumns:\n" +
			"    - columnName: columnName\n" +
			"      identifierType: identifierType\n" +
			"  tableName: tableName\n" +
			"  timestampColumnName: timestampColumnName\n" +
			"  valueColumnName: valueColumnName\n" +
			"tags:\n" +
			"  - string\n" +
			"type: binomial\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"metrics", "update",
			"--id", "id",
		)
	})
}

func TestMetricsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"metrics", "list",
			"--datasource-id", "datasourceId",
			"--limit", "1",
			"--offset", "0",
			"--project-id", "projectId",
		)
	})
}

func TestMetricsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"metrics", "delete",
			"--id", "id",
		)
	})
}
