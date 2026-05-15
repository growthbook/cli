// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/growthbook/cli/internal/mocktest"
	"github.com/growthbook/cli/internal/requestflag"
)

func TestProductAnalyticsCreateDataSourceExploration(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"product-analytics", "create-data-source-exploration",
			"--chart-type", "line",
			"--dataset", "{columnTypes: {foo: string}, path: path, table: table, timestampColumn: timestampColumn, type: data_source, values: [{name: name, rowFilters: [{operator: '=', column: column, values: [string]}], type: data_source, unit: unit, valueColumn: valueColumn, valueType: unit_count}]}",
			"--datasource", "datasource",
			"--date-range", "{predefined: today, endDate: endDate, lookbackUnit: hour, lookbackValue: 0, startDate: startDate}",
			"--dimension", "{column: column, dateGranularity: auto, dimensionType: date}",
			"--type", "data_source",
			"--cache", "preferred",
			"--show-as", "total",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(productAnalyticsCreateDataSourceExploration)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"product-analytics", "create-data-source-exploration",
			"--chart-type", "line",
			"--dataset.column-types", "{foo: string}",
			"--dataset.path", "path",
			"--dataset.table", "table",
			"--dataset.timestamp-column", "timestampColumn",
			"--dataset.type", "data_source",
			"--dataset.values", "[{name: name, rowFilters: [{operator: '=', column: column, values: [string]}], type: data_source, unit: unit, valueColumn: valueColumn, valueType: unit_count}]",
			"--datasource", "datasource",
			"--date-range.predefined", "today",
			"--date-range.end-date", "endDate",
			"--date-range.lookback-unit", "hour",
			"--date-range.lookback-value", "0",
			"--date-range.start-date", "startDate",
			"--dimension", "{column: column, dateGranularity: auto, dimensionType: date}",
			"--type", "data_source",
			"--cache", "preferred",
			"--show-as", "total",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"chartType: line\n" +
			"dataset:\n" +
			"  columnTypes:\n" +
			"    foo: string\n" +
			"  path: path\n" +
			"  table: table\n" +
			"  timestampColumn: timestampColumn\n" +
			"  type: data_source\n" +
			"  values:\n" +
			"    - name: name\n" +
			"      rowFilters:\n" +
			"        - operator: '='\n" +
			"          column: column\n" +
			"          values:\n" +
			"            - string\n" +
			"      type: data_source\n" +
			"      unit: unit\n" +
			"      valueColumn: valueColumn\n" +
			"      valueType: unit_count\n" +
			"datasource: datasource\n" +
			"dateRange:\n" +
			"  predefined: today\n" +
			"  endDate: endDate\n" +
			"  lookbackUnit: hour\n" +
			"  lookbackValue: 0\n" +
			"  startDate: startDate\n" +
			"dimensions:\n" +
			"  - column: column\n" +
			"    dateGranularity: auto\n" +
			"    dimensionType: date\n" +
			"type: data_source\n" +
			"showAs: total\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"product-analytics", "create-data-source-exploration",
			"--cache", "preferred",
		)
	})
}

