// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
	"github.com/stainless-sdks/growthbook-cli/internal/requestflag"
)

func TestExperimentTemplatesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-templates", "create",
			"--datasource", "datasource",
			"--exposure-query-id", "exposureQueryId",
			"--stats-engine", "bayesian",
			"--targeting", "{condition: condition, coverage: 0, prerequisites: [{id: id, condition: condition}], savedGroups: [{ids: [string], match: all}]}",
			"--template-metadata", "{name: name, description: description}",
			"--type", "standard",
			"--activation-metric", "activationMetric",
			"--custom-fields", "{foo: string}",
			"--custom-metric-slice", "{slices: [{column: column, levels: [string]}]}",
			"--description", "description",
			"--disable-sticky-bucketing=true",
			"--fallback-attribute", "fallbackAttribute",
			"--goal-metric", "string",
			"--guardrail-metric", "string",
			"--hash-attribute", "hashAttribute",
			"--hypothesis", "hypothesis",
			"--project", "project",
			"--secondary-metric", "string",
			"--segment", "segment",
			"--skip-partial-data=true",
			"--tag", "string",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(experimentTemplatesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-templates", "create",
			"--datasource", "datasource",
			"--exposure-query-id", "exposureQueryId",
			"--stats-engine", "bayesian",
			"--targeting.condition", "condition",
			"--targeting.coverage", "0",
			"--targeting.prerequisites", "[{id: id, condition: condition}]",
			"--targeting.saved-groups", "[{ids: [string], match: all}]",
			"--template-metadata.name", "name",
			"--template-metadata.description", "description",
			"--type", "standard",
			"--activation-metric", "activationMetric",
			"--custom-fields", "{foo: string}",
			"--custom-metric-slice.slices", "[{column: column, levels: [string]}]",
			"--description", "description",
			"--disable-sticky-bucketing=true",
			"--fallback-attribute", "fallbackAttribute",
			"--goal-metric", "string",
			"--guardrail-metric", "string",
			"--hash-attribute", "hashAttribute",
			"--hypothesis", "hypothesis",
			"--project", "project",
			"--secondary-metric", "string",
			"--segment", "segment",
			"--skip-partial-data=true",
			"--tag", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"datasource: datasource\n" +
			"exposureQueryId: exposureQueryId\n" +
			"statsEngine: bayesian\n" +
			"targeting:\n" +
			"  condition: condition\n" +
			"  coverage: 0\n" +
			"  prerequisites:\n" +
			"    - id: id\n" +
			"      condition: condition\n" +
			"  savedGroups:\n" +
			"    - ids:\n" +
			"        - string\n" +
			"      match: all\n" +
			"templateMetadata:\n" +
			"  name: name\n" +
			"  description: description\n" +
			"type: standard\n" +
			"activationMetric: activationMetric\n" +
			"customFields:\n" +
			"  foo: string\n" +
			"customMetricSlices:\n" +
			"  - slices:\n" +
			"      - column: column\n" +
			"        levels:\n" +
			"          - string\n" +
			"description: description\n" +
			"disableStickyBucketing: true\n" +
			"fallbackAttribute: fallbackAttribute\n" +
			"goalMetrics:\n" +
			"  - string\n" +
			"guardrailMetrics:\n" +
			"  - string\n" +
			"hashAttribute: hashAttribute\n" +
			"hypothesis: hypothesis\n" +
			"project: project\n" +
			"secondaryMetrics:\n" +
			"  - string\n" +
			"segment: segment\n" +
			"skipPartialData: true\n" +
			"tags:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiment-templates", "create",
		)
	})
}

func TestExperimentTemplatesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-templates", "retrieve",
			"--id", "id",
		)
	})
}

func TestExperimentTemplatesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-templates", "update",
			"--id", "id",
			"--activation-metric", "activationMetric",
			"--custom-fields", "{foo: string}",
			"--custom-metric-slice", "{slices: [{column: column, levels: [string]}]}",
			"--datasource", "datasource",
			"--description", "description",
			"--disable-sticky-bucketing=true",
			"--exposure-query-id", "exposureQueryId",
			"--fallback-attribute", "fallbackAttribute",
			"--goal-metric", "string",
			"--guardrail-metric", "string",
			"--hash-attribute", "hashAttribute",
			"--hypothesis", "hypothesis",
			"--project", "project",
			"--secondary-metric", "string",
			"--segment", "segment",
			"--skip-partial-data=true",
			"--stats-engine", "bayesian",
			"--tag", "string",
			"--targeting", "{condition: condition, coverage: 0, prerequisites: [{id: id, condition: condition}], savedGroups: [{ids: [string], match: all}]}",
			"--template-metadata", "{name: name, description: description}",
			"--type", "standard",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(experimentTemplatesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-templates", "update",
			"--id", "id",
			"--activation-metric", "activationMetric",
			"--custom-fields", "{foo: string}",
			"--custom-metric-slice.slices", "[{column: column, levels: [string]}]",
			"--datasource", "datasource",
			"--description", "description",
			"--disable-sticky-bucketing=true",
			"--exposure-query-id", "exposureQueryId",
			"--fallback-attribute", "fallbackAttribute",
			"--goal-metric", "string",
			"--guardrail-metric", "string",
			"--hash-attribute", "hashAttribute",
			"--hypothesis", "hypothesis",
			"--project", "project",
			"--secondary-metric", "string",
			"--segment", "segment",
			"--skip-partial-data=true",
			"--stats-engine", "bayesian",
			"--tag", "string",
			"--targeting.condition", "condition",
			"--targeting.coverage", "0",
			"--targeting.prerequisites", "[{id: id, condition: condition}]",
			"--targeting.saved-groups", "[{ids: [string], match: all}]",
			"--template-metadata.name", "name",
			"--template-metadata.description", "description",
			"--type", "standard",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"activationMetric: activationMetric\n" +
			"customFields:\n" +
			"  foo: string\n" +
			"customMetricSlices:\n" +
			"  - slices:\n" +
			"      - column: column\n" +
			"        levels:\n" +
			"          - string\n" +
			"datasource: datasource\n" +
			"description: description\n" +
			"disableStickyBucketing: true\n" +
			"exposureQueryId: exposureQueryId\n" +
			"fallbackAttribute: fallbackAttribute\n" +
			"goalMetrics:\n" +
			"  - string\n" +
			"guardrailMetrics:\n" +
			"  - string\n" +
			"hashAttribute: hashAttribute\n" +
			"hypothesis: hypothesis\n" +
			"project: project\n" +
			"secondaryMetrics:\n" +
			"  - string\n" +
			"segment: segment\n" +
			"skipPartialData: true\n" +
			"statsEngine: bayesian\n" +
			"tags:\n" +
			"  - string\n" +
			"targeting:\n" +
			"  condition: condition\n" +
			"  coverage: 0\n" +
			"  prerequisites:\n" +
			"    - id: id\n" +
			"      condition: condition\n" +
			"  savedGroups:\n" +
			"    - ids:\n" +
			"        - string\n" +
			"      match: all\n" +
			"templateMetadata:\n" +
			"  name: name\n" +
			"  description: description\n" +
			"type: standard\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiment-templates", "update",
			"--id", "id",
		)
	})
}

func TestExperimentTemplatesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-templates", "list",
			"--project-id", "projectId",
		)
	})
}

func TestExperimentTemplatesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-templates", "delete",
			"--id", "id",
		)
	})
}

func TestExperimentTemplatesBulkImport(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-templates", "bulk-import",
			"--template", "{id: id, data: {datasource: datasource, exposureQueryId: exposureQueryId, statsEngine: bayesian, targeting: {condition: condition, coverage: 0, prerequisites: [{id: id, condition: condition}], savedGroups: [{ids: [string], match: all}]}, templateMetadata: {name: name, description: description}, type: standard, activationMetric: activationMetric, customFields: {foo: string}, customMetricSlices: [{slices: [{column: column, levels: [string]}]}], description: description, disableStickyBucketing: true, fallbackAttribute: fallbackAttribute, goalMetrics: [string], guardrailMetrics: [string], hashAttribute: hashAttribute, hypothesis: hypothesis, project: project, secondaryMetrics: [string], segment: segment, skipPartialData: true, tags: [string]}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(experimentTemplatesBulkImport)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiment-templates", "bulk-import",
			"--template.id", "id",
			"--template.data", "{datasource: datasource, exposureQueryId: exposureQueryId, statsEngine: bayesian, targeting: {condition: condition, coverage: 0, prerequisites: [{id: id, condition: condition}], savedGroups: [{ids: [string], match: all}]}, templateMetadata: {name: name, description: description}, type: standard, activationMetric: activationMetric, customFields: {foo: string}, customMetricSlices: [{slices: [{column: column, levels: [string]}]}], description: description, disableStickyBucketing: true, fallbackAttribute: fallbackAttribute, goalMetrics: [string], guardrailMetrics: [string], hashAttribute: hashAttribute, hypothesis: hypothesis, project: project, secondaryMetrics: [string], segment: segment, skipPartialData: true, tags: [string]}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"templates:\n" +
			"  - id: id\n" +
			"    data:\n" +
			"      datasource: datasource\n" +
			"      exposureQueryId: exposureQueryId\n" +
			"      statsEngine: bayesian\n" +
			"      targeting:\n" +
			"        condition: condition\n" +
			"        coverage: 0\n" +
			"        prerequisites:\n" +
			"          - id: id\n" +
			"            condition: condition\n" +
			"        savedGroups:\n" +
			"          - ids:\n" +
			"              - string\n" +
			"            match: all\n" +
			"      templateMetadata:\n" +
			"        name: name\n" +
			"        description: description\n" +
			"      type: standard\n" +
			"      activationMetric: activationMetric\n" +
			"      customFields:\n" +
			"        foo: string\n" +
			"      customMetricSlices:\n" +
			"        - slices:\n" +
			"            - column: column\n" +
			"              levels:\n" +
			"                - string\n" +
			"      description: description\n" +
			"      disableStickyBucketing: true\n" +
			"      fallbackAttribute: fallbackAttribute\n" +
			"      goalMetrics:\n" +
			"        - string\n" +
			"      guardrailMetrics:\n" +
			"        - string\n" +
			"      hashAttribute: hashAttribute\n" +
			"      hypothesis: hypothesis\n" +
			"      project: project\n" +
			"      secondaryMetrics:\n" +
			"        - string\n" +
			"      segment: segment\n" +
			"      skipPartialData: true\n" +
			"      tags:\n" +
			"        - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiment-templates", "bulk-import",
		)
	})
}
