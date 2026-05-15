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

var teamsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a single team",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "description",
			Required: true,
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "role",
			Usage:    "The global role for members of this team",
			Required: true,
			BodyPath: "role",
		},
		&requestflag.Flag[string]{
			Name:     "created-by",
			BodyPath: "createdBy",
		},
		&requestflag.Flag[string]{
			Name:     "default-project",
			BodyPath: "defaultProject",
		},
		&requestflag.Flag[[]string]{
			Name:     "environment",
			Usage:    "An empty array means 'all environments'",
			BodyPath: "environments",
		},
		&requestflag.Flag[bool]{
			Name:     "limit-access-by-environment",
			BodyPath: "limitAccessByEnvironment",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "managed-by",
			BodyPath: "managedBy",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "project-role",
			BodyPath: "projectRoles",
		},
	},
	Action:          handleTeamsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"managed-by": {
		&requestflag.InnerFlag[string]{
			Name:       "managed-by.resource-id",
			InnerField: "resourceId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "managed-by.type",
			Usage:      `Allowed values: "vercel".`,
			InnerField: "type",
		},
	},
	"project-role": {
		&requestflag.InnerFlag[[]string]{
			Name:       "project-role.environments",
			InnerField: "environments",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "project-role.limit-access-by-environment",
			InnerField: "limitAccessByEnvironment",
		},
		&requestflag.InnerFlag[string]{
			Name:       "project-role.project",
			InnerField: "project",
		},
		&requestflag.InnerFlag[string]{
			Name:       "project-role.role",
			InnerField: "role",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "project-role.teams",
			InnerField: "teams",
		},
	},
})

var teamsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single team",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleTeamsRetrieve,
	HideHelpCommand: true,
}

var teamsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a single team",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "created-by",
			BodyPath: "createdBy",
		},
		&requestflag.Flag[string]{
			Name:     "default-project",
			BodyPath: "defaultProject",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[[]string]{
			Name:     "environment",
			Usage:    "An empty array means 'all environments'",
			BodyPath: "environments",
		},
		&requestflag.Flag[bool]{
			Name:     "limit-access-by-environment",
			BodyPath: "limitAccessByEnvironment",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "managed-by",
			BodyPath: "managedBy",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "project-role",
			BodyPath: "projectRoles",
		},
		&requestflag.Flag[string]{
			Name:     "role",
			Usage:    "The global role for members of this team",
			BodyPath: "role",
		},
	},
	Action:          handleTeamsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"managed-by": {
		&requestflag.InnerFlag[string]{
			Name:       "managed-by.resource-id",
			InnerField: "resourceId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "managed-by.type",
			Usage:      `Allowed values: "vercel".`,
			InnerField: "type",
		},
	},
	"project-role": {
		&requestflag.InnerFlag[[]string]{
			Name:       "project-role.environments",
			InnerField: "environments",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "project-role.limit-access-by-environment",
			InnerField: "limitAccessByEnvironment",
		},
		&requestflag.InnerFlag[string]{
			Name:       "project-role.project",
			InnerField: "project",
		},
		&requestflag.InnerFlag[string]{
			Name:       "project-role.role",
			InnerField: "role",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "project-role.teams",
			InnerField: "teams",
		},
	},
})

var teamsList = cli.Command{
	Name:            "list",
	Usage:           "Get all teams",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleTeamsList,
	HideHelpCommand: true,
}

var teamsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a single team",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "delete-members",
			Usage:     "When 'true', enables deleting a team that contains members",
			QueryPath: "deleteMembers",
		},
	},
	Action:          handleTeamsDelete,
	HideHelpCommand: true,
}

func handleTeamsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.TeamNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Teams.New(ctx, params, options...)
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
		Title:          "teams create",
		Transform:      transform,
	})
}

func handleTeamsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Teams.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "teams retrieve",
		Transform:      transform,
	})
}

func handleTeamsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.TeamUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Teams.Update(
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
		Title:          "teams update",
		Transform:      transform,
	})
}

func handleTeamsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Teams.List(ctx, options...)
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
		Title:          "teams list",
		Transform:      transform,
	})
}

func handleTeamsDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.TeamDeleteParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Teams.Delete(
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
		Title:          "teams delete",
		Transform:      transform,
	})
}
