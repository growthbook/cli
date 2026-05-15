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

var experimentTemplatesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a single experimentTemplate",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "datasource",
			Required: true,
			BodyPath: "datasource",
		},
		&requestflag.Flag[string]{
			Name:     "exposure-query-id",
			Required: true,
			BodyPath: "exposureQueryId",
		},
		&requestflag.Flag[string]{
			Name:     "stats-engine",
			Usage:    `Allowed values: "bayesian", "frequentist".`,
			Required: true,
			BodyPath: "statsEngine",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "targeting",
			Required: true,
			BodyPath: "targeting",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "template-metadata",
			Required: true,
			BodyPath: "templateMetadata",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "standard".`,
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "activation-metric",
			BodyPath: "activationMetric",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "custom-fields",
			BodyPath: "customFields",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "custom-metric-slice",
			BodyPath: "customMetricSlices",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[bool]{
			Name:     "disable-sticky-bucketing",
			BodyPath: "disableStickyBucketing",
		},
		&requestflag.Flag[string]{
			Name:     "fallback-attribute",
			BodyPath: "fallbackAttribute",
		},
		&requestflag.Flag[[]string]{
			Name:     "goal-metric",
			BodyPath: "goalMetrics",
		},
		&requestflag.Flag[[]string]{
			Name:     "guardrail-metric",
			BodyPath: "guardrailMetrics",
		},
		&requestflag.Flag[string]{
			Name:     "hash-attribute",
			BodyPath: "hashAttribute",
		},
		&requestflag.Flag[string]{
			Name:     "hypothesis",
			BodyPath: "hypothesis",
		},
		&requestflag.Flag[string]{
			Name:     "project",
			BodyPath: "project",
		},
		&requestflag.Flag[[]string]{
			Name:     "secondary-metric",
			BodyPath: "secondaryMetrics",
		},
		&requestflag.Flag[string]{
			Name:     "segment",
			BodyPath: "segment",
		},
		&requestflag.Flag[bool]{
			Name:     "skip-partial-data",
			BodyPath: "skipPartialData",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
	},
	Action:          handleExperimentTemplatesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"targeting": {
		&requestflag.InnerFlag[string]{
			Name:       "targeting.condition",
			InnerField: "condition",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "targeting.coverage",
			InnerField: "coverage",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "targeting.prerequisites",
			InnerField: "prerequisites",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "targeting.saved-groups",
			InnerField: "savedGroups",
		},
	},
	"template-metadata": {
		&requestflag.InnerFlag[string]{
			Name:       "template-metadata.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "template-metadata.description",
			InnerField: "description",
		},
	},
	"custom-metric-slice": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "custom-metric-slice.slices",
			InnerField: "slices",
		},
	},
})

var experimentTemplatesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single experimentTemplate",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleExperimentTemplatesRetrieve,
	HideHelpCommand: true,
}

var experimentTemplatesUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a single experimentTemplate",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "activation-metric",
			BodyPath: "activationMetric",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "custom-fields",
			BodyPath: "customFields",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "custom-metric-slice",
			BodyPath: "customMetricSlices",
		},
		&requestflag.Flag[string]{
			Name:     "datasource",
			BodyPath: "datasource",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[bool]{
			Name:     "disable-sticky-bucketing",
			BodyPath: "disableStickyBucketing",
		},
		&requestflag.Flag[string]{
			Name:     "exposure-query-id",
			BodyPath: "exposureQueryId",
		},
		&requestflag.Flag[string]{
			Name:     "fallback-attribute",
			BodyPath: "fallbackAttribute",
		},
		&requestflag.Flag[[]string]{
			Name:     "goal-metric",
			BodyPath: "goalMetrics",
		},
		&requestflag.Flag[[]string]{
			Name:     "guardrail-metric",
			BodyPath: "guardrailMetrics",
		},
		&requestflag.Flag[string]{
			Name:     "hash-attribute",
			BodyPath: "hashAttribute",
		},
		&requestflag.Flag[string]{
			Name:     "hypothesis",
			BodyPath: "hypothesis",
		},
		&requestflag.Flag[string]{
			Name:     "project",
			BodyPath: "project",
		},
		&requestflag.Flag[[]string]{
			Name:     "secondary-metric",
			BodyPath: "secondaryMetrics",
		},
		&requestflag.Flag[string]{
			Name:     "segment",
			BodyPath: "segment",
		},
		&requestflag.Flag[bool]{
			Name:     "skip-partial-data",
			BodyPath: "skipPartialData",
		},
		&requestflag.Flag[string]{
			Name:     "stats-engine",
			Usage:    `Allowed values: "bayesian", "frequentist".`,
			BodyPath: "statsEngine",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			BodyPath: "tags",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "targeting",
			BodyPath: "targeting",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "template-metadata",
			BodyPath: "templateMetadata",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "standard".`,
			BodyPath: "type",
		},
	},
	Action:          handleExperimentTemplatesUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"custom-metric-slice": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "custom-metric-slice.slices",
			InnerField: "slices",
		},
	},
	"targeting": {
		&requestflag.InnerFlag[string]{
			Name:       "targeting.condition",
			InnerField: "condition",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "targeting.coverage",
			InnerField: "coverage",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "targeting.prerequisites",
			InnerField: "prerequisites",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "targeting.saved-groups",
			InnerField: "savedGroups",
		},
	},
	"template-metadata": {
		&requestflag.InnerFlag[string]{
			Name:       "template-metadata.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "template-metadata.description",
			InnerField: "description",
		},
	},
})

var experimentTemplatesList = cli.Command{
	Name:    "list",
	Usage:   "Get all experimentTemplates",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "project-id",
			QueryPath: "projectId",
		},
	},
	Action:          handleExperimentTemplatesList,
	HideHelpCommand: true,
}

var experimentTemplatesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a single experimentTemplate",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleExperimentTemplatesDelete,
	HideHelpCommand: true,
}

var experimentTemplatesBulkImport = requestflag.WithInnerFlags(cli.Command{
	Name:    "bulk-import",
	Usage:   "Bulk create or update experiment templates",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "template",
			Required: true,
			BodyPath: "templates",
		},
	},
	Action:          handleExperimentTemplatesBulkImport,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"template": {
		&requestflag.InnerFlag[string]{
			Name:       "template.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "template.data",
			InnerField: "data",
		},
	},
})

func handleExperimentTemplatesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.ExperimentTemplateNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ExperimentTemplates.New(ctx, params, options...)
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
		Title:          "experiment-templates create",
		Transform:      transform,
	})
}

func handleExperimentTemplatesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ExperimentTemplates.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "experiment-templates retrieve",
		Transform:      transform,
	})
}

func handleExperimentTemplatesUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.ExperimentTemplateUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ExperimentTemplates.Update(
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
		Title:          "experiment-templates update",
		Transform:      transform,
	})
}

func handleExperimentTemplatesList(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.ExperimentTemplateListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ExperimentTemplates.List(ctx, params, options...)
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
		Title:          "experiment-templates list",
		Transform:      transform,
	})
}

func handleExperimentTemplatesDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ExperimentTemplates.Delete(ctx, cmd.Value("id").(string), options...)
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
		Title:          "experiment-templates delete",
		Transform:      transform,
	})
}

func handleExperimentTemplatesBulkImport(ctx context.Context, cmd *cli.Command) error {
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

	params := growthbook.ExperimentTemplateBulkImportParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ExperimentTemplates.BulkImport(ctx, params, options...)
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
		Title:          "experiment-templates bulk-import",
		Transform:      transform,
	})
}
