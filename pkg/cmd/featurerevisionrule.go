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

var featuresRevisionsRulesUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Patches fields on an existing rule (identified by `ruleId`). The rule `type`\ncannot be changed. Scope can be updated via `allEnvironments` / `environments`\npatch fields.",
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
		&requestflag.Flag[map[string]any]{
			Name:     "rule",
			Required: true,
			BodyPath: "rule",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "ramp-schedule",
			BodyPath: "rampSchedule",
		},
		&requestflag.Flag[string]{
			Name:     "revision-comment",
			BodyPath: "revisionComment",
		},
		&requestflag.Flag[string]{
			Name:     "revision-title",
			BodyPath: "revisionTitle",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "schedule",
			BodyPath: "schedule",
		},
	},
	Action:          handleFeaturesRevisionsRulesUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"rule": {
		&requestflag.InnerFlag[bool]{
			Name:       "rule.all-environments",
			InnerField: "allEnvironments",
		},
		&requestflag.InnerFlag[string]{
			Name:       "rule.condition",
			InnerField: "condition",
		},
		&requestflag.InnerFlag[string]{
			Name:       "rule.control-value",
			InnerField: "controlValue",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "rule.coverage",
			InnerField: "coverage",
		},
		&requestflag.InnerFlag[string]{
			Name:       "rule.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "rule.enabled",
			InnerField: "enabled",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "rule.environments",
			InnerField: "environments",
		},
		&requestflag.InnerFlag[string]{
			Name:       "rule.experiment-id",
			InnerField: "experimentId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "rule.hash-attribute",
			InnerField: "hashAttribute",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "rule.prerequisites",
			InnerField: "prerequisites",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "rule.saved-groups",
			InnerField: "savedGroups",
		},
		&requestflag.InnerFlag[any]{
			Name:       "rule.schedule-rules",
			InnerField: "scheduleRules",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "rule.schedule-type",
			Usage:      `Allowed values: "none", "schedule", "ramp".`,
			InnerField: "scheduleType",
		},
		&requestflag.InnerFlag[string]{
			Name:       "rule.seed",
			InnerField: "seed",
		},
		&requestflag.InnerFlag[string]{
			Name:       "rule.type",
			Usage:      `Allowed values: "force", "rollout", "experiment-ref", "safe-rollout".`,
			InnerField: "type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "rule.value",
			InnerField: "value",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "rule.variations",
			InnerField: "variations",
		},
		&requestflag.InnerFlag[string]{
			Name:       "rule.variation-value",
			InnerField: "variationValue",
		},
	},
	"ramp-schedule": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "ramp-schedule.end-actions",
			InnerField: "endActions",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "ramp-schedule.end-condition",
			InnerField: "endCondition",
		},
		&requestflag.InnerFlag[string]{
			Name:       "ramp-schedule.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "ramp-schedule.start-date",
			InnerField: "startDate",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "ramp-schedule.steps",
			InnerField: "steps",
		},
		&requestflag.InnerFlag[string]{
			Name:       "ramp-schedule.template-id",
			InnerField: "templateId",
		},
	},
	"schedule": {
		&requestflag.InnerFlag[*string]{
			Name:       "schedule.end-date",
			InnerField: "endDate",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "schedule.start-date",
			InnerField: "startDate",
		},
	},
})

var featuresRevisionsRulesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Removes the rule from the revision. Any pending ramp actions for this rule are\nalso cleared.",
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
	Action:          handleFeaturesRevisionsRulesDelete,
	HideHelpCommand: true,
}

var featuresRevisionsRulesAdd = requestflag.WithInnerFlags(cli.Command{
	Name:    "add",
	Usage:   "Appends a new rule to the revision's rule list. Supply `allEnvironments: true`\nto target all environments, or `environments: [...]` to scope to specific ones.\nUse `rampSchedule` for ramp configuration or `schedule` for a simple start/end\nwindow.",
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
			Name:     "rule",
			Required: true,
			BodyPath: "rule",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "ramp-schedule",
			BodyPath: "rampSchedule",
		},
		&requestflag.Flag[string]{
			Name:     "revision-comment",
			BodyPath: "revisionComment",
		},
		&requestflag.Flag[string]{
			Name:     "revision-title",
			BodyPath: "revisionTitle",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "schedule",
			BodyPath: "schedule",
		},
	},
	Action:          handleFeaturesRevisionsRulesAdd,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"ramp-schedule": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "ramp-schedule.end-actions",
			InnerField: "endActions",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "ramp-schedule.end-condition",
			InnerField: "endCondition",
		},
		&requestflag.InnerFlag[string]{
			Name:       "ramp-schedule.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "ramp-schedule.start-date",
			InnerField: "startDate",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "ramp-schedule.steps",
			InnerField: "steps",
		},
		&requestflag.InnerFlag[string]{
			Name:       "ramp-schedule.template-id",
			InnerField: "templateId",
		},
	},
	"schedule": {
		&requestflag.InnerFlag[*string]{
			Name:       "schedule.end-date",
			InnerField: "endDate",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "schedule.start-date",
			InnerField: "startDate",
		},
	},
})

var featuresRevisionsRulesReorder = cli.Command{
	Name:    "reorder",
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
	Action:          handleFeaturesRevisionsRulesReorder,
	HideHelpCommand: true,
}

func handleFeaturesRevisionsRulesUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.FeatureRevisionRuleUpdateParams{
		ID:      cmd.Value("id").(string),
		Version: growthbook.FeatureRevisionRuleUpdateParamsVersionUnion(cmd.Value("version").(any)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Features.Revisions.Rules.Update(
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
		Title:          "features:revisions:rules update",
		Transform:      transform,
	})
}

func handleFeaturesRevisionsRulesDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.FeatureRevisionRuleDeleteParams{
		ID:      cmd.Value("id").(string),
		Version: growthbook.FeatureRevisionRuleDeleteParamsVersionUnion(cmd.Value("version").(any)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Features.Revisions.Rules.Delete(
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
		Title:          "features:revisions:rules delete",
		Transform:      transform,
	})
}

func handleFeaturesRevisionsRulesAdd(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.FeatureRevisionRuleAddParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Features.Revisions.Rules.Add(
		ctx,
		growthbook.FeatureRevisionRuleAddParamsVersionUnion(cmd.Value("version").(any)),
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
		Title:          "features:revisions:rules add",
		Transform:      transform,
	})
}

func handleFeaturesRevisionsRulesReorder(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.FeatureRevisionRuleReorderParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Features.Revisions.Rules.Reorder(
		ctx,
		growthbook.FeatureRevisionRuleReorderParamsVersionUnion(cmd.Value("version").(any)),
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
		Title:          "features:revisions:rules reorder",
		Transform:      transform,
	})
}
