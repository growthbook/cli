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

var rampSchedulesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new ramp schedule, optionally attaching it to a published feature\nrule.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "end-action",
			Usage:    "Actions applied when the ramp completes. `targetId` and `patch.ruleId` are auto-injected when `featureId`+`ruleId` are provided.",
			BodyPath: "endActions",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "end-condition",
			Usage:    "Optional hard deadline",
			BodyPath: "endCondition",
		},
		&requestflag.Flag[string]{
			Name:     "environment",
			Usage:    "Deprecated. Legacy disambiguator for pre-v2 rules whose `ruleId` could repeat across envs. Omit on new schedules — the resolver uses `rule.id` directly.",
			BodyPath: "environment",
		},
		&requestflag.Flag[string]{
			Name:     "feature-id",
			Usage:    "Feature that anchors this schedule. Required when `ruleId` is set.",
			BodyPath: "featureId",
		},
		&requestflag.Flag[string]{
			Name:     "rule-id",
			Usage:    "Rule to attach as the initial target. Requires `featureId`. Post-v2 `rule.id` is uniquely sufficient; `environment` is optional and deprecated.",
			BodyPath: "ruleId",
		},
		&requestflag.Flag[any]{
			Name:     "start-date",
			Usage:    "When to start. Absent/null = immediately on start action.",
			BodyPath: "startDate",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "step",
			Usage:    "Ordered ramp steps. When `featureId`+`ruleId` are provided,\n`targetId` and `patch.ruleId` in actions are auto-injected — only\nsupply the patch fields you want to change.",
			BodyPath: "steps",
		},
		&requestflag.Flag[string]{
			Name:     "template-id",
			Usage:    "Load steps and endActions from a saved template (featureId+ruleId must also be set for auto-injection)",
			BodyPath: "templateId",
		},
	},
	Action:          handleRampSchedulesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"end-action": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "end-action.patch",
			Usage:      "Sparse patch — only fields present are applied; absent fields accumulate from previous steps",
			InnerField: "patch",
		},
		&requestflag.InnerFlag[string]{
			Name:       "end-action.target-id",
			Usage:      "Auto-injected when featureId+ruleId+environment are provided",
			InnerField: "targetId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "end-action.target-type",
			Usage:      "Omit when using featureId+ruleId+environment (auto-injected)",
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

var rampSchedulesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single rampSchedule",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleRampSchedulesRetrieve,
	HideHelpCommand: true,
}

var rampSchedulesUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates the name, steps, endActions, startDate, or endCondition of a ramp\nschedule.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
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
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "start-date",
			BodyPath: "startDate",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "step",
			BodyPath: "steps",
		},
	},
	Action:          handleRampSchedulesUpdate,
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

var rampSchedulesList = cli.Command{
	Name:    "list",
	Usage:   "Returns all ramp schedules for the organization, with optional filters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "feature-id",
			QueryPath: "featureId",
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
			Name:      "status",
			Usage:     "Filter by schedule status",
			QueryPath: "status",
		},
	},
	Action:          handleRampSchedulesList,
	HideHelpCommand: true,
}

var rampSchedulesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently deletes a ramp schedule. This does not undo any rule patches that\nwere already applied by completed steps.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleRampSchedulesDelete,
	HideHelpCommand: true,
}

func handleRampSchedulesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.RampScheduleNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.RampSchedules.New(ctx, params, options...)
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
		Title:          "ramp-schedules create",
		Transform:      transform,
	})
}

func handleRampSchedulesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.RampSchedules.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "ramp-schedules retrieve",
		Transform:      transform,
	})
}

func handleRampSchedulesUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.RampScheduleUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.RampSchedules.Update(
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
		Title:          "ramp-schedules update",
		Transform:      transform,
	})
}

func handleRampSchedulesList(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.RampScheduleListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.RampSchedules.List(ctx, params, options...)
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
		Title:          "ramp-schedules list",
		Transform:      transform,
	})
}

func handleRampSchedulesDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.RampSchedules.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "ramp-schedules delete",
		Transform:      transform,
	})
}
