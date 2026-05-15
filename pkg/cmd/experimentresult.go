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

var experimentsResultsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get results for an experiment",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "dimension",
			QueryPath: "dimension",
		},
		&requestflag.Flag[string]{
			Name:      "phase",
			QueryPath: "phase",
		},
	},
	Action:          handleExperimentsResultsRetrieve,
	HideHelpCommand: true,
}

var experimentsResultsList = cli.Command{
	Name:    "list",
	Usage:   "Returns the latest non-dimension snapshot for each experiment matching the\nfilters. Use this to scan results across a portfolio in one call.",
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
	Action:          handleExperimentsResultsList,
	HideHelpCommand: true,
}

func handleExperimentsResultsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.ExperimentResultGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.Results.Get(
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
		Title:          "experiments:results retrieve",
		Transform:      transform,
	})
}

func handleExperimentsResultsList(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.ExperimentResultListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.Results.List(ctx, params, options...)
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
		Title:          "experiments:results list",
		Transform:      transform,
	})
}
