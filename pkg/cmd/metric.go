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

var metricsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a single metric",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "datasource-id",
			Usage:    "ID for the [DataSource](#tag/DataSource_model)",
			Required: true,
			BodyPath: "datasourceId",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the metric",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Type of metric. See [Metrics documentation](/app/metrics/legacy)",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[bool]{
			Name:     "archived",
			BodyPath: "archived",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "behavior",
			BodyPath: "behavior",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of the metric",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "managed-by",
			Usage:    `Where this metric must be managed from. If not set (empty string), it can be managed from anywhere. If set to "api", it can be managed via the API only.`,
			BodyPath: "managedBy",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "mixpanel",
			Usage:    "Only use for MixPanel (non-SQL) Data Sources. Only one of `sql`, `sqlBuilder` or `mixpanel` allowed, and at least one must be specified.",
			BodyPath: "mixpanel",
		},
		&requestflag.Flag[string]{
			Name:     "owner",
			Usage:    "The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.",
			BodyPath: "owner",
		},
		&requestflag.Flag[[]string]{
			Name:     "project",
			Usage:    "List of project IDs for projects that can access this metric",
			BodyPath: "projects",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sql",
			Usage:    "Preferred way to define SQL. Only one of `sql`, `sqlBuilder` or `mixpanel` allowed, and at least one must be specified.",
			BodyPath: "sql",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sql-builder",
			Usage:    "An alternative way to specify a SQL metric, rather than a full query. Using `sql` is preferred to `sqlBuilder`. Only one of `sql`, `sqlBuilder` or `mixpanel` allowed, and at least one must be specified.",
			BodyPath: "sqlBuilder",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "List of tags",
			BodyPath: "tags",
		},
	},
	Action:          handleMetricsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"behavior": {
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.cap",
			Usage:      "(deprecated, use cappingSettings instead) This should be non-negative",
			InnerField: "cap",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "behavior.capping",
			Usage:      "(deprecated, use cappingSettings instead) Used in conjunction with `capValue` to set the capping (winsorization). Do not specify or set to null for no capping. \"absolute\" will cap user values at the `capValue` if it is greater than 0. \"percentile\" will cap user values at the percentile of user values in an experiment using the `capValue` for the percentile, if greater than 0. <br/>  If `behavior.capping` is non-null, you must specify `behavior.capValue`.",
			InnerField: "capping",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "behavior.capping-settings",
			Usage:      "Controls how outliers are handled",
			InnerField: "cappingSettings",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.cap-value",
			Usage:      "(deprecated, use cappingSettings instead) This should be non-negative. <br/> Must specify `behavior.capping` when setting `behavior.capValue`.",
			InnerField: "capValue",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.conversion-window-end",
			Usage:      "The end of a [Conversion Window](/app/metrics/legacy/#conversion-window) relative to the exposure date, in hours. This is equivalent to the [Conversion Delay](/app/metrics/legacy/#conversion-delay) + Conversion Window Hours settings in the UI. In other words, if you want a 48 hour window starting after 24 hours, you would set conversionWindowStart to 24 and conversionWindowEnd to 72 (24+48). <br/> Must specify both `behavior.conversionWindowStart` and `behavior.conversionWindowEnd` or neither.",
			InnerField: "conversionWindowEnd",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.conversion-window-start",
			Usage:      "The start of a Conversion Window relative to the exposure date, in hours. This is equivalent to the [Conversion Delay](/app/metrics/legacy/#conversion-delay). <br/> Must specify both `behavior.conversionWindowStart` and `behavior.conversionWindowEnd` or neither.",
			InnerField: "conversionWindowStart",
		},
		&requestflag.InnerFlag[string]{
			Name:       "behavior.goal",
			Usage:      `Allowed values: "increase", "decrease".`,
			InnerField: "goal",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.max-percent-change",
			Usage:      "Maximum percent change to consider uplift significant, as a proportion (e.g. put 0.5 for 50%)",
			InnerField: "maxPercentChange",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.min-percent-change",
			Usage:      "Minimum percent change to consider uplift significant, as a proportion (e.g. put 0.005 for 0.5%)",
			InnerField: "minPercentChange",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.min-sample-size",
			InnerField: "minSampleSize",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "behavior.prior-settings",
			Usage:      "Controls the bayesian prior for the metric. If omitted, organization defaults will be used.",
			InnerField: "priorSettings",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.risk-threshold-danger",
			Usage:      "No longer used. Threshold for Risk to be considered too high, as a proportion (e.g. put 0.0125 for 1.25%). <br/> Must be a non-negative number.",
			InnerField: "riskThresholdDanger",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.risk-threshold-success",
			Usage:      "No longer used. Threshold for Risk to be considered low enough, as a proportion (e.g. put 0.0025 for 0.25%). <br/> Must be a non-negative number and must not be higher than `riskThresholdDanger`.",
			InnerField: "riskThresholdSuccess",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.target-mde",
			Usage:      `The percentage change that you want to reliably detect before ending an experiment, as a proportion (e.g. put 0.1 for 10%). This is used to estimate the "Days Left" for running experiments.`,
			InnerField: "targetMDE",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "behavior.window-settings",
			Usage:      "Controls the conversion window for the metric",
			InnerField: "windowSettings",
		},
	},
	"mixpanel": {
		&requestflag.InnerFlag[string]{
			Name:       "mixpanel.event-name",
			InnerField: "eventName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "mixpanel.user-aggregation",
			InnerField: "userAggregation",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "mixpanel.conditions",
			InnerField: "conditions",
		},
		&requestflag.InnerFlag[string]{
			Name:       "mixpanel.event-value",
			InnerField: "eventValue",
		},
	},
	"sql": {
		&requestflag.InnerFlag[string]{
			Name:       "sql.conversion-sql",
			InnerField: "conversionSQL",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "sql.identifier-types",
			InnerField: "identifierTypes",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sql.denominator-metric-id",
			Usage:      "The metric ID for a [denominator metric for funnel and ratio metrics](/app/metrics/legacy/#denominator-ratio--funnel-metrics)",
			InnerField: "denominatorMetricId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sql.user-aggregation-sql",
			Usage:      "Custom user level aggregation for your metric (default: `SUM(value)`)",
			InnerField: "userAggregationSQL",
		},
	},
	"sql-builder": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "sql-builder.identifier-type-columns",
			InnerField: "identifierTypeColumns",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sql-builder.table-name",
			InnerField: "tableName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sql-builder.timestamp-column-name",
			InnerField: "timestampColumnName",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "sql-builder.conditions",
			InnerField: "conditions",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sql-builder.value-column-name",
			InnerField: "valueColumnName",
		},
	},
})

var metricsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single metric",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleMetricsRetrieve,
	HideHelpCommand: true,
}

var metricsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a metric",
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
			Name:     "behavior",
			BodyPath: "behavior",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of the metric",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "managed-by",
			Usage:    `Where this metric must be managed from. If not set (empty string), it can be managed from anywhere. If set to "api", it can be managed via the API only. Please note that we have deprecated support for setting the managedBy property to "admin". Your existing Legacy Metrics with this value will continue to work, but we suggest migrating to Fact Metrics instead.`,
			BodyPath: "managedBy",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "mixpanel",
			Usage:    "Only use for MixPanel (non-SQL) Data Sources. Only one of `sql`, `sqlBuilder` or `mixpanel` allowed.",
			BodyPath: "mixpanel",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the metric",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "owner",
			Usage:    "The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.",
			BodyPath: "owner",
		},
		&requestflag.Flag[[]string]{
			Name:     "project",
			Usage:    "List of project IDs for projects that can access this metric",
			BodyPath: "projects",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sql",
			Usage:    "Preferred way to define SQL. Only one of `sql`, `sqlBuilder` or `mixpanel` allowed.",
			BodyPath: "sql",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sql-builder",
			Usage:    "An alternative way to specify a SQL metric, rather than a full query. Using `sql` is preferred to `sqlBuilder`. Only one of `sql`, `sqlBuilder` or `mixpanel` allowed",
			BodyPath: "sqlBuilder",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "List of tags",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Type of metric. See [Metrics documentation](/app/metrics/legacy)",
			BodyPath: "type",
		},
	},
	Action:          handleMetricsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"behavior": {
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.cap",
			Usage:      "(deprecated, use cappingSettings instead) This should be non-negative",
			InnerField: "cap",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "behavior.capping",
			Usage:      "(deprecated, use cappingSettings instead) Used in conjunction with `capValue` to set the capping (winsorization). Do not specify or set to null for no capping. \"absolute\" will cap user values at the `capValue` if it is greater than 0. \"percentile\" will cap user values at the percentile of user values in an experiment using the `capValue` for the percentile, if greater than 0. <br/>  If `behavior.capping` is non-null, you must specify `behavior.capValue`.",
			InnerField: "capping",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "behavior.capping-settings",
			Usage:      "Controls how outliers are handled",
			InnerField: "cappingSettings",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.cap-value",
			Usage:      "(deprecated, use cappingSettings instead) This should be non-negative. <br/> Must specify `behavior.capping` when setting `behavior.capValue`.",
			InnerField: "capValue",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.conversion-window-end",
			Usage:      "The end of a [Conversion Window](/app/metrics/legacy/#conversion-window) relative to the exposure date, in hours. This is equivalent to the [Conversion Delay](/app/metrics/legacy/#conversion-delay) + Conversion Window Hours settings in the UI. In other words, if you want a 48 hour window starting after 24 hours, you would set conversionWindowStart to 24 and conversionWindowEnd to 72 (24+48). <br/> Must specify both `behavior.conversionWindowStart` and `behavior.conversionWindowEnd` or neither.",
			InnerField: "conversionWindowEnd",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.conversion-window-start",
			Usage:      "The start of a Conversion Window relative to the exposure date, in hours. This is equivalent to the [Conversion Delay](/app/metrics/legacy/#conversion-delay). <br/> Must specify both `behavior.conversionWindowStart` and `behavior.conversionWindowEnd` or neither.",
			InnerField: "conversionWindowStart",
		},
		&requestflag.InnerFlag[string]{
			Name:       "behavior.goal",
			Usage:      `Allowed values: "increase", "decrease".`,
			InnerField: "goal",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.max-percent-change",
			Usage:      "Maximum percent change to consider uplift significant, as a proportion (e.g. put 0.5 for 50%)",
			InnerField: "maxPercentChange",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.min-percent-change",
			Usage:      "Minimum percent change to consider uplift significant, as a proportion (e.g. put 0.005 for 0.5%)",
			InnerField: "minPercentChange",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.min-sample-size",
			InnerField: "minSampleSize",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "behavior.prior-settings",
			Usage:      "Controls the bayesian prior for the metric. If omitted, organization defaults will be used.",
			InnerField: "priorSettings",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.risk-threshold-danger",
			Usage:      "No longer used. Threshold for Risk to be considered too high, as a proportion (e.g. put 0.0125 for 1.25%). <br/> Must be a non-negative number.",
			InnerField: "riskThresholdDanger",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.risk-threshold-success",
			Usage:      "No longer used. Threshold for Risk to be considered low enough, as a proportion (e.g. put 0.0025 for 0.25%). <br/> Must be a non-negative number and must not be higher than `riskThresholdDanger`.",
			InnerField: "riskThresholdSuccess",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "behavior.target-mde",
			Usage:      `The percentage change that you want to reliably detect before ending an experiment, as a proportion (e.g. put 0.1 for 10%). This is used to estimate the "Days Left" for running experiments.`,
			InnerField: "targetMDE",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "behavior.window-settings",
			Usage:      "Controls the conversion window for the metric",
			InnerField: "windowSettings",
		},
	},
	"mixpanel": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "mixpanel.conditions",
			InnerField: "conditions",
		},
		&requestflag.InnerFlag[string]{
			Name:       "mixpanel.event-name",
			InnerField: "eventName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "mixpanel.event-value",
			InnerField: "eventValue",
		},
		&requestflag.InnerFlag[string]{
			Name:       "mixpanel.user-aggregation",
			InnerField: "userAggregation",
		},
	},
	"sql": {
		&requestflag.InnerFlag[string]{
			Name:       "sql.conversion-sql",
			InnerField: "conversionSQL",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sql.denominator-metric-id",
			Usage:      "The metric ID for a [denominator metric for funnel and ratio metrics](/app/metrics/legacy/#denominator-ratio--funnel-metrics)",
			InnerField: "denominatorMetricId",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "sql.identifier-types",
			InnerField: "identifierTypes",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sql.user-aggregation-sql",
			Usage:      "Custom user level aggregation for your metric (default: `SUM(value)`)",
			InnerField: "userAggregationSQL",
		},
	},
	"sql-builder": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "sql-builder.conditions",
			InnerField: "conditions",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "sql-builder.identifier-type-columns",
			InnerField: "identifierTypeColumns",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sql-builder.table-name",
			InnerField: "tableName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sql-builder.timestamp-column-name",
			InnerField: "timestampColumnName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sql-builder.value-column-name",
			InnerField: "valueColumnName",
		},
	},
})

var metricsList = cli.Command{
	Name:    "list",
	Usage:   "Get all metrics",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "datasource-id",
			Usage:     "Filter by Data Source",
			QueryPath: "datasourceId",
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
	Action:          handleMetricsList,
	HideHelpCommand: true,
}

var metricsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a metric",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleMetricsDelete,
	HideHelpCommand: true,
}

func handleMetricsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.MetricNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Metrics.New(ctx, params, options...)
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
		Title:          "metrics create",
		Transform:      transform,
	})
}

func handleMetricsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Metrics.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "metrics retrieve",
		Transform:      transform,
	})
}

func handleMetricsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.MetricUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Metrics.Update(
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
		Title:          "metrics update",
		Transform:      transform,
	})
}

func handleMetricsList(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.MetricListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Metrics.List(ctx, params, options...)
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
		Title:          "metrics list",
		Transform:      transform,
	})
}

func handleMetricsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Metrics.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "metrics delete",
		Transform:      transform,
	})
}
