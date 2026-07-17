## growthbook

GrowthBook REST API: A command-line interface for GrowthBook — manage feature flags, experiments, metrics, and more from your terminal

### Synopsis

GrowthBook REST API: A command-line interface for GrowthBook — manage feature flags, experiments, metrics, and more from your terminal.

Endpoints are versioned by path prefix: `/v1` (stable) and `/v2` (improved shapes). Each command group targets the newest version of its endpoint; superseded versions remain available under a `-vN` suffix.

Authenticate with a Secret Key or Personal Access Token via `--bearer-auth` (or the `GBCLI_BEARER_AUTH` environment variable). Run `growthbook configure` to store credentials, or `growthbook whoami` to check the active configuration.

```
growthbook [flags]
```

### Options

```
      --agent-mode                    Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth string            Bearer auth token: your Secret Key or Personal Access Token, sent as an Authorization Bearer header.
      --color string                  Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                         Log request and response diagnostics to stderr
      --domain string                 Server template variable: domain
      --dry-run                       Preview the request that would be sent without executing it (output to stderr)
  -H, --header stringArray            Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
  -h, --help                          help for growthbook
      --include-headers               Include HTTP response headers in the output
  -q, --jq string                     Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive                Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
      --no-update-check               Disable the once-a-day check for a newer CLI version
  -o, --output-format string          Specify the output format. Options: pretty, json, yaml, toon. (default "pretty")
      --password string               HTTP Basic auth: use your GrowthBook Secret Key as the username and leave the password empty. password
      --profile growthbook profiles   Use a named credential/server profile (manage with growthbook profiles)
      --server string                 Select a server by index (for indexed servers) or name (for named servers)
      --server-url string             Override the default server URL
      --timeout string                HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                         Print the CLI Usage schema in KDL format
      --username string               HTTP Basic auth: use your GrowthBook Secret Key as the username and leave the password empty. username
```

### SEE ALSO

