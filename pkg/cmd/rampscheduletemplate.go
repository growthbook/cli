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

var rampScheduleTemplatesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a single rampScheduleTemplate",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "step",
			Required: true,
			BodyPath: "steps",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "end-patch",
			BodyPath: "endPatch",
		},
		&requestflag.Flag[bool]{
			Name:     "official",
			BodyPath: "official",
		},
	},
	Action:          handleRampScheduleTemplatesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"step": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "step.actions",
			InnerField: "actions",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "step.trigger",
			InnerField: "trigger",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "step.approval-notes",
			InnerField: "approvalNotes",
		},
	},
	"end-patch": {
		&requestflag.InnerFlag[string]{
			Name:       "end-patch.condition",
			InnerField: "condition",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "end-patch.coverage",
			InnerField: "coverage",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "end-patch.prerequisites",
			InnerField: "prerequisites",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "end-patch.saved-groups",
			InnerField: "savedGroups",
		},
	},
})

var rampScheduleTemplatesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single rampScheduleTemplate",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleRampScheduleTemplatesRetrieve,
	HideHelpCommand: true,
}

var rampScheduleTemplatesUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a single rampScheduleTemplate",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "end-patch",
			BodyPath: "endPatch",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[bool]{
			Name:     "official",
			BodyPath: "official",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "step",
			BodyPath: "steps",
		},
	},
	Action:          handleRampScheduleTemplatesUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"end-patch": {
		&requestflag.InnerFlag[string]{
			Name:       "end-patch.condition",
			InnerField: "condition",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "end-patch.coverage",
			InnerField: "coverage",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "end-patch.prerequisites",
			InnerField: "prerequisites",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "end-patch.saved-groups",
			InnerField: "savedGroups",
		},
	},
	"step": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "step.actions",
			InnerField: "actions",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "step.trigger",
			InnerField: "trigger",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "step.approval-notes",
			InnerField: "approvalNotes",
		},
	},
})

var rampScheduleTemplatesList = cli.Command{
	Name:            "list",
	Usage:           "Get all rampScheduleTemplates",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleRampScheduleTemplatesList,
	HideHelpCommand: true,
}

var rampScheduleTemplatesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a single rampScheduleTemplate",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleRampScheduleTemplatesDelete,
	HideHelpCommand: true,
}

func handleRampScheduleTemplatesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.RampScheduleTemplateNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.RampScheduleTemplates.New(ctx, params, options...)
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
		Title:          "ramp-schedule-templates create",
		Transform:      transform,
	})
}

func handleRampScheduleTemplatesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.RampScheduleTemplates.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "ramp-schedule-templates retrieve",
		Transform:      transform,
	})
}

func handleRampScheduleTemplatesUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.RampScheduleTemplateUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.RampScheduleTemplates.Update(
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
		Title:          "ramp-schedule-templates update",
		Transform:      transform,
	})
}

func handleRampScheduleTemplatesList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.RampScheduleTemplates.List(ctx, options...)
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
		Title:          "ramp-schedule-templates list",
		Transform:      transform,
	})
}

func handleRampScheduleTemplatesDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.RampScheduleTemplates.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "ramp-schedule-templates delete",
		Transform:      transform,
	})
}
