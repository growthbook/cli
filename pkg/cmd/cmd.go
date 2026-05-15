// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/stainless-sdks/growthbook-cli/internal/autocomplete"
	"github.com/stainless-sdks/growthbook-cli/internal/requestflag"
	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command            *cli.Command
	CommandErrorBuffer bytes.Buffer
)

func init() {
	Command = &cli.Command{
		Name:      "growthbook",
		Usage:     "CLI for the growthbook API",
		Suggest:   true,
		Version:   Version,
		ErrWriter: &CommandErrorBuffer,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
				Validator: func(baseURL string) error {
					return ValidateBaseURL(baseURL, "--base-url")
				},
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
			&cli.BoolFlag{
				Name:    "raw-output",
				Aliases: []string{"r"},
				Usage:   "If the result is a string, print it without JSON quotes. This can be useful for making output transforms talk to non-JSON-based systems.",
			},
			&requestflag.Flag[string]{
				Name:    "api-key",
				Usage:   "If using Bearer auth, pass the Secret Key as the token:\n```bash\ncurl https://api.growthbook.io/api/v1/features   -H \"Authorization: Bearer secret_abc123DEF456\"\n```\n",
				Sources: cli.EnvVars("GROWTHBOOK_API_KEY"),
			},
			&requestflag.Flag[string]{
				Name:    "username",
				Usage:   "If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:\n```bash\ncurl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:\n# The \":\" at the end stops curl from asking for a password\n```\n",
				Sources: cli.EnvVars("GROWTHBOOK_USERNAME"),
			},
			&requestflag.Flag[string]{
				Name:    "password",
				Usage:   "If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:\n```bash\ncurl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:\n# The \":\" at the end stops curl from asking for a password\n```\n",
				Sources: cli.EnvVars("GROWTHBOOK_PASSWORD"),
			},
			&requestflag.Flag[string]{
				Name:    "domain",
				Sources: cli.EnvVars("GROWTHBOOK_DOMAIN"),
			},
			&cli.StringFlag{
				Name:  "environment",
				Usage: "Set the environment for API requests",
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "features",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresCreate,
					&featuresRetrieve,
					&featuresUpdate,
					&featuresList,
					&featuresDelete,
				},
			},
			{
				Name:     "features:toggle",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresToggleToggle,
				},
			},
			{
				Name:     "features:revert",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresRevertRevert,
				},
			},
			{
				Name:     "features:revisions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresRevisionsCreate,
					&featuresRevisionsRetrieve,
					&featuresRevisionsList,
				},
			},
			{
				Name:     "features:revisions:latest",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresRevisionsLatestList,
				},
			},
			{
				Name:     "features:revisions:metadata",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresRevisionsMetadataCreate,
				},
			},
			{
				Name:     "features:revisions:default-value",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresRevisionsDefaultValueUpdateDefaultValue,
				},
			},
			{
				Name:     "features:revisions:prerequisites",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresRevisionsPrerequisitesCreate,
				},
			},
			{
				Name:     "features:revisions:holdout",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresRevisionsHoldoutCreate,
				},
			},
			{
				Name:     "features:revisions:archive",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresRevisionsArchiveCreate,
				},
			},
			{
				Name:     "features:revisions:toggle",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresRevisionsToggleCreate,
				},
			},
			{
				Name:     "features:revisions:rules",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresRevisionsRulesCreate,
					&featuresRevisionsRulesUpdate,
					&featuresRevisionsRulesDelete,
				},
			},
			{
				Name:     "features:revisions:rules:reorder",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresRevisionsRulesReorderCreate,
				},
			},
			{
				Name:     "features:revisions:rules:ramp-schedule",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresRevisionsRulesRampScheduleDeleteRampSchedule,
					&featuresRevisionsRulesRampScheduleUpdateRampSchedule,
				},
			},
			{
				Name:     "features:revisions:request-review",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresRevisionsRequestReviewRequestReview,
				},
			},
			{
				Name:     "features:revisions:submit-review",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresRevisionsSubmitReviewSubmitReview,
				},
			},
			{
				Name:     "features:revisions:merge-status",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresRevisionsMergeStatusRetrieveMergeStatus,
				},
			},
			{
				Name:     "features:revisions:rebase",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresRevisionsRebaseCreate,
				},
			},
			{
				Name:     "features:revisions:publish",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresRevisionsPublishCreate,
				},
			},
			{
				Name:     "features:revisions:discard",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresRevisionsDiscardCreate,
				},
			},
			{
				Name:     "features:revisions:revert",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&featuresRevisionsRevertCreate,
				},
			},
			{
				Name:     "revisions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&revisionsList,
				},
			},
			{
				Name:     "archetypes",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&archetypesCreate,
					&archetypesRetrieve,
					&archetypesUpdate,
					&archetypesList,
					&archetypesDelete,
				},
			},
			{
				Name:     "experiments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&experimentsCreate,
					&experimentsRetrieve,
					&experimentsUpdate,
					&experimentsList,
					&experimentsCreateSnapshot,
					&experimentsModifyTemporaryRollout,
					&experimentsStart,
					&experimentsStop,
				},
			},
			{
				Name:     "experiments:results",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&experimentsResultsRetrieve,
					&experimentsResultsList,
				},
			},
			{
				Name:     "experiments:variation:screenshot",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&experimentsVariationScreenshotDelete,
					&experimentsVariationScreenshotUpload,
				},
			},
			{
				Name:     "experiments:visual-changesets",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&experimentsVisualChangesetsCreate,
					&experimentsVisualChangesetsList,
				},
			},
			{
				Name:     "experiment-names",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&experimentNamesList,
				},
			},
			{
				Name:     "snapshots",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&snapshotsRetrieve,
				},
			},
			{
				Name:     "metrics",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&metricsCreate,
					&metricsRetrieve,
					&metricsUpdate,
					&metricsList,
					&metricsDelete,
				},
			},
			{
				Name:     "usage",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usageRetrieveMetrics,
				},
			},
			{
				Name:     "segments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&segmentsCreate,
					&segmentsRetrieve,
					&segmentsUpdate,
					&segmentsList,
					&segmentsDelete,
				},
			},
			{
				Name:     "dimensions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&dimensionsCreate,
					&dimensionsRetrieve,
					&dimensionsUpdate,
					&dimensionsList,
					&dimensionsDelete,
				},
			},
			{
				Name:     "projects",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&projectsCreate,
					&projectsRetrieve,
					&projectsUpdate,
					&projectsList,
					&projectsDelete,
				},
			},
			{
				Name:     "environments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&environmentsCreate,
					&environmentsUpdate,
					&environmentsList,
					&environmentsDelete,
				},
			},
			{
				Name:     "attributes",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&attributesCreate,
					&attributesUpdate,
					&attributesList,
					&attributesDelete,
				},
			},
			{
				Name:     "sdk-connections",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&sdkConnectionsCreate,
					&sdkConnectionsRetrieve,
					&sdkConnectionsUpdate,
					&sdkConnectionsList,
					&sdkConnectionsDelete,
					&sdkConnectionsLookup,
				},
			},
			{
				Name:     "data-sources",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&dataSourcesRetrieve,
					&dataSourcesList,
					&dataSourcesGetInformationSchema,
				},
			},
			{
				Name:     "visual-changesets",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&visualChangesetsRetrieve,
					&visualChangesetsUpdate,
				},
			},
			{
				Name:     "visual-changesets:visual-change",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&visualChangesetsVisualChangeCreate,
					&visualChangesetsVisualChangeUpdate,
				},
			},
			{
				Name:     "saved-groups",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&savedGroupsCreate,
					&savedGroupsRetrieve,
					&savedGroupsUpdate,
					&savedGroupsList,
					&savedGroupsDelete,
					&savedGroupsArchive,
					&savedGroupsUnarchive,
				},
			},
			{
				Name:     "organizations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&organizationsCreate,
					&organizationsUpdate,
					&organizationsList,
				},
			},
			{
				Name:     "sdk-payload",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&sdkPayloadRetrieve,
				},
			},
			{
				Name:     "fact-tables",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&factTablesCreate,
					&factTablesRetrieve,
					&factTablesUpdate,
					&factTablesList,
					&factTablesDelete,
				},
			},
			{
				Name:     "fact-tables:filters",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&factTablesFiltersCreate,
					&factTablesFiltersRetrieve,
					&factTablesFiltersUpdate,
					&factTablesFiltersList,
					&factTablesFiltersDelete,
				},
			},
			{
				Name:     "fact-metrics",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&factMetricsCreate,
					&factMetricsRetrieve,
					&factMetricsUpdate,
					&factMetricsList,
					&factMetricsDelete,
					&factMetricsCreateAnalysis,
				},
			},
			{
				Name:     "bulk-import",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&bulkImportCreateFacts,
				},
			},
			{
				Name:     "code-refs",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&codeRefsCreate,
					&codeRefsRetrieve,
					&codeRefsList,
				},
			},
			{
				Name:     "members",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&membersList,
					&membersDelete,
					&membersUpdateRole,
				},
			},
			{
				Name:     "queries",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&queriesRetrieve,
				},
			},
			{
				Name:     "settings",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&settingsRetrieve,
				},
			},
			{
				Name:     "information-schema-tables",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&informationSchemaTablesRetrieve,
				},
			},
			{
				Name:     "ramp-schedules",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&rampSchedulesCreate,
					&rampSchedulesRetrieve,
					&rampSchedulesUpdate,
					&rampSchedulesList,
					&rampSchedulesDelete,
				},
			},
			{
				Name:     "ramp-schedules:actions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&rampSchedulesActionsAddTarget,
					&rampSchedulesActionsApproveStep,
					&rampSchedulesActionsComplete,
					&rampSchedulesActionsEjectTarget,
					&rampSchedulesActionsJump,
					&rampSchedulesActionsPause,
					&rampSchedulesActionsResume,
					&rampSchedulesActionsRollback,
					&rampSchedulesActionsStart,
				},
			},
			{
				Name:     "namespaces",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&namespacesCreate,
					&namespacesRetrieve,
					&namespacesUpdate,
					&namespacesList,
					&namespacesDelete,
					&namespacesGetMemberships,
					&namespacesRotateSeed,
				},
			},
			{
				Name:     "transform-copy",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&transformCopyCreate,
				},
			},
			{
				Name:     "dashboards",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&dashboardsCreate,
					&dashboardsRetrieve,
					&dashboardsUpdate,
					&dashboardsList,
					&dashboardsDelete,
					&dashboardsListByExperiment,
				},
			},
			{
				Name:     "custom-fields",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&customFieldsCreate,
					&customFieldsRetrieve,
					&customFieldsUpdate,
					&customFieldsList,
					&customFieldsDelete,
				},
			},
			{
				Name:     "metric-groups",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&metricGroupsCreate,
					&metricGroupsRetrieve,
					&metricGroupsUpdate,
					&metricGroupsList,
					&metricGroupsDelete,
				},
			},
			{
				Name:     "teams",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&teamsCreate,
					&teamsRetrieve,
					&teamsUpdate,
					&teamsList,
					&teamsDelete,
				},
			},
			{
				Name:     "teams:members",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&teamsMembersAdd,
					&teamsMembersRemove,
				},
			},
			{
				Name:     "experiment-templates",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&experimentTemplatesCreate,
					&experimentTemplatesRetrieve,
					&experimentTemplatesUpdate,
					&experimentTemplatesList,
					&experimentTemplatesDelete,
					&experimentTemplatesBulkImport,
				},
			},
			{
				Name:     "product-analytics",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&productAnalyticsCreateDataSourceExploration,
					&productAnalyticsCreateMetricExploration,
					&productAnalyticsRunFactTableExploration,
				},
			},
			{
				Name:     "ramp-schedule-templates",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&rampScheduleTemplatesCreate,
					&rampScheduleTemplatesRetrieve,
					&rampScheduleTemplatesUpdate,
					&rampScheduleTemplatesList,
					&rampScheduleTemplatesDelete,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "growthbook @manpages [-o growthbook.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
		},
		HideHelpCommand: true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "growthbook.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "growthbook.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
