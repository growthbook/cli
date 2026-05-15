// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/growthbook-cli/internal/apiquery"
	"github.com/stainless-sdks/growthbook-cli/internal/requestflag"
	"github.com/stainless-sdks/growthbook-go"
	"github.com/stainless-sdks/growthbook-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var experimentsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a single experiment",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the experiment",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "tracking-key",
			Required: true,
			BodyPath: "trackingKey",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "variation",
			Required: true,
			BodyPath: "variations",
		},
		&requestflag.Flag[string]{
			Name:     "activation-metric",
			Usage:    "Users must convert on this metric before being included",
			BodyPath: "activationMetric",
		},
		&requestflag.Flag[bool]{
			Name:     "archived",
			BodyPath: "archived",
		},
		&requestflag.Flag[string]{
			Name:     "assignment-query-id",
			Usage:    "The ID property of one of the assignment query objects associated with the datasource. Can only be set if a templateId is not provided.",
			BodyPath: "assignmentQueryId",
		},
		&requestflag.Flag[string]{
			Name:     "attribution-model",
			Usage:    "Setting attribution model to `\"experimentDuration\"` is the same as selecting \"Ignore Conversion Windows\" for the Conversion Window Override. Setting it to `\"lookbackOverride\"` requires a `lookbackOverride` object to be provided.",
			BodyPath: "attributionModel",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-refresh",
			BodyPath: "autoRefresh",
		},
		&requestflag.Flag[string]{
			Name:     "bandit-burn-in-unit",
			Usage:    `Allowed values: "days", "hours".`,
			BodyPath: "banditBurnInUnit",
		},
		&requestflag.Flag[float64]{
			Name:     "bandit-burn-in-value",
			BodyPath: "banditBurnInValue",
		},
		&requestflag.Flag[string]{
			Name:     "bandit-conversion-window-unit",
			Usage:    `Allowed values: "days", "hours".`,
			BodyPath: "banditConversionWindowUnit",
		},
		&requestflag.Flag[float64]{
			Name:     "bandit-conversion-window-value",
			BodyPath: "banditConversionWindowValue",
		},
		&requestflag.Flag[string]{
			Name:     "bandit-schedule-unit",
			Usage:    `Allowed values: "days", "hours".`,
			BodyPath: "banditScheduleUnit",
		},
		&requestflag.Flag[float64]{
			Name:     "bandit-schedule-value",
			BodyPath: "banditScheduleValue",
		},
		&requestflag.Flag[float64]{
			Name:     "bucket-version",
			BodyPath: "bucketVersion",
		},
		&requestflag.Flag[bool]{
			Name:     "bypass-duplicate-key-check",
			Usage:    "If true, allow creating an experiment even if another experiment with the same tracking key already exists. This is ignored if the organization requires unique tracking keys as a rule.",
			BodyPath: "bypassDuplicateKeyCheck",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "custom-fields",
			BodyPath: "customFields",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "custom-metric-slice",
			Usage:    "Custom slices that apply to ALL applicable metrics in the experiment",
			BodyPath: "customMetricSlices",
		},
		&requestflag.Flag[string]{
			Name:     "datasource-id",
			Usage:    "ID for the [DataSource](#tag/DataSource_model). Can only be set if a templateId is not provided.",
			BodyPath: "datasourceId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "decision-framework-settings",
			Usage:    "Controls the decision framework and metric overrides for the experiment. Replaces the entire stored object on update (does not patch individual fields).",
			BodyPath: "decisionFrameworkSettings",
		},
		&requestflag.Flag[string]{
			Name:     "default-dashboard-id",
			Usage:    "ID of the default dashboard for this experiment.",
			BodyPath: "defaultDashboardId",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of the experiment",
			BodyPath: "description",
		},
		&requestflag.Flag[bool]{
			Name:     "disable-sticky-bucketing",
			BodyPath: "disableStickyBucketing",
		},
		&requestflag.Flag[bool]{
			Name:     "exclude-from-payload",
			BodyPath: "excludeFromPayload",
		},
		&requestflag.Flag[string]{
			Name:     "fallback-attribute",
			BodyPath: "fallbackAttribute",
		},
		&requestflag.Flag[[]string]{
			Name:     "guardrail-metric",
			BodyPath: "guardrailMetrics",
		},
		&requestflag.Flag[string]{
			Name:     "hash-attribute",
			BodyPath: "hashAttribute",
		},
		&requestflag.Flag[float64]{
			Name:     "hash-version",
			Usage:    "Allowed values: 1, 2.",
			BodyPath: "hashVersion",
		},
		&requestflag.Flag[string]{
			Name:     "hypothesis",
			Usage:    "Hypothesis of the experiment",
			BodyPath: "hypothesis",
		},
		&requestflag.Flag[string]{
			Name:     "in-progress-conversions",
			Usage:    `Allowed values: "loose", "strict".`,
			BodyPath: "inProgressConversions",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "lookback-override",
			Usage:    `Controls the lookback override for the experiment. For type "window", value must be a non-negative number and valueUnit is required.`,
			BodyPath: "lookbackOverride",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "metric-override",
			Usage:    "Per-metric analysis overrides for this experiment. Replaces the entire stored array (does not patch individual entries).",
			BodyPath: "metricOverrides",
		},
		&requestflag.Flag[[]string]{
			Name:     "metric",
			BodyPath: "metrics",
		},
		&requestflag.Flag[float64]{
			Name:     "min-bucket-version",
			BodyPath: "minBucketVersion",
		},
		&requestflag.Flag[string]{
			Name:     "owner",
			Usage:    "The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.",
			BodyPath: "owner",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "phase",
			BodyPath: "phases",
		},
		&requestflag.Flag[*bool]{
			Name:     "post-stratification-enabled",
			Usage:    "When null, the organization default is used.",
			BodyPath: "postStratificationEnabled",
		},
		&requestflag.Flag[string]{
			Name:     "project",
			Usage:    "Project ID which the experiment belongs to",
			BodyPath: "project",
		},
		&requestflag.Flag[string]{
			Name:     "query-filter",
			Usage:    "WHERE clause to add to the default experiment query",
			BodyPath: "queryFilter",
		},
		&requestflag.Flag[bool]{
			Name:     "regression-adjustment-enabled",
			Usage:    "Controls whether regression adjustment (CUPED) is enabled for experiment analyses",
			BodyPath: "regressionAdjustmentEnabled",
		},
		&requestflag.Flag[string]{
			Name:     "released-variation-id",
			BodyPath: "releasedVariationId",
		},
		&requestflag.Flag[[]string]{
			Name:     "secondary-metric",
			BodyPath: "secondaryMetrics",
		},
		&requestflag.Flag[string]{
			Name:     "segment-id",
			Usage:    "Only users in this segment will be included",
			BodyPath: "segmentId",
		},
		&requestflag.Flag[bool]{
			Name:     "sequential-testing-enabled",
			Usage:    "Only applicable to frequentist analyses",
			BodyPath: "sequentialTestingEnabled",
		},
		&requestflag.Flag[float64]{
			Name:     "sequential-testing-tuning-parameter",
			BodyPath: "sequentialTestingTuningParameter",
		},
		&requestflag.Flag[string]{
			Name:     "share-level",
			Usage:    `Allowed values: "public", "organization".`,
			BodyPath: "shareLevel",
		},
		&requestflag.Flag[string]{
			Name:     "stats-engine",
			Usage:    `Allowed values: "bayesian", "frequentist".`,
			BodyPath: "statsEngine",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    `Allowed values: "draft", "running", "stopped".`,
			BodyPath: "status",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "template-id",
			Usage:    "ID of the [ExperimentTemplate](#tag/ExperimentTemplate_model) this experiment was created from. Template fields are applied by default and overridden by explicitly provided payload fields.",
			BodyPath: "templateId",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "standard", "multi-armed-bandit".`,
			BodyPath: "type",
		},
	},
	Action:          handleExperimentsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"variation": {
		&requestflag.InnerFlag[string]{
			Name:       "variation.key",
			InnerField: "key",
		},
		&requestflag.InnerFlag[string]{
			Name:       "variation.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "variation.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "variation.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "variation.screenshots",
			InnerField: "screenshots",
		},
	},
	"custom-metric-slice": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "custom-metric-slice.slices",
			InnerField: "slices",
		},
	},
	"decision-framework-settings": {
		&requestflag.InnerFlag[string]{
			Name:       "decision-framework-settings.decision-criteria-id",
			InnerField: "decisionCriteriaId",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "decision-framework-settings.decision-framework-metric-overrides",
			InnerField: "decisionFrameworkMetricOverrides",
		},
	},
	"lookback-override": {
		&requestflag.InnerFlag[string]{
			Name:       "lookback-override.type",
			Usage:      `Allowed values: "date", "window".`,
			InnerField: "type",
		},
		&requestflag.InnerFlag[any]{
			Name:       "lookback-override.value",
			Usage:      `For "window" type - non-negative numeric value (e.g. 7 for 7 days). For "date" type a date string.`,
			InnerField: "value",
		},
		&requestflag.InnerFlag[string]{
			Name:       "lookback-override.value-unit",
			Usage:      `Used when type is "window". Defaults to "days".`,
			InnerField: "valueUnit",
		},
	},
	"metric-override": {
		&requestflag.InnerFlag[string]{
			Name:       "metric-override.id",
			Usage:      "ID of the metric to override settings for.",
			InnerField: "id",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "metric-override.delay-hours",
			InnerField: "delayHours",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "metric-override.proper-prior-enabled",
			InnerField: "properPriorEnabled",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "metric-override.proper-prior-mean",
			InnerField: "properPriorMean",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "metric-override.proper-prior-override",
			Usage:      "Must be true for the override to take effect. If true, the other proper prior settings in this object will be used if present.",
			InnerField: "properPriorOverride",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "metric-override.proper-prior-std-dev",
			InnerField: "properPriorStdDev",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "metric-override.regression-adjustment-days",
			InnerField: "regressionAdjustmentDays",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "metric-override.regression-adjustment-enabled",
			InnerField: "regressionAdjustmentEnabled",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "metric-override.regression-adjustment-override",
			Usage:      "Must be true for the override to take effect. If true, the other regression adjustment settings in this object will be used if present.",
			InnerField: "regressionAdjustmentOverride",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "metric-override.window-hours",
			InnerField: "windowHours",
		},
		&requestflag.InnerFlag[string]{
			Name:       "metric-override.window-type",
			Usage:      `Allowed values: "conversion", "lookback", "".`,
			InnerField: "windowType",
		},
	},
	"phase": {
		&requestflag.InnerFlag[any]{
			Name:       "phase.date-started",
			InnerField: "dateStarted",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phase.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phase.condition",
			InnerField: "condition",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "phase.coverage",
			InnerField: "coverage",
		},
		&requestflag.InnerFlag[any]{
			Name:       "phase.date-ended",
			InnerField: "dateEnded",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "phase.namespace",
			InnerField: "namespace",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "phase.prerequisites",
			InnerField: "prerequisites",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phase.reason",
			InnerField: "reason",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phase.reason-for-stopping",
			InnerField: "reasonForStopping",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "phase.saved-group-targeting",
			InnerField: "savedGroupTargeting",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phase.seed",
			InnerField: "seed",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phase.targeting-condition",
			InnerField: "targetingCondition",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "phase.traffic-split",
			InnerField: "trafficSplit",
		},
		&requestflag.InnerFlag[[]float64]{
			Name:       "phase.variation-weights",
			InnerField: "variationWeights",
		},
	},
})

var experimentsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single experiment",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleExperimentsRetrieve,
	HideHelpCommand: true,
}

var experimentsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a single experiment",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "activation-metric",
			Usage:    "Users must convert on this metric before being included",
			BodyPath: "activationMetric",
		},
		&requestflag.Flag[string]{
			Name:     "analysis",
			Usage:    "Analysis summary or conclusions for the experiment. Maps to resultSummary.conclusions in the GET response.",
			BodyPath: "analysis",
		},
		&requestflag.Flag[bool]{
			Name:     "archived",
			BodyPath: "archived",
		},
		&requestflag.Flag[string]{
			Name:     "assignment-query-id",
			BodyPath: "assignmentQueryId",
		},
		&requestflag.Flag[string]{
			Name:     "attribution-model",
			Usage:    "Setting attribution model to `\"experimentDuration\"` is the same as selecting \"Ignore Conversion Windows\" for the Conversion Window Override. Setting it to `\"lookbackOverride\"` requires a `lookbackOverride` object to be provided.",
			BodyPath: "attributionModel",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-refresh",
			BodyPath: "autoRefresh",
		},
		&requestflag.Flag[string]{
			Name:     "bandit-burn-in-unit",
			Usage:    `Allowed values: "days", "hours".`,
			BodyPath: "banditBurnInUnit",
		},
		&requestflag.Flag[float64]{
			Name:     "bandit-burn-in-value",
			BodyPath: "banditBurnInValue",
		},
		&requestflag.Flag[string]{
			Name:     "bandit-conversion-window-unit",
			Usage:    `Allowed values: "days", "hours".`,
			BodyPath: "banditConversionWindowUnit",
		},
		&requestflag.Flag[float64]{
			Name:     "bandit-conversion-window-value",
			BodyPath: "banditConversionWindowValue",
		},
		&requestflag.Flag[string]{
			Name:     "bandit-schedule-unit",
			Usage:    `Allowed values: "days", "hours".`,
			BodyPath: "banditScheduleUnit",
		},
		&requestflag.Flag[float64]{
			Name:     "bandit-schedule-value",
			BodyPath: "banditScheduleValue",
		},
		&requestflag.Flag[float64]{
			Name:     "bucket-version",
			BodyPath: "bucketVersion",
		},
		&requestflag.Flag[bool]{
			Name:     "bypass-duplicate-key-check",
			Usage:    "If true, allow updating the tracking key even if another experiment with the same tracking key already exist. This is ignored if the organization requires unique tracking keys as a rule.",
			BodyPath: "bypassDuplicateKeyCheck",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "custom-fields",
			BodyPath: "customFields",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "custom-metric-slice",
			Usage:    "Custom slices that apply to ALL applicable metrics in the experiment",
			BodyPath: "customMetricSlices",
		},
		&requestflag.Flag[string]{
			Name:     "datasource-id",
			Usage:    "Can only be set if existing experiment does not have a datasource",
			BodyPath: "datasourceId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "decision-framework-settings",
			Usage:    "Controls the decision framework and metric overrides for the experiment. Replaces the entire stored object on update (does not patch individual fields).",
			BodyPath: "decisionFrameworkSettings",
		},
		&requestflag.Flag[string]{
			Name:     "default-dashboard-id",
			Usage:    "ID of the default dashboard for this experiment.",
			BodyPath: "defaultDashboardId",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of the experiment",
			BodyPath: "description",
		},
		&requestflag.Flag[bool]{
			Name:     "disable-sticky-bucketing",
			BodyPath: "disableStickyBucketing",
		},
		&requestflag.Flag[bool]{
			Name:     "exclude-from-payload",
			Usage:    "If true, the experiment is excluded from the SDK payload. Maps to resultSummary.excludeFromPayload in the GET response.",
			BodyPath: "excludeFromPayload",
		},
		&requestflag.Flag[string]{
			Name:     "fallback-attribute",
			BodyPath: "fallbackAttribute",
		},
		&requestflag.Flag[[]string]{
			Name:     "guardrail-metric",
			BodyPath: "guardrailMetrics",
		},
		&requestflag.Flag[string]{
			Name:     "hash-attribute",
			BodyPath: "hashAttribute",
		},
		&requestflag.Flag[float64]{
			Name:     "hash-version",
			Usage:    "Allowed values: 1, 2.",
			BodyPath: "hashVersion",
		},
		&requestflag.Flag[string]{
			Name:     "hypothesis",
			Usage:    "Hypothesis of the experiment",
			BodyPath: "hypothesis",
		},
		&requestflag.Flag[string]{
			Name:     "in-progress-conversions",
			Usage:    `Allowed values: "loose", "strict".`,
			BodyPath: "inProgressConversions",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "lookback-override",
			Usage:    `Controls the lookback override for the experiment. For type "window", value must be a non-negative number and valueUnit is required.`,
			BodyPath: "lookbackOverride",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "metric-override",
			Usage:    "Per-metric analysis overrides for this experiment. Replaces the entire stored array (does not patch individual entries).",
			BodyPath: "metricOverrides",
		},
		&requestflag.Flag[[]string]{
			Name:     "metric",
			BodyPath: "metrics",
		},
		&requestflag.Flag[float64]{
			Name:     "min-bucket-version",
			BodyPath: "minBucketVersion",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the experiment",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "owner",
			Usage:    "The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.",
			BodyPath: "owner",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "phase",
			BodyPath: "phases",
		},
		&requestflag.Flag[*bool]{
			Name:     "post-stratification-enabled",
			Usage:    "When null, the organization default is used.",
			BodyPath: "postStratificationEnabled",
		},
		&requestflag.Flag[string]{
			Name:     "project",
			Usage:    "Project ID which the experiment belongs to",
			BodyPath: "project",
		},
		&requestflag.Flag[string]{
			Name:     "query-filter",
			Usage:    "WHERE clause to add to the default experiment query",
			BodyPath: "queryFilter",
		},
		&requestflag.Flag[bool]{
			Name:     "regression-adjustment-enabled",
			Usage:    "Controls whether regression adjustment (CUPED) is enabled for experiment analyses",
			BodyPath: "regressionAdjustmentEnabled",
		},
		&requestflag.Flag[string]{
			Name:     "released-variation-id",
			Usage:    "The ID of the released variation. Maps to resultSummary.releasedVariationId in the GET response.",
			BodyPath: "releasedVariationId",
		},
		&requestflag.Flag[string]{
			Name:     "results",
			Usage:    "The result status of the experiment. Maps to resultSummary.status in the GET response.",
			BodyPath: "results",
		},
		&requestflag.Flag[[]string]{
			Name:     "secondary-metric",
			BodyPath: "secondaryMetrics",
		},
		&requestflag.Flag[string]{
			Name:     "segment-id",
			Usage:    "Only users in this segment will be included",
			BodyPath: "segmentId",
		},
		&requestflag.Flag[bool]{
			Name:     "sequential-testing-enabled",
			Usage:    "Only applicable to frequentist analyses",
			BodyPath: "sequentialTestingEnabled",
		},
		&requestflag.Flag[float64]{
			Name:     "sequential-testing-tuning-parameter",
			BodyPath: "sequentialTestingTuningParameter",
		},
		&requestflag.Flag[string]{
			Name:     "share-level",
			Usage:    `Allowed values: "public", "organization".`,
			BodyPath: "shareLevel",
		},
		&requestflag.Flag[string]{
			Name:     "stats-engine",
			Usage:    `Allowed values: "bayesian", "frequentist".`,
			BodyPath: "statsEngine",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    `Allowed values: "draft", "running", "stopped".`,
			BodyPath: "status",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "tracking-key",
			BodyPath: "trackingKey",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "standard", "multi-armed-bandit".`,
			BodyPath: "type",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "variation",
			BodyPath: "variations",
		},
		&requestflag.Flag[float64]{
			Name:     "winner",
			Usage:    "The index of the winning variation (0-indexed). Maps to resultSummary.winner (variation ID) in the GET response.",
			BodyPath: "winner",
		},
	},
	Action:          handleExperimentsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"custom-metric-slice": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "custom-metric-slice.slices",
			InnerField: "slices",
		},
	},
	"decision-framework-settings": {
		&requestflag.InnerFlag[string]{
			Name:       "decision-framework-settings.decision-criteria-id",
			InnerField: "decisionCriteriaId",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "decision-framework-settings.decision-framework-metric-overrides",
			InnerField: "decisionFrameworkMetricOverrides",
		},
	},
	"lookback-override": {
		&requestflag.InnerFlag[string]{
			Name:       "lookback-override.type",
			Usage:      `Allowed values: "date", "window".`,
			InnerField: "type",
		},
		&requestflag.InnerFlag[any]{
			Name:       "lookback-override.value",
			Usage:      `For "window" type - non-negative numeric value (e.g. 7 for 7 days). For "date" type a date string.`,
			InnerField: "value",
		},
		&requestflag.InnerFlag[string]{
			Name:       "lookback-override.value-unit",
			Usage:      `Used when type is "window". Defaults to "days".`,
			InnerField: "valueUnit",
		},
	},
	"metric-override": {
		&requestflag.InnerFlag[string]{
			Name:       "metric-override.id",
			Usage:      "ID of the metric to override settings for.",
			InnerField: "id",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "metric-override.delay-hours",
			InnerField: "delayHours",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "metric-override.proper-prior-enabled",
			InnerField: "properPriorEnabled",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "metric-override.proper-prior-mean",
			InnerField: "properPriorMean",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "metric-override.proper-prior-override",
			Usage:      "Must be true for the override to take effect. If true, the other proper prior settings in this object will be used if present.",
			InnerField: "properPriorOverride",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "metric-override.proper-prior-std-dev",
			InnerField: "properPriorStdDev",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "metric-override.regression-adjustment-days",
			InnerField: "regressionAdjustmentDays",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "metric-override.regression-adjustment-enabled",
			InnerField: "regressionAdjustmentEnabled",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "metric-override.regression-adjustment-override",
			Usage:      "Must be true for the override to take effect. If true, the other regression adjustment settings in this object will be used if present.",
			InnerField: "regressionAdjustmentOverride",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "metric-override.window-hours",
			InnerField: "windowHours",
		},
		&requestflag.InnerFlag[string]{
			Name:       "metric-override.window-type",
			Usage:      `Allowed values: "conversion", "lookback", "".`,
			InnerField: "windowType",
		},
	},
	"phase": {
		&requestflag.InnerFlag[any]{
			Name:       "phase.date-started",
			InnerField: "dateStarted",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phase.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phase.condition",
			InnerField: "condition",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "phase.coverage",
			InnerField: "coverage",
		},
		&requestflag.InnerFlag[any]{
			Name:       "phase.date-ended",
			InnerField: "dateEnded",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "phase.namespace",
			InnerField: "namespace",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "phase.prerequisites",
			InnerField: "prerequisites",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phase.reason",
			InnerField: "reason",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phase.reason-for-stopping",
			InnerField: "reasonForStopping",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "phase.saved-group-targeting",
			InnerField: "savedGroupTargeting",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phase.seed",
			InnerField: "seed",
		},
		&requestflag.InnerFlag[string]{
			Name:       "phase.targeting-condition",
			InnerField: "targetingCondition",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "phase.traffic-split",
			Usage:      "Deprecated and unused. Use variationWeights instead.",
			InnerField: "trafficSplit",
		},
		&requestflag.InnerFlag[[]float64]{
			Name:       "phase.variation-weights",
			InnerField: "variationWeights",
		},
	},
	"variation": {
		&requestflag.InnerFlag[string]{
			Name:       "variation.key",
			InnerField: "key",
		},
		&requestflag.InnerFlag[string]{
			Name:       "variation.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "variation.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "variation.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "variation.screenshots",
			InnerField: "screenshots",
		},
	},
})

var experimentsList = cli.Command{
	Name:    "list",
	Usage:   "Get all experiments",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "datasource-id",
			Usage:     "Filter by Data Source",
			QueryPath: "datasourceId",
		},
		&requestflag.Flag[string]{
			Name:      "experiment-id",
			Usage:     "Filter the returned list by the experiment tracking key (not the internal experiment ID). Note, this was deprecated to help reduce confusion, consider using `trackingKey` instead, which is functionally identical. You cannot use both params at the same time.",
			QueryPath: "experimentId",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "The number of items to return",
			Default:   10,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "offset",
			Usage:     "How many items to skip (use in conjunction with limit for pagination)",
			Default:   0,
			QueryPath: "offset",
		},
		&requestflag.Flag[string]{
			Name:      "project-id",
			Usage:     "Filter by project id",
			QueryPath: "projectId",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     `Allowed values: "draft", "running", "stopped".`,
			QueryPath: "status",
		},
		&requestflag.Flag[string]{
			Name:      "tracking-key",
			Usage:     "Filter by experiment tracking key",
			QueryPath: "trackingKey",
		},
	},
	Action:          handleExperimentsList,
	HideHelpCommand: true,
}

var experimentsCreateSnapshot = cli.Command{
	Name:    "create-snapshot",
	Usage:   "Create Experiment Snapshot",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The experiment id of the experiment to update",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "dimension",
			Usage:    `Dimension to break results down by. For Unit Dimensions, use the dimension id (e.g. "dim_abc123"). For Experiment Dimensions, use "exp:<dimensionName>" (e.g. "exp:country"). Built-in pre-exposure dimensions include "pre:date" and, when configured, "pre:activation". Omit this field to create a standard snapshot.`,
			BodyPath: "dimension",
		},
		&requestflag.Flag[int64]{
			Name:     "phase",
			Usage:    "Zero-based phase index to snapshot, where 0 is the first experiment phase. Defaults to the latest phase.",
			BodyPath: "phase",
		},
		&requestflag.Flag[string]{
			Name:     "triggered-by",
			Usage:    `Set to "schedule" if you want this request to trigger notifications and other events as it if were a scheduled update. Defaults to manual.`,
			BodyPath: "triggeredBy",
		},
	},
	Action:          handleExperimentsCreateSnapshot,
	HideHelpCommand: true,
}

var experimentsModifyTemporaryRollout = cli.Command{
	Name:    "modify-temporary-rollout",
	Usage:   "Modify temporary rollout status for a stopped experiment",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[bool]{
			Name:     "enable-temporary-rollout",
			Usage:    "If true, keep the stopped experiment in SDK payload and force traffic to the winner variation. If false, end temporary rollout and remove from SDK payload.",
			Required: true,
			BodyPath: "enableTemporaryRollout",
		},
		&requestflag.Flag[string]{
			Name:     "released-variation-id",
			Usage:    "Variation ID (e.g. var_abc123) to release to 100% of traffic eligible for this experiment. Required if enableTemporaryRollout is true.",
			BodyPath: "releasedVariationId",
		},
	},
	Action:          handleExperimentsModifyTemporaryRollout,
	HideHelpCommand: true,
}

var experimentsStart = cli.Command{
	Name:    "start",
	Usage:   "Start an experiment",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[bool]{
			Name:     "skip-checklist",
			Usage:    "If true, skips validating the experiment satisifies all pre-launch checklist items",
			BodyPath: "skipChecklist",
		},
	},
	Action:          handleExperimentsStart,
	HideHelpCommand: true,
}

var experimentsStop = cli.Command{
	Name:    "stop",
	Usage:   "Stop an experiment",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "results",
			Usage:    "The experiment conclusion status.",
			Required: true,
			BodyPath: "results",
		},
		&requestflag.Flag[string]{
			Name:     "analysis",
			Usage:    "Optional markdown summary displayed on the experiment results page.",
			BodyPath: "analysis",
		},
		&requestflag.Flag[string]{
			Name:     "date-ended",
			Usage:    "Optional ISO datetime for ending the latest phase. Defaults to the current date and time.",
			BodyPath: "dateEnded",
		},
		&requestflag.Flag[bool]{
			Name:     "enable-temporary-rollout",
			Usage:    "If true, include this stopped experiment in SDK payload and force the release variation (`releasedVariationId`) to all traffic.",
			BodyPath: "enableTemporaryRollout",
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Usage:    "Optional reason for ending the phase stored on the latest phase metadata.",
			BodyPath: "reason",
		},
		&requestflag.Flag[string]{
			Name:     "released-variation-id",
			Usage:    "Required if enableTemporaryRollout is true. Variation ID (e.g. var_abc123) to release to 100% of traffic eligible for this experiment.",
			BodyPath: "releasedVariationId",
		},
		&requestflag.Flag[string]{
			Name:     "winner-variation-id",
			Usage:    "Variation ID (e.g. var_abc123) of the winning variation. Used only as metadata. Required if results is 'won' and there are multiple test variations. Otherwise, defaults to the test variation when results is 'won' and to the baseline variation for other results.",
			BodyPath: "winnerVariationId",
		},
	},
	Action:          handleExperimentsStop,
	HideHelpCommand: true,
}

func handleExperimentsCreate(ctx context.Context, cmd *cli.Command) error {
	client := growthbook.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := growthbook.ExperimentNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "experiments create",
		Transform:      transform,
	})
}

func handleExperimentsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := growthbook.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "experiments retrieve",
		Transform:      transform,
	})
}

func handleExperimentsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := growthbook.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := growthbook.ExperimentUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.Update(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "experiments update",
		Transform:      transform,
	})
}

func handleExperimentsList(ctx context.Context, cmd *cli.Command) error {
	client := growthbook.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := growthbook.ExperimentListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "experiments list",
		Transform:      transform,
	})
}

func handleExperimentsCreateSnapshot(ctx context.Context, cmd *cli.Command) error {
	client := growthbook.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := growthbook.ExperimentNewSnapshotParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.NewSnapshot(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "experiments create-snapshot",
		Transform:      transform,
	})
}

func handleExperimentsModifyTemporaryRollout(ctx context.Context, cmd *cli.Command) error {
	client := growthbook.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := growthbook.ExperimentModifyTemporaryRolloutParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.ModifyTemporaryRollout(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "experiments modify-temporary-rollout",
		Transform:      transform,
	})
}

func handleExperimentsStart(ctx context.Context, cmd *cli.Command) error {
	client := growthbook.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := growthbook.ExperimentStartParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.Start(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "experiments start",
		Transform:      transform,
	})
}

func handleExperimentsStop(ctx context.Context, cmd *cli.Command) error {
	client := growthbook.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := growthbook.ExperimentStopParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.Stop(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "experiments stop",
		Transform:      transform,
	})
}
