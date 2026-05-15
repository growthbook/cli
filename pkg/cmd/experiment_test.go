// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/growthbook-cli/internal/mocktest"
	"github.com/stainless-sdks/growthbook-cli/internal/requestflag"
)

func TestExperimentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "create",
			"--name", "name",
			"--tracking-key", "trackingKey",
			"--variation", "{key: key, name: name, id: id, description: description, screenshots: [{path: path, description: description, height: 0, width: 0}]}",
			"--variation", "{key: key, name: name, id: id, description: description, screenshots: [{path: path, description: description, height: 0, width: 0}]}",
			"--activation-metric", "activationMetric",
			"--archived=true",
			"--assignment-query-id", "assignmentQueryId",
			"--attribution-model", "firstExposure",
			"--auto-refresh=true",
			"--bandit-burn-in-unit", "days",
			"--bandit-burn-in-value", "0",
			"--bandit-conversion-window-unit", "days",
			"--bandit-conversion-window-value", "0",
			"--bandit-schedule-unit", "days",
			"--bandit-schedule-value", "0",
			"--bucket-version", "0",
			"--bypass-duplicate-key-check=true",
			"--custom-fields", "{foo: string}",
			"--custom-metric-slice", "{slices: [{column: column, levels: [string]}]}",
			"--datasource-id", "datasourceId",
			"--decision-framework-settings", "{decisionCriteriaId: decisionCriteriaId, decisionFrameworkMetricOverrides: [{id: id, targetMDE: 1}]}",
			"--default-dashboard-id", "defaultDashboardId",
			"--description", "description",
			"--disable-sticky-bucketing=true",
			"--exclude-from-payload=true",
			"--fallback-attribute", "fallbackAttribute",
			"--guardrail-metric", "string",
			"--hash-attribute", "hashAttribute",
			"--hash-version", "1",
			"--hypothesis", "hypothesis",
			"--in-progress-conversions", "loose",
			"--lookback-override", "{type: date, value: 0, valueUnit: minutes}",
			"--metric-override", "{id: id, delayHours: 0, properPriorEnabled: true, properPriorMean: 0, properPriorOverride: true, properPriorStdDev: 0, regressionAdjustmentDays: 0, regressionAdjustmentEnabled: true, regressionAdjustmentOverride: true, windowHours: 0, windowType: conversion}",
			"--metric", "string",
			"--min-bucket-version", "0",
			"--owner", "owner",
			"--phase", "{dateStarted: '2019-12-27T18:11:19.117Z', name: name, condition: condition, coverage: 0, dateEnded: '2019-12-27T18:11:19.117Z', namespace: {namespaceId: namespaceId, enabled: true, range: [0, 0], ranges: [[{}]]}, prerequisites: [{id: id, condition: condition}], reason: reason, reasonForStopping: reasonForStopping, savedGroupTargeting: [{matchType: all, savedGroups: [string]}], seed: seed, targetingCondition: targetingCondition, trafficSplit: [{variationId: variationId, weight: 0}], variationWeights: [0]}",
			"--post-stratification-enabled=true",
			"--project", "project",
			"--query-filter", "queryFilter",
			"--regression-adjustment-enabled=true",
			"--released-variation-id", "releasedVariationId",
			"--secondary-metric", "string",
			"--segment-id", "segmentId",
			"--sequential-testing-enabled=true",
			"--sequential-testing-tuning-parameter", "0",
			"--share-level", "public",
			"--stats-engine", "bayesian",
			"--status", "draft",
			"--tag", "string",
			"--template-id", "templateId",
			"--type", "standard",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(experimentsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "create",
			"--name", "name",
			"--tracking-key", "trackingKey",
			"--variation.key", "key",
			"--variation.name", "name",
			"--variation.id", "id",
			"--variation.description", "description",
			"--variation.screenshots", "[{path: path, description: description, height: 0, width: 0}]",
			"--variation.key", "key",
			"--variation.name", "name",
			"--variation.id", "id",
			"--variation.description", "description",
			"--variation.screenshots", "[{path: path, description: description, height: 0, width: 0}]",
			"--activation-metric", "activationMetric",
			"--archived=true",
			"--assignment-query-id", "assignmentQueryId",
			"--attribution-model", "firstExposure",
			"--auto-refresh=true",
			"--bandit-burn-in-unit", "days",
			"--bandit-burn-in-value", "0",
			"--bandit-conversion-window-unit", "days",
			"--bandit-conversion-window-value", "0",
			"--bandit-schedule-unit", "days",
			"--bandit-schedule-value", "0",
			"--bucket-version", "0",
			"--bypass-duplicate-key-check=true",
			"--custom-fields", "{foo: string}",
			"--custom-metric-slice.slices", "[{column: column, levels: [string]}]",
			"--datasource-id", "datasourceId",
			"--decision-framework-settings.decision-criteria-id", "decisionCriteriaId",
			"--decision-framework-settings.decision-framework-metric-overrides", "[{id: id, targetMDE: 1}]",
			"--default-dashboard-id", "defaultDashboardId",
			"--description", "description",
			"--disable-sticky-bucketing=true",
			"--exclude-from-payload=true",
			"--fallback-attribute", "fallbackAttribute",
			"--guardrail-metric", "string",
			"--hash-attribute", "hashAttribute",
			"--hash-version", "1",
			"--hypothesis", "hypothesis",
			"--in-progress-conversions", "loose",
			"--lookback-override.type", "date",
			"--lookback-override.value", "0",
			"--lookback-override.value-unit", "minutes",
			"--metric-override.id", "id",
			"--metric-override.delay-hours", "0",
			"--metric-override.proper-prior-enabled=true",
			"--metric-override.proper-prior-mean", "0",
			"--metric-override.proper-prior-override=true",
			"--metric-override.proper-prior-std-dev", "0",
			"--metric-override.regression-adjustment-days", "0",
			"--metric-override.regression-adjustment-enabled=true",
			"--metric-override.regression-adjustment-override=true",
			"--metric-override.window-hours", "0",
			"--metric-override.window-type", "conversion",
			"--metric", "string",
			"--min-bucket-version", "0",
			"--owner", "owner",
			"--phase.date-started", "2019-12-27T18:11:19.117Z",
			"--phase.name", "name",
			"--phase.condition", "condition",
			"--phase.coverage", "0",
			"--phase.date-ended", "2019-12-27T18:11:19.117Z",
			"--phase.namespace", "{namespaceId: namespaceId, enabled: true, range: [0, 0], ranges: [[{}]]}",
			"--phase.prerequisites", "[{id: id, condition: condition}]",
			"--phase.reason", "reason",
			"--phase.reason-for-stopping", "reasonForStopping",
			"--phase.saved-group-targeting", "[{matchType: all, savedGroups: [string]}]",
			"--phase.seed", "seed",
			"--phase.targeting-condition", "targetingCondition",
			"--phase.traffic-split", "[{variationId: variationId, weight: 0}]",
			"--phase.variation-weights", "[0]",
			"--post-stratification-enabled=true",
			"--project", "project",
			"--query-filter", "queryFilter",
			"--regression-adjustment-enabled=true",
			"--released-variation-id", "releasedVariationId",
			"--secondary-metric", "string",
			"--segment-id", "segmentId",
			"--sequential-testing-enabled=true",
			"--sequential-testing-tuning-parameter", "0",
			"--share-level", "public",
			"--stats-engine", "bayesian",
			"--status", "draft",
			"--tag", "string",
			"--template-id", "templateId",
			"--type", "standard",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"trackingKey: trackingKey\n" +
			"variations:\n" +
			"  - key: key\n" +
			"    name: name\n" +
			"    id: id\n" +
			"    description: description\n" +
			"    screenshots:\n" +
			"      - path: path\n" +
			"        description: description\n" +
			"        height: 0\n" +
			"        width: 0\n" +
			"  - key: key\n" +
			"    name: name\n" +
			"    id: id\n" +
			"    description: description\n" +
			"    screenshots:\n" +
			"      - path: path\n" +
			"        description: description\n" +
			"        height: 0\n" +
			"        width: 0\n" +
			"activationMetric: activationMetric\n" +
			"archived: true\n" +
			"assignmentQueryId: assignmentQueryId\n" +
			"attributionModel: firstExposure\n" +
			"autoRefresh: true\n" +
			"banditBurnInUnit: days\n" +
			"banditBurnInValue: 0\n" +
			"banditConversionWindowUnit: days\n" +
			"banditConversionWindowValue: 0\n" +
			"banditScheduleUnit: days\n" +
			"banditScheduleValue: 0\n" +
			"bucketVersion: 0\n" +
			"bypassDuplicateKeyCheck: true\n" +
			"customFields:\n" +
			"  foo: string\n" +
			"customMetricSlices:\n" +
			"  - slices:\n" +
			"      - column: column\n" +
			"        levels:\n" +
			"          - string\n" +
			"datasourceId: datasourceId\n" +
			"decisionFrameworkSettings:\n" +
			"  decisionCriteriaId: decisionCriteriaId\n" +
			"  decisionFrameworkMetricOverrides:\n" +
			"    - id: id\n" +
			"      targetMDE: 1\n" +
			"defaultDashboardId: defaultDashboardId\n" +
			"description: description\n" +
			"disableStickyBucketing: true\n" +
			"excludeFromPayload: true\n" +
			"fallbackAttribute: fallbackAttribute\n" +
			"guardrailMetrics:\n" +
			"  - string\n" +
			"hashAttribute: hashAttribute\n" +
			"hashVersion: 1\n" +
			"hypothesis: hypothesis\n" +
			"inProgressConversions: loose\n" +
			"lookbackOverride:\n" +
			"  type: date\n" +
			"  value: 0\n" +
			"  valueUnit: minutes\n" +
			"metricOverrides:\n" +
			"  - id: id\n" +
			"    delayHours: 0\n" +
			"    properPriorEnabled: true\n" +
			"    properPriorMean: 0\n" +
			"    properPriorOverride: true\n" +
			"    properPriorStdDev: 0\n" +
			"    regressionAdjustmentDays: 0\n" +
			"    regressionAdjustmentEnabled: true\n" +
			"    regressionAdjustmentOverride: true\n" +
			"    windowHours: 0\n" +
			"    windowType: conversion\n" +
			"metrics:\n" +
			"  - string\n" +
			"minBucketVersion: 0\n" +
			"owner: owner\n" +
			"phases:\n" +
			"  - dateStarted: '2019-12-27T18:11:19.117Z'\n" +
			"    name: name\n" +
			"    condition: condition\n" +
			"    coverage: 0\n" +
			"    dateEnded: '2019-12-27T18:11:19.117Z'\n" +
			"    namespace:\n" +
			"      namespaceId: namespaceId\n" +
			"      enabled: true\n" +
			"      range:\n" +
			"        - 0\n" +
			"        - 0\n" +
			"      ranges:\n" +
			"        - - {}\n" +
			"    prerequisites:\n" +
			"      - id: id\n" +
			"        condition: condition\n" +
			"    reason: reason\n" +
			"    reasonForStopping: reasonForStopping\n" +
			"    savedGroupTargeting:\n" +
			"      - matchType: all\n" +
			"        savedGroups:\n" +
			"          - string\n" +
			"    seed: seed\n" +
			"    targetingCondition: targetingCondition\n" +
			"    trafficSplit:\n" +
			"      - variationId: variationId\n" +
			"        weight: 0\n" +
			"    variationWeights:\n" +
			"      - 0\n" +
			"postStratificationEnabled: true\n" +
			"project: project\n" +
			"queryFilter: queryFilter\n" +
			"regressionAdjustmentEnabled: true\n" +
			"releasedVariationId: releasedVariationId\n" +
			"secondaryMetrics:\n" +
			"  - string\n" +
			"segmentId: segmentId\n" +
			"sequentialTestingEnabled: true\n" +
			"sequentialTestingTuningParameter: 0\n" +
			"shareLevel: public\n" +
			"statsEngine: bayesian\n" +
			"status: draft\n" +
			"tags:\n" +
			"  - string\n" +
			"templateId: templateId\n" +
			"type: standard\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiments", "create",
		)
	})
}

func TestExperimentsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "retrieve",
			"--id", "id",
		)
	})
}

func TestExperimentsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "update",
			"--id", "id",
			"--activation-metric", "activationMetric",
			"--analysis", "analysis",
			"--archived=true",
			"--assignment-query-id", "assignmentQueryId",
			"--attribution-model", "firstExposure",
			"--auto-refresh=true",
			"--bandit-burn-in-unit", "days",
			"--bandit-burn-in-value", "0",
			"--bandit-conversion-window-unit", "days",
			"--bandit-conversion-window-value", "0",
			"--bandit-schedule-unit", "days",
			"--bandit-schedule-value", "0",
			"--bucket-version", "0",
			"--bypass-duplicate-key-check=true",
			"--custom-fields", "{foo: string}",
			"--custom-metric-slice", "{slices: [{column: column, levels: [string]}]}",
			"--datasource-id", "datasourceId",
			"--decision-framework-settings", "{decisionCriteriaId: decisionCriteriaId, decisionFrameworkMetricOverrides: [{id: id, targetMDE: 1}]}",
			"--default-dashboard-id", "defaultDashboardId",
			"--description", "description",
			"--disable-sticky-bucketing=true",
			"--exclude-from-payload=true",
			"--fallback-attribute", "fallbackAttribute",
			"--guardrail-metric", "string",
			"--hash-attribute", "hashAttribute",
			"--hash-version", "1",
			"--hypothesis", "hypothesis",
			"--in-progress-conversions", "loose",
			"--lookback-override", "{type: date, value: 0, valueUnit: minutes}",
			"--metric-override", "{id: id, delayHours: 0, properPriorEnabled: true, properPriorMean: 0, properPriorOverride: true, properPriorStdDev: 0, regressionAdjustmentDays: 0, regressionAdjustmentEnabled: true, regressionAdjustmentOverride: true, windowHours: 0, windowType: conversion}",
			"--metric", "string",
			"--min-bucket-version", "0",
			"--name", "name",
			"--owner", "owner",
			"--phase", "{dateStarted: '2019-12-27T18:11:19.117Z', name: name, condition: condition, coverage: 0, dateEnded: '2019-12-27T18:11:19.117Z', namespace: {namespaceId: namespaceId, range: [0, 0], enabled: true}, prerequisites: [{id: id, condition: condition}], reason: reason, reasonForStopping: reasonForStopping, savedGroupTargeting: [{matchType: all, savedGroups: [string]}], seed: seed, targetingCondition: targetingCondition, trafficSplit: [{variationId: variationId, weight: 0}], variationWeights: [0]}",
			"--post-stratification-enabled=true",
			"--project", "project",
			"--query-filter", "queryFilter",
			"--regression-adjustment-enabled=true",
			"--released-variation-id", "releasedVariationId",
			"--results", "dnf",
			"--secondary-metric", "string",
			"--segment-id", "segmentId",
			"--sequential-testing-enabled=true",
			"--sequential-testing-tuning-parameter", "0",
			"--share-level", "public",
			"--stats-engine", "bayesian",
			"--status", "draft",
			"--tag", "string",
			"--tracking-key", "trackingKey",
			"--type", "standard",
			"--variation", "{key: key, name: name, id: id, description: description, screenshots: [{path: path, description: description, height: 0, width: 0}]}",
			"--variation", "{key: key, name: name, id: id, description: description, screenshots: [{path: path, description: description, height: 0, width: 0}]}",
			"--winner", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(experimentsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "update",
			"--id", "id",
			"--activation-metric", "activationMetric",
			"--analysis", "analysis",
			"--archived=true",
			"--assignment-query-id", "assignmentQueryId",
			"--attribution-model", "firstExposure",
			"--auto-refresh=true",
			"--bandit-burn-in-unit", "days",
			"--bandit-burn-in-value", "0",
			"--bandit-conversion-window-unit", "days",
			"--bandit-conversion-window-value", "0",
			"--bandit-schedule-unit", "days",
			"--bandit-schedule-value", "0",
			"--bucket-version", "0",
			"--bypass-duplicate-key-check=true",
			"--custom-fields", "{foo: string}",
			"--custom-metric-slice.slices", "[{column: column, levels: [string]}]",
			"--datasource-id", "datasourceId",
			"--decision-framework-settings.decision-criteria-id", "decisionCriteriaId",
			"--decision-framework-settings.decision-framework-metric-overrides", "[{id: id, targetMDE: 1}]",
			"--default-dashboard-id", "defaultDashboardId",
			"--description", "description",
			"--disable-sticky-bucketing=true",
			"--exclude-from-payload=true",
			"--fallback-attribute", "fallbackAttribute",
			"--guardrail-metric", "string",
			"--hash-attribute", "hashAttribute",
			"--hash-version", "1",
			"--hypothesis", "hypothesis",
			"--in-progress-conversions", "loose",
			"--lookback-override.type", "date",
			"--lookback-override.value", "0",
			"--lookback-override.value-unit", "minutes",
			"--metric-override.id", "id",
			"--metric-override.delay-hours", "0",
			"--metric-override.proper-prior-enabled=true",
			"--metric-override.proper-prior-mean", "0",
			"--metric-override.proper-prior-override=true",
			"--metric-override.proper-prior-std-dev", "0",
			"--metric-override.regression-adjustment-days", "0",
			"--metric-override.regression-adjustment-enabled=true",
			"--metric-override.regression-adjustment-override=true",
			"--metric-override.window-hours", "0",
			"--metric-override.window-type", "conversion",
			"--metric", "string",
			"--min-bucket-version", "0",
			"--name", "name",
			"--owner", "owner",
			"--phase.date-started", "2019-12-27T18:11:19.117Z",
			"--phase.name", "name",
			"--phase.condition", "condition",
			"--phase.coverage", "0",
			"--phase.date-ended", "2019-12-27T18:11:19.117Z",
			"--phase.namespace", "{namespaceId: namespaceId, range: [0, 0], enabled: true}",
			"--phase.prerequisites", "[{id: id, condition: condition}]",
			"--phase.reason", "reason",
			"--phase.reason-for-stopping", "reasonForStopping",
			"--phase.saved-group-targeting", "[{matchType: all, savedGroups: [string]}]",
			"--phase.seed", "seed",
			"--phase.targeting-condition", "targetingCondition",
			"--phase.traffic-split", "[{variationId: variationId, weight: 0}]",
			"--phase.variation-weights", "[0]",
			"--post-stratification-enabled=true",
			"--project", "project",
			"--query-filter", "queryFilter",
			"--regression-adjustment-enabled=true",
			"--released-variation-id", "releasedVariationId",
			"--results", "dnf",
			"--secondary-metric", "string",
			"--segment-id", "segmentId",
			"--sequential-testing-enabled=true",
			"--sequential-testing-tuning-parameter", "0",
			"--share-level", "public",
			"--stats-engine", "bayesian",
			"--status", "draft",
			"--tag", "string",
			"--tracking-key", "trackingKey",
			"--type", "standard",
			"--variation.key", "key",
			"--variation.name", "name",
			"--variation.id", "id",
			"--variation.description", "description",
			"--variation.screenshots", "[{path: path, description: description, height: 0, width: 0}]",
			"--variation.key", "key",
			"--variation.name", "name",
			"--variation.id", "id",
			"--variation.description", "description",
			"--variation.screenshots", "[{path: path, description: description, height: 0, width: 0}]",
			"--winner", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"activationMetric: activationMetric\n" +
			"analysis: analysis\n" +
			"archived: true\n" +
			"assignmentQueryId: assignmentQueryId\n" +
			"attributionModel: firstExposure\n" +
			"autoRefresh: true\n" +
			"banditBurnInUnit: days\n" +
			"banditBurnInValue: 0\n" +
			"banditConversionWindowUnit: days\n" +
			"banditConversionWindowValue: 0\n" +
			"banditScheduleUnit: days\n" +
			"banditScheduleValue: 0\n" +
			"bucketVersion: 0\n" +
			"bypassDuplicateKeyCheck: true\n" +
			"customFields:\n" +
			"  foo: string\n" +
			"customMetricSlices:\n" +
			"  - slices:\n" +
			"      - column: column\n" +
			"        levels:\n" +
			"          - string\n" +
			"datasourceId: datasourceId\n" +
			"decisionFrameworkSettings:\n" +
			"  decisionCriteriaId: decisionCriteriaId\n" +
			"  decisionFrameworkMetricOverrides:\n" +
			"    - id: id\n" +
			"      targetMDE: 1\n" +
			"defaultDashboardId: defaultDashboardId\n" +
			"description: description\n" +
			"disableStickyBucketing: true\n" +
			"excludeFromPayload: true\n" +
			"fallbackAttribute: fallbackAttribute\n" +
			"guardrailMetrics:\n" +
			"  - string\n" +
			"hashAttribute: hashAttribute\n" +
			"hashVersion: 1\n" +
			"hypothesis: hypothesis\n" +
			"inProgressConversions: loose\n" +
			"lookbackOverride:\n" +
			"  type: date\n" +
			"  value: 0\n" +
			"  valueUnit: minutes\n" +
			"metricOverrides:\n" +
			"  - id: id\n" +
			"    delayHours: 0\n" +
			"    properPriorEnabled: true\n" +
			"    properPriorMean: 0\n" +
			"    properPriorOverride: true\n" +
			"    properPriorStdDev: 0\n" +
			"    regressionAdjustmentDays: 0\n" +
			"    regressionAdjustmentEnabled: true\n" +
			"    regressionAdjustmentOverride: true\n" +
			"    windowHours: 0\n" +
			"    windowType: conversion\n" +
			"metrics:\n" +
			"  - string\n" +
			"minBucketVersion: 0\n" +
			"name: name\n" +
			"owner: owner\n" +
			"phases:\n" +
			"  - dateStarted: '2019-12-27T18:11:19.117Z'\n" +
			"    name: name\n" +
			"    condition: condition\n" +
			"    coverage: 0\n" +
			"    dateEnded: '2019-12-27T18:11:19.117Z'\n" +
			"    namespace:\n" +
			"      namespaceId: namespaceId\n" +
			"      range:\n" +
			"        - 0\n" +
			"        - 0\n" +
			"      enabled: true\n" +
			"    prerequisites:\n" +
			"      - id: id\n" +
			"        condition: condition\n" +
			"    reason: reason\n" +
			"    reasonForStopping: reasonForStopping\n" +
			"    savedGroupTargeting:\n" +
			"      - matchType: all\n" +
			"        savedGroups:\n" +
			"          - string\n" +
			"    seed: seed\n" +
			"    targetingCondition: targetingCondition\n" +
			"    trafficSplit:\n" +
			"      - variationId: variationId\n" +
			"        weight: 0\n" +
			"    variationWeights:\n" +
			"      - 0\n" +
			"postStratificationEnabled: true\n" +
			"project: project\n" +
			"queryFilter: queryFilter\n" +
			"regressionAdjustmentEnabled: true\n" +
			"releasedVariationId: releasedVariationId\n" +
			"results: dnf\n" +
			"secondaryMetrics:\n" +
			"  - string\n" +
			"segmentId: segmentId\n" +
			"sequentialTestingEnabled: true\n" +
			"sequentialTestingTuningParameter: 0\n" +
			"shareLevel: public\n" +
			"statsEngine: bayesian\n" +
			"status: draft\n" +
			"tags:\n" +
			"  - string\n" +
			"trackingKey: trackingKey\n" +
			"type: standard\n" +
			"variations:\n" +
			"  - key: key\n" +
			"    name: name\n" +
			"    id: id\n" +
			"    description: description\n" +
			"    screenshots:\n" +
			"      - path: path\n" +
			"        description: description\n" +
			"        height: 0\n" +
			"        width: 0\n" +
			"  - key: key\n" +
			"    name: name\n" +
			"    id: id\n" +
			"    description: description\n" +
			"    screenshots:\n" +
			"      - path: path\n" +
			"        description: description\n" +
			"        height: 0\n" +
			"        width: 0\n" +
			"winner: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiments", "update",
			"--id", "id",
		)
	})
}

func TestExperimentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "list",
			"--datasource-id", "datasourceId",
			"--experiment-id", "experimentId",
			"--limit", "1",
			"--offset", "0",
			"--project-id", "projectId",
			"--status", "draft",
			"--tracking-key", "trackingKey",
		)
	})
}

func TestExperimentsCreateSnapshot(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "create-snapshot",
			"--id", "id",
			"--dimension", "dimension",
			"--phase", "0",
			"--triggered-by", "manual",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"dimension: dimension\n" +
			"phase: 0\n" +
			"triggeredBy: manual\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiments", "create-snapshot",
			"--id", "id",
		)
	})
}

func TestExperimentsModifyTemporaryRollout(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "modify-temporary-rollout",
			"--id", "id",
			"--enable-temporary-rollout=true",
			"--released-variation-id", "releasedVariationId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"enableTemporaryRollout: true\n" +
			"releasedVariationId: releasedVariationId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiments", "modify-temporary-rollout",
			"--id", "id",
		)
	})
}

func TestExperimentsStart(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "start",
			"--id", "id",
			"--skip-checklist=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("skipChecklist: true")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiments", "start",
			"--id", "id",
		)
	})
}

func TestExperimentsStop(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"experiments", "stop",
			"--id", "id",
			"--results", "dnf",
			"--analysis", "analysis",
			"--date-ended", "dateEnded",
			"--enable-temporary-rollout=true",
			"--reason", "reason",
			"--released-variation-id", "releasedVariationId",
			"--winner-variation-id", "winnerVariationId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"results: dnf\n" +
			"analysis: analysis\n" +
			"dateEnded: dateEnded\n" +
			"enableTemporaryRollout: true\n" +
			"reason: reason\n" +
			"releasedVariationId: releasedVariationId\n" +
			"winnerVariationId: winnerVariationId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"experiments", "stop",
			"--id", "id",
		)
	})
}
