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

var visualChangesetsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single visual changeset",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[int64]{
			Name:      "include-experiment",
			Usage:     "Include the associated experiment in payload",
			QueryPath: "includeExperiment",
		},
	},
	Action:          handleVisualChangesetsRetrieve,
	HideHelpCommand: true,
}

var visualChangesetsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a visual changeset",
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
			BodyPath: "editorUrl",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "url-pattern",
			Usage:    "URL patterns that determine which pages this visual changeset applies to",
			BodyPath: "urlPatterns",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "visual-change",
			BodyPath: "visualChanges",
		},
	},
	Action:          handleVisualChangesetsUpdate,
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
	"visual-change": {
		&requestflag.InnerFlag[string]{
			Name:       "visual-change.variation",
			InnerField: "variation",
		},
		&requestflag.InnerFlag[string]{
			Name:       "visual-change.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "visual-change.css",
			InnerField: "css",
		},
		&requestflag.InnerFlag[string]{
			Name:       "visual-change.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "visual-change.dom-mutations",
			InnerField: "domMutations",
		},
		&requestflag.InnerFlag[string]{
			Name:       "visual-change.js",
			InnerField: "js",
		},
	},
})

func handleVisualChangesetsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.VisualChangesetGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.VisualChangesets.Get(
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
		Title:          "visual-changesets retrieve",
		Transform:      transform,
	})
}

func handleVisualChangesetsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.VisualChangesetUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.VisualChangesets.Update(
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
		Title:          "visual-changesets update",
		Transform:      transform,
	})
}
