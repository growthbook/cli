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

var visualChangesetsVisualChangeCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a visual change for a visual changeset",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "variation",
			Required: true,
			BodyPath: "variation",
		},
		&requestflag.Flag[string]{
			Name:     "id",
			BodyPath: "id",
		},
		&requestflag.Flag[string]{
			Name:     "css",
			BodyPath: "css",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "dom-mutation",
			BodyPath: "domMutations",
		},
		&requestflag.Flag[string]{
			Name:     "js",
			BodyPath: "js",
		},
	},
	Action:          handleVisualChangesetsVisualChangeCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"dom-mutation": {
		&requestflag.InnerFlag[string]{
			Name:       "dom-mutation.action",
			Usage:      `Allowed values: "append", "set", "remove".`,
			InnerField: "action",
		},
		&requestflag.InnerFlag[string]{
			Name:       "dom-mutation.attribute",
			InnerField: "attribute",
		},
		&requestflag.InnerFlag[string]{
			Name:       "dom-mutation.selector",
			InnerField: "selector",
		},
		&requestflag.InnerFlag[string]{
			Name:       "dom-mutation.insert-before-selector",
			InnerField: "insertBeforeSelector",
		},
		&requestflag.InnerFlag[string]{
			Name:       "dom-mutation.parent-selector",
			InnerField: "parentSelector",
		},
		&requestflag.InnerFlag[string]{
			Name:       "dom-mutation.value",
			InnerField: "value",
		},
	},
})

var visualChangesetsVisualChangeUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a visual change for a visual changeset",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "visual-change-id",
			Usage:     "Specify a specific visual change",
			Required:  true,
			PathParam: "visualChangeId",
		},
		&requestflag.Flag[string]{
			Name:     "id",
			BodyPath: "id",
		},
		&requestflag.Flag[string]{
			Name:     "css",
			BodyPath: "css",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "dom-mutation",
			BodyPath: "domMutations",
		},
		&requestflag.Flag[string]{
			Name:     "js",
			BodyPath: "js",
		},
		&requestflag.Flag[string]{
			Name:     "variation",
			BodyPath: "variation",
		},
	},
	Action:          handleVisualChangesetsVisualChangeUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"dom-mutation": {
		&requestflag.InnerFlag[string]{
			Name:       "dom-mutation.action",
			Usage:      `Allowed values: "append", "set", "remove".`,
			InnerField: "action",
		},
		&requestflag.InnerFlag[string]{
			Name:       "dom-mutation.attribute",
			InnerField: "attribute",
		},
		&requestflag.InnerFlag[string]{
			Name:       "dom-mutation.selector",
			InnerField: "selector",
		},
		&requestflag.InnerFlag[string]{
			Name:       "dom-mutation.insert-before-selector",
			InnerField: "insertBeforeSelector",
		},
		&requestflag.InnerFlag[string]{
			Name:       "dom-mutation.parent-selector",
			InnerField: "parentSelector",
		},
		&requestflag.InnerFlag[string]{
			Name:       "dom-mutation.value",
			InnerField: "value",
		},
	},
})

func handleVisualChangesetsVisualChangeCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.VisualChangesetVisualChangeNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.VisualChangesets.VisualChange.New(
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
		Title:          "visual-changesets:visual-change create",
		Transform:      transform,
	})
}

func handleVisualChangesetsVisualChangeUpdate(ctx context.Context, cmd *cli.Command) error {
	client := growthbook.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("visual-change-id") && len(unusedArgs) > 0 {
		cmd.Set("visual-change-id", unusedArgs[0])
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

	params := growthbook.VisualChangesetVisualChangeUpdateParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.VisualChangesets.VisualChange.Update(
		ctx,
		cmd.Value("visual-change-id").(string),
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
		Title:          "visual-changesets:visual-change update",
		Transform:      transform,
	})
}
