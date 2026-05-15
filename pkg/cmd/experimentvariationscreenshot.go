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

var experimentsVariationScreenshotDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a variation screenshot",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "variation-id",
			Required:  true,
			PathParam: "variationId",
		},
		&requestflag.Flag[string]{
			Name:     "path",
			Usage:    "The screenshot path/URL to delete (from upload response)",
			Required: true,
			BodyPath: "path",
		},
	},
	Action:          handleExperimentsVariationScreenshotDelete,
	HideHelpCommand: true,
}

var experimentsVariationScreenshotUpload = cli.Command{
	Name:    "upload",
	Usage:   "Upload a variation screenshot",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "variation-id",
			Required:  true,
			PathParam: "variationId",
		},
		&requestflag.Flag[string]{
			Name:     "content-type",
			Usage:    "MIME type of the screenshot",
			Required: true,
			BodyPath: "contentType",
		},
		&requestflag.Flag[string]{
			Name:     "screenshot",
			Usage:    "Base64-encoded screenshot data",
			Required: true,
			BodyPath: "screenshot",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Optional description for the screenshot",
			BodyPath: "description",
		},
	},
	Action:          handleExperimentsVariationScreenshotUpload,
	HideHelpCommand: true,
}

func handleExperimentsVariationScreenshotDelete(ctx context.Context, cmd *cli.Command) error {
	client := growthbook.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("variation-id") && len(unusedArgs) > 0 {
		cmd.Set("variation-id", unusedArgs[0])
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

	params := growthbook.ExperimentVariationScreenshotDeleteParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.Variation.Screenshot.Delete(
		ctx,
		cmd.Value("variation-id").(string),
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
		Title:          "experiments:variation:screenshot delete",
		Transform:      transform,
	})
}

func handleExperimentsVariationScreenshotUpload(ctx context.Context, cmd *cli.Command) error {
	client := growthbook.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("variation-id") && len(unusedArgs) > 0 {
		cmd.Set("variation-id", unusedArgs[0])
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

	params := growthbook.ExperimentVariationScreenshotUploadParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Experiments.Variation.Screenshot.Upload(
		ctx,
		cmd.Value("variation-id").(string),
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
		Title:          "experiments:variation:screenshot upload",
		Transform:      transform,
	})
}
