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

var productAnalyticsCreateDataSourceExploration = requestflag.WithInnerFlags(cli.Command{
	Name:    "create-data-source-exploration",
	Usage:   "Create a Data Source based visualization",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "chart-type",
			Usage:    `Allowed values: "line", "area", "timeseries-table", "table", "bar", "stackedBar", "horizontalBar", "stackedHorizontalBar", "bigNumber".`,
			Required: true,
			BodyPath: "chartType",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "dataset",
			Required: true,
			BodyPath: "dataset",
		},
		&requestflag.Flag[string]{
			Name:     "datasource",
			Usage:    "ID of the datasource to query",
			Required: true,
			BodyPath: "datasource",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "date-range",
			Required: true,
			BodyPath: "dateRange",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "dimension",
			Required: true,
			BodyPath: "dimensions",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "data_source".`,
			Default:  "data_source",
			Const:    true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:      "cache",
			Usage:     "Controls cache behavior for this exploration: `preferred` (default) returns a cached result if one exists, otherwise runs a new query; `never` always runs a new query, ignoring any cached results; `required` only returns a cached result, if none exists returns exploration: null with a message",
			QueryPath: "cache",
		},
		&requestflag.Flag[string]{
			Name:     "show-as",
			Usage:    `Allowed values: "total", "per_unit".`,
			BodyPath: "showAs",
		},
	},
	Action:          handleProductAnalyticsCreateDataSourceExploration,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"dataset": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "dataset.column-types",
			InnerField: "columnTypes",
		},
		&requestflag.InnerFlag[string]{
			Name:       "dataset.path",
			InnerField: "path",
		},
		&requestflag.InnerFlag[string]{
			Name:       "dataset.table",
			InnerField: "table",
		},
		&requestflag.InnerFlag[string]{
			Name:       "dataset.timestamp-column",
			InnerField: "timestampColumn",
		},
		&requestflag.InnerFlag[string]{
			Name:       "dataset.type",
			Usage:      `Allowed values: "data_source".`,
			InnerField: "type",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "dataset.values",
			InnerField: "values",
		},
	},
	"date-range": {
		&requestflag.InnerFlag[string]{
			Name:       "date-range.predefined",
			Usage:      `Allowed values: "today", "last7Days", "last30Days", "last90Days", "customLookback", "customDateRange".`,
			InnerField: "predefined",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "date-range.end-date",
			InnerField: "endDate",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "date-range.lookback-unit",
			Usage:      `Allowed values: "hour", "day", "week", "month".`,
			InnerField: "lookbackUnit",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "date-range.lookback-value",
			InnerField: "lookbackValue",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "date-range.start-date",
			InnerField: "startDate",
		},
	},
})

var productAnalyticsCreateMetricExploration = requestflag.WithInnerFlags(cli.Command{
	Name:    "create-metric-exploration",
	Usage:   "Create a Metric based visualization",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "chart-type",
			Usage:    `Allowed values: "line", "area", "timeseries-table", "table", "bar", "stackedBar", "horizontalBar", "stackedHorizontalBar", "bigNumber".`,
			Required: true,
			BodyPath: "chartType",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "dataset",
			Required: true,
			BodyPath: "dataset",
		},
		&requestflag.Flag[string]{
			Name:     "datasource",
			Usage:    "ID of the datasource to query",
			Required: true,
			BodyPath: "datasource",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "date-range",
			Required: true,
			BodyPath: "dateRange",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "dimension",
			Required: true,
			BodyPath: "dimensions",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "metric".`,
			Default:  "metric",
			Const:    true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:      "cache",
			Usage:     "Controls cache behavior for this exploration: `preferred` (default) returns a cached result if one exists, otherwise runs a new query; `never` always runs a new query, ignoring any cached results; `required` only returns a cached result, if none exists returns exploration: null with a message",
			QueryPath: "cache",
		},
		&requestflag.Flag[string]{
			Name:     "show-as",
			Usage:    `Allowed values: "total", "per_unit".`,
			BodyPath: "showAs",
		},
	},
	Action:          handleProductAnalyticsCreateMetricExploration,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"dataset": {
		&requestflag.InnerFlag[string]{
			Name:       "dataset.type",
			Usage:      `Allowed values: "metric".`,
			InnerField: "type",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "dataset.values",
			InnerField: "values",
		},
	},
	"date-range": {
		&requestflag.InnerFlag[string]{
			Name:       "date-range.predefined",
			Usage:      `Allowed values: "today", "last7Days", "last30Days", "last90Days", "customLookback", "customDateRange".`,
			InnerField: "predefined",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "date-range.end-date",
			InnerField: "endDate",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "date-range.lookback-unit",
			Usage:      `Allowed values: "hour", "day", "week", "month".`,
			InnerField: "lookbackUnit",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "date-range.lookback-value",
			InnerField: "lookbackValue",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "date-range.start-date",
			InnerField: "startDate",
		},
	},
})

