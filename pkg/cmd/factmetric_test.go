// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
	"github.com/stainless-sdks/growthbook-cli/internal/requestflag"
)

func TestFactMetricsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"fact-metrics", "create",
			"--metric-type", "proportion",
			"--name", "name",
			"--numerator", "{factTableId: factTableId, aggregateFilter: aggregateFilter, aggregateFilterColumn: aggregateFilterColumn, aggregation: sum, column: column, filters: [string], inlineFilters: {foo: [string]}, rowFilters: [{operator: '=', column: column, values: [string]}]}",
			"--capping-settings", "{type: none, ignoreZeros: true, value: 0}",
			"--denominator", "{column: column, factTableId: factTableId, aggregation: sum, filters: [string], inlineFilters: {foo: [string]}, rowFilters: [{operator: '=', column: column, values: [string]}]}",
			"--description", "description",
			"--display-as-percentage=true",
			"--inverse=true",
			"--managed-by", "",
			"--max-percent-change", "0",
			"--metric-auto-slice", "string",
			"--min-percent-change", "0",
			"--min-sample-size", "0",
			"--owner", "owner",
			"--prior-settings", "{mean: 0, override: true, proper: true, stddev: 1}",
			"--project", "string",
			"--quantile-settings", "{ignoreZeros: true, quantile: 0.001, type: event, quantileEventCountColumn: quantileEventCountColumn}",
			"--regression-adjustment-settings", "{override: true, days: 0, enabled: true}",
			"--risk-threshold-danger", "0",
			"--risk-threshold-success", "0",
			"--tag", "string",
			"--target-mde", "0",
			"--window-settings", "{type: none, delayHours: 0, delayUnit: minutes, delayValue: 0, windowUnit: minutes, windowValue: 0}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(factMetricsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"fact-metrics", "create",
			"--metric-type", "proportion",
			"--name", "name",
			"--numerator.fact-table-id", "factTableId",
			"--numerator.aggregate-filter", "aggregateFilter",
			"--numerator.aggregate-filter-column", "aggregateFilterColumn",
			"--numerator.aggregation", "sum",
			"--numerator.column", "column",
			"--numerator.filters", "[string]",
			"--numerator.inline-filters", "{foo: [string]}",
			"--numerator.row-filters", "[{operator: '=', column: column, values: [string]}]",
			"--capping-settings.type", "none",
			"--capping-settings.ignore-zeros=true",
			"--capping-settings.value", "0",
			"--denominator.column", "column",
			"--denominator.fact-table-id", "factTableId",
			"--denominator.aggregation", "sum",
			"--denominator.filters", "[string]",
			"--denominator.inline-filters", "{foo: [string]}",
			"--denominator.row-filters", "[{operator: '=', column: column, values: [string]}]",
			"--description", "description",
			"--display-as-percentage=true",
			"--inverse=true",
			"--managed-by", "",
			"--max-percent-change", "0",
			"--metric-auto-slice", "string",
			"--min-percent-change", "0",
			"--min-sample-size", "0",
			"--owner", "owner",
			"--prior-settings.mean", "0",
			"--prior-settings.override=true",
			"--prior-settings.proper=true",
			"--prior-settings.stddev", "1",
			"--project", "string",
			"--quantile-settings.ignore-zeros=true",
			"--quantile-settings.quantile", "0.001",
			"--quantile-settings.type", "event",
			"--quantile-settings.quantile-event-count-column", "quantileEventCountColumn",
			"--regression-adjustment-settings.override=true",
			"--regression-adjustment-settings.days", "0",
			"--regression-adjustment-settings.enabled=true",
			"--risk-threshold-danger", "0",
			"--risk-threshold-success", "0",
			"--tag", "string",
			"--target-mde", "0",
			"--window-settings.type", "none",
			"--window-settings.delay-hours", "0",
			"--window-settings.delay-unit", "minutes",
			"--window-settings.delay-value", "0",
			"--window-settings.window-unit", "minutes",
			"--window-settings.window-value", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"metricType: proportion\n" +
			"name: name\n" +
			"numerator:\n" +
			"  factTableId: factTableId\n" +
			"  aggregateFilter: aggregateFilter\n" +
			"  aggregateFilterColumn: aggregateFilterColumn\n" +
			"  aggregation: sum\n" +
			"  column: column\n" +
			"  filters:\n" +
			"    - string\n" +
			"  inlineFilters:\n" +
			"    foo:\n" +
			"      - string\n" +
			"  rowFilters:\n" +
			"    - operator: '='\n" +
			"      column: column\n" +
			"      values:\n" +
			"        - string\n" +
			"cappingSettings:\n" +
			"  type: none\n" +
			"  ignoreZeros: true\n" +
			"  value: 0\n" +
			"denominator:\n" +
			"  column: column\n" +
			"  factTableId: factTableId\n" +
			"  aggregation: sum\n" +
			"  filters:\n" +
			"    - string\n" +
			"  inlineFilters:\n" +
			"    foo:\n" +
			"      - string\n" +
			"  rowFilters:\n" +
			"    - operator: '='\n" +
			"      column: column\n" +
			"      values:\n" +
			"        - string\n" +
			"description: description\n" +
			"displayAsPercentage: true\n" +
			"inverse: true\n" +
			"managedBy: ''\n" +
			"maxPercentChange: 0\n" +
			"metricAutoSlices:\n" +
			"  - string\n" +
			"minPercentChange: 0\n" +
			"minSampleSize: 0\n" +
			"owner: owner\n" +
			"priorSettings:\n" +
			"  mean: 0\n" +
			"  override: true\n" +
			"  proper: true\n" +
			"  stddev: 1\n" +
			"projects:\n" +
			"  - string\n" +
			"quantileSettings:\n" +
			"  ignoreZeros: true\n" +
			"  quantile: 0.001\n" +
			"  type: event\n" +
			"  quantileEventCountColumn: quantileEventCountColumn\n" +
			"regressionAdjustmentSettings:\n" +
			"  override: true\n" +
			"  days: 0\n" +
			"  enabled: true\n" +
			"riskThresholdDanger: 0\n" +
			"riskThresholdSuccess: 0\n" +
			"tags:\n" +
			"  - string\n" +
			"targetMDE: 0\n" +
			"windowSettings:\n" +
			"  type: none\n" +
			"  delayHours: 0\n" +
			"  delayUnit: minutes\n" +
			"  delayValue: 0\n" +
			"  windowUnit: minutes\n" +
			"  windowValue: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"fact-metrics", "create",
		)
	})
}

func TestFactMetricsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"fact-metrics", "retrieve",
			"--id", "id",
		)
	})
}

func TestFactMetricsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"fact-metrics", "update",
			"--id", "id",
			"--archived=true",
			"--capping-settings", "{type: none, ignoreZeros: true, value: 0}",
			"--denominator", "{column: column, factTableId: factTableId, aggregation: sum, filters: [string], inlineFilters: {foo: [string]}, rowFilters: [{operator: '=', column: column, values: [string]}]}",
			"--description", "description",
			"--display-as-percentage=true",
			"--inverse=true",
			"--managed-by", "",
			"--max-percent-change", "0",
			"--metric-auto-slice", "string",
			"--metric-type", "proportion",
			"--min-percent-change", "0",
			"--min-sample-size", "0",
			"--name", "name",
			"--numerator", "{factTableId: factTableId, aggregateFilter: aggregateFilter, aggregateFilterColumn: aggregateFilterColumn, aggregation: sum, column: column, filters: [string], inlineFilters: {foo: [string]}, rowFilters: [{operator: '=', column: column, values: [string]}]}",
			"--owner", "owner",
			"--prior-settings", "{mean: 0, override: true, proper: true, stddev: 1}",
			"--project", "string",
			"--quantile-settings", "{ignoreZeros: true, quantile: 0.001, type: event, quantileEventCountColumn: quantileEventCountColumn}",
			"--regression-adjustment-settings", "{override: true, days: 0, enabled: true}",
			"--risk-threshold-danger", "0",
			"--risk-threshold-success", "0",
			"--tag", "string",
			"--target-mde", "0",
			"--window-settings", "{type: none, delayHours: 0, delayUnit: minutes, delayValue: 0, windowUnit: minutes, windowValue: 0}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(factMetricsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"fact-metrics", "update",
			"--id", "id",
			"--archived=true",
			"--capping-settings.type", "none",
			"--capping-settings.ignore-zeros=true",
			"--capping-settings.value", "0",
			"--denominator.column", "column",
			"--denominator.fact-table-id", "factTableId",
			"--denominator.aggregation", "sum",
			"--denominator.filters", "[string]",
			"--denominator.inline-filters", "{foo: [string]}",
			"--denominator.row-filters", "[{operator: '=', column: column, values: [string]}]",
			"--description", "description",
			"--display-as-percentage=true",
			"--inverse=true",
			"--managed-by", "",
			"--max-percent-change", "0",
			"--metric-auto-slice", "string",
			"--metric-type", "proportion",
			"--min-percent-change", "0",
			"--min-sample-size", "0",
			"--name", "name",
			"--numerator.fact-table-id", "factTableId",
			"--numerator.aggregate-filter", "aggregateFilter",
			"--numerator.aggregate-filter-column", "aggregateFilterColumn",
			"--numerator.aggregation", "sum",
			"--numerator.column", "column",
			"--numerator.filters", "[string]",
			"--numerator.inline-filters", "{foo: [string]}",
			"--numerator.row-filters", "[{operator: '=', column: column, values: [string]}]",
			"--owner", "owner",
			"--prior-settings.mean", "0",
			"--prior-settings.override=true",
			"--prior-settings.proper=true",
			"--prior-settings.stddev", "1",
			"--project", "string",
			"--quantile-settings.ignore-zeros=true",
			"--quantile-settings.quantile", "0.001",
			"--quantile-settings.type", "event",
			"--quantile-settings.quantile-event-count-column", "quantileEventCountColumn",
			"--regression-adjustment-settings.override=true",
			"--regression-adjustment-settings.days", "0",
			"--regression-adjustment-settings.enabled=true",
			"--risk-threshold-danger", "0",
			"--risk-threshold-success", "0",
			"--tag", "string",
			"--target-mde", "0",
			"--window-settings.type", "none",
			"--window-settings.delay-hours", "0",
			"--window-settings.delay-unit", "minutes",
			"--window-settings.delay-value", "0",
			"--window-settings.window-unit", "minutes",
			"--window-settings.window-value", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"archived: true\n" +
			"cappingSettings:\n" +
			"  type: none\n" +
			"  ignoreZeros: true\n" +
			"  value: 0\n" +
			"denominator:\n" +
			"  column: column\n" +
			"  factTableId: factTableId\n" +
			"  aggregation: sum\n" +
			"  filters:\n" +
			"    - string\n" +
			"  inlineFilters:\n" +
			"    foo:\n" +
			"      - string\n" +
			"  rowFilters:\n" +
			"    - operator: '='\n" +
			"      column: column\n" +
			"      values:\n" +
			"        - string\n" +
			"description: description\n" +
			"displayAsPercentage: true\n" +
			"inverse: true\n" +
			"managedBy: ''\n" +
			"maxPercentChange: 0\n" +
			"metricAutoSlices:\n" +
			"  - string\n" +
			"metricType: proportion\n" +
			"minPercentChange: 0\n" +
			"minSampleSize: 0\n" +
			"name: name\n" +
			"numerator:\n" +
			"  factTableId: factTableId\n" +
			"  aggregateFilter: aggregateFilter\n" +
			"  aggregateFilterColumn: aggregateFilterColumn\n" +
			"  aggregation: sum\n" +
			"  column: column\n" +
			"  filters:\n" +
			"    - string\n" +
			"  inlineFilters:\n" +
			"    foo:\n" +
			"      - string\n" +
			"  rowFilters:\n" +
			"    - operator: '='\n" +
			"      column: column\n" +
			"      values:\n" +
			"        - string\n" +
			"owner: owner\n" +
			"priorSettings:\n" +
			"  mean: 0\n" +
			"  override: true\n" +
			"  proper: true\n" +
			"  stddev: 1\n" +
			"projects:\n" +
			"  - string\n" +
			"quantileSettings:\n" +
			"  ignoreZeros: true\n" +
			"  quantile: 0.001\n" +
			"  type: event\n" +
			"  quantileEventCountColumn: quantileEventCountColumn\n" +
			"regressionAdjustmentSettings:\n" +
			"  override: true\n" +
			"  days: 0\n" +
			"  enabled: true\n" +
			"riskThresholdDanger: 0\n" +
			"riskThresholdSuccess: 0\n" +
			"tags:\n" +
			"  - string\n" +
			"targetMDE: 0\n" +
			"windowSettings:\n" +
			"  type: none\n" +
			"  delayHours: 0\n" +
			"  delayUnit: minutes\n" +
			"  delayValue: 0\n" +
			"  windowUnit: minutes\n" +
			"  windowValue: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"fact-metrics", "update",
			"--id", "id",
		)
	})
}

func TestFactMetricsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"fact-metrics", "list",
			"--datasource-id", "datasourceId",
			"--fact-table-id", "factTableId",
			"--limit", "1",
			"--offset", "0",
			"--project-id", "projectId",
		)
	})
}

func TestFactMetricsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"fact-metrics", "delete",
			"--id", "id",
		)
	})
}

func TestFactMetricsCreateAnalysis(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"fact-metrics", "create-analysis",
			"--id", "id",
			"--additional-denominator-filter", "string",
			"--additional-numerator-filter", "string",
			"--lookback-days", "1",
			"--population-id", "populationId",
			"--population-type", "factTable",
			"--use-cache=true",
			"--user-id-type", "userIdType",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"additionalDenominatorFilters:\n" +
			"  - string\n" +
			"additionalNumeratorFilters:\n" +
			"  - string\n" +
			"lookbackDays: 1\n" +
			"populationId: populationId\n" +
			"populationType: factTable\n" +
			"useCache: true\n" +
			"userIdType: userIdType\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"fact-metrics", "create-analysis",
			"--id", "id",
		)
	})
}
