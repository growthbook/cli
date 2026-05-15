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

var dimensionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a single dimension",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "datasource-id",
			Usage:    "ID of the datasource this dimension belongs to",
			Required: true,
			BodyPath: "datasourceId",
		},
		&requestflag.Flag[string]{
			Name:     "identifier-type",
			Usage:    "Type of identifier (user, anonymous, etc.)",
			Required: true,
			BodyPath: "identifierType",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the dimension",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "SQL query or equivalent for the dimension",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of the dimension",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "managed-by",
			Usage:    "Where this dimension must be managed from. If not set (empty string), it can be managed from anywhere.",
			BodyPath: "managedBy",
		},
		&requestflag.Flag[string]{
			Name:     "owner",
			Usage:    "The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.",
			BodyPath: "owner",
		},
	},
	Action:          handleDimensionsCreate,
	HideHelpCommand: true,
}

var dimensionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single dimension",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleDimensionsRetrieve,
	HideHelpCommand: true,
}

var dimensionsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a single dimension",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "datasource-id",
			Usage:    "ID of the datasource this dimension belongs to",
			BodyPath: "datasourceId",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of the dimension",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "identifier-type",
			Usage:    "Type of identifier (user, anonymous, etc.)",
			BodyPath: "identifierType",
		},
		&requestflag.Flag[string]{
			Name:     "managed-by",
			Usage:    "Where this dimension must be managed from. If not set (empty string), it can be managed from anywhere.",
			BodyPath: "managedBy",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the dimension",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "owner",
			Usage:    "The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.",
			BodyPath: "owner",
		},
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "SQL query or equivalent for the dimension",
			BodyPath: "query",
		},
	},
	Action:          handleDimensionsUpdate,
	HideHelpCommand: true,
}

var dimensionsList = cli.Command{
	Name:    "list",
	Usage:   "Get all dimensions",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "datasource-id",
			Usage:     "Filter by Data Source",
			QueryPath: "datasourceId",
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
	},
	Action:          handleDimensionsList,
	HideHelpCommand: true,
}

var dimensionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a single dimension",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleDimensionsDelete,
	HideHelpCommand: true,
}

func handleDimensionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.DimensionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Dimensions.New(ctx, params, options...)
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
		Title:          "dimensions create",
		Transform:      transform,
	})
}

func handleDimensionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Dimensions.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "dimensions retrieve",
		Transform:      transform,
	})
}

func handleDimensionsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.DimensionUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Dimensions.Update(
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
		Title:          "dimensions update",
		Transform:      transform,
	})
}

func handleDimensionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.DimensionListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Dimensions.List(ctx, params, options...)
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
		Title:          "dimensions list",
		Transform:      transform,
	})
}

func handleDimensionsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Dimensions.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "dimensions delete",
		Transform:      transform,
	})
}
