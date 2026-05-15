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

var featuresRevisionsRulesReorderCreate = cli.Command{
	Name:    "create",
	Usage:   "Replaces the flat global rule order. `ruleIds` must contain **exactly** the set\nof all existing rule IDs in the revision — no additions, omissions, or\nduplicates.",
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
		&requestflag.Flag[[]string]{
			Name:     "rule-id",
			Required: true,
			BodyPath: "ruleIds",
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
	Action:          handleFeaturesRevisionsRulesReorderCreate,
	HideHelpCommand: true,
}

func handleFeaturesRevisionsRulesReorderCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.FeatureRevisionRuleReorderNewParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Features.Revisions.Rules.Reorder.New(
		ctx,
		growthbook.FeatureRevisionRuleReorderNewParamsVersionUnion(cmd.Value("version").(any)),
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
		Title:          "features:revisions:rules:reorder create",
		Transform:      transform,
	})
}
