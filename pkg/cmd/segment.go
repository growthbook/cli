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

var segmentsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a single segment",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "datasource-id",
			Usage:    "ID of the datasource this segment belongs to",
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
			Usage:    "Name of the segment",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "GrowthBook supports two types of Segments, SQL and FACT. SQL segments are defined by a SQL query, and FACT segments are defined by a fact table and filters.",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of the segment",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "fact-table-id",
			Usage:    "ID of the fact table this segment belongs to. This is required for FACT segments.",
			BodyPath: "factTableId",
		},
		&requestflag.Flag[[]string]{
			Name:     "filter",
			Usage:    "Optional array of fact table filter ids that can further define the Fact Table based Segment.",
			BodyPath: "filters",
		},
		&requestflag.Flag[string]{
			Name:     "managed-by",
			Usage:    "Where this Segment must be managed from. If not set (empty string), it can be managed from anywhere.",
			BodyPath: "managedBy",
		},
		&requestflag.Flag[string]{
			Name:     "owner",
			Usage:    "The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.",
			BodyPath: "owner",
		},
		&requestflag.Flag[[]string]{
			Name:     "project",
			Usage:    "List of project IDs for projects that can access this segment",
			BodyPath: "projects",
		},
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "SQL query that defines the Segment. This is required for SQL segments.",
			BodyPath: "query",
		},
	},
	Action:          handleSegmentsCreate,
	HideHelpCommand: true,
}

var segmentsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single segment",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleSegmentsRetrieve,
	HideHelpCommand: true,
}

var segmentsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a single segment",
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
			Usage:    "ID of the datasource this segment belongs to",
			BodyPath: "datasourceId",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of the segment",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "fact-table-id",
			Usage:    "ID of the fact table this segment belongs to. This is required for FACT segments.",
			BodyPath: "factTableId",
		},
		&requestflag.Flag[[]string]{
			Name:     "filter",
			Usage:    "Optional array of fact table filter ids that can further define the Fact Table based Segment.",
			BodyPath: "filters",
		},
		&requestflag.Flag[string]{
			Name:     "identifier-type",
			Usage:    "Type of identifier (user, anonymous, etc.)",
			BodyPath: "identifierType",
		},
		&requestflag.Flag[string]{
			Name:     "managed-by",
			Usage:    "Where this Segment must be managed from. If not set (empty string), it can be managed from anywhere.",
			BodyPath: "managedBy",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the segment",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "owner",
			Usage:    "The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.",
			BodyPath: "owner",
		},
		&requestflag.Flag[[]string]{
			Name:     "project",
			Usage:    "List of project IDs for projects that can access this segment",
			BodyPath: "projects",
		},
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "SQL query that defines the Segment. This is required for SQL segments.",
			BodyPath: "query",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "GrowthBook supports two types of Segments, SQL and FACT. SQL segments are defined by a SQL query, and FACT segments are defined by a fact table and filters.",
			BodyPath: "type",
		},
	},
	Action:          handleSegmentsUpdate,
	HideHelpCommand: true,
}

var segmentsList = cli.Command{
	Name:    "list",
	Usage:   "Get all segments",
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
	Action:          handleSegmentsList,
	HideHelpCommand: true,
}

var segmentsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a single segment",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleSegmentsDelete,
	HideHelpCommand: true,
}

func handleSegmentsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.SegmentNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Segments.New(ctx, params, options...)
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
		Title:          "segments create",
		Transform:      transform,
	})
}

func handleSegmentsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Segments.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "segments retrieve",
		Transform:      transform,
	})
}

func handleSegmentsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.SegmentUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Segments.Update(
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
		Title:          "segments update",
		Transform:      transform,
	})
}

func handleSegmentsList(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.SegmentListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Segments.List(ctx, params, options...)
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
		Title:          "segments list",
		Transform:      transform,
	})
}

func handleSegmentsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Segments.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "segments delete",
		Transform:      transform,
	})
}
