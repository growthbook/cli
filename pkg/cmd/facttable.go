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

var factTablesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a single fact table",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "datasource",
			Usage:    "The datasource id",
			Required: true,
			BodyPath: "datasource",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "sql",
			Usage:    "The SQL query for this fact table",
			Required: true,
			BodyPath: "sql",
		},
		&requestflag.Flag[[]string]{
			Name:     "user-id-type",
			Usage:    `List of identifier columns in this table. For example, "id" or "anonymous_id"`,
			Required: true,
			BodyPath: "userIdTypes",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of the fact table",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "event-name",
			Usage:    "The event name used in SQL template variables",
			BodyPath: "eventName",
		},
		&requestflag.Flag[string]{
			Name:     "managed-by",
			Usage:    `Set this to "api" to disable editing in the GrowthBook UI`,
			BodyPath: "managedBy",
		},
		&requestflag.Flag[string]{
			Name:     "owner",
			Usage:    "The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.",
			BodyPath: "owner",
		},
		&requestflag.Flag[[]string]{
			Name:     "project",
			Usage:    "List of associated project ids",
			BodyPath: "projects",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "List of associated tags",
			BodyPath: "tags",
		},
	},
	Action:          handleFactTablesCreate,
	HideHelpCommand: true,
}

var factTablesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single fact table",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleFactTablesRetrieve,
	HideHelpCommand: true,
}

var factTablesUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a single fact table",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[bool]{
			Name:     "archived",
			BodyPath: "archived",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "column",
			Usage:    "Optional array of columns that you want to update. Only allows updating properties of existing columns. Cannot create new columns or delete existing ones. Columns cannot be added or deleted; column structure is determined by SQL parsing. Slice-related properties require an enterprise license.",
			BodyPath: "columns",
		},
		&requestflag.Flag[*string]{
			Name:     "columns-error",
			Usage:    "Error message if there was an issue parsing the SQL schema",
			BodyPath: "columnsError",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of the fact table",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "event-name",
			Usage:    "The event name used in SQL template variables",
			BodyPath: "eventName",
		},
		&requestflag.Flag[string]{
			Name:     "managed-by",
			Usage:    `Set this to "api" to disable editing in the GrowthBook UI`,
			BodyPath: "managedBy",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "owner",
			Usage:    "The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.",
			BodyPath: "owner",
		},
		&requestflag.Flag[[]string]{
			Name:     "project",
			Usage:    "List of associated project ids",
			BodyPath: "projects",
		},
		&requestflag.Flag[string]{
			Name:     "sql",
			Usage:    "The SQL query for this fact table",
			BodyPath: "sql",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "List of associated tags",
			BodyPath: "tags",
		},
		&requestflag.Flag[[]string]{
			Name:     "user-id-type",
			Usage:    `List of identifier columns in this table. For example, "id" or "anonymous_id"`,
			BodyPath: "userIdTypes",
		},
	},
	Action:          handleFactTablesUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"column": {
		&requestflag.InnerFlag[string]{
			Name:       "column.column",
			Usage:      "The actual column name in the database/SQL query",
			InnerField: "column",
		},
		&requestflag.InnerFlag[string]{
			Name:       "column.datatype",
			Usage:      `Allowed values: "number", "string", "date", "boolean", "json", "binary", "other", "".`,
			InnerField: "datatype",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "column.always-inline-filter",
			Usage:      "Whether this column should always be included as an inline filter in queries",
			InnerField: "alwaysInlineFilter",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "column.auto-slices",
			Usage:      "Specific slices to automatically analyze for this column.",
			InnerField: "autoSlices",
		},
		&requestflag.InnerFlag[any]{
			Name:       "column.date-created",
			InnerField: "dateCreated",
		},
		&requestflag.InnerFlag[any]{
			Name:       "column.date-updated",
			InnerField: "dateUpdated",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "column.deleted",
			InnerField: "deleted",
		},
		&requestflag.InnerFlag[string]{
			Name:       "column.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "column.is-auto-slice-column",
			Usage:      "Whether this column can be used for auto slice analysis. This is an enterprise feature.",
			InnerField: "isAutoSliceColumn",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "column.json-fields",
			Usage:      "For JSON columns, defines the structure of nested fields",
			InnerField: "jsonFields",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "column.locked-auto-slices",
			Usage:      "Locked slices that are protected from automatic updates. These will always be included in the slice levels even if they're not in the top values query results.",
			InnerField: "lockedAutoSlices",
		},
		&requestflag.InnerFlag[string]{
			Name:       "column.name",
			Usage:      "Display name for the column (can be different from the actual column name)",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "column.number-format",
			Usage:      `Allowed values: "", "currency", "time:seconds", "memory:bytes", "memory:kilobytes".`,
			InnerField: "numberFormat",
		},
	},
})

var factTablesList = cli.Command{
	Name:    "list",
	Usage:   "Get all fact tables",
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
		&requestflag.Flag[string]{
			Name:      "project-id",
			Usage:     "Filter by project id",
			QueryPath: "projectId",
		},
	},
	Action:          handleFactTablesList,
	HideHelpCommand: true,
}

var factTablesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a single fact table",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleFactTablesDelete,
	HideHelpCommand: true,
}

func handleFactTablesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.FactTableNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.FactTables.New(ctx, params, options...)
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
		Title:          "fact-tables create",
		Transform:      transform,
	})
}

func handleFactTablesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.FactTables.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "fact-tables retrieve",
		Transform:      transform,
	})
}

func handleFactTablesUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.FactTableUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.FactTables.Update(
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
		Title:          "fact-tables update",
		Transform:      transform,
	})
}

func handleFactTablesList(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.FactTableListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.FactTables.List(ctx, params, options...)
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
		Title:          "fact-tables list",
		Transform:      transform,
	})
}

func handleFactTablesDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.FactTables.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "fact-tables delete",
		Transform:      transform,
	})
}
