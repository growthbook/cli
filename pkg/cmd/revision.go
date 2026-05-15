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

var revisionsList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of feature revisions across all features in the\norganization. Revision `rules` is a flat array with per-rule scope.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "author",
			QueryPath: "author",
		},
		&requestflag.Flag[string]{
			Name:      "feature-id",
			QueryPath: "featureId",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "The number of items to return",
			Default:   10,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "mine",
			Usage:     "If true, return only revisions authored by or contributed to by the calling user.",
			QueryPath: "mine",
		},
		&requestflag.Flag[int64]{
			Name:      "offset",
			Usage:     "How many items to skip (use in conjunction with limit for pagination)",
			Default:   0,
			QueryPath: "offset",
		},
		&requestflag.Flag[any]{
			Name:      "skip-pagination",
			Usage:     "If true, return all matching items and ignore limit/offset.\nSelf-hosted only. Has no effect unless API_ALLOW_SKIP_PAGINATION is set to true or 1.",
			Default:   false,
			QueryPath: "skipPagination",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     `Allowed values: "draft", "published", "discarded", "approved", "changes-requested", "pending-review", "pending-parent".`,
			QueryPath: "status",
		},
	},
	Action:          handleRevisionsList,
	HideHelpCommand: true,
}

func handleRevisionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.RevisionListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Revisions.List(ctx, params, options...)
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
		Title:          "revisions list",
		Transform:      transform,
	})
}
