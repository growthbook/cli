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

var dashboardsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a single dashboard",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "block",
			Required: true,
			BodyPath: "blocks",
		},
		&requestflag.Flag[string]{
			Name:     "edit-level",
			Usage:    `Dashboards that are "published" are editable by organization members with appropriate permissions`,
			Required: true,
			BodyPath: "editLevel",
		},
		&requestflag.Flag[bool]{
			Name:     "enable-auto-updates",
			Usage:    "If enabled for a General Dashboard, also requires an updateSchedule",
			Required: true,
			BodyPath: "enableAutoUpdates",
		},
		&requestflag.Flag[string]{
			Name:     "share-level",
			Usage:    `General Dashboards only. Dashboards that are "published" are viewable by organization members with appropriate permissions`,
			Required: true,
			BodyPath: "shareLevel",
		},
		&requestflag.Flag[string]{
			Name:     "title",
			Usage:    "The display name of the Dashboard",
			Required: true,
			BodyPath: "title",
		},
		&requestflag.Flag[string]{
			Name:     "experiment-id",
			Usage:    "The parent experiment for an Experiment Dashboard, or undefined for a general dashboard",
			BodyPath: "experimentId",
		},
		&requestflag.Flag[[]string]{
			Name:     "project",
			Usage:    "General Dashboards only, Experiment Dashboards use the experiment's projects",
			BodyPath: "projects",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "update-schedule",
			Usage:    "General Dashboards only. Experiment Dashboards update based on the parent experiment instead",
			BodyPath: "updateSchedule",
		},
	},
	Action:          handleDashboardsCreate,
	HideHelpCommand: true,
}

var dashboardsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single dashboard",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleDashboardsRetrieve,
	HideHelpCommand: true,
}

var dashboardsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a single dashboard",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "block",
			BodyPath: "blocks",
		},
		&requestflag.Flag[string]{
			Name:     "edit-level",
			Usage:    `Dashboards that are "published" are editable by organization members with appropriate permissions`,
			BodyPath: "editLevel",
		},
		&requestflag.Flag[bool]{
			Name:     "enable-auto-updates",
			Usage:    "If enabled for a General Dashboard, also requires an updateSchedule",
			BodyPath: "enableAutoUpdates",
		},
		&requestflag.Flag[[]string]{
			Name:     "project",
			Usage:    "General Dashboards only, Experiment Dashboards use the experiment's projects",
			BodyPath: "projects",
		},
		&requestflag.Flag[string]{
			Name:     "share-level",
			Usage:    `General Dashboards only. Dashboards that are "published" are viewable by organization members with appropriate permissions`,
			BodyPath: "shareLevel",
		},
		&requestflag.Flag[string]{
			Name:     "title",
			Usage:    "The display name of the Dashboard",
			BodyPath: "title",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "update-schedule",
			Usage:    "General Dashboards only. Experiment Dashboards update based on the parent experiment instead",
			BodyPath: "updateSchedule",
		},
	},
	Action:          handleDashboardsUpdate,
	HideHelpCommand: true,
}

var dashboardsList = cli.Command{
	Name:            "list",
	Usage:           "Get all dashboards",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleDashboardsList,
	HideHelpCommand: true,
}

var dashboardsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a single dashboard",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleDashboardsDelete,
	HideHelpCommand: true,
}

var dashboardsListByExperiment = cli.Command{
	Name:    "list-by-experiment",
	Usage:   "Get all dashboards for an experiment",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "experiment-id",
			Required:  true,
			PathParam: "experimentId",
		},
	},
	Action:          handleDashboardsListByExperiment,
	HideHelpCommand: true,
}

func handleDashboardsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.DashboardNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Dashboards.New(ctx, params, options...)
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
		Title:          "dashboards create",
		Transform:      transform,
	})
}

func handleDashboardsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Dashboards.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "dashboards retrieve",
		Transform:      transform,
	})
}

func handleDashboardsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.DashboardUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Dashboards.Update(
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
		Title:          "dashboards update",
		Transform:      transform,
	})
}

func handleDashboardsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Dashboards.List(ctx, options...)
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
		Title:          "dashboards list",
		Transform:      transform,
	})
}

func handleDashboardsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Dashboards.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "dashboards delete",
		Transform:      transform,
	})
}

func handleDashboardsListByExperiment(ctx context.Context, cmd *cli.Command) error {
	client := growthbook.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("experiment-id") && len(unusedArgs) > 0 {
		cmd.Set("experiment-id", unusedArgs[0])
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
	_, err = client.Dashboards.ListByExperiment(ctx, cmd.Value("experiment-id").(string), options...)
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
		Title:          "dashboards list-by-experiment",
		Transform:      transform,
	})
}