func TestProductAnalyticsCreateMetricExploration(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"product-analytics", "create-metric-exploration",
			"--chart-type", "line",
			"--dataset", "{type: metric, values: [{denominatorUnit: denominatorUnit, metricId: metricId, name: name, rowFilters: [{operator: '=', column: column, values: [string]}], type: metric, unit: unit}]}",
			"--datasource", "datasource",
			"--date-range", "{predefined: today, endDate: endDate, lookbackUnit: hour, lookbackValue: 0, startDate: startDate}",
			"--dimension", "{column: column, dateGranularity: auto, dimensionType: date}",
			"--type", "metric",
			"--cache", "preferred",
			"--show-as", "total",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(productAnalyticsCreateMetricExploration)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"product-analytics", "create-metric-exploration",
			"--chart-type", "line",
			"--dataset.type", "metric",
			"--dataset.values", "[{denominatorUnit: denominatorUnit, metricId: metricId, name: name, rowFilters: [{operator: '=', column: column, values: [string]}], type: metric, unit: unit}]",
			"--datasource", "datasource",
			"--date-range.predefined", "today",
			"--date-range.end-date", "endDate",
			"--date-range.lookback-unit", "hour",
			"--date-range.lookback-value", "0",
			"--date-range.start-date", "startDate",
			"--dimension", "{column: column, dateGranularity: auto, dimensionType: date}",
			"--type", "metric",
			"--cache", "preferred",
			"--show-as", "total",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"chartType: line\n" +
			"dataset:\n" +
			"  type: metric\n" +
			"  values:\n" +
			"    - denominatorUnit: denominatorUnit\n" +
			"      metricId: metricId\n" +
			"      name: name\n" +
			"      rowFilters:\n" +
			"        - operator: '='\n" +
			"          column: column\n" +
			"          values:\n" +
			"            - string\n" +
			"      type: metric\n" +
			"      unit: unit\n" +
			"datasource: datasource\n" +
			"dateRange:\n" +
			"  predefined: today\n" +
			"  endDate: endDate\n" +
			"  lookbackUnit: hour\n" +
			"  lookbackValue: 0\n" +
			"  startDate: startDate\n" +
			"dimensions:\n" +
			"  - column: column\n" +
			"    dateGranularity: auto\n" +
			"    dimensionType: date\n" +
			"type: metric\n" +
			"showAs: total\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"product-analytics", "create-metric-exploration",
			"--cache", "preferred",
		)
	})
}

func TestProductAnalyticsRunFactTableExploration(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"product-analytics", "run-fact-table-exploration",
			"--chart-type", "line",
			"--dataset", "{factTableId: factTableId, type: fact_table, values: [{name: name, rowFilters: [{operator: '=', column: column, values: [string]}], type: fact_table, unit: unit, valueColumn: valueColumn, valueType: unit_count}]}",
			"--datasource", "datasource",
			"--date-range", "{predefined: today, endDate: endDate, lookbackUnit: hour, lookbackValue: 0, startDate: startDate}",
			"--dimension", "{column: column, dateGranularity: auto, dimensionType: date}",
			"--type", "fact_table",
			"--cache", "preferred",
			"--show-as", "total",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(productAnalyticsRunFactTableExploration)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"product-analytics", "run-fact-table-exploration",
			"--chart-type", "line",
			"--dataset.fact-table-id", "factTableId",
			"--dataset.type", "fact_table",
			"--dataset.values", "[{name: name, rowFilters: [{operator: '=', column: column, values: [string]}], type: fact_table, unit: unit, valueColumn: valueColumn, valueType: unit_count}]",
			"--datasource", "datasource",
			"--date-range.predefined", "today",
			"--date-range.end-date", "endDate",
			"--date-range.lookback-unit", "hour",
			"--date-range.lookback-value", "0",
			"--date-range.start-date", "startDate",
			"--dimension", "{column: column, dateGranularity: auto, dimensionType: date}",
			"--type", "fact_table",
			"--cache", "preferred",
			"--show-as", "total",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"chartType: line\n" +
			"dataset:\n" +
			"  factTableId: factTableId\n" +
			"  type: fact_table\n" +
			"  values:\n" +
			"    - name: name\n" +
			"      rowFilters:\n" +
			"        - operator: '='\n" +
			"          column: column\n" +
			"          values:\n" +
			"            - string\n" +
			"      type: fact_table\n" +
			"      unit: unit\n" +
			"      valueColumn: valueColumn\n" +
			"      valueType: unit_count\n" +
			"datasource: datasource\n" +
			"dateRange:\n" +
			"  predefined: today\n" +
			"  endDate: endDate\n" +
			"  lookbackUnit: hour\n" +
			"  lookbackValue: 0\n" +
			"  startDate: startDate\n" +
			"dimensions:\n" +
			"  - column: column\n" +
			"    dateGranularity: auto\n" +
			"    dimensionType: date\n" +
			"type: fact_table\n" +
			"showAs: total\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"product-analytics", "run-fact-table-exploration",
			"--cache", "preferred",
		)
	})
}
