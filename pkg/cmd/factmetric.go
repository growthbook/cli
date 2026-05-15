// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/growthbook/cli/internal/apiquery"
	"github.com/growthbook/cli/internal/requestflag"
	"github.com/stainless-sdks/growthbook-go"
	"github.com/stainless-sdks/growthbook-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var factMetricsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a single fact metric",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "metric-type",
			Usage:    `Allowed values: "proportion", "retention", "mean", "quantile", "ratio", "dailyParticipation".`,
			Required: true,
			BodyPath: "metricType",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "numerator",
			Required: true,
			BodyPath: "numerator",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "capping-settings",
			Usage:    "Controls how outliers are handled",
			BodyPath: "cappingSettings",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "denominator",
			Usage:    "Only when metricType is 'ratio'",
			BodyPath: "denominator",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[bool]{
			Name:     "display-as-percentage",
			Usage:    "If true and the metric is a ratio or dailyParticipation metric, variation means will be displayed as a percentage. Defaults to true for dailyParticipation metrics and false for ratio metrics.",
			BodyPath: "displayAsPercentage",
		},
		&requestflag.Flag[bool]{
			Name:     "inverse",
			Usage:    "Set to true for things like Bounce Rate, where you want the metric to decrease",
			BodyPath: "inverse",
		},
		&requestflag.Flag[string]{
			Name:     "managed-by",
			Usage:    `Set this to "api" to disable editing in the GrowthBook UI`,
			BodyPath: "managedBy",
		},
		&requestflag.Flag[float64]{
			Name:     "max-percent-change",
			Usage:    "Maximum percent change to consider uplift significant, as a proportion (e.g. put 0.5 for 50%)",
			BodyPath: "maxPercentChange",
		},
		&requestflag.Flag[[]string]{
			Name:     "metric-auto-slice",
			Usage:    "Array of slice column names that will be automatically included in metric analysis. This is an enterprise feature.",
			BodyPath: "metricAutoSlices",
		},
		&requestflag.Flag[float64]{
			Name:     "min-percent-change",
			Usage:    "Minimum percent change to consider uplift significant, as a proportion (e.g. put 0.005 for 0.5%)",
			BodyPath: "minPercentChange",
		},
		&requestflag.Flag[float64]{
			Name:     "min-sample-size",
			BodyPath: "minSampleSize",
		},
		&requestflag.Flag[string]{
			Name:     "owner",
			Usage:    "The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.",
			BodyPath: "owner",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "prior-settings",
			Usage:    "Controls the bayesian prior for the metric. If omitted, organization defaults will be used.",
			BodyPath: "priorSettings",
		},
		&requestflag.Flag[[]string]{
			Name:     "project",
			BodyPath: "projects",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "quantile-settings",
			Usage:    `Controls the settings for quantile metrics (mandatory if metricType is "quantile")`,
			BodyPath: "quantileSettings",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "regression-adjustment-settings",
			Usage:    "Controls the regression adjustment (CUPED) settings for the metric",
			BodyPath: "regressionAdjustmentSettings",
		},
		&requestflag.Flag[float64]{
			Name:     "risk-threshold-danger",
			Usage:    "No longer used. Threshold for Risk to be considered too high, as a proportion (e.g. put 0.0125 for 1.25%). <br/> Must be a non-negative number.",
			BodyPath: "riskThresholdDanger",
		},
		&requestflag.Flag[float64]{
			Name:     "risk-threshold-success",
			Usage:    "No longer used. Threshold for Risk to be considered low enough, as a proportion (e.g. put 0.0025 for 0.25%). <br/> Must be a non-negative number and must not be higher than `riskThresholdDanger`.",
			BodyPath: "riskThresholdSuccess",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[float64]{
			Name:     "target-mde",
			Usage:    `The percentage change that you want to reliably detect before ending an experiment, as a proportion (e.g. put 0.1 for 10%). This is used to estimate the "Days Left" for running experiments.`,
			BodyPath: "targetMDE",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "window-settings",
			Usage:    "Controls the conversion window for the metric",
			BodyPath: "windowSettings",
		},
	},
	Action:          handleFactMetricsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"numerator": {
		&requestflag.InnerFlag[string]{
			Name:       "numerator.fact-table-id",
			InnerField: "factTableId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "numerator.aggregate-filter",
			Usage:      "Simple comparison operator and value to apply after aggregation (e.g. '= 10' or '>= 1'). Requires `aggregateFilterColumn`.",
			InnerField: "aggregateFilter",
		},
		&requestflag.InnerFlag[string]{
			Name:       "numerator.aggregate-filter-column",
			Usage:      "Column to use to filter users after aggregation. Either '$$count' of rows or the name of a numeric column that will be summed by user. Must specify `aggregateFilter` if using this. Only can be used with 'retention' and 'proportion' metrics.",
			InnerField: "aggregateFilterColumn",
		},
		&requestflag.InnerFlag[string]{
			Name:       "numerator.aggregation",
			Usage:      "User aggregation of selected column. Either sum or max for numeric columns; count distinct for string columns; hll merge / kll merge for pre-built sketch columns (requires data-source support); ignored for special columns. Default: sum. If you specify a string column you must explicitly specify count distinct. Not used for proportion metrics; for event quantile metrics only kll merge is applicable.",
			InnerField: "aggregation",
		},
		&requestflag.InnerFlag[string]{
			Name:       "numerator.column",
			Usage:      "Must be empty for proportion metrics and dailyParticipation metrics. Otherwise, the column name or one of the special values: '$$distinctUsers' or '$$count' (or '$$distinctDates' if metricType is 'mean' or 'ratio' or 'quantile' and quantileSettings.type is 'unit')",
			InnerField: "column",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "numerator.filters",
			Usage:      "Array of Fact Table Filter Ids. Deprecated, use rowFilters instead.",
			InnerField: "filters",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "numerator.inline-filters",
			Usage:      "Inline filters to apply to the fact table. Keys are column names, values are arrays of values to filter by. Deprecated, use rowFilters instead.",
			InnerField: "inlineFilters",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "numerator.row-filters",
			Usage:      "Filters to apply to the rows of the fact table before aggregation.",
			InnerField: "rowFilters",
		},
	},
	"capping-settings": {
		&requestflag.InnerFlag[string]{
			Name:       "capping-settings.type",
			Usage:      `Allowed values: "none", "absolute", "percentile".`,
			InnerField: "type",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "capping-settings.ignore-zeros",
			Usage:      "If true and capping is `percentile`, zeros will be ignored when calculating the percentile.",
			InnerField: "ignoreZeros",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "capping-settings.value",
			Usage:      "When type is absolute, this is the absolute value. When type is percentile, this is the percentile value (from 0.0 to 1.0).",
			InnerField: "value",
		},
	},
	"denominator": {
		&requestflag.InnerFlag[string]{
			Name:       "denominator.column",
			Usage:      "The column name or one of the special values: '$$distinctUsers' or '$$count' (or '$$distinctDates' if metricType is 'mean' or 'ratio' or 'quantile' and quantileSettings.type is 'unit')",
			InnerField: "column",
		},
		&requestflag.InnerFlag[string]{
			Name:       "denominator.fact-table-id",
			InnerField: "factTableId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "denominator.aggregation",
			Usage:      "User aggregation of selected column. Either sum or max for numeric columns; count distinct for string columns; hll merge / kll merge for pre-built sketch columns (requires data-source support); ignored for special columns. Default: sum. If you specify a string column you must explicitly specify count distinct. Not used for proportion metrics; for event quantile metrics only kll merge is applicable.",
			InnerField: "aggregation",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "denominator.filters",
			Usage:      "Array of Fact Table Filter Ids. Deprecated, use rowFilters instead.",
			InnerField: "filters",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "denominator.inline-filters",
			Usage:      "Inline filters to apply to the fact table. Keys are column names, values are arrays of values to filter by. Deprecated, use rowFilters instead.",
			InnerField: "inlineFilters",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "denominator.row-filters",
			Usage:      "Filters to apply to the rows of the fact table before aggregation.",
			InnerField: "rowFilters",
		},
	},
	"prior-settings": {
		&requestflag.InnerFlag[float64]{
			Name:       "prior-settings.mean",
			Usage:      "The mean of the prior distribution of relative effects in proportion terms (e.g. 0.01 is 1%)",
			InnerField: "mean",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "prior-settings.override",
			Usage:      "If false, the organization default settings will be used instead of the other settings in this object",
			InnerField: "override",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "prior-settings.proper",
			Usage:      "If true, the `mean` and `stddev` will be used, otherwise we will use an improper flat prior.",
			InnerField: "proper",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "prior-settings.stddev",
			Usage:      "Must be > 0. The standard deviation of the prior distribution of relative effects in proportion terms.",
			InnerField: "stddev",
		},
	},
	"quantile-settings": {
		&requestflag.InnerFlag[bool]{
			Name:       "quantile-settings.ignore-zeros",
			Usage:      "If true, zero values will be ignored when calculating the quantile",
			InnerField: "ignoreZeros",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "quantile-settings.quantile",
			Usage:      "The quantile value (from 0.001 to 0.999)",
			InnerField: "quantile",
		},
		&requestflag.InnerFlag[string]{
			Name:       "quantile-settings.type",
			Usage:      "Whether the quantile is over unit aggregations or raw event values",
			InnerField: "type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "quantile-settings.quantile-event-count-column",
			Usage:      "Optional override for the source-column name used to recover per-row event counts when numerator.aggregation is 'kll merge'. Defaults to '<numerator.column>_n_events'. Only valid for event-quantile metrics with a 'kll merge' numerator.",
			InnerField: "quantileEventCountColumn",
		},
	},
	"regression-adjustment-settings": {
		&requestflag.InnerFlag[bool]{
			Name:       "regression-adjustment-settings.override",
			Usage:      "If false, the organization default settings will be used",
			InnerField: "override",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "regression-adjustment-settings.days",
			Usage:      "Number of pre-exposure days to use for the regression adjustment",
			InnerField: "days",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "regression-adjustment-settings.enabled",
			Usage:      "Controls whether or not regression adjustment is applied to the metric",
			InnerField: "enabled",
		},
	},
	"window-settings": {
		&requestflag.InnerFlag[string]{
			Name:       "window-settings.type",
			Usage:      `Allowed values: "none", "conversion", "lookback".`,
			InnerField: "type",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "window-settings.delay-hours",
			Usage:      "Wait this many hours after experiment exposure before counting conversions. Ignored if delayValue is set.",
			InnerField: "delayHours",
		},
		&requestflag.InnerFlag[string]{
			Name:       "window-settings.delay-unit",
			Usage:      "Default `hours`.",
			InnerField: "delayUnit",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "window-settings.delay-value",
			Usage:      "Wait this long after experiment exposure before counting conversions.",
			InnerField: "delayValue",
		},
		&requestflag.InnerFlag[string]{
			Name:       "window-settings.window-unit",
			Usage:      "Default `hours`.",
			InnerField: "windowUnit",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "window-settings.window-value",
			InnerField: "windowValue",
		},
	},
})

var factMetricsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single fact metric",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleFactMetricsRetrieve,
	HideHelpCommand: true,
}

var factMetricsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a single fact metric",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[bool]{
			Name:     "archived",
			BodyPath: "archived",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "capping-settings",
			Usage:    "Controls how outliers are handled",
			BodyPath: "cappingSettings",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "denominator",
			Usage:    "Only when metricType is 'ratio'",
			BodyPath: "denominator",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[bool]{
			Name:     "display-as-percentage",
			Usage:    "If true and the metric is a ratio or dailyParticipation metric, variation means will be displayed as a percentage. Defaults to true for dailyParticipation metrics and false for ratio metrics.",
			BodyPath: "displayAsPercentage",
		},
		&requestflag.Flag[bool]{
			Name:     "inverse",
			Usage:    "Set to true for things like Bounce Rate, where you want the metric to decrease",
			BodyPath: "inverse",
		},
		&requestflag.Flag[string]{
			Name:     "managed-by",
			Usage:    `Set this to "api" to disable editing in the GrowthBook UI`,
			BodyPath: "managedBy",
		},
		&requestflag.Flag[float64]{
			Name:     "max-percent-change",
			Usage:    "Maximum percent change to consider uplift significant, as a proportion (e.g. put 0.5 for 50%)",
			BodyPath: "maxPercentChange",
		},
		&requestflag.Flag[[]string]{
			Name:     "metric-auto-slice",
			Usage:    "Array of slice column names that will be automatically included in metric analysis. This is an enterprise feature.",
			BodyPath: "metricAutoSlices",
		},
		&requestflag.Flag[string]{
			Name:     "metric-type",
			Usage:    `Allowed values: "proportion", "retention", "mean", "quantile", "ratio", "dailyParticipation".`,
			BodyPath: "metricType",
		},
		&requestflag.Flag[float64]{
			Name:     "min-percent-change",
			Usage:    "Minimum percent change to consider uplift significant, as a proportion (e.g. put 0.005 for 0.5%)",
			BodyPath: "minPercentChange",
		},
		&requestflag.Flag[float64]{
			Name:     "min-sample-size",
			BodyPath: "minSampleSize",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "numerator",
			BodyPath: "numerator",
		},
		&requestflag.Flag[string]{
			Name:     "owner",
			Usage:    "The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.",
			BodyPath: "owner",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "prior-settings",
			Usage:    "Controls the bayesian prior for the metric. If omitted, organization defaults will be used.",
			BodyPath: "priorSettings",
		},
		&requestflag.Flag[[]string]{
			Name:     "project",
			BodyPath: "projects",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "quantile-settings",
			Usage:    `Controls the settings for quantile metrics (mandatory if metricType is "quantile")`,
			BodyPath: "quantileSettings",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "regression-adjustment-settings",
			Usage:    "Controls the regression adjustment (CUPED) settings for the metric",
			BodyPath: "regressionAdjustmentSettings",
		},
		&requestflag.Flag[float64]{
			Name:     "risk-threshold-danger",
			Usage:    "No longer used. Threshold for Risk to be considered too high, as a proportion (e.g. put 0.0125 for 1.25%). <br/> Must be a non-negative number.",
			BodyPath: "riskThresholdDanger",
		},
		&requestflag.Flag[float64]{
			Name:     "risk-threshold-success",
			Usage:    "No longer used. Threshold for Risk to be considered low enough, as a proportion (e.g. put 0.0025 for 0.25%). <br/> Must be a non-negative number and must not be higher than `riskThresholdDanger`.",
			BodyPath: "riskThresholdSuccess",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[float64]{
			Name:     "target-mde",
			BodyPath: "targetMDE",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "window-settings",
			Usage:    "Controls the conversion window for the metric",
			BodyPath: "windowSettings",
		},
	},
	Action:          handleFactMetricsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"capping-settings": {
		&requestflag.InnerFlag[string]{
			Name:       "capping-settings.type",
			Usage:      `Allowed values: "none", "absolute", "percentile".`,
			InnerField: "type",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "capping-settings.ignore-zeros",
			Usage:      "If true and capping is `percentile`, zeros will be ignored when calculating the percentile.",
			InnerField: "ignoreZeros",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "capping-settings.value",
			Usage:      "When type is absolute, this is the absolute value. When type is percentile, this is the percentile value (from 0.0 to 1.0).",
			InnerField: "value",
		},
	},
	"denominator": {
		&requestflag.InnerFlag[string]{
			Name:       "denominator.column",
			Usage:      "The column name or one of the special values: '$$distinctUsers' or '$$count' (or '$$distinctDates' if metricType is 'mean' or 'ratio' or 'quantile' and quantileSettings.type is 'unit')",
			InnerField: "column",
		},
		&requestflag.InnerFlag[string]{
			Name:       "denominator.fact-table-id",
			InnerField: "factTableId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "denominator.aggregation",
			Usage:      "User aggregation of selected column. Either sum or max for numeric columns; count distinct for string columns; hll merge / kll merge for pre-built sketch columns (requires data-source support); ignored for special columns. Default: sum. If you specify a string column you must explicitly specify count distinct. Not used for proportion metrics; for event quantile metrics only kll merge is applicable.",
			InnerField: "aggregation",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "denominator.filters",
			Usage:      "Array of Fact Table Filter Ids. Deprecated, use rowFilters instead.",
			InnerField: "filters",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "denominator.inline-filters",
			Usage:      "Inline filters to apply to the fact table. Keys are column names, values are arrays of values to filter by. Deprecated, use rowFilters instead.",
			InnerField: "inlineFilters",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "denominator.row-filters",
			Usage:      "Filters to apply to the rows of the fact table before aggregation.",
			InnerField: "rowFilters",
		},
	},
	"numerator": {
		&requestflag.InnerFlag[string]{
			Name:       "numerator.fact-table-id",
			InnerField: "factTableId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "numerator.aggregate-filter",
			Usage:      "Simple comparison operator and value to apply after aggregation (e.g. '= 10' or '>= 1'). Requires `aggregateFilterColumn`.",
			InnerField: "aggregateFilter",
		},
		&requestflag.InnerFlag[string]{
			Name:       "numerator.aggregate-filter-column",
			Usage:      "Column to use to filter users after aggregation. Either '$$count' of rows or the name of a numeric column that will be summed by user. Must specify `aggregateFilter` if using this. Only can be used with 'retention' and 'proportion' metrics.",
			InnerField: "aggregateFilterColumn",
		},
		&requestflag.InnerFlag[string]{
			Name:       "numerator.aggregation",
			Usage:      "User aggregation of selected column. Either sum or max for numeric columns; count distinct for string columns; hll merge / kll merge for pre-built sketch columns (requires data-source support); ignored for special columns. Default: sum. If you specify a string column you must explicitly specify count distinct. Not used for proportion metrics; for event quantile metrics only kll merge is applicable.",
			InnerField: "aggregation",
		},
		&requestflag.InnerFlag[string]{
			Name:       "numerator.column",
			Usage:      "Must be empty for proportion metrics and dailyParticipation metrics. Otherwise, the column name or one of the special values: '$$distinctUsers' or '$$count' (or '$$distinctDates' if metricType is 'mean' or 'ratio' or 'quantile' and quantileSettings.type is 'unit')",
			InnerField: "column",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "numerator.filters",
			Usage:      "Array of Fact Table Filter Ids. Deprecated, use rowFilters instead.",
			InnerField: "filters",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "numerator.inline-filters",
			Usage:      "Inline filters to apply to the fact table. Keys are column names, values are arrays of values to filter by. Deprecated, use rowFilters instead.",
			InnerField: "inlineFilters",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "numerator.row-filters",
			Usage:      "Filters to apply to the rows of the fact table before aggregation.",
			InnerField: "rowFilters",
		},
	},
	"prior-settings": {
		&requestflag.InnerFlag[float64]{
			Name:       "prior-settings.mean",
			Usage:      "The mean of the prior distribution of relative effects in proportion terms (e.g. 0.01 is 1%)",
			InnerField: "mean",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "prior-settings.override",
			Usage:      "If false, the organization default settings will be used instead of the other settings in this object",
			InnerField: "override",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "prior-settings.proper",
			Usage:      "If true, the `mean` and `stddev` will be used, otherwise we will use an improper flat prior.",
			InnerField: "proper",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "prior-settings.stddev",
			Usage:      "Must be > 0. The standard deviation of the prior distribution of relative effects in proportion terms.",
			InnerField: "stddev",
		},
	},
	"quantile-settings": {
		&requestflag.InnerFlag[bool]{
			Name:       "quantile-settings.ignore-zeros",
			Usage:      "If true, zero values will be ignored when calculating the quantile",
			InnerField: "ignoreZeros",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "quantile-settings.quantile",
			Usage:      "The quantile value (from 0.001 to 0.999)",
			InnerField: "quantile",
		},
		&requestflag.InnerFlag[string]{
			Name:       "quantile-settings.type",
			Usage:      "Whether the quantile is over unit aggregations or raw event values",
			InnerField: "type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "quantile-settings.quantile-event-count-column",
			Usage:      "Optional override for the source-column name used to recover per-row event counts when numerator.aggregation is 'kll merge'. Defaults to '<numerator.column>_n_events'. Only valid for event-quantile metrics with a 'kll merge' numerator.",
			InnerField: "quantileEventCountColumn",
		},
	},
	"regression-adjustment-settings": {
		&requestflag.InnerFlag[bool]{
			Name:       "regression-adjustment-settings.override",
			Usage:      "If false, the organization default settings will be used",
			InnerField: "override",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "regression-adjustment-settings.days",
			Usage:      "Number of pre-exposure days to use for the regression adjustment",
			InnerField: "days",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "regression-adjustment-settings.enabled",
			Usage:      "Controls whether or not regression adjustment is applied to the metric",
			InnerField: "enabled",
		},
	},
	"window-settings": {
		&requestflag.InnerFlag[string]{
			Name:       "window-settings.type",
			Usage:      `Allowed values: "none", "conversion", "lookback".`,
			InnerField: "type",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "window-settings.delay-hours",
			Usage:      "Wait this many hours after experiment exposure before counting conversions. Ignored if delayValue is set.",
			InnerField: "delayHours",
		},
		&requestflag.InnerFlag[string]{
			Name:       "window-settings.delay-unit",
			Usage:      "Default `hours`.",
			InnerField: "delayUnit",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "window-settings.delay-value",
			Usage:      "Wait this long after experiment exposure before counting conversions.",
			InnerField: "delayValue",
		},
		&requestflag.InnerFlag[string]{
			Name:       "window-settings.window-unit",
			Usage:      "Default `hours`.",
			InnerField: "windowUnit",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "window-settings.window-value",
			InnerField: "windowValue",
		},
	},
})

