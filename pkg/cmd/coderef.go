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

var codeRefsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Submit list of code references",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "branch",
			Required: true,
			BodyPath: "branch",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "ref",
			Required: true,
			BodyPath: "refs",
		},
		&requestflag.Flag[string]{
			Name:     "repo-name",
			Required: true,
			BodyPath: "repoName",
		},
		&requestflag.Flag[string]{
			Name:      "delete-missing",
			Usage:     "Whether to delete code references that are no longer present in the submitted data",
			Default:   "false",
			QueryPath: "deleteMissing",
		},
	},
	Action:          handleCodeRefsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"ref": {
		&requestflag.InnerFlag[string]{
			Name:       "ref.content-hash",
			InnerField: "contentHash",
		},
		&requestflag.InnerFlag[string]{
			Name:       "ref.file-path",
			InnerField: "filePath",
		},
		&requestflag.InnerFlag[string]{
			Name:       "ref.flag-key",
			InnerField: "flagKey",
		},
		&requestflag.InnerFlag[string]{
			Name:       "ref.lines",
			InnerField: "lines",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "ref.starting-line-number",
			InnerField: "startingLineNumber",
		},
	},
})

var codeRefsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get list of code references for a single feature id",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleCodeRefsRetrieve,
	HideHelpCommand: true,
}

var codeRefsList = cli.Command{
	Name:    "list",
	Usage:   "Get list of all code references for the current organization",
	Suggest: true,
	Flags: []cli.Flag{
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
	Action:          handleCodeRefsList,
	HideHelpCommand: true,
}

func handleCodeRefsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.CodeRefNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.CodeRefs.New(ctx, params, options...)
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
		Title:          "code-refs create",
		Transform:      transform,
	})
}

func handleCodeRefsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.CodeRefs.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "code-refs retrieve",
		Transform:      transform,
	})
}

func handleCodeRefsList(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.CodeRefListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.CodeRefs.List(ctx, params, options...)
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
		Title:          "code-refs list",
		Transform:      transform,
	})
}
