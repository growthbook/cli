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

var featuresRevisionsMetadataCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Update revision metadata",
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
			Name:     "comment",
			BodyPath: "comment",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "custom-fields",
			BodyPath: "customFields",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "json-schema",
			BodyPath: "jsonSchema",
		},
		&requestflag.Flag[bool]{
			Name:     "never-stale",
			BodyPath: "neverStale",
		},
		&requestflag.Flag[string]{
			Name:     "owner",
			Usage:    "The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.",
			BodyPath: "owner",
		},
		&requestflag.Flag[string]{
			Name:     "project",
			BodyPath: "project",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "title",
			BodyPath: "title",
		},
	},
	Action:          handleFeaturesRevisionsMetadataCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"json-schema": {
		&requestflag.InnerFlag[any]{
			Name:       "json-schema.date",
			InnerField: "date",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "json-schema.enabled",
			InnerField: "enabled",
		},
		&requestflag.InnerFlag[string]{
			Name:       "json-schema.schema",
			InnerField: "schema",
		},
		&requestflag.InnerFlag[string]{
			Name:       "json-schema.schema-type",
			Usage:      `Allowed values: "schema", "simple".`,
			InnerField: "schemaType",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "json-schema.simple",
			InnerField: "simple",
		},
	},
})

func handleFeaturesRevisionsMetadataCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.FeatureRevisionMetadataNewParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Features.Revisions.Metadata.New(
		ctx,
		growthbook.FeatureRevisionMetadataNewParamsVersionUnion(cmd.Value("version").(any)),
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
		Title:          "features:revisions:metadata create",
		Transform:      transform,
	})
}
