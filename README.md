# growthbook

Command-line interface for the *GrowthBook REST* API.

[![Built by Speakeasy](https://img.shields.io/badge/Built_by-SPEAKEASY-374151?style=for-the-badge&labelColor=f3f4f6)](https://www.speakeasy.com/?utm_source=github-com/growthbook/cli&utm_campaign=cli)
[![License: MIT](https://img.shields.io/badge/LICENSE_//_MIT-3b5bdb?style=for-the-badge&labelColor=eff6ff)](https://opensource.org/licenses/MIT)


<br /><br />
> [!IMPORTANT]
> This CLI is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/growthbook/growthbook). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

GrowthBook REST API: GrowthBook offers a full REST API for interacting with the application.

Request data can use either JSON or Form data encoding (with proper `Content-Type` headers). All response bodies are JSON-encoded.

The API base URL for GrowthBook Cloud is `https://api.growthbook.io/api`. For self-hosted deployments, it is the same as your API_HOST environment variable (defaults to `http://localhost:3100/api`). The rest of these docs will assume you are using GrowthBook Cloud.

## Versioning

Endpoints are versioned by path prefix:

- `/v1/...` — stable, widely-supported endpoints
- `/v2/...` — updated endpoints with improved shapes (e.g. unified per-rule environment scope for feature flags)

New integrations should prefer v2 where available.

## Authentication

We support both the HTTP Basic and Bearer authentication schemes for convenience.

You first need to generate a new API Key in GrowthBook. Different keys have different permissions:

- **Personal Access Tokens**: These are sensitive and provide the same level of access as the user has to an organization. These can be created by going to `Personal Access Tokens` under the your user menu.
- **Secret Keys**: These are sensitive and provide the level of access for the role, which currently is either `admin` or `readonly`. Only Admins with the `manageApiKeys` permission can manage Secret Keys on behalf of an organization. These can be created by going to `Settings -> API Keys`

If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank (when using curl, add `:` at the end of the secret to indicate an empty password)

```bash
curl https://api.growthbook.io/api/v1/features \
  -u secret_abc123DEF456:
```

If using Bearer auth, pass the Secret Key as the token:

```bash
curl https://api.growthbook.io/api/v1/features \
-H "Authorization: Bearer secret_abc123DEF456"
```

## Errors

The API may return the following error status codes:

- **400** - Bad Request - Often due to a missing required parameter
- **401** - Unauthorized - No valid API key provided
- **402** - Request Failed - The parameters are valid, but the request failed
- **403** - Forbidden - Provided API key does not have the required access
- **404** - Not Found - Unknown API route or requested resource
- **429** - Too Many Requests - You exceeded the rate limit of 60 requests per minute. Try again later.
- **5XX** - Server Error - Something went wrong on GrowthBook's end (these are rare)

The response body will be a JSON object with the following properties:

- **message** - Information about the error
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [growthbook](#growthbook)
  * [Versioning](#versioning)
  * [Authentication](#authentication)
  * [Errors](#errors)
  * [CLI Installation](#cli-installation)
  * [Shell Completion](#shell-completion)
  * [CLI Example Usage](#cli-example-usage)
  * [Authentication](#authentication-1)
  * [Available Commands](#available-commands)
  * [Request Body Input](#request-body-input)
  * [Server Selection](#server-selection)
  * [Output Formats](#output-formats)
  * [Error Handling](#error-handling)
  * [Diagnostics](#diagnostics)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start CLI Installation [installation] -->
## CLI Installation

### Quick Install (Linux/macOS)

```bash
curl -fsSL https://raw.githubusercontent.com/growthbook/cli/main/scripts/install.sh | bash
```

### Quick Install (Windows PowerShell)

```powershell
iwr -useb https://raw.githubusercontent.com/growthbook/cli/main/scripts/install.ps1 | iex
```

### Go Install

Alternatively, install directly via Go:

```bash
go install github.com/growthbook/cli/cmd/growthbook@latest
```

### Manual Download

Download pre-built binaries for your platform from the [releases page](https://github.com/growthbook/cli/releases).
<!-- End CLI Installation [installation] -->

<!-- Start Shell Completion [completion] -->
## Shell Completion

Shell completions are available for Bash, Zsh, Fish, and PowerShell.

### Bash

```bash
# Add to ~/.bashrc:
source <(growthbook completion bash)

# Or install permanently:
growthbook completion bash > /etc/bash_completion.d/growthbook
```

### Zsh

```zsh
# Add to ~/.zshrc:
source <(growthbook completion zsh)

# Or install permanently:
growthbook completion zsh > "${fpath[1]}/_growthbook"
```

### Fish

```fish
growthbook completion fish | source

# Or install permanently:
growthbook completion fish > ~/.config/fish/completions/growthbook.fish
```

### PowerShell

```powershell
growthbook completion powershell | Out-String | Invoke-Expression
```
<!-- End Shell Completion [completion] -->

<!-- Start CLI Example Usage [usage] -->
## CLI Example Usage

### Example

```bash
growthbook get-SDK-payload --bearer-auth 'Bearer test_token' --key '<key>'

```
<!-- End CLI Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

Authentication credentials can be configured in four ways (in order of priority):

### 1. Command-line flags

Pass credentials directly as flags to any command:

```bash
growthbook --bearer-auth <value> --username <value> --password <value> <command> [arguments]
```

### 2. Environment variables

Set credentials via environment variables:

| Variable | Description |
|----------|-------------|
| `GB_BEARER_AUTH` | If using Bearer auth, pass the Secret Key as the token:
```bash
curl https://api.growthbook.io/api/v1/features   -H "Authorization: Bearer secret_abc123DEF456"
``` |
| `GB_USERNAME` | If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
```bash
curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
# The ":" at the end stops curl from asking for a password
```
 username |
| `GB_PASSWORD` | If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
```bash
curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
# The ":" at the end stops curl from asking for a password
```
 password |

### 3. OS Keychain (recommended for workstations)

Credentials are stored securely in your operating system's keychain when you run:

```bash
growthbook configure
```

Secret credentials (tokens, API keys, passwords) are automatically stored in:
- **macOS**: Keychain
- **Linux**: GNOME Keyring / KWallet (via D-Bus Secret Service)
- **Windows**: Windows Credential Locker

If no keychain is available (e.g., in CI environments), credentials fall back to the config file.

### 4. Configuration file

Run the interactive `configure` command to store non-secret settings:

```bash
growthbook configure
```

Configuration is stored in `~/.config/growthbook/config.yaml`.
<!-- End Authentication [security] -->

<!-- Start Available Commands [operations] -->
## Available Commands

<details open>
<summary>Available commands</summary>

* [`get-SDK-payload`](docs/growthbook_get-SDK-payload.md) - Get a SDK payload
* [`post-copy-transform`](docs/growthbook_post-copy-transform.md)

### [features-v2](docs/growthbook_features-v2.md)

* [`list`](docs/growthbook_features-v2_list.md) - Get all features
* [`post-feature-v2`](docs/growthbook_features-v2_post-feature-v2.md) - Create a single feature
* [`get-feature-v2`](docs/growthbook_features-v2_get-feature-v2.md) - Get a single feature
* [`update-feature-v2`](docs/growthbook_features-v2_update-feature-v2.md) - Partially update a feature
* [`delete-feature-v2`](docs/growthbook_features-v2_delete-feature-v2.md) - Deletes a single feature
* [`toggle-feature-v2`](docs/growthbook_features-v2_toggle-feature-v2.md) - Toggle a feature in one or more environments
* [`revert-feature-v2`](docs/growthbook_features-v2_revert-feature-v2.md) - Revert a feature to a specific revision
* [`get-feature-keys-v2`](docs/growthbook_features-v2_get-feature-keys-v2.md) - Get list of feature keys
* [`get-feature-stale-v2`](docs/growthbook_features-v2_get-feature-stale-v2.md) - Get stale status for one or more features

### [feature-revisions-v2](docs/growthbook_feature-revisions-v2.md)

* [`list-revisions-v2`](docs/growthbook_feature-revisions-v2_list-revisions-v2.md) - List feature revisions
* [`get`](docs/growthbook_feature-revisions-v2_get.md) - List revisions for a feature
* [`post-feature-revision-v2`](docs/growthbook_feature-revisions-v2_post-feature-revision-v2.md) - Create a draft revision
* [`get-feature-revision-latest-v2`](docs/growthbook_feature-revisions-v2_get-feature-revision-latest-v2.md) - Get the most recent active draft revision
* [`get-feature-revision-v2`](docs/growthbook_feature-revisions-v2_get-feature-revision-v2.md) - Get a single feature revision
* [`post-feature-revision-request-review-v2`](docs/growthbook_feature-revisions-v2_post-feature-revision-request-review-v2.md) - Request review for a draft revision
* [`post-feature-revision-submit-review-v2`](docs/growthbook_feature-revisions-v2_post-feature-revision-submit-review-v2.md) - Submit a review on a draft revision
* [`get-feature-revision-merge-status-v2`](docs/growthbook_feature-revisions-v2_get-feature-revision-merge-status-v2.md) - Get merge status for a draft revision
* [`post-feature-revision-rebase-v2`](docs/growthbook_feature-revisions-v2_post-feature-revision-rebase-v2.md) - Rebase a draft revision onto the current live version
* [`post-feature-revision-publish-v2`](docs/growthbook_feature-revisions-v2_post-feature-revision-publish-v2.md) - Publish a draft revision
* [`post-feature-revision-discard-v2`](docs/growthbook_feature-revisions-v2_post-feature-revision-discard-v2.md) - Discard a draft revision
* [`post-feature-revision-revert-v2`](docs/growthbook_feature-revisions-v2_post-feature-revision-revert-v2.md) - Revert the feature to a prior revision

### [archetypes](docs/growthbook_archetypes.md)

* [`list`](docs/growthbook_archetypes_list.md) - Get the organization's archetypes
* [`post`](docs/growthbook_archetypes_post.md) - Create a single archetype
* [`get`](docs/growthbook_archetypes_get.md) - Get a single archetype
* [`put`](docs/growthbook_archetypes_put.md) - Update a single archetype
* [`delete`](docs/growthbook_archetypes_delete.md) - Deletes a single archetype

### [experiments](docs/growthbook_experiments.md)

* [`list`](docs/growthbook_experiments_list.md) - Get all experiments
* [`post`](docs/growthbook_experiments_post.md) - Create a single experiment
* [`list-experiment-results`](docs/growthbook_experiments_list-experiment-results.md) - Get latest results for many experiments
* [`get`](docs/growthbook_experiments_get.md) - Get a single experiment
* [`update`](docs/growthbook_experiments_update.md) - Update a single experiment
* [`get-experiment-results`](docs/growthbook_experiments_get-experiment-results.md) - Get results for an experiment
* [`post-experiment-start`](docs/growthbook_experiments_post-experiment-start.md) - Start an experiment
* [`post-experiment-stop`](docs/growthbook_experiments_post-experiment-stop.md) - Stop an experiment
* [`post-experiment-modify-temporary-rollout`](docs/growthbook_experiments_post-experiment-modify-temporary-rollout.md) - Modify temporary rollout status for a stopped experiment
* [`post-experiment-snapshot`](docs/growthbook_experiments_post-experiment-snapshot.md) - Create Experiment Snapshot
* [`post-variation-image-upload`](docs/growthbook_experiments_post-variation-image-upload.md) - Upload a variation screenshot
* [`delete-variation-screenshot`](docs/growthbook_experiments_delete-variation-screenshot.md) - Delete a variation screenshot
* [`get-experiment-names`](docs/growthbook_experiments_get-experiment-names.md) - Get a list of experiments with names and ids

### [snapshots](docs/growthbook_snapshots.md)

* [`post-experiment`](docs/growthbook_snapshots_post-experiment.md) - Create Experiment Snapshot
* [`get-experiment`](docs/growthbook_snapshots_get-experiment.md) - Get an experiment snapshot status

### [visual-changesets](docs/growthbook_visual-changesets.md)

* [`list`](docs/growthbook_visual-changesets_list.md) - Get all visual changesets
* [`post`](docs/growthbook_visual-changesets_post.md) - Create a visual changeset for an experiment
* [`get`](docs/growthbook_visual-changesets_get.md) - Get a single visual changeset
* [`put`](docs/growthbook_visual-changesets_put.md) - Update a visual changeset
* [`post-visual-change`](docs/growthbook_visual-changesets_post-visual-change.md) - Create a visual change for a visual changeset
* [`put-visual-change`](docs/growthbook_visual-changesets_put-visual-change.md) - Update a visual change for a visual changeset

### [metrics](docs/growthbook_metrics.md)

* [`list`](docs/growthbook_metrics_list.md) - Get all metrics
* [`post`](docs/growthbook_metrics_post.md) - Create a single metric
* [`get`](docs/growthbook_metrics_get.md) - Get a single metric
* [`put`](docs/growthbook_metrics_put.md) - Update a metric
* [`delete`](docs/growthbook_metrics_delete.md) - Deletes a metric

### [usage](docs/growthbook_usage.md)

* [`get-metric`](docs/growthbook_usage_get-metric.md) - Get metric usage across experiments

### [segments](docs/growthbook_segments.md)

* [`list`](docs/growthbook_segments_list.md) - Get all segments
* [`post`](docs/growthbook_segments_post.md) - Create a single segment
* [`get`](docs/growthbook_segments_get.md) - Get a single segment
* [`update`](docs/growthbook_segments_update.md) - Update a single segment
* [`delete`](docs/growthbook_segments_delete.md) - Deletes a single segment

### [dimensions](docs/growthbook_dimensions.md)

* [`list`](docs/growthbook_dimensions_list.md) - Get all dimensions
* [`post`](docs/growthbook_dimensions_post.md) - Create a single dimension
* [`get`](docs/growthbook_dimensions_get.md) - Get a single dimension
* [`update`](docs/growthbook_dimensions_update.md) - Update a single dimension
* [`delete`](docs/growthbook_dimensions_delete.md) - Deletes a single dimension

### [projects](docs/growthbook_projects.md)

* [`list`](docs/growthbook_projects_list.md) - Get all projects
* [`post`](docs/growthbook_projects_post.md) - Create a single project
* [`get`](docs/growthbook_projects_get.md) - Get a single project
* [`put`](docs/growthbook_projects_put.md) - Edit a single project
* [`delete`](docs/growthbook_projects_delete.md) - Deletes a single project

### [environments](docs/growthbook_environments.md)

* [`list`](docs/growthbook_environments_list.md) - Get the organization's environments
* [`post`](docs/growthbook_environments_post.md) - Create a new environment
* [`put`](docs/growthbook_environments_put.md) - Update an environment
* [`delete`](docs/growthbook_environments_delete.md) - Deletes a single environment

### [attributes](docs/growthbook_attributes.md)

* [`list`](docs/growthbook_attributes_list.md) - Get the organization's attributes
* [`post`](docs/growthbook_attributes_post.md) - Create a new attribute
* [`put`](docs/growthbook_attributes_put.md) - Update an attribute
* [`delete`](docs/growthbook_attributes_delete.md) - Deletes a single attribute

### [SDK-connections](docs/growthbook_SDK-connections.md)

* [`list`](docs/growthbook_SDK-connections_list.md) - Get all sdk connections
* [`post`](docs/growthbook_SDK-connections_post.md) - Create a single sdk connection
* [`get`](docs/growthbook_SDK-connections_get.md) - Get a single sdk connection
* [`put`](docs/growthbook_SDK-connections_put.md) - Update a single sdk connection
* [`delete`](docs/growthbook_SDK-connections_delete.md) - Deletes a single SDK connection
* [`lookup-SDK-connection-by-key`](docs/growthbook_SDK-connections_lookup-SDK-connection-by-key.md) - Find a single sdk connection by its key

### [data-sources](docs/growthbook_data-sources.md)

* [`list`](docs/growthbook_data-sources_list.md) - Get all data sources
* [`get`](docs/growthbook_data-sources_get.md) - Get a single data source
* [`get-information-schema`](docs/growthbook_data-sources_get-information-schema.md) - Get a Data Source's Information Schema
* [`get-information-schema-table`](docs/growthbook_data-sources_get-information-schema-table.md) - Get a single Information Schema Table by id

### [saved-groups](docs/growthbook_saved-groups.md)

* [`list`](docs/growthbook_saved-groups_list.md) - Get all saved group
* [`post`](docs/growthbook_saved-groups_post.md) - Create a single saved group
* [`get`](docs/growthbook_saved-groups_get.md) - Get a single saved group
* [`update`](docs/growthbook_saved-groups_update.md) - Partially update a single saved group
* [`delete`](docs/growthbook_saved-groups_delete.md) - Deletes a single saved group
* [`archive`](docs/growthbook_saved-groups_archive.md) - Archive a single saved group
* [`unarchive`](docs/growthbook_saved-groups_unarchive.md) - Unarchive a single saved group

### [organizations](docs/growthbook_organizations.md)

* [`list`](docs/growthbook_organizations_list.md) - Get all organizations (only for super admins on multi-org Enterprise Plan only)
* [`post`](docs/growthbook_organizations_post.md) - Create a single organization (only for super admins on multi-org Enterprise Plan only)
* [`put`](docs/growthbook_organizations_put.md) - Edit a single organization (only for super admins on multi-org Enterprise Plan only)

### [fact-tables](docs/growthbook_fact-tables.md)

* [`list`](docs/growthbook_fact-tables_list.md) - Get all fact tables
* [`post`](docs/growthbook_fact-tables_post.md) - Create a single fact table
* [`get`](docs/growthbook_fact-tables_get.md) - Get a single fact table
* [`update`](docs/growthbook_fact-tables_update.md) - Update a single fact table
* [`delete`](docs/growthbook_fact-tables_delete.md) - Deletes a single fact table
* [`list-fact-table-filters`](docs/growthbook_fact-tables_list-fact-table-filters.md) - Get all filters for a fact table
* [`post-fact-table-filter`](docs/growthbook_fact-tables_post-fact-table-filter.md) - Create a single fact table filter
* [`get-fact-table-filter`](docs/growthbook_fact-tables_get-fact-table-filter.md) - Get a single fact filter
* [`update-fact-table-filter`](docs/growthbook_fact-tables_update-fact-table-filter.md) - Update a single fact table filter
* [`delete-fact-table-filter`](docs/growthbook_fact-tables_delete-fact-table-filter.md) - Deletes a single fact table filter
* [`post-bulk-import-facts`](docs/growthbook_fact-tables_post-bulk-import-facts.md) - Bulk import fact tables, filters, and metrics

### [fact-metrics](docs/growthbook_fact-metrics.md)

* [`list`](docs/growthbook_fact-metrics_list.md) - Get all fact metrics
* [`post`](docs/growthbook_fact-metrics_post.md) - Create a single fact metric
* [`get`](docs/growthbook_fact-metrics_get.md) - Get a single fact metric
* [`update`](docs/growthbook_fact-metrics_update.md) - Update a single fact metric
* [`delete`](docs/growthbook_fact-metrics_delete.md) - Deletes a single fact metric
* [`post-fact-metric-analysis`](docs/growthbook_fact-metrics_post-fact-metric-analysis.md) - Create a fact metric analysis

### [code-references](docs/growthbook_code-references.md)

* [`post-code-refs`](docs/growthbook_code-references_post-code-refs.md) - Submit list of code references
* [`list-code-refs`](docs/growthbook_code-references_list-code-refs.md) - Get list of all code references for the current organization
* [`get-code-refs`](docs/growthbook_code-references_get-code-refs.md) - Get list of code references for a single feature id

### [members](docs/growthbook_members.md)

* [`list`](docs/growthbook_members_list.md) - Get all organization members
* [`update-member-role`](docs/growthbook_members_update-member-role.md) - Update a member's global role (including any enviroment restrictions, if applicable). Can also update a member's project roles if your plan supports it.
* [`delete`](docs/growthbook_members_delete.md) - Removes a single user from an organization

### [queries](docs/growthbook_queries.md)

* [`get-query`](docs/growthbook_queries_get-query.md) - Get a single query

### [settings](docs/growthbook_settings.md)

* [`get`](docs/growthbook_settings_get.md) - Get organization settings

### [ramp-schedules](docs/growthbook_ramp-schedules.md)

* [`list`](docs/growthbook_ramp-schedules_list.md) - Get all rampSchedules
* [`create`](docs/growthbook_ramp-schedules_create.md) - Create a single rampSchedule
* [`start`](docs/growthbook_ramp-schedules_start.md) - Start a ramp schedule
* [`pause`](docs/growthbook_ramp-schedules_pause.md) - Pause a ramp schedule
* [`resume`](docs/growthbook_ramp-schedules_resume.md) - Resume a paused ramp schedule
* [`rollback`](docs/growthbook_ramp-schedules_rollback.md) - Roll back a ramp schedule
* [`jump`](docs/growthbook_ramp-schedules_jump.md) - Jump to a specific step
* [`complete`](docs/growthbook_ramp-schedules_complete.md) - Complete a ramp schedule immediately
* [`approve-step`](docs/growthbook_ramp-schedules_approve-step.md) - Approve the current pending-approval step
* [`add-target`](docs/growthbook_ramp-schedules_add-target.md) - Add a target rule to a ramp schedule
* [`eject-target`](docs/growthbook_ramp-schedules_eject-target.md) - Remove a target rule from a ramp schedule
* [`get`](docs/growthbook_ramp-schedules_get.md) - Get a single rampSchedule
* [`delete`](docs/growthbook_ramp-schedules_delete.md) - Delete a single rampSchedule
* [`update`](docs/growthbook_ramp-schedules_update.md) - Update a single rampSchedule

### [namespaces](docs/growthbook_namespaces.md)

* [`list`](docs/growthbook_namespaces_list.md) - Get all namespaces
* [`post`](docs/growthbook_namespaces_post.md) - Create a namespace
* [`get`](docs/growthbook_namespaces_get.md) - Get a single namespace
* [`put`](docs/growthbook_namespaces_put.md) - Update a namespace
* [`delete`](docs/growthbook_namespaces_delete.md) - Delete a namespace
* [`get-namespace-memberships`](docs/growthbook_namespaces_get-namespace-memberships.md) - Get namespace membership
* [`post-namespace-rotate-seed`](docs/growthbook_namespaces_post-namespace-rotate-seed.md) - Rotate namespace seed

### [dashboards](docs/growthbook_dashboards.md)

* [`get`](docs/growthbook_dashboards_get.md) - Get a single dashboard
* [`delete`](docs/growthbook_dashboards_delete.md) - Delete a single dashboard
* [`update`](docs/growthbook_dashboards_update.md) - Update a single dashboard
* [`create`](docs/growthbook_dashboards_create.md) - Create a single dashboard
* [`list`](docs/growthbook_dashboards_list.md) - Get all dashboards
* [`get-dashboards-for-experiment`](docs/growthbook_dashboards_get-dashboards-for-experiment.md) - Get all dashboards for an experiment

### [custom-fields](docs/growthbook_custom-fields.md)

* [`create`](docs/growthbook_custom-fields_create.md) - Create a single customField
* [`list`](docs/growthbook_custom-fields_list.md) - Get all custom fields
* [`delete`](docs/growthbook_custom-fields_delete.md) - Delete a single customField
* [`get`](docs/growthbook_custom-fields_get.md) - Get a single customField
* [`update`](docs/growthbook_custom-fields_update.md) - Update a single customField

### [metric-groups](docs/growthbook_metric-groups.md)

* [`get`](docs/growthbook_metric-groups_get.md) - Get a single metricGroup
* [`delete`](docs/growthbook_metric-groups_delete.md) - Delete a single metricGroup
* [`update`](docs/growthbook_metric-groups_update.md) - Update a single metricGroup
* [`create`](docs/growthbook_metric-groups_create.md) - Create a single metricGroup
* [`list`](docs/growthbook_metric-groups_list.md) - Get all metricGroups

### [teams](docs/growthbook_teams.md)

* [`get`](docs/growthbook_teams_get.md) - Get a single team
* [`update`](docs/growthbook_teams_update.md) - Update a single team
* [`delete`](docs/growthbook_teams_delete.md) - Delete a single team
* [`create`](docs/growthbook_teams_create.md) - Create a single team
* [`list`](docs/growthbook_teams_list.md) - Get all teams
* [`add-team-members`](docs/growthbook_teams_add-team-members.md) - Add members to team
* [`remove-team-member`](docs/growthbook_teams_remove-team-member.md) - Remove members from team

### [experiment-templates](docs/growthbook_experiment-templates.md)

* [`get`](docs/growthbook_experiment-templates_get.md) - Get a single experimentTemplate
* [`delete`](docs/growthbook_experiment-templates_delete.md) - Delete a single experimentTemplate
* [`update`](docs/growthbook_experiment-templates_update.md) - Update a single experimentTemplate
* [`create`](docs/growthbook_experiment-templates_create.md) - Create a single experimentTemplate
* [`list`](docs/growthbook_experiment-templates_list.md) - Get all experimentTemplates
* [`bulk-import`](docs/growthbook_experiment-templates_bulk-import.md) - Bulk create or update experiment templates

### [analytics-explorations](docs/growthbook_analytics-explorations.md)

* [`post-metric-exploration`](docs/growthbook_analytics-explorations_post-metric-exploration.md) - Create a Metric based visualization
* [`post-fact-table-exploration`](docs/growthbook_analytics-explorations_post-fact-table-exploration.md) - Run a Fact Table based visualization
* [`post-data-source-exploration`](docs/growthbook_analytics-explorations_post-data-source-exploration.md) - Create a Data Source based visualization

### [ramp-schedule-templates](docs/growthbook_ramp-schedule-templates.md)

* [`get`](docs/growthbook_ramp-schedule-templates_get.md) - Get a single rampScheduleTemplate
* [`delete`](docs/growthbook_ramp-schedule-templates_delete.md) - Delete a single rampScheduleTemplate
* [`update`](docs/growthbook_ramp-schedule-templates_update.md) - Update a single rampScheduleTemplate
* [`create`](docs/growthbook_ramp-schedule-templates_create.md) - Create a single rampScheduleTemplate
* [`list`](docs/growthbook_ramp-schedule-templates_list.md) - Get all rampScheduleTemplates

</details>
<!-- End Available Commands [operations] -->

<!-- Start Request Body Input [stdinpiping] -->
## Request Body Input

Operations that accept a request body support three input methods, with a clear priority chain:

### Individual flags (highest priority)

```bash
growthbook <command> --name "Jane" --age 30
```

### `--body` flag

Provide the entire request body as a JSON string:

```bash
growthbook <command> --body '{"name": "John", "age": 30}'
```

Individual flags override `--body` values:

```bash
# Result: {name: "Jane", age: 30}
growthbook <command> --body '{"name": "John", "age": 30}' --name "Jane"
```

### Stdin piping (lowest priority)

Pipe JSON into any command that accepts a request body:

```bash
echo '{"name": "John", "age": 30}' | growthbook <command>
```

Individual flags override stdin values:

```bash
# Result: {name: "Jane", age: 30}
echo '{"name": "John", "age": 30}' | growthbook <command> --name "Jane"
```

This is useful for chaining commands, reading from files, or scripting:

```bash
# Read body from a file
growthbook <command> < request.json

# Pipe from another command
curl -s https://example.com/data.json | growthbook <command>
```

### Priority

When multiple input methods are used, the priority is:

| Priority | Source | Description |
|----------|--------|-------------|
| 1 (highest) | Individual flags | `--name "Jane"` always wins |
| 2 | `--body` flag | Whole-body JSON via flag |
| 3 (lowest) | Stdin | Piped JSON input |
<!-- End Request Body Input [stdinpiping] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

Use `--server <index>` to select a server by its zero-based index (default: `0`):

| #   | Server                          | Variables | Description            |
| --- | ------------------------------- | --------- | ---------------------- |
| 0   | `https://api.growthbook.io/api` |           | GrowthBook Cloud       |
| 1   | `https://{domain}/api`          | `domain`  | Self-hosted GrowthBook |

```bash
growthbook --server <index> <command> [arguments]
```

### Server Variables

Some server URLs contain template variables (e.g., `https://{hostname}:{port}/v1`). Set these via dedicated flags:

| Variable | Flag               | Default            | Description                                                                                                                |
| -------- | ------------------ | ------------------ | -------------------------------------------------------------------------------------------------------------------------- |
| `domain` | `--domain <value>` | `"localhost:3100"` | The host (and optional port) of your self-hosted GrowthBook API server, e.g. `growthbook.example.com` or `localhost:3100`. |

```bash
growthbook --domain value <command> [arguments]
```

Server variable flags are combined with the selected server URL to produce the final endpoint.
### Override Server URL

Use `--server-url` to override the server URL entirely, bypassing any named or indexed server selection:

```bash
growthbook --server-url https://custom-api.example.com <command> [arguments]
```

**Precedence**: `--server-url` > `--server` > default
<!-- End Server Selection [server] -->

<!-- Start Output Formats [output-formats] -->
## Output Formats

Every command supports a `--output-format` flag that controls how the response is rendered to stdout.

### Available formats

| Format | Flag | Description |
|--------|------|-------------|
| Pretty | `--output-format pretty` (default) | Aligned key-value pairs with color, nested indentation. Human-readable at a glance. |
| JSON | `--output-format json` | JSON output. Passthrough when the response is already JSON (preserves original field order and numeric precision). Falls back to typed marshaling otherwise. |
| YAML | `--output-format yaml` | YAML output via standard marshaling. |
| Table | `--output-format table` | Tabular output for array responses. |
| TOON | `--output-format toon` | [Token-Oriented Object Notation](https://github.com/toon-format/spec) — a compact, line-oriented format that typically uses 30–60% fewer tokens than JSON. Well-suited for piping responses into LLM prompts. |

```bash
# Default pretty output
growthbook <command>

# Machine-readable JSON
growthbook <command> --output-format json

# TOON for LLM-friendly compact output
growthbook <command> --output-format toon

# Pipe JSON to jq without using --output-format
growthbook <command> --output-format json | jq '.fieldName'
```

### jq filtering

Use `--jq` to filter or transform the response inline using a [jq](https://jqlang.org) expression. This always outputs JSON and overrides `--output-format`:

```bash
# Extract a single field
growthbook <command> --jq '.name'

# Filter an array
growthbook <command> --jq '.items[] | select(.active == true)'
```

### Color control

Use `--color` to control terminal colors:

| Value | Behavior |
|-------|----------|
| `auto` (default) | Color when stdout is a TTY, plain text otherwise |
| `always` | Always colorize |
| `never` | Never colorize |

The `NO_COLOR` and `FORCE_COLOR` environment variables are also respected.

### Streaming and pagination

When using `--all` (pagination) or streaming operations, output is written incrementally as items arrive:

| Format | Streaming behavior |
|--------|-------------------|
| `json` | One compact JSON object per line ([NDJSON](https://github.com/ndjson/ndjson-spec)) |
| `yaml` | YAML documents separated by `---` |
| `toon` | One TOON-encoded object per block, separated by blank lines |
| `pretty` (default) | Pretty-printed items separated by blank lines |
<!-- End Output Formats [output-formats] -->

<!-- Start Error Handling [errors] -->
## Error Handling

The CLI uses standard exit codes to indicate success or failure:

| Exit Code | Meaning |
|-----------|---------|
| `0` | Success |
| `1` | Error (API error, invalid input, etc.) |

On success, the response data is printed to **stdout** as JSON. On failure, error details are printed to **stderr**.

```bash
# Capture output and handle errors
growthbook ... > output.json 2> error.log
if [ $? -ne 0 ]; then
  echo "Error occurred, see error.log"
fi
```
<!-- End Error Handling [errors] -->

<!-- Start Diagnostics [diagnostics] -->
## Diagnostics

The CLI includes two diagnostic flags available on all commands:

### Dry Run

Preview what would be sent without making any network calls:

```bash
growthbook <command> --dry-run
```

Output goes to stderr and includes:
- HTTP method and URL
- Request headers (sensitive values redacted)
- Request body preview (sensitive fields redacted)

The command exits successfully without contacting the API. This is useful for verifying request construction before executing.

### Debug

Log request and response diagnostics while running normally:

```bash
growthbook <command> --debug
```

Debug output goes to stderr and includes:
- Request method, URL, headers, and body preview
- Response status, headers, and body preview
- Transport errors (if any)

The command still executes normally and produces its regular output on stdout.

### Flag Precedence

If both `--dry-run` and `--debug` are set, `--dry-run` takes precedence and no network calls are made.

### Security

Sensitive information is automatically redacted in diagnostic output:
- **Headers**: `Authorization`, `Cookie`, `Set-Cookie`, `X-API-Key`, and other security headers show `[REDACTED]`
- **Body**: JSON fields named `password`, `secret`, `token`, `api_key`, `client_secret`, etc. show `[REDACTED]`

Diagnostic output should still be treated as potentially sensitive operational data.
<!-- End Diagnostics [diagnostics] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This CLI is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this CLI, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### CLI Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/growthbook/cli&utm_campaign=cli)