* [growthbook SDK-connections](growthbook_SDK-connections.md)	 - Client keys and settings for connecting SDKs to a GrowthBook instance
* [growthbook analytics-explorations](growthbook_analytics-explorations.md)	 - Operations for analytics-explorations
* [growthbook archetypes](growthbook_archetypes.md)	 - Archetypes allow you to simulate the result of targeting rules on pre-set user attributes
* [growthbook attributes](growthbook_attributes.md)	 - Used when targeting feature flags and experiments
* [growthbook auth](growthbook_auth.md)	 - Manage authentication credentials
* [growthbook code-references](growthbook_code-references.md)	 - Intended for use with our code reference CI utility, [`gb-find-code-refs`](https://github
* [growthbook config-revisions](growthbook_config-revisions.md)	 - Draft revisions for configs, including value and schema edits, schema import (JSON Schema / TypeScript / inferred), approvals, and lifecycle (publish, discard, revert)
* [growthbook configs](growthbook_configs.md)	 - Reusable, typed, inheritable JSON objects referenced from feature flag values as `@config:key`
* [growthbook configure](growthbook_configure.md)	 - Configure authentication credentials and preferences
* [growthbook constant-revisions](growthbook_constant-revisions.md)	 - Draft revisions for constants, including pending changes, approvals, and lifecycle (publish, discard, revert)
* [growthbook constants](growthbook_constants.md)	 - Reusable named values referenced from feature flag values as `@const:key` and resolved into the SDK payload at build time
* [growthbook contextual-bandit-queries](growthbook_contextual-bandit-queries.md)	 - Operations for contextual-bandit-queries
* [growthbook contextual-bandits](growthbook_contextual-bandits.md)	 - Operations for contextual-bandits
* [growthbook copy-transform](growthbook_copy-transform.md)	 - Copy Transform
* [growthbook custom-fields](growthbook_custom-fields.md)	 - Operations for custom-fields
* [growthbook custom-hooks](growthbook_custom-hooks.md)	 - Sandboxed JavaScript validation hooks that run when features, configs, or their revisions are saved or published
* [growthbook dashboards](growthbook_dashboards.md)	 - Operations for dashboards
* [growthbook data-sources](growthbook_data-sources.md)	 - How GrowthBook connects and queries your data, including cached database schema metadata (information schemas) for tables and columns
* [growthbook dimensions](growthbook_dimensions.md)	 - Dimensions used during experiment analysis
* [growthbook environments](growthbook_environments.md)	 - GrowthBook comes with one environment by default (production), but you can add as many as you need
* [growthbook experiment-templates](growthbook_experiment-templates.md)	 - Operations for experiment-templates
* [growthbook experiments](growthbook_experiments.md)	 - Experiments (A/B Tests)
* [growthbook explore](growthbook_explore.md)	 - Interactively browse and run commands
* [growthbook fact-metrics](growthbook_fact-metrics.md)	 - Fact Metrics are metrics built on top of Fact Table definitions
* [growthbook fact-tables](growthbook_fact-tables.md)	 - Fact Tables describe the shape of your data warehouse tables
* [growthbook feature-revisions](growthbook_feature-revisions.md)	 - Draft revisions for feature flags, including rules, scheduling, and approval workflows
* [growthbook feature-revisions-v1](growthbook_feature-revisions-v1.md)	 - Draft revisions for feature flags, including rules, scheduling, and approval workflows
* [growthbook features](growthbook_features.md)	 - Control your feature flags programatically
* [growthbook features-v1](growthbook_features-v1.md)	 - Control your feature flags programatically
* [growthbook generate-types](growthbook_generate-types.md)	 - Generate TypeScript types for all your features
* [growthbook get-SDK-payload](growthbook_get-SDK-payload.md)	 - Get a SDK payload
* [growthbook members](growthbook_members.md)	 - Members are users who have been invited to an organization
* [growthbook meta](growthbook_meta.md)	 - Server metadata, including the running build's version and commit for version-skew checks
* [growthbook metric-groups](growthbook_metric-groups.md)	 - Operations for metric-groups
* [growthbook metrics](growthbook_metrics.md)	 - Metrics used as goals and guardrails for experiments
* [growthbook namespaces](growthbook_namespaces.md)	 - Namespaces partition your user population into buckets so that experiments using the same hash attribute do not overlap unintentionally
* [growthbook organizations](growthbook_organizations.md)	 - Organizations are used for multi-org deployments where different teams can run their own isolated feature flags and experiments
* [growthbook profiles](growthbook_profiles.md)	 - Manage named credential/server profiles
* [growthbook projects](growthbook_projects.md)	 - Projects are used to organize your feature flags and experiments
* [growthbook queries](growthbook_queries.md)	 - Retrieve queries used in experiments to calculate results
* [growthbook ramp-schedule-templates](growthbook_ramp-schedule-templates.md)	 - Reusable step configurations for ramp schedules
* [growthbook ramp-schedules](growthbook_ramp-schedules.md)	 - Multi-step rollout schedules that gradually increase feature rule traffic over time, with optional real-time monitoring
* [growthbook reports](growthbook_reports.md)	 - Custom analysis reports built on top of experiment snapshots
* [growthbook saved-group-revisions](growthbook_saved-group-revisions.md)	 - Draft revisions for saved groups, including pending changes, approvals, and lifecycle (publish, discard, revert)
* [growthbook saved-groups](growthbook_saved-groups.md)	 - Defined sets of attribute values which can be used with feature rules for targeting features at particular users
* [growthbook segments](growthbook_segments.md)	 - Segments used during experiment analysis
* [growthbook settings](growthbook_settings.md)	 - Get the organization settings
* [growthbook snapshots](growthbook_snapshots.md)	 - Experiment Snapshots (the individual updates of an experiment)
* [growthbook teams](growthbook_teams.md)	 - Operations for teams
* [growthbook usage-metrics](growthbook_usage-metrics.md)	 - Operations for usage-metrics
* [growthbook version](growthbook_version.md)	 - Print the CLI version
* [growthbook visual-changesets](growthbook_visual-changesets.md)	 - Groups of visual changes made by the visual editor to a single page
* [growthbook whoami](growthbook_whoami.md)	 - Display current authentication configuration
