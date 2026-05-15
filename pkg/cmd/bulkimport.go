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

var bulkImportCreateFacts = requestflag.WithInnerFlags(cli.Command{
	Name:    "create-facts",
	Usage:   "Bulk import fact tables, filters, and metrics",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "fact-metric",
			BodyPath: "factMetrics",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "fact-table-filter",
			BodyPath: "factTableFilters",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "fact-table",
			BodyPath: "factTables",
		},
	},
	Action:          handleBulkImportCreateFacts,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"fact-metric": {
		&requestflag.InnerFlag[string]{
			Name:       "fact-metric.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "fact-metric.data",
			InnerField: "data",
		},
	},
	"fact-table-filter": {
		&requestflag.InnerFlag[string]{
			Name:       "fact-table-filter.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "fact-table-filter.data",
			InnerField: "data",
		},
		&requestflag.InnerFlag[string]{
			Name:       "fact-table-filter.fact-table-id",
			InnerField: "factTableId",
		},
	},
	"fact-table": {
		&requestflag.InnerFlag[string]{
			Name:       "fact-table.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "fact-table.data",
			InnerField: "data",
		},
	},
})

func handleBulkImportCreateFacts(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.BulkImportNewFactsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.BulkImport.NewFacts(ctx, params, options...)
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
		Title:          "bulk-import create-facts",
		Transform:      transform,
	})
}