var productAnalyticsRunFactTableExploration = requestflag.WithInnerFlags(cli.Command{
	Name:    "run-fact-table-exploration",
	Usage:   "Run a Fact Table based visualization",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "chart-type",
			Usage:    `Allowed values: "line", "area", "timeseries-table", "table", "bar", "stackedBar", "horizontalBar", "stackedHorizontalBar", "bigNumber".`,
			Required: true,
			BodyPath: "chartType",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "dataset",
			Required: true,
			BodyPath: "dataset",
		},
		&requestflag.Flag[string]{
			Name:     "datasource",
			Usage:    "ID of the datasource to query",
			Required: true,
			BodyPath: "datasource",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "date-range",
			Required: true,
			BodyPath: "dateRange",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "dimension",
			Required: true,
			BodyPath: "dimensions",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "fact_table".`,
			Default:  "fact_table",
			Const:    true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:      "cache",
			Usage:     "Controls cache behavior for this exploration: `preferred` (default) returns a cached result if one exists, otherwise runs a new query; `never` always runs a new query, ignoring any cached results; `required` only returns a cached result, if none exists returns exploration: null with a message",
			QueryPath: "cache",
		},
		&requestflag.Flag[string]{
			Name:     "show-as",
			Usage:    `Allowed values: "total", "per_unit".`,
			BodyPath: "showAs",
		},
	},
	Action:          handleProductAnalyticsRunFactTableExploration,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"dataset": {
		&requestflag.InnerFlag[*string]{
			Name:       "dataset.fact-table-id",
			InnerField: "factTableId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "dataset.type",
			Usage:      `Allowed values: "fact_table".`,
			InnerField: "type",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "dataset.values",
			InnerField: "values",
		},
	},
	"date-range": {
		&requestflag.InnerFlag[string]{
			Name:       "date-range.predefined",
			Usage:      `Allowed values: "today", "last7Days", "last30Days", "last90Days", "customLookback", "customDateRange".`,
			InnerField: "predefined",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "date-range.end-date",
			InnerField: "endDate",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "date-range.lookback-unit",
			Usage:      `Allowed values: "hour", "day", "week", "month".`,
			InnerField: "lookbackUnit",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "date-range.lookback-value",
			InnerField: "lookbackValue",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "date-range.start-date",
			InnerField: "startDate",
		},
	},
})

func handleProductAnalyticsCreateDataSourceExploration(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.ProductAnalyticsNewDataSourceExplorationParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ProductAnalytics.NewDataSourceExploration(ctx, params, options...)
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
		Title:          "product-analytics create-data-source-exploration",
		Transform:      transform,
	})
}

func handleProductAnalyticsCreateMetricExploration(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.ProductAnalyticsNewMetricExplorationParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ProductAnalytics.NewMetricExploration(ctx, params, options...)
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
		Title:          "product-analytics create-metric-exploration",
		Transform:      transform,
	})
}

func handleProductAnalyticsRunFactTableExploration(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.ProductAnalyticsRunFactTableExplorationParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ProductAnalytics.RunFactTableExploration(ctx, params, options...)
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
		Title:          "product-analytics run-fact-table-exploration",
		Transform:      transform,
	})
}
