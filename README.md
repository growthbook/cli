# growthbook

Command-line interface for the *GrowthBook REST* API.

[![Built by Speakeasy](https://img.shields.io/badge/Built_by-SPEAKEASY-374151?style=for-the-badge&labelColor=f3f4f6)](https://www.speakeasy.com/?utm_source=github-com/growthbook/cli&utm_campaign=cli)
[![License: MIT](https://img.shields.io/badge/LICENSE_//_MIT-3b5bdb?style=for-the-badge&labelColor=eff6ff)](https://opensource.org/licenses/MIT)

## Installation

The recommended way to install is via **npm**, which installs a small launcher that downloads the prebuilt binary for your platform:

```bash
npm i -g growthbook
```

Prefer a standalone binary with no Node.js dependency? An install script, `go install`, and prebuilt downloads are covered under [CLI Installation](#cli-installation) below.

## Versioning and stability

This CLI follows [semantic versioning](https://semver.org): breaking changes ship only in a **major** release and are called out in the changelog. Additive changes (new commands, flags, or response fields) are **minor** releases; fixes are **patch** releases. We recommend pinning to a version you have tested.

The CLI checks for a newer release at most once a day and prints a one-line notice to stderr (never stdout) when one is available and compatible with your server. Disable it with `--no-update-check` or `GBCLI_NO_UPDATE_CHECK=1`; it is also off automatically in CI and non-interactive shells.

### Command versioning

Each command group targets the **newest** version of its API endpoint — e.g. `features` is the latest (v2) feature API. When a command is superseded, the previous version stays available under a version suffix (e.g. `features-v1`), deprecated and printing a one-line notice to stderr. When a newer API version lands, the base group advances to it and the prior version becomes the next `-vN`.

Because the bare command then returns the newer version's response shape, **re-pointing a base command is a breaking change: it ships as a major release with the change called out in the changelog** — so you can pin the prior major or move to the explicit `-vN` command on your own schedule.

<!-- Start Summary [summary] -->
## Summary

GrowthBook REST API: A command-line interface for GrowthBook — manage feature flags, experiments, metrics, and more from your terminal.

Endpoints are versioned by path prefix: `/v1` (stable) and `/v2` (improved shapes). Each command group targets the newest version of its endpoint; superseded versions remain available under a `-vN` suffix.

Authenticate with a Secret Key or Personal Access Token via `--bearer-auth` (or the `GBCLI_BEARER_AUTH` environment variable). Run `growthbook configure` to store credentials, or `growthbook whoami` to check the active configuration.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [growthbook](#growthbook)
  * [Installation](#installation)
  * [Versioning and stability](#versioning-and-stability)
  * [CLI Installation](#cli-installation)
  * [Shell Completion](#shell-completion)
  * [CLI Example Usage](#cli-example-usage)
  * [Authentication](#authentication)
  * [Profiles](#profiles)
  * [Generating TypeScript types](#generating-typescript-types)
  * [Available Commands](#available-commands)
  * [Request Body Input](#request-body-input)
  * [Server Selection](#server-selection)
  * [Output Formats](#output-formats)
  * [Pagination](#pagination)
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
| `GBCLI_BEARER_AUTH` | Bearer auth token: your Secret Key or Personal Access Token, sent as an Authorization Bearer header. |
| `GBCLI_USERNAME` | HTTP Basic auth: use your GrowthBook Secret Key as the username and leave the password empty. username |
| `GBCLI_PASSWORD` | HTTP Basic auth: use your GrowthBook Secret Key as the username and leave the password empty. password |

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

## Profiles

Profiles let you switch between GrowthBook environments (e.g. cloud, staging, self-hosted), each with its own server URL and API key. Keys are stored in your OS keychain, not in plaintext config.

```bash
# Create or update a profile (the key comes from --bearer-auth, stored in the keychain)
growthbook profiles set staging --server-url https://gb.staging.example.com/api --bearer-auth secret_xxx

growthbook profiles use staging      # make it the active profile
growthbook profiles list             # show configured profiles
growthbook profiles remove staging   # delete one
```

Select a profile for a single command with `--profile <name>`, or set `GBCLI_PROFILE` in your environment. (Upgrading from the legacy CLI? An existing `~/.growthbook/config.toml` is imported into profiles automatically on first run — see [MIGRATION.md](MIGRATION.md).)

## Generating TypeScript types

`generate-types` fetches all of your features and writes a TypeScript `AppFeatures` definition, giving you type-safe feature keys and values in the GrowthBook SDK:

```bash
growthbook generate-types                                   # → ./growthbook-types/app-features.ts
growthbook generate-types --output ./src/types --filename growthbook.ts
growthbook generate-types --project prj_123                 # limit to one project
```

<!-- Start Available Commands [operations] -->
## Available Commands

<details open>
<summary>Available commands</summary>

* [`get-SDK-payload`](docs/growthbook_get-SDK-payload.md) - Get a SDK payload
* [`post-copy-transform`](docs/growthbook_post-copy-transform.md)

### [features-v1](docs/growthbook_features-v1.md)

* [`~~list~~`](docs/growthbook_features-v1_list.md) - Get all features :warning: **Deprecated**
* [`~~create~~`](docs/growthbook_features-v1_create.md) - Create a single feature :warning: **Deprecated**
* [`~~get~~`](docs/growthbook_features-v1_get.md) - Get a single feature :warning: **Deprecated**
* [`~~update~~`](docs/growthbook_features-v1_update.md) - Partially update a feature :warning: **Deprecated**
* [`~~delete~~`](docs/growthbook_features-v1_delete.md) - Deletes a single feature :warning: **Deprecated**
* [`~~toggle~~`](docs/growthbook_features-v1_toggle.md) - Toggle a feature in one or more environments :warning: **Deprecated**
* [`~~revert~~`](docs/growthbook_features-v1_revert.md) - Revert a feature to a specific revision :warning: **Deprecated**
* [`~~get-feature-keys~~`](docs/growthbook_features-v1_get-feature-keys.md) - Get list of feature keys :warning: **Deprecated**
* [`~~get-feature-stale~~`](docs/growthbook_features-v1_get-feature-stale.md) - Get stale status for one or more features :warning: **Deprecated**

### [feature-revisions-v1](docs/growthbook_feature-revisions-v1.md)

* [`~~list-revisions~~`](docs/growthbook_feature-revisions-v1_list-revisions.md) - List feature revisions :warning: **Deprecated**
* [`~~list-for-feature~~`](docs/growthbook_feature-revisions-v1_list-for-feature.md) - List revisions for a feature :warning: **Deprecated**
* [`~~create~~`](docs/growthbook_feature-revisions-v1_create.md) - Create a draft revision :warning: **Deprecated**
* [`~~get-feature-revision-latest~~`](docs/growthbook_feature-revisions-v1_get-feature-revision-latest.md) - Get the most recent active draft revision :warning: **Deprecated**
* [`~~get-feature-revision~~`](docs/growthbook_feature-revisions-v1_get-feature-revision.md) - Get a single feature revision :warning: **Deprecated**
* [`~~put-feature-revision-metadata~~`](docs/growthbook_feature-revisions-v1_put-feature-revision-metadata.md) - Update revision metadata (comment, title, feature metadata) :warning: **Deprecated**
* [`~~put-feature-revision-default-value~~`](docs/growthbook_feature-revisions-v1_put-feature-revision-default-value.md) - Set the default value in a draft revision :warning: **Deprecated**
* [`~~put-feature-revision-prerequisites~~`](docs/growthbook_feature-revisions-v1_put-feature-revision-prerequisites.md) - Set feature-level prerequisites in a draft revision :warning: **Deprecated**
* [`~~put-feature-revision-holdout~~`](docs/growthbook_feature-revisions-v1_put-feature-revision-holdout.md) - Set holdout in a draft revision :warning: **Deprecated**
* [`~~put-feature-revision-archive~~`](docs/growthbook_feature-revisions-v1_put-feature-revision-archive.md) - Set archived state in a draft revision :warning: **Deprecated**
* [`~~post-feature-revision-toggle~~`](docs/growthbook_feature-revisions-v1_post-feature-revision-toggle.md) - Toggle an environment on/off in a draft revision :warning: **Deprecated**
* [`~~post-feature-revision-rule-add~~`](docs/growthbook_feature-revisions-v1_post-feature-revision-rule-add.md) - Add a rule to a draft revision :warning: **Deprecated**
* [`~~put-feature-revision-rule~~`](docs/growthbook_feature-revisions-v1_put-feature-revision-rule.md) - Update a rule in a draft revision :warning: **Deprecated**
* [`~~delete-feature-revision-rule~~`](docs/growthbook_feature-revisions-v1_delete-feature-revision-rule.md) - Delete a rule from a draft revision :warning: **Deprecated**
* [`~~post-feature-revision-rules-reorder~~`](docs/growthbook_feature-revisions-v1_post-feature-revision-rules-reorder.md) - Reorder rules in an environment :warning: **Deprecated**
* [`~~put-feature-revision-rule-ramp-schedule~~`](docs/growthbook_feature-revisions-v1_put-feature-revision-rule-ramp-schedule.md) - Set ramp schedule for a rule :warning: **Deprecated**
* [`~~delete-feature-revision-rule-ramp-schedule~~`](docs/growthbook_feature-revisions-v1_delete-feature-revision-rule-ramp-schedule.md) - Remove ramp schedule from a rule :warning: **Deprecated**
* [`~~post-feature-revision-request-review~~`](docs/growthbook_feature-revisions-v1_post-feature-revision-request-review.md) - Request review for a draft revision :warning: **Deprecated**
* [`~~post-feature-revision-submit-review~~`](docs/growthbook_feature-revisions-v1_post-feature-revision-submit-review.md) - Submit a review on a draft revision :warning: **Deprecated**
* [`~~get-feature-revision-merge-status~~`](docs/growthbook_feature-revisions-v1_get-feature-revision-merge-status.md) - Get merge status for a draft revision :warning: **Deprecated**
* [`~~post-feature-revision-rebase~~`](docs/growthbook_feature-revisions-v1_post-feature-revision-rebase.md) - Rebase a draft revision onto the current live version :warning: **Deprecated**
* [`~~post-feature-revision-publish~~`](docs/growthbook_feature-revisions-v1_post-feature-revision-publish.md) - Publish a draft revision :warning: **Deprecated**
* [`~~post-feature-revision-discard~~`](docs/growthbook_feature-revisions-v1_post-feature-revision-discard.md) - Discard a draft revision :warning: **Deprecated**
* [`~~post-feature-revision-revert~~`](docs/growthbook_feature-revisions-v1_post-feature-revision-revert.md) - Revert the feature to a prior revision :warning: **Deprecated**

### [features](docs/growthbook_features.md)

* [`list`](docs/growthbook_features_list.md) - Get all features
* [`create`](docs/growthbook_features_create.md) - Create a single feature
* [`get`](docs/growthbook_features_get.md) - Get a single feature
* [`update`](docs/growthbook_features_update.md) - Partially update a feature
* [`delete`](docs/growthbook_features_delete.md) - Deletes a single feature
* [`toggle`](docs/growthbook_features_toggle.md) - Toggle a feature in one or more environments
* [`revert`](docs/growthbook_features_revert.md) - Revert a feature to a specific revision
* [`get-feature-keys`](docs/growthbook_features_get-feature-keys.md) - Get list of feature keys
* [`get-feature-stale`](docs/growthbook_features_get-feature-stale.md) - Get stale status for one or more features

### [feature-revisions](docs/growthbook_feature-revisions.md)

* [`list-revisions`](docs/growthbook_feature-revisions_list-revisions.md) - List revisions across all features
* [`list-for-feature`](docs/growthbook_feature-revisions_list-for-feature.md) - List revisions for a feature
* [`create`](docs/growthbook_feature-revisions_create.md) - Create a draft revision
* [`latest`](docs/growthbook_feature-revisions_latest.md) - Get the most recent active draft revision
* [`get`](docs/growthbook_feature-revisions_get.md) - Get a single feature revision
* [`diff`](docs/growthbook_feature-revisions_diff.md) - Diff a revision against another revision
* [`set-metadata`](docs/growthbook_feature-revisions_set-metadata.md) - Update revision metadata
* [`set-default-value`](docs/growthbook_feature-revisions_set-default-value.md) - Set the default value in a draft revision
* [`set-prerequisites`](docs/growthbook_feature-revisions_set-prerequisites.md) - Set feature-level prerequisites in a draft revision
* [`set-holdout`](docs/growthbook_feature-revisions_set-holdout.md) - Set holdout in a draft revision
* [`archive`](docs/growthbook_feature-revisions_archive.md) - Set archived state in a draft revision
* [`toggle`](docs/growthbook_feature-revisions_toggle.md) - Toggle an environment on/off in a draft revision
* [`add-rule`](docs/growthbook_feature-revisions_add-rule.md) - Add a rule to a draft revision
* [`update-rule`](docs/growthbook_feature-revisions_update-rule.md) - Update a rule in a draft revision
* [`delete-rule`](docs/growthbook_feature-revisions_delete-rule.md) - Delete a rule from a draft revision
* [`reorder-rules`](docs/growthbook_feature-revisions_reorder-rules.md) - Reorder rules in the revision
* [`set-rule-ramp-schedule`](docs/growthbook_feature-revisions_set-rule-ramp-schedule.md) - Set ramp schedule for a rule
* [`delete-rule-ramp-schedule`](docs/growthbook_feature-revisions_delete-rule-ramp-schedule.md) - Remove ramp schedule from a rule
* [`request-review`](docs/growthbook_feature-revisions_request-review.md) - Request review for a draft revision
* [`schedule-publish`](docs/growthbook_feature-revisions_schedule-publish.md) - Schedule (or cancel) a deferred publish for a draft revision
* [`submit-review`](docs/growthbook_feature-revisions_submit-review.md) - Submit a review on a draft revision
* [`recall-review`](docs/growthbook_feature-revisions_recall-review.md) - Recall a review request (revert to draft)
* [`undo-review`](docs/growthbook_feature-revisions_undo-review.md) - Undo a reviewer's own review verdict
* [`log`](docs/growthbook_feature-revisions_log.md) - List the activity log for a revision
* [`add-log-comment`](docs/growthbook_feature-revisions_add-log-comment.md) - Edit the comment text of an owned log entry
* [`delete-log-entry`](docs/growthbook_feature-revisions_delete-log-entry.md) - Delete an owned revision Comment entry
* [`merge-status`](docs/growthbook_feature-revisions_merge-status.md) - Get merge status for a draft revision
* [`rebase-preview`](docs/growthbook_feature-revisions_rebase-preview.md) - Preview a rebase without applying it
* [`rebase`](docs/growthbook_feature-revisions_rebase.md) - Rebase a draft revision onto the current live version
* [`publish`](docs/growthbook_feature-revisions_publish.md) - Publish a draft revision
* [`discard`](docs/growthbook_feature-revisions_discard.md) - Discard a draft revision
* [`reopen`](docs/growthbook_feature-revisions_reopen.md) - Reopen a discarded revision as a draft
* [`revert`](docs/growthbook_feature-revisions_revert.md) - Revert the feature to a prior revision

### [archetypes](docs/growthbook_archetypes.md)

* [`list`](docs/growthbook_archetypes_list.md) - Get the organization's archetypes
* [`create`](docs/growthbook_archetypes_create.md) - Create a single archetype
* [`get`](docs/growthbook_archetypes_get.md) - Get a single archetype
* [`update`](docs/growthbook_archetypes_update.md) - Update a single archetype
* [`delete`](docs/growthbook_archetypes_delete.md) - Deletes a single archetype

### [experiments](docs/growthbook_experiments.md)

* [`list`](docs/growthbook_experiments_list.md) - Get all experiments
* [`create`](docs/growthbook_experiments_create.md) - Create a single experiment
* [`list-experiment-results`](docs/growthbook_experiments_list-experiment-results.md) - Get latest results for many experiments
* [`get`](docs/growthbook_experiments_get.md) - Get a single experiment
* [`update`](docs/growthbook_experiments_update.md) - Update a single experiment
* [`get-experiment-start-checklist`](docs/growthbook_experiments_get-experiment-start-checklist.md) - Get an experiment pre-launch checklist status
* [`get-experiment-results`](docs/growthbook_experiments_get-experiment-results.md) - Get results for an experiment
* [`post-experiment-start`](docs/growthbook_experiments_post-experiment-start.md) - Start/Stage an experiment
* [`post-experiment-start-checklist-manual-complete`](docs/growthbook_experiments_post-experiment-start-checklist-manual-complete.md) - Mark manual pre-launch checklist items complete
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
* [`create`](docs/growthbook_visual-changesets_create.md) - Create a visual changeset for an experiment
* [`get`](docs/growthbook_visual-changesets_get.md) - Get a single visual changeset
* [`update`](docs/growthbook_visual-changesets_update.md) - Update a visual changeset
* [`post-visual-change`](docs/growthbook_visual-changesets_post-visual-change.md) - Create a visual change for a visual changeset
* [`put-visual-change`](docs/growthbook_visual-changesets_put-visual-change.md) - Update a visual change for a visual changeset

### [contextual-bandits](docs/growthbook_contextual-bandits.md)

* [`get-contextual-bandit-current-weights`](docs/growthbook_contextual-bandits_get-contextual-bandit-current-weights.md) - Get current Contextual Bandit leaf weights and latest event
* [`list-contextual-bandit-snapshots`](docs/growthbook_contextual-bandits_list-contextual-bandit-snapshots.md) - List Contextual Bandit snapshots
* [`get-contextual-bandit-snapshot`](docs/growthbook_contextual-bandits_get-contextual-bandit-snapshot.md) - Get a single Contextual Bandit snapshot
* [`list-contextual-bandit-events`](docs/growthbook_contextual-bandits_list-contextual-bandit-events.md) - List Contextual Bandit weight-update events
* [`get-contextual-bandit-event`](docs/growthbook_contextual-bandits_get-contextual-bandit-event.md) - Get a single Contextual Bandit weight-update event
* [`get-contextual-bandit-results`](docs/growthbook_contextual-bandits_get-contextual-bandit-results.md) - Get latest Contextual Bandit results
* [`get-contextual-bandit-linked-features`](docs/growthbook_contextual-bandits_get-contextual-bandit-linked-features.md) - Get features linked to a Contextual Bandit
* [`delete-contextual-bandit-linked-feature`](docs/growthbook_contextual-bandits_delete-contextual-bandit-linked-feature.md) - Unlink a feature from a Contextual Bandit
* [`get`](docs/growthbook_contextual-bandits_get.md) - Get a single contextualBandit
* [`update`](docs/growthbook_contextual-bandits_update.md) - Update a single contextualBandit
* [`create`](docs/growthbook_contextual-bandits_create.md) - Create a single contextualBandit
* [`list`](docs/growthbook_contextual-bandits_list.md) - Get all contextualBandits
* [`start`](docs/growthbook_contextual-bandits_start.md) - Start a Contextual Bandit
* [`stop`](docs/growthbook_contextual-bandits_stop.md) - Stop a Contextual Bandit
* [`refresh`](docs/growthbook_contextual-bandits_refresh.md) - Trigger a Contextual Bandit snapshot refresh

### [metrics](docs/growthbook_metrics.md)

* [`list`](docs/growthbook_metrics_list.md) - Get all metrics
* [`create`](docs/growthbook_metrics_create.md) - Create a single metric
* [`get`](docs/growthbook_metrics_get.md) - Get a single metric
* [`update`](docs/growthbook_metrics_update.md) - Update a metric
* [`delete`](docs/growthbook_metrics_delete.md) - Deletes a metric

### [usage-metrics](docs/growthbook_usage-metrics.md)

* [`get`](docs/growthbook_usage-metrics_get.md) - Get metric usage across experiments

### [segments](docs/growthbook_segments.md)

* [`list`](docs/growthbook_segments_list.md) - Get all segments
* [`create`](docs/growthbook_segments_create.md) - Create a single segment
* [`get`](docs/growthbook_segments_get.md) - Get a single segment
* [`update`](docs/growthbook_segments_update.md) - Update a single segment
* [`delete`](docs/growthbook_segments_delete.md) - Deletes a single segment

### [dimensions](docs/growthbook_dimensions.md)

* [`list`](docs/growthbook_dimensions_list.md) - Get all dimensions
* [`create`](docs/growthbook_dimensions_create.md) - Create a single dimension
* [`get`](docs/growthbook_dimensions_get.md) - Get a single dimension
* [`update`](docs/growthbook_dimensions_update.md) - Update a single dimension
* [`delete`](docs/growthbook_dimensions_delete.md) - Deletes a single dimension

### [projects](docs/growthbook_projects.md)

* [`list`](docs/growthbook_projects_list.md) - Get all projects
* [`create`](docs/growthbook_projects_create.md) - Create a single project
* [`get`](docs/growthbook_projects_get.md) - Get a single project
* [`update`](docs/growthbook_projects_update.md) - Edit a single project
* [`delete`](docs/growthbook_projects_delete.md) - Deletes a single project

### [environments](docs/growthbook_environments.md)

* [`list`](docs/growthbook_environments_list.md) - Get the organization's environments
* [`create`](docs/growthbook_environments_create.md) - Create a new environment
* [`update`](docs/growthbook_environments_update.md) - Update an environment
* [`delete`](docs/growthbook_environments_delete.md) - Deletes a single environment

### [attributes](docs/growthbook_attributes.md)

* [`list`](docs/growthbook_attributes_list.md) - Get the organization's attributes
* [`create`](docs/growthbook_attributes_create.md) - Create a new attribute
* [`update`](docs/growthbook_attributes_update.md) - Update an attribute
* [`delete`](docs/growthbook_attributes_delete.md) - Deletes a single attribute

### [SDK-connections](docs/growthbook_SDK-connections.md)

* [`list`](docs/growthbook_SDK-connections_list.md) - Get all sdk connections
* [`create`](docs/growthbook_SDK-connections_create.md) - Create a single sdk connection
* [`get`](docs/growthbook_SDK-connections_get.md) - Get a single sdk connection
* [`update`](docs/growthbook_SDK-connections_update.md) - Update a single sdk connection
* [`delete`](docs/growthbook_SDK-connections_delete.md) - Deletes a single SDK connection
* [`lookup-SDK-connection-by-key`](docs/growthbook_SDK-connections_lookup-SDK-connection-by-key.md) - Find a single sdk connection by its key

### [data-sources](docs/growthbook_data-sources.md)

* [`list`](docs/growthbook_data-sources_list.md) - Get all data sources
* [`get`](docs/growthbook_data-sources_get.md) - Get a single data source
* [`get-information-schema`](docs/growthbook_data-sources_get-information-schema.md) - Get a Data Source's Information Schema
* [`get-information-schema-table`](docs/growthbook_data-sources_get-information-schema-table.md) - Get a single Information Schema Table by id

### [saved-groups](docs/growthbook_saved-groups.md)

* [`list`](docs/growthbook_saved-groups_list.md) - Get all saved group
* [`create`](docs/growthbook_saved-groups_create.md) - Create a single saved group
* [`get`](docs/growthbook_saved-groups_get.md) - Get a single saved group
* [`update`](docs/growthbook_saved-groups_update.md) - Partially update a single saved group
* [`delete`](docs/growthbook_saved-groups_delete.md) - Deletes a single saved group
* [`archive`](docs/growthbook_saved-groups_archive.md) - Archive a single saved group
* [`unarchive`](docs/growthbook_saved-groups_unarchive.md) - Unarchive a single saved group
* [`get-saved-group-references`](docs/growthbook_saved-groups_get-saved-group-references.md) - Get features, experiments, and saved groups that reference this saved group

### [saved-group-revisions](docs/growthbook_saved-group-revisions.md)

* [`list`](docs/growthbook_saved-group-revisions_list.md) - List saved-group revisions across the organization
* [`list-for-saved-group`](docs/growthbook_saved-group-revisions_list-for-saved-group.md) - List revisions for a saved group
* [`create`](docs/growthbook_saved-group-revisions_create.md) - Create a draft revision
* [`get-saved-group-revision-latest`](docs/growthbook_saved-group-revisions_get-saved-group-revision-latest.md) - Get the most recent active draft revision
* [`get`](docs/growthbook_saved-group-revisions_get.md) - Get a single saved group revision
* [`put-saved-group-revision-metadata`](docs/growthbook_saved-group-revisions_put-saved-group-revision-metadata.md) - Update saved group metadata in a draft revision
* [`put-saved-group-revision-condition`](docs/growthbook_saved-group-revisions_put-saved-group-revision-condition.md) - Update the condition of a condition saved group draft revision
* [`put-saved-group-revision-values`](docs/growthbook_saved-group-revisions_put-saved-group-revision-values.md) - Replace the values list in a list saved group draft revision
* [`put-saved-group-revision-archive`](docs/growthbook_saved-group-revisions_put-saved-group-revision-archive.md) - Stage an archive/unarchive in a draft revision
* [`post-saved-group-revision-items-add`](docs/growthbook_saved-group-revisions_post-saved-group-revision-items-add.md) - Append items to a list saved group draft revision
* [`post-saved-group-revision-items-remove`](docs/growthbook_saved-group-revisions_post-saved-group-revision-items-remove.md) - Remove items from a list saved group draft revision
* [`post-saved-group-revision-request-review`](docs/growthbook_saved-group-revisions_post-saved-group-revision-request-review.md) - Request review for a draft revision
* [`post-saved-group-revision-submit-review`](docs/growthbook_saved-group-revisions_post-saved-group-revision-submit-review.md) - Submit a review on a draft revision
* [`get-saved-group-revision-merge-status`](docs/growthbook_saved-group-revisions_get-saved-group-revision-merge-status.md) - Get merge status for a draft revision
* [`post-saved-group-revision-rebase`](docs/growthbook_saved-group-revisions_post-saved-group-revision-rebase.md) - Rebase a draft revision onto the current live saved group
* [`post-saved-group-revision-publish`](docs/growthbook_saved-group-revisions_post-saved-group-revision-publish.md) - Publish a draft revision
* [`post-saved-group-revision-discard`](docs/growthbook_saved-group-revisions_post-saved-group-revision-discard.md) - Discard a draft revision
* [`post-saved-group-revision-revert`](docs/growthbook_saved-group-revisions_post-saved-group-revision-revert.md) - Revert the saved group to a prior revision

### [constants](docs/growthbook_constants.md)

* [`list`](docs/growthbook_constants_list.md) - Get all constants
* [`create`](docs/growthbook_constants_create.md) - Create a single constant
* [`get-constant-references`](docs/growthbook_constants_get-constant-references.md) - Get features and constants that reference this constant
* [`get`](docs/growthbook_constants_get.md) - Get a single constant
* [`update`](docs/growthbook_constants_update.md) - Partially update a single constant
* [`delete`](docs/growthbook_constants_delete.md) - Delete a single constant
* [`archive`](docs/growthbook_constants_archive.md) - Archive a single constant
* [`unarchive`](docs/growthbook_constants_unarchive.md) - Unarchive a single constant

### [constant-revisions](docs/growthbook_constant-revisions.md)

* [`list`](docs/growthbook_constant-revisions_list.md) - List constant revisions across the organization
* [`list-for-constant`](docs/growthbook_constant-revisions_list-for-constant.md) - List revisions for a constant
* [`create`](docs/growthbook_constant-revisions_create.md) - Create a draft revision
* [`get-constant-revision-latest`](docs/growthbook_constant-revisions_get-constant-revision-latest.md) - Get the most recent active draft revision
* [`get`](docs/growthbook_constant-revisions_get.md) - Get a single constant revision
* [`put-constant-revision-metadata`](docs/growthbook_constant-revisions_put-constant-revision-metadata.md) - Update constant metadata in a draft revision
* [`put-constant-revision-value`](docs/growthbook_constant-revisions_put-constant-revision-value.md) - Update the value of a constant draft revision
* [`put-constant-revision-archive`](docs/growthbook_constant-revisions_put-constant-revision-archive.md) - Stage an archive/unarchive in a draft revision
* [`post-constant-revision-request-review`](docs/growthbook_constant-revisions_post-constant-revision-request-review.md) - Request review for a draft revision
* [`post-constant-revision-submit-review`](docs/growthbook_constant-revisions_post-constant-revision-submit-review.md) - Submit a review on a draft revision
* [`get-constant-revision-merge-status`](docs/growthbook_constant-revisions_get-constant-revision-merge-status.md) - Get merge status for a draft revision
* [`post-constant-revision-rebase`](docs/growthbook_constant-revisions_post-constant-revision-rebase.md) - Rebase a draft revision onto the current live constant
* [`post-constant-revision-publish`](docs/growthbook_constant-revisions_post-constant-revision-publish.md) - Publish a draft revision
* [`post-constant-revision-discard`](docs/growthbook_constant-revisions_post-constant-revision-discard.md) - Discard a draft revision
* [`post-constant-revision-revert`](docs/growthbook_constant-revisions_post-constant-revision-revert.md) - Revert the constant to a prior revision

### [organizations](docs/growthbook_organizations.md)

* [`list`](docs/growthbook_organizations_list.md) - Get all organizations (only for super admins on multi-org Enterprise Plan only)
* [`create`](docs/growthbook_organizations_create.md) - Create a single organization (only for super admins on multi-org Enterprise Plan only)
* [`update`](docs/growthbook_organizations_update.md) - Edit a single organization (only for super admins on multi-org Enterprise Plan only)

### [fact-tables](docs/growthbook_fact-tables.md)

* [`list`](docs/growthbook_fact-tables_list.md) - Get all fact tables
* [`create`](docs/growthbook_fact-tables_create.md) - Create a single fact table
* [`get`](docs/growthbook_fact-tables_get.md) - Get a single fact table
* [`update`](docs/growthbook_fact-tables_update.md) - Update a single fact table
* [`delete`](docs/growthbook_fact-tables_delete.md) - Deletes a single fact table
* [`list-fact-table-filters`](docs/growthbook_fact-tables_list-fact-table-filters.md) - Get all filters for a fact table
* [`post-fact-table-filter`](docs/growthbook_fact-tables_post-fact-table-filter.md) - Create a single fact table filter
* [`get-fact-table-filter`](docs/growthbook_fact-tables_get-fact-table-filter.md) - Get a single fact filter
* [`update-fact-table-filter`](docs/growthbook_fact-tables_update-fact-table-filter.md) - Update a single fact table filter
* [`delete-fact-table-filter`](docs/growthbook_fact-tables_delete-fact-table-filter.md) - Deletes a single fact table filter
* [`get-aggregated`](docs/growthbook_fact-tables_get-aggregated.md) - Get the materialization status of a fact table's shared daily aggregated tables
* [`refresh-aggregated`](docs/growthbook_fact-tables_refresh-aggregated.md) - Force a refresh or full restate of a fact table's shared daily aggregated tables
* [`list-aggregated-table-runs`](docs/growthbook_fact-tables_list-aggregated-table-runs.md) - List aggregated table runs
* [`get-aggregated-table-run`](docs/growthbook_fact-tables_get-aggregated-table-run.md) - Get a single aggregated table run
* [`post-bulk-import-facts`](docs/growthbook_fact-tables_post-bulk-import-facts.md) - Bulk import fact tables, filters, and metrics

### [fact-metrics](docs/growthbook_fact-metrics.md)

* [`list`](docs/growthbook_fact-metrics_list.md) - Get all fact metrics
* [`create`](docs/growthbook_fact-metrics_create.md) - Create a single fact metric
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

### [meta](docs/growthbook_meta.md)

* [`get-version`](docs/growthbook_meta_get-version.md) - Get the GrowthBook server version and build info

### [ramp-schedules](docs/growthbook_ramp-schedules.md)

* [`list`](docs/growthbook_ramp-schedules_list.md) - Get all rampSchedules
* [`create`](docs/growthbook_ramp-schedules_create.md) - Create a ramp schedule
* [`start`](docs/growthbook_ramp-schedules_start.md) - Start a ramp schedule
* [`pause`](docs/growthbook_ramp-schedules_pause.md) - Pause a ramp schedule
* [`resume`](docs/growthbook_ramp-schedules_resume.md) - Resume a paused ramp schedule
* [`rollback`](docs/growthbook_ramp-schedules_rollback.md) - Roll back a ramp schedule
* [`restart`](docs/growthbook_ramp-schedules_restart.md) - Restart a terminal ramp schedule
* [`jump`](docs/growthbook_ramp-schedules_jump.md) - Jump to a specific step
* [`complete`](docs/growthbook_ramp-schedules_complete.md) - Complete a ramp schedule immediately
* [`approve-step`](docs/growthbook_ramp-schedules_approve-step.md) - Approve the current step
* [`add-target`](docs/growthbook_ramp-schedules_add-target.md) - Add a target rule to a ramp schedule
* [`eject-target`](docs/growthbook_ramp-schedules_eject-target.md) - Remove a target rule from a ramp schedule
* [`api-advance`](docs/growthbook_ramp-schedules_api-advance.md) - Advance to the next step, overriding any holds
* [`get-ramp-schedule-status`](docs/growthbook_ramp-schedules_get-ramp-schedule-status.md) - Get ramp schedule status summary
* [`set-monitoring-mode`](docs/growthbook_ramp-schedules_set-monitoring-mode.md) - Set ramp monitoring mode
* [`set-auto-update`](docs/growthbook_ramp-schedules_set-auto-update.md) - Toggle automatic monitoring updates
* [`update-ramp-schedule-monitoring`](docs/growthbook_ramp-schedules_update-ramp-schedule-monitoring.md) - Update ramp monitoring configuration
* [`update-ramp-schedule-lockdown`](docs/growthbook_ramp-schedules_update-ramp-schedule-lockdown.md) - Update ramp lockdown configuration
* [`update-ramp-schedule-steps`](docs/growthbook_ramp-schedules_update-ramp-schedule-steps.md) - Update ramp schedule steps
* [`refresh-monitoring`](docs/growthbook_ramp-schedules_refresh-monitoring.md) - Trigger a manual monitoring update
* [`get`](docs/growthbook_ramp-schedules_get.md) - Get a single rampSchedule
* [`delete`](docs/growthbook_ramp-schedules_delete.md) - Delete a single rampSchedule
* [`update`](docs/growthbook_ramp-schedules_update.md) - Update a single rampSchedule

### [reports](docs/growthbook_reports.md)

* [`list`](docs/growthbook_reports_list.md) - Get all reports
* [`create`](docs/growthbook_reports_create.md) - Create a new report
* [`get`](docs/growthbook_reports_get.md) - Get a single report
* [`post-report-refresh`](docs/growthbook_reports_post-report-refresh.md) - Refresh a report by re-running its analysis
* [`put-report-metadata`](docs/growthbook_reports_put-report-metadata.md) - Update report metadata (title, description, visibility)
* [`put-report-settings`](docs/growthbook_reports_put-report-settings.md) - Update report analysis settings

### [namespaces](docs/growthbook_namespaces.md)

* [`list`](docs/growthbook_namespaces_list.md) - Get all namespaces
* [`create`](docs/growthbook_namespaces_create.md) - Create a namespace
* [`get`](docs/growthbook_namespaces_get.md) - Get a single namespace
* [`update`](docs/growthbook_namespaces_update.md) - Update a namespace
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

### [contextual-bandit-queries](docs/growthbook_contextual-bandit-queries.md)

* [`get-contextual-bandit-query`](docs/growthbook_contextual-bandit-queries_get-contextual-bandit-query.md) - Get a single contextualBanditQuery
* [`delete-contextual-bandit-query`](docs/growthbook_contextual-bandit-queries_delete-contextual-bandit-query.md) - Delete a single contextualBanditQuery
* [`update-contextual-bandit-query`](docs/growthbook_contextual-bandit-queries_update-contextual-bandit-query.md) - Update a single contextualBanditQuery
* [`create-contextual-bandit-query`](docs/growthbook_contextual-bandit-queries_create-contextual-bandit-query.md) - Create a single contextualBanditQuery
* [`list`](docs/growthbook_contextual-bandit-queries_list.md) - Get all contextualBanditQueries

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

| Variable | Flag               | Default            | Description                                 |
| -------- | ------------------ | ------------------ | ------------------------------------------- |
| `domain` | `--domain <value>` | `"localhost:3100"` | Your self-hosted GrowthBook host (and port) |

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

<!-- Start Pagination [pagination] -->
## Pagination

Some operations in this CLI support automatic pagination. These operations accept `--all` to automatically fetch all pages and stream results incrementally.

### Basic usage

```bash
# Fetch a single page (default behavior)
growthbook <command> --page 1

# Automatically fetch all pages
growthbook <command> --all
```

### Limiting pages

Use `--max-pages` to cap the number of pages fetched:

```bash
# Fetch at most 5 pages
growthbook <command> --all --max-pages 5
```

### Output formats

When using `--all`, results are streamed as each page is fetched:

| Format | Behavior |
|--------|----------|
| `--output-format json` | One JSON object per line ([NDJSON](https://github.com/ndjson/ndjson-spec)) |
| `--output-format yaml` | YAML documents separated by `---` |
| `--output-format toon` | One TOON-encoded block per item, separated by blank lines |
| Default (pretty) | Pretty-printed items separated by blank lines |

```bash
# Stream all results as NDJSON
growthbook <command> --all --output-format json

# Pipe to jq for further processing
growthbook <command> --all --output-format json | jq '.fieldName'

# Use the built-in --jq flag
growthbook <command> --all --jq '.fieldName'
```

### How it works

Under the hood, `--all` calls the operation once, then follows the underlying `Next()` pagination closure to fetch subsequent pages. Results are written to stdout as they arrive rather than buffered in memory, so this works well even with large result sets.

Without `--all`, paginated operations behave like any other command — pass cursor, page, offset, or limit flags manually and get a single page of results.
<!-- End Pagination [pagination] -->

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

This CLI follows semantic versioning and a stable command-versioning convention — see **Versioning and stability** at the top of this README. We recommend pinning to a version you have tested.

## Contributions

While we value open-source contributions to this CLI, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### CLI Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/growthbook/cli&utm_campaign=cli)
