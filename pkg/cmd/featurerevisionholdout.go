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

var featuresRevisionsHoldoutCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Set holdout in a draft revision",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[any]{
			Name:      "version",
			Required:  true,
			PathParam: "version",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "holdout",
			Required: true,
			BodyPath: "holdout",
		},
		&requestflag.Flag[string]{
			Name:     "revision-comment",
			BodyPath: "revisionComment",
		},
		&requestflag.Flag[string]{
			Name:     "revision-title",
			BodyPath: "revisionTitle",
		},
	},
	Action:          handleFeaturesRevisionsHoldoutCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"holdout": {
		&requestflag.InnerFlag[string]{
			Name:       "holdout.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "holdout.value",
			InnerField: "value",
		},
	},
})

func handleFeaturesRevisionsHoldoutCreate(ctx context.Context, cmd *cli.Command) error {
	client := growthbook.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("version") && len(unusedArgs) > 0 {
		cmd.Set("version", unusedArgs[0])
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

	params := growthbook.FeatureRevisionHoldoutNewParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Features.Revisions.Holdout.New(
		ctx,
		growthbook.FeatureRevisionHoldoutNewParamsVersionUnion(cmd.Value("version").(any)),
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
		Title:          "features:revisions:holdout create",
		Transform:      transform,
	})
}