var factMetricsList = cli.Command{
	Name:    "list",
	Usage:   "Get all fact metrics",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "datasource-id",
			Usage:     "Filter by Data Source",
			QueryPath: "datasourceId",
		},
		&requestflag.Flag[string]{
			Name:      "fact-table-id",
			Usage:     "Filter by Fact Table Id (for ratio metrics, we only look at the numerator)",
			QueryPath: "factTableId",
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
	},
	Action:          handleFactMetricsList,
	HideHelpCommand: true,
}

var factMetricsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a single fact metric",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleFactMetricsDelete,
	HideHelpCommand: true,
}

var factMetricsCreateAnalysis = cli.Command{
	Name:    "create-analysis",
	Usage:   "Create a fact metric analysis",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The fact metric id to analyze",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[[]string]{
			Name:     "additional-denominator-filter",
			Usage:    "We support passing in adhoc filters for an analysis that don't live on the metric itself. These are in addition to the metric's filters. To use this, you can pass in an array of Fact Table Filter Ids.",
			BodyPath: "additionalDenominatorFilters",
		},
		&requestflag.Flag[[]string]{
			Name:     "additional-numerator-filter",
			Usage:    "We support passing in adhoc filters for an analysis that don't live on the metric itself. These are in addition to the metric's filters. To use this, you can pass in an array of Fact Table Filter Ids.",
			BodyPath: "additionalNumeratorFilters",
		},
		&requestflag.Flag[float64]{
			Name:     "lookback-days",
			Usage:    "Number of days to look back for the analysis. Defaults to 30.",
			BodyPath: "lookbackDays",
		},
		&requestflag.Flag[*string]{
			Name:     "population-id",
			Usage:    "The ID of the population (e.g., segment ID) when populationType is not 'factTable'. Defaults to null.",
			BodyPath: "populationId",
		},
		&requestflag.Flag[string]{
			Name:     "population-type",
			Usage:    "The type of population to analyze. Defaults to 'factTable', meaning the analysis will return the metric value for all units found in the fact table.",
			BodyPath: "populationType",
		},
		&requestflag.Flag[bool]{
			Name:     "use-cache",
			Usage:    "Whether to use a cached query if one exists. Defaults to true.",
			BodyPath: "useCache",
		},
		&requestflag.Flag[string]{
			Name:     "user-id-type",
			Usage:    "The identifier type to use for the analysis. If not provided, defaults to the first available identifier type in the fact table.",
			BodyPath: "userIdType",
		},
	},
	Action:          handleFactMetricsCreateAnalysis,
	HideHelpCommand: true,
}

func handleFactMetricsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.FactMetricNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.FactMetrics.New(ctx, params, options...)
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
		Title:          "fact-metrics create",
		Transform:      transform,
	})
}

func handleFactMetricsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.FactMetrics.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "fact-metrics retrieve",
		Transform:      transform,
	})
}

func handleFactMetricsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.FactMetricUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.FactMetrics.Update(
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
		Title:          "fact-metrics update",
		Transform:      transform,
	})
}

func handleFactMetricsList(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.FactMetricListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.FactMetrics.List(ctx, params, options...)
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
		Title:          "fact-metrics list",
		Transform:      transform,
	})
}

func handleFactMetricsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.FactMetrics.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "fact-metrics delete",
		Transform:      transform,
	})
}

func handleFactMetricsCreateAnalysis(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.FactMetricNewAnalysisParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.FactMetrics.NewAnalysis(
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
		Title:          "fact-metrics create-analysis",
		Transform:      transform,
	})
}
