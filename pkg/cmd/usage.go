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

var usageRetrieveMetrics = cli.Command{
	Name:    "retrieve-metrics",
	Usage:   "Returns usage information for one or more legacy or fact metrics, showing which\nexperiments use each metric and some usage statistics. If a metric is part of a\nmetric group, then usage of that metric group counts as usage of all metrics in\nthe group. Warning: only includes experiments that you have access to! If you do\nnot have admin access or read access to experiments across all projects, this\nendpoint may not return the latest usage data across all experiments.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "ids",
			Usage:     "List of comma-separated metric IDs (both fact and legacy) to get usage for, e.g. ids=met_123,fact_456",
			Required:  true,
			QueryPath: "ids",
		},
	},
	Action:          handleUsageRetrieveMetrics,
	HideHelpCommand: true,
}

func handleUsageRetrieveMetrics(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.UsageGetMetricsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Usage.GetMetrics(ctx, params, options...)
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
		Title:          "usage retrieve-metrics",
		Transform:      transform,
	})
}
