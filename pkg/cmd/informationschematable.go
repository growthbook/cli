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

var informationSchemaTablesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns cached metadata for a specific table in the Data Source, including\ncolumns and their data types. Not all data source types support information\nschemas.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "table-id",
			Usage:     "The id of the information schema table",
			Required:  true,
			PathParam: "tableId",
		},
	},
	Action:          handleInformationSchemaTablesRetrieve,
	HideHelpCommand: true,
}

func handleInformationSchemaTablesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := growthbook.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("table-id") && len(unusedArgs) > 0 {
		cmd.Set("table-id", unusedArgs[0])
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
	_, err = client.InformationSchemaTables.Get(ctx, cmd.Value("table-id").(string), options...)
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
		Title:          "information-schema-tables retrieve",
		Transform:      transform,
	})
}
