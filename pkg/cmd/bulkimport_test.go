// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
	"github.com/stainless-sdks/growthbook-cli/internal/requestflag"
)

func TestBulkImportCreateFacts(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"bulk-import", "create-facts",
			"--fact-metric", "{id: id, data: {metricType: proportion, name: name, numerator: {factTableId: factTableId, aggregateFilter: aggregateFilter, aggregateFilterColumn: aggregateFilterColumn, aggregation: sum, column: column, filters: [string], inlineFilters: {foo: [string]}, rowFilters: [{operator: '=', column: column, values: [string]}]}, archived: true, cappingSettings: {type: none, ignoreZeros: true, value: 0}, denominator: {column: column, factTableId: factTableId, aggregation: sum, filters: [string], inlineFilters: {foo: [string]}, rowFilters: [{operator: '=', column: column, values: [string]}]}, description: description, displayAsPercentage: true, inverse: true, managedBy: '', maxPercentChange: 0, metricAutoSlices: [string], minPercentChange: 0, minSampleSize: 0, owner: owner, priorSettings: {mean: 0, override: true, proper: true, stddev: 1}, projects: [string], quantileSettings: {ignoreZeros: true, quantile: 0.001, type: event, quantileEventCountColumn: quantileEventCountColumn}, regressionAdjustmentSettings: {override: true, days: 0, enabled: true}, riskThresholdDanger: 0, riskThresholdSuccess: 0, tags: [string], targetMDE: 0, windowSettings: {type: none, delayHours: 0, delayUnit: minutes, delayValue: 0, windowUnit: minutes, windowValue: 0}}}",
			"--fact-table-filter", "{id: id, data: {name: name, value: country = 'US', description: description, managedBy: ''}, factTableId: factTableId}",
			"--fact-table", "{id: id, data: {datasource: datasource, name: name, sql: sql, userIdTypes: [string], description: description, eventName: eventName, managedBy: '', owner: owner, projects: [string], tags: [string]}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(bulkImportCreateFacts)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--domain", "string",
			"bulk-import", "create-facts",
			"--fact-metric.id", "id",
			"--fact-metric.data", "{metricType: proportion, name: name, numerator: {factTableId: factTableId, aggregateFilter: aggregateFilter, aggregateFilterColumn: aggregateFilterColumn, aggregation: sum, column: column, filters: [string], inlineFilters: {foo: [string]}, rowFilters: [{operator: '=', column: column, values: [string]}]}, archived: true, cappingSettings: {type: none, ignoreZeros: true, value: 0}, denominator: {column: column, factTableId: factTableId, aggregation: sum, filters: [string], inlineFilters: {foo: [string]}, rowFilters: [{operator: '=', column: column, values: [string]}]}, description: description, displayAsPercentage: true, inverse: true, managedBy: '', maxPercentChange: 0, metricAutoSlices: [string], minPercentChange: 0, minSampleSize: 0, owner: owner, priorSettings: {mean: 0, override: true, proper: true, stddev: 1}, projects: [string], quantileSettings: {ignoreZeros: true, quantile: 0.001, type: event, quantileEventCountColumn: quantileEventCountColumn}, regressionAdjustmentSettings: {override: true, days: 0, enabled: true}, riskThresholdDanger: 0, riskThresholdSuccess: 0, tags: [string], targetMDE: 0, windowSettings: {type: none, delayHours: 0, delayUnit: minutes, delayValue: 0, windowUnit: minutes, windowValue: 0}}",
			"--fact-table-filter.id", "id",
			"--fact-table-filter.data", "{name: name, value: country = 'US', description: description, managedBy: ''}",
			"--fact-table-filter.fact-table-id", "factTableId",
			"--fact-table.id", "id",
			"--fact-table.data", "{datasource: datasource, name: name, sql: sql, userIdTypes: [string], description: description, eventName: eventName, managedBy: '', owner: owner, projects: [string], tags: [string]}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"factMetrics:\n" +
			"  - id: id\n" +
			"    data:\n" +
			"      metricType: proportion\n" +
			"      name: name\n" +
			"      numerator:\n" +
			"        factTableId: factTableId\n" +
			"        aggregateFilter: aggregateFilter\n" +
			"        aggregateFilterColumn: aggregateFilterColumn\n" +
			"        aggregation: sum\n" +
			"        column: column\n" +
			"        filters:\n" +
			"          - string\n" +
			"        inlineFilters:\n" +
			"          foo:\n" +
			"            - string\n" +
			"        rowFilters:\n" +
			"          - operator: '='\n" +
			"            column: column\n" +
			"            values:\n" +
			"              - string\n" +
			"      archived: true\n" +
			"      cappingSettings:\n" +
			"        type: none\n" +
			"        ignoreZeros: true\n" +
			"        value: 0\n" +
			"      denominator:\n" +
			"        column: column\n" +
			"        factTableId: factTableId\n" +
			"        aggregation: sum\n" +
			"        filters:\n" +
			"          - string\n" +
			"        inlineFilters:\n" +
			"          foo:\n" +
			"            - string\n" +
			"        rowFilters:\n" +
			"          - operator: '='\n" +
			"            column: column\n" +
			"            values:\n" +
			"              - string\n" +
			"      description: description\n" +
			"      displayAsPercentage: true\n" +
			"      inverse: true\n" +
			"      managedBy: ''\n" +
			"      maxPercentChange: 0\n" +
			"      metricAutoSlices:\n" +
			"        - string\n" +
			"      minPercentChange: 0\n" +
			"      minSampleSize: 0\n" +
			"      owner: owner\n" +
			"      priorSettings:\n" +
			"        mean: 0\n" +
			"        override: true\n" +
			"        proper: true\n" +
			"        stddev: 1\n" +
			"      projects:\n" +
			"        - string\n" +
			"      quantileSettings:\n" +
			"        ignoreZeros: true\n" +
			"        quantile: 0.001\n" +
			"        type: event\n" +
			"        quantileEventCountColumn: quantileEventCountColumn\n" +
			"      regressionAdjustmentSettings:\n" +
			"        override: true\n" +
			"        days: 0\n" +
			"        enabled: true\n" +
			"      riskThresholdDanger: 0\n" +
			"      riskThresholdSuccess: 0\n" +
			"      tags:\n" +
			"        - string\n" +
			"      targetMDE: 0\n" +
			"      windowSettings:\n" +
			"        type: none\n" +
			"        delayHours: 0\n" +
			"        delayUnit: minutes\n" +
			"        delayValue: 0\n" +
			"        windowUnit: minutes\n" +
			"        windowValue: 0\n" +
			"factTableFilters:\n" +
			"  - id: id\n" +
			"    data:\n" +
			"      name: name\n" +
			"      value: country = 'US'\n" +
			"      description: description\n" +
			"      managedBy: ''\n" +
			"    factTableId: factTableId\n" +
			"factTables:\n" +
			"  - id: id\n" +
			"    data:\n" +
			"      datasource: datasource\n" +
			"      name: name\n" +
			"      sql: sql\n" +
			"      userIdTypes:\n" +
			"        - string\n" +
			"      description: description\n" +
			"      eventName: eventName\n" +
			"      managedBy: ''\n" +
			"      owner: owner\n" +
			"      projects:\n" +
			"        - string\n" +
			"      tags:\n" +
			"        - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--domain", "string",
			"bulk-import", "create-facts",
		)
	})
}
