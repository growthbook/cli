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

var featuresCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new feature. Rules are supplied as a top-level `rules` array; each\nrule includes `allEnvironments` / `environments` scope fields.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "A unique key name for the feature. Feature keys can only include letters, numbers, hyphens, and underscores.",
			Required: true,
			BodyPath: "id",
		},
		&requestflag.Flag[string]{
			Name:     "default-value",
			Usage:    "Default value when feature is enabled. Type must match `valueType`.",
			Required: true,
			BodyPath: "defaultValue",
		},
		&requestflag.Flag[string]{
			Name:     "owner",
			Usage:    "The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.",
			Required: true,
			BodyPath: "owner",
		},
		&requestflag.Flag[string]{
			Name:     "value-type",
			Usage:    "The data type of the feature payload. Boolean by default.",
			Required: true,
			BodyPath: "valueType",
		},
		&requestflag.Flag[bool]{
			Name:     "archived",
			BodyPath: "archived",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "custom-fields",
			BodyPath: "customFields",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of the feature",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "environments",
			Usage:    "Per-environment enabled state. V2 rules are specified on the top-level `rules` field.",
			BodyPath: "environments",
		},
		&requestflag.Flag[string]{
			Name:     "json-schema",
			Usage:    "Use JSON schema to validate the payload of a JSON-type feature value (enterprise only).",
			BodyPath: "jsonSchema",
		},
		&requestflag.Flag[[]string]{
			Name:     "prerequisite",
			Usage:    "Feature IDs. Each feature must evaluate to `true`",
			BodyPath: "prerequisites",
		},
		&requestflag.Flag[string]{
			Name:     "project",
			Usage:    "An associated project ID",
			BodyPath: "project",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "rule",
			Usage:    "Feature rules. Each rule carries its own environment scope via `allEnvironments` / `environments`.",
			BodyPath: "rules",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "List of associated tags",
			BodyPath: "tags",
		},
	},
	Action:          handleFeaturesCreate,
	HideHelpCommand: true,
}

var featuresRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single feature",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "with-revisions",
			Usage:     "Also return feature revisions (all, draft, or published statuses)",
			QueryPath: "withRevisions",
		},
	},
	Action:          handleFeaturesRetrieve,
	HideHelpCommand: true,
}

var featuresUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates any combination of a feature's metadata, default value, environment\nstate, and rules. Other top-level fields are patch-merged: omit a field to leave\nit unchanged. The `rules` field, when supplied, replaces the entire `rules`\narray atomically in a single revision (v1 PUT applied per-environment patches;\nv2 swaps the full flat array). To preserve existing rules during a partial edit,\nGET the feature first, mutate the returned `rules` array, and PUT the full array\nback. Safe-rollout rules round-trip via their `safeRolloutId`; use\n`POST /v2/features/:id/revisions/:version/rules` to create new ones. Returns 403\nif approval rules are enabled for an affected environment and the bypass setting\nis off.",
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
		&requestflag.Flag[map[string]any]{
			Name:     "custom-fields",
			BodyPath: "customFields",
		},
		&requestflag.Flag[string]{
			Name:     "default-value",
			BodyPath: "defaultValue",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of the feature",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "environments",
			Usage:    "Per-environment enabled state. V2 rules are specified on the top-level `rules` field.",
			BodyPath: "environments",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "holdout",
			Usage:    "Holdout to assign this feature to. Pass `null` to remove the feature from its current holdout. Omit the field entirely to leave the holdout unchanged.",
			BodyPath: "holdout",
		},
		&requestflag.Flag[string]{
			Name:     "json-schema",
			Usage:    "Use JSON schema to validate the payload of a JSON-type feature value (enterprise only).",
			BodyPath: "jsonSchema",
		},
		&requestflag.Flag[string]{
			Name:     "owner",
			Usage:    "The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.",
			BodyPath: "owner",
		},
		&requestflag.Flag[[]string]{
			Name:     "prerequisite",
			Usage:    "Feature IDs. Each feature must evaluate to `true`",
			BodyPath: "prerequisites",
		},
		&requestflag.Flag[string]{
			Name:     "project",
			Usage:    "An associated project ID",
			BodyPath: "project",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "rule",
			Usage:    "Replaces all feature rules atomically. Behavior differs from v1: v1 PUT applies per-environment patches, v2 PUT swaps the entire `rules` array in one revision. To preserve existing rules during a partial edit, GET the feature first, mutate the returned `rules` array, and PUT the full array back. Safe-rollout rules round-trip via their `safeRolloutId` (creation requires `POST /v2/features/:id/revisions/:version/rules`).",
			BodyPath: "rules",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "List of associated tags. Will override tags completely with submitted list",
			BodyPath: "tags",
		},
	},
	Action:          handleFeaturesUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"holdout": {
		&requestflag.InnerFlag[string]{
			Name:       "holdout.id",
			Usage:      "Holdout ID",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "holdout.value",
			Usage:      "The feature value assigned to users in the holdout treatment group",
			InnerField: "value",
		},
	},
})

var featuresList = cli.Command{
	Name:    "list",
	Usage:   "Returns features with pagination. Rules are returned as a unified top-level\narray with per-rule environment scope.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "client-key",
			Usage:     "Filter by a SDK connection's client key",
			QueryPath: "clientKey",
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
		&requestflag.Flag[any]{
			Name:      "skip-pagination",
			Usage:     "If true, return all matching items and ignore limit/offset.\nSelf-hosted only. Has no effect unless API_ALLOW_SKIP_PAGINATION is set to true or 1.",
			Default:   false,
			QueryPath: "skipPagination",
		},
	},
	Action:          handleFeaturesList,
	HideHelpCommand: true,
}

var featuresDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently deletes a feature and all of its revisions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleFeaturesDelete,
	HideHelpCommand: true,
}

var featuresRevert = cli.Command{
	Name:    "revert",
	Usage:   "Revert a feature to a specific revision",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[float64]{
			Name:     "revision",
			Required: true,
			BodyPath: "revision",
		},
		&requestflag.Flag[string]{
			Name:     "comment",
			BodyPath: "comment",
		},
	},
	Action:          handleFeaturesRevert,
	HideHelpCommand: true,
}

var featuresToggle = cli.Command{
	Name:    "toggle",
	Usage:   "Enables or disables a feature in one or more environments simultaneously.\nAccepts a map of environment name → boolean.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "environments",
			Required: true,
			BodyPath: "environments",
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			BodyPath: "reason",
		},
	},
	Action:          handleFeaturesToggle,
	HideHelpCommand: true,
}

func handleFeaturesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.FeatureNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Features.New(ctx, params, options...)
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
		Title:          "features create",
		Transform:      transform,
	})
}

func handleFeaturesRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.FeatureGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Features.Get(
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
		Title:          "features retrieve",
		Transform:      transform,
	})
}

func handleFeaturesUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.FeatureUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Features.Update(
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
		Title:          "features update",
		Transform:      transform,
	})
}

func handleFeaturesList(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.FeatureListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Features.List(ctx, params, options...)
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
		Title:          "features list",
		Transform:      transform,
	})
}

func handleFeaturesDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Features.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "features delete",
		Transform:      transform,
	})
}

func handleFeaturesRevert(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.FeatureRevertParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Features.Revert(
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
		Title:          "features revert",
		Transform:      transform,
	})
}

func handleFeaturesToggle(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.FeatureToggleParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Features.Toggle(
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
		Title:          "features toggle",
		Transform:      transform,
	})
}
