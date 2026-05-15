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

var attributesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new attribute",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "datatype",
			Usage:    "The attribute datatype",
			Required: true,
			BodyPath: "datatype",
		},
		&requestflag.Flag[string]{
			Name:     "property",
			Usage:    "The attribute property",
			Required: true,
			BodyPath: "property",
		},
		&requestflag.Flag[bool]{
			Name:     "archived",
			Usage:    "The attribute is archived",
			BodyPath: "archived",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "The description of the new attribute",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "enum",
			BodyPath: "enum",
		},
		&requestflag.Flag[string]{
			Name:     "format",
			Usage:    "The attribute's format",
			BodyPath: "format",
		},
		&requestflag.Flag[bool]{
			Name:     "hash-attribute",
			Usage:    "Shall the attribute be hashed",
			BodyPath: "hashAttribute",
		},
		&requestflag.Flag[[]string]{
			Name:     "project",
			BodyPath: "projects",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
	},
	Action:          handleAttributesCreate,
	HideHelpCommand: true,
}

var attributesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update an attribute",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "property",
			Usage:     "The attribute property",
			Required:  true,
			PathParam: "property",
		},
		&requestflag.Flag[bool]{
			Name:     "archived",
			Usage:    "The attribute is archived",
			BodyPath: "archived",
		},
		&requestflag.Flag[string]{
			Name:     "datatype",
			Usage:    "The attribute datatype",
			BodyPath: "datatype",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "The description of the new attribute",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "enum",
			BodyPath: "enum",
		},
		&requestflag.Flag[string]{
			Name:     "format",
			Usage:    "The attribute's format",
			BodyPath: "format",
		},
		&requestflag.Flag[bool]{
			Name:     "hash-attribute",
			Usage:    "Shall the attribute be hashed",
			BodyPath: "hashAttribute",
		},
		&requestflag.Flag[[]string]{
			Name:     "project",
			BodyPath: "projects",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
	},
	Action:          handleAttributesUpdate,
	HideHelpCommand: true,
}

var attributesList = cli.Command{
	Name:            "list",
	Usage:           "Get the organization's attributes",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleAttributesList,
	HideHelpCommand: true,
}

var attributesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a single attribute",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "property",
			Usage:     "The attribute property",
			Required:  true,
			PathParam: "property",
		},
	},
	Action:          handleAttributesDelete,
	HideHelpCommand: true,
}

func handleAttributesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.AttributeNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Attributes.New(ctx, params, options...)
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
		Title:          "attributes create",
		Transform:      transform,
	})
}

func handleAttributesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := growthbook.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("property") && len(unusedArgs) > 0 {
		cmd.Set("property", unusedArgs[0])
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

	params := growthbook.AttributeUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Attributes.Update(
		ctx,
		cmd.Value("property").(string),
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
		Title:          "attributes update",
		Transform:      transform,
	})
}

func handleAttributesList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Attributes.List(ctx, options...)
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
		Title:          "attributes list",
		Transform:      transform,
	})
}

func handleAttributesDelete(ctx context.Context, cmd *cli.Command) error {
	client := growthbook.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("property") && len(unusedArgs) > 0 {
		cmd.Set("property", unusedArgs[0])
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
	_, err = client.Attributes.Delete(ctx, cmd.Value("property").(string), options...)
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
		Title:          "attributes delete",
		Transform:      transform,
	})
}
