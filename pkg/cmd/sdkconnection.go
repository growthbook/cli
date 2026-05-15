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

var sdkConnectionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a single sdk connection",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "environment",
			Required: true,
			BodyPath: "environment",
		},
		&requestflag.Flag[string]{
			Name:     "language",
			Required: true,
			BodyPath: "language",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "allowed-custom-fields-in-metadata",
			BodyPath: "allowedCustomFieldsInMetadata",
		},
		&requestflag.Flag[bool]{
			Name:     "encrypt-payload",
			BodyPath: "encryptPayload",
		},
		&requestflag.Flag[bool]{
			Name:     "hash-secure-attributes",
			BodyPath: "hashSecureAttributes",
		},
		&requestflag.Flag[bool]{
			Name:     "include-custom-fields-in-metadata",
			BodyPath: "includeCustomFieldsInMetadata",
		},
		&requestflag.Flag[bool]{
			Name:     "include-draft-experiments",
			BodyPath: "includeDraftExperiments",
		},
		&requestflag.Flag[bool]{
			Name:     "include-experiment-names",
			BodyPath: "includeExperimentNames",
		},
		&requestflag.Flag[bool]{
			Name:     "include-project-id-in-metadata",
			BodyPath: "includeProjectIdInMetadata",
		},
		&requestflag.Flag[bool]{
			Name:     "include-redirect-experiments",
			BodyPath: "includeRedirectExperiments",
		},
		&requestflag.Flag[bool]{
			Name:     "include-rule-ids",
			BodyPath: "includeRuleIds",
		},
		&requestflag.Flag[bool]{
			Name:     "include-tags-in-metadata",
			BodyPath: "includeTagsInMetadata",
		},
		&requestflag.Flag[bool]{
			Name:     "include-visual-experiments",
			BodyPath: "includeVisualExperiments",
		},
		&requestflag.Flag[[]string]{
			Name:     "project",
			BodyPath: "projects",
		},
		&requestflag.Flag[bool]{
			Name:     "proxy-enabled",
			BodyPath: "proxyEnabled",
		},
		&requestflag.Flag[string]{
			Name:     "proxy-host",
			BodyPath: "proxyHost",
		},
		&requestflag.Flag[bool]{
			Name:     "remote-eval-enabled",
			BodyPath: "remoteEvalEnabled",
		},
		&requestflag.Flag[bool]{
			Name:     "saved-group-references-enabled",
			BodyPath: "savedGroupReferencesEnabled",
		},
		&requestflag.Flag[string]{
			Name:     "sdk-version",
			BodyPath: "sdkVersion",
		},
	},
	Action:          handleSDKConnectionsCreate,
	HideHelpCommand: true,
}

var sdkConnectionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single sdk connection",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleSDKConnectionsRetrieve,
	HideHelpCommand: true,
}

var sdkConnectionsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a single sdk connection",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[[]string]{
			Name:     "allowed-custom-fields-in-metadata",
			BodyPath: "allowedCustomFieldsInMetadata",
		},
		&requestflag.Flag[bool]{
			Name:     "encrypt-payload",
			BodyPath: "encryptPayload",
		},
		&requestflag.Flag[string]{
			Name:     "environment",
			BodyPath: "environment",
		},
		&requestflag.Flag[bool]{
			Name:     "hash-secure-attributes",
			BodyPath: "hashSecureAttributes",
		},
		&requestflag.Flag[bool]{
			Name:     "include-custom-fields-in-metadata",
			BodyPath: "includeCustomFieldsInMetadata",
		},
		&requestflag.Flag[bool]{
			Name:     "include-draft-experiments",
			BodyPath: "includeDraftExperiments",
		},
		&requestflag.Flag[bool]{
			Name:     "include-experiment-names",
			BodyPath: "includeExperimentNames",
		},
		&requestflag.Flag[bool]{
			Name:     "include-project-id-in-metadata",
			BodyPath: "includeProjectIdInMetadata",
		},
		&requestflag.Flag[bool]{
			Name:     "include-redirect-experiments",
			BodyPath: "includeRedirectExperiments",
		},
		&requestflag.Flag[bool]{
			Name:     "include-rule-ids",
			BodyPath: "includeRuleIds",
		},
		&requestflag.Flag[bool]{
			Name:     "include-tags-in-metadata",
			BodyPath: "includeTagsInMetadata",
		},
		&requestflag.Flag[bool]{
			Name:     "include-visual-experiments",
			BodyPath: "includeVisualExperiments",
		},
		&requestflag.Flag[string]{
			Name:     "language",
			BodyPath: "language",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "project",
			BodyPath: "projects",
		},
		&requestflag.Flag[bool]{
			Name:     "proxy-enabled",
			BodyPath: "proxyEnabled",
		},
		&requestflag.Flag[string]{
			Name:     "proxy-host",
			BodyPath: "proxyHost",
		},
		&requestflag.Flag[bool]{
			Name:     "remote-eval-enabled",
			BodyPath: "remoteEvalEnabled",
		},
		&requestflag.Flag[bool]{
			Name:     "saved-group-references-enabled",
			BodyPath: "savedGroupReferencesEnabled",
		},
		&requestflag.Flag[string]{
			Name:     "sdk-version",
			BodyPath: "sdkVersion",
		},
	},
	Action:          handleSDKConnectionsUpdate,
	HideHelpCommand: true,
}

var sdkConnectionsList = cli.Command{
	Name:    "list",
	Usage:   "Get all sdk connections",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "The number of items to return",
			Default:   10,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "multi-org",
			QueryPath: "multiOrg",
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
		&requestflag.Flag[string]{
			Name:      "with-proxy",
			QueryPath: "withProxy",
		},
	},
	Action:          handleSDKConnectionsList,
	HideHelpCommand: true,
}

var sdkConnectionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a single SDK connection",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The id of the requested resource",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleSDKConnectionsDelete,
	HideHelpCommand: true,
}

var sdkConnectionsLookup = cli.Command{
	Name:    "lookup",
	Usage:   "Find a single sdk connection by its key",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "key",
			Usage:     "The key of the requested sdkConnection",
			Required:  true,
			PathParam: "key",
		},
	},
	Action:          handleSDKConnectionsLookup,
	HideHelpCommand: true,
}

func handleSDKConnectionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.SDKConnectionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.SDKConnections.New(ctx, params, options...)
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
		Title:          "sdk-connections create",
		Transform:      transform,
	})
}

func handleSDKConnectionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.SDKConnections.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "sdk-connections retrieve",
		Transform:      transform,
	})
}

func handleSDKConnectionsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.SDKConnectionUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.SDKConnections.Update(
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
		Title:          "sdk-connections update",
		Transform:      transform,
	})
}

func handleSDKConnectionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.SDKConnectionListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.SDKConnections.List(ctx, params, options...)
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
		Title:          "sdk-connections list",
		Transform:      transform,
	})
}

func handleSDKConnectionsDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.SDKConnections.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "sdk-connections delete",
		Transform:      transform,
	})
}

func handleSDKConnectionsLookup(ctx context.Context, cmd *cli.Command) error {
	client := growthbook.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("key") && len(unusedArgs) > 0 {
		cmd.Set("key", unusedArgs[0])
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
	_, err = client.SDKConnections.Lookup(ctx, cmd.Value("key").(string), options...)
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
		Title:          "sdk-connections lookup",
		Transform:      transform,
	})
}
