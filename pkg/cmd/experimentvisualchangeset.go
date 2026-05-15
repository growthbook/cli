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

var experimentsVisualChangesetsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a visual changeset for an experiment",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "editor-url",
			Usage:    "URL of the page opened in the visual editor when creating this changeset",
			Required: true,
			BodyPath: "editorUrl",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "url-pattern",
			Usage:    "URL patterns that determine which pages this visual changeset applies to",
			Required: true,
			BodyPath: "urlPatterns",
		},
	},
	Action:          handleExperimentsVisualChangesetsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"url-pattern": {
		&requestflag.InnerFlag[string]{
			Name:       "url-pattern.pattern",
			InnerField: "pattern",
		},
		&requestflag.InnerFlag[string]{
			Name:       "url-pattern.type",
			Usage:      `Allowed values: "simple", "regex".`,
			InnerField: "type",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "url-pattern.include",
			InnerField: "include",
		},
	},
})

var experimentsVisualChangesetsList = cli.Command{
	Name:    "list",
	Usage:   "Get all visual changesets",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The experiment id the visual changesets belong to",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleExperimentsVisualChangesetsList,
	HideHelpCommand: true,
}

func handleExperimentsVisualChangesetsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.ExperimentVisualChangesetNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.VisualChangesets.New(
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
		Title:          "experiments:visual-changesets create",
		Transform:      transform,
	})
}

func handleExperimentsVisualChangesetsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Experiments.VisualChangesets.List(ctx, cmd.Value("id").(string), options...)
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
		Title:          "experiments:visual-changesets list",
		Transform:      transform,
	})
}
