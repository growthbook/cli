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

var featuresRevisionsRulesRampScheduleUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Set ramp schedule for a rule",
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
		&requestflag.Flag[string]{
			Name:      "rule-id",
			Required:  true,
			PathParam: "ruleId",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "end-action",
			BodyPath: "endActions",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "end-condition",
			BodyPath: "endCondition",
		},
		&requestflag.Flag[string]{
			Name:     "environment",
			BodyPath: "environment",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "revision-comment",
			BodyPath: "revisionComment",
		},
		&requestflag.Flag[string]{
			Name:     "revision-title",
			BodyPath: "revisionTitle",
		},
		&requestflag.Flag[*string]{
			Name:     "start-date",
			BodyPath: "startDate",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "step",
			BodyPath: "steps",
		},
		&requestflag.Flag[string]{
			Name:     "template-id",
			BodyPath: "templateId",
		},
	},
	Action:          handleFeaturesRevisionsRulesRampScheduleUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"end-action": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "end-action.patch",
			InnerField: "patch",
		},
		&requestflag.InnerFlag[string]{
			Name:       "end-action.target-id",
			InnerField: "targetId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "end-action.target-type",
			Usage:      `Allowed values: "feature-rule".`,
			InnerField: "targetType",
		},
	},
	"end-condition": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "end-condition.trigger",
			InnerField: "trigger",
		},
	},
	"step": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "step.trigger",
			InnerField: "trigger",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "step.actions",
			InnerField: "actions",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "step.approval-notes",
			InnerField: "approvalNotes",
		},
	},
})

var featuresRevisionsRulesRampScheduleDelete = cli.Command{
	Name:    "delete",
	Usage:   "Remove ramp schedule from a rule",
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
		&requestflag.Flag[string]{
			Name:      "rule-id",
			Required:  true,
			PathParam: "ruleId",
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
	Action:          handleFeaturesRevisionsRulesRampScheduleDelete,
	HideHelpCommand: true,
}

func handleFeaturesRevisionsRulesRampScheduleUpdate(ctx context.Context, cmd *cli.Command) error {
	client := growthbook.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("rule-id") && len(unusedArgs) > 0 {
		cmd.Set("rule-id", unusedArgs[0])
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

	params := growthbook.FeatureRevisionRuleRampScheduleUpdateParams{
		ID:      cmd.Value("id").(string),
		Version: growthbook.FeatureRevisionRuleRampScheduleUpdateParamsVersionUnion(cmd.Value("version").(any)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Features.Revisions.Rules.RampSchedule.Update(
		ctx,
		cmd.Value("rule-id").(string),
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
		Title:          "features:revisions:rules:ramp-schedule update",
		Transform:      transform,
	})
}

func handleFeaturesRevisionsRulesRampScheduleDelete(ctx context.Context, cmd *cli.Command) error {
	client := growthbook.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("rule-id") && len(unusedArgs) > 0 {
		cmd.Set("rule-id", unusedArgs[0])
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

	params := growthbook.FeatureRevisionRuleRampScheduleDeleteParams{
		ID:      cmd.Value("id").(string),
		Version: growthbook.FeatureRevisionRuleRampScheduleDeleteParamsVersionUnion(cmd.Value("version").(any)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Features.Revisions.Rules.RampSchedule.Delete(
		ctx,
		cmd.Value("rule-id").(string),
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
		Title:          "features:revisions:rules:ramp-schedule delete",
		Transform:      transform,
	})
}
