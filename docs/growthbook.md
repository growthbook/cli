## growthbook

GrowthBook REST API: GrowthBook offers a full REST API for interacting with the application

### Synopsis

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
- **422** - Soft Warning - The request failed, but can be re-submitted with `?ignoreWarnings=true` to proceed anyway.
- **429** - Too Many Requests - You exceeded the rate limit of 60 requests per minute. Try again later.
- **5XX** - Server Error - Something went wrong on GrowthBook's end (these are rare)

The response body will be a JSON object with the following properties:

- **message** - Information about the error

```
growthbook [flags]
```

### Options

```
      --agent-mode                    Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth                   If using Bearer auth, pass the Secret Key as the token:
                                      `bash
                                      curl https://api.growthbook.io/api/v1/features   -H "Authorization: Bearer secret_abc123DEF456"
                                      ```
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
  -o, --output-format string          Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --password                      If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                                      `bash
                                      curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                                      # The ":" at the end stops curl from asking for a password
                                      ```
                                       password
      --profile growthbook profiles   Use a named credential/server profile (manage with growthbook profiles)
      --server string                 Select a server by index (for indexed servers) or name (for named servers)
      --server-url string             Override the default server URL
      --timeout string                HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                         Print the CLI Usage schema in KDL format
      --username                      If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                                      `bash
                                      curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                                      # The ":" at the end stops curl from asking for a password
                                      ```
                                       username
```

### SEE ALSO

* [growthbook SDK-connections](growthbook_SDK-connections.md)	 - Client keys and settings for connecting SDKs to a GrowthBook instance
* [growthbook analytics-explorations](growthbook_analytics-explorations.md)	 - Operations for analytics-explorations
* [growthbook archetypes](growthbook_archetypes.md)	 - Archetypes allow you to simulate the result of targeting rules on pre-set user attributes
* [growthbook attributes](growthbook_attributes.md)	 - Used when targeting feature flags and experiments
* [growthbook auth](growthbook_auth.md)	 - Manage authentication credentials
* [growthbook code-references](growthbook_code-references.md)	 - Intended for use with our code reference CI utility, [`gb-find-code-refs`](https://github
* [growthbook configure](growthbook_configure.md)	 - Configure authentication credentials and preferences
* [growthbook custom-fields](growthbook_custom-fields.md)	 - Operations for custom-fields
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
* [growthbook feature-revisions-v2](growthbook_feature-revisions-v2.md)	 - Draft revisions for feature flags, including rules, scheduling, and approval workflows
* [growthbook features](growthbook_features.md)	 - Control your feature flags programatically
* [growthbook features-v1](growthbook_features-v1.md)	 - Control your feature flags programatically
* [growthbook generate-types](growthbook_generate-types.md)	 - Generate TypeScript types for all your features
* [growthbook get-SDK-payload](growthbook_get-SDK-payload.md)	 - Get a SDK payload
* [growthbook get-visual-editor-bootstrap](growthbook_get-visual-editor-bootstrap.md)	 - Get Visual Editor Bootstrap
* [growthbook get-visual-editor-library-images](growthbook_get-visual-editor-library-images.md)	 - Get Visual Editor Library Images
* [growthbook members](growthbook_members.md)	 - Members are users who have been invited to an organization
* [growthbook meta](growthbook_meta.md)	 - Server metadata, including the running build's version and commit for version-skew checks
* [growthbook metric-groups](growthbook_metric-groups.md)	 - Operations for metric-groups
* [growthbook metrics](growthbook_metrics.md)	 - Metrics used as goals and guardrails for experiments
* [growthbook namespaces](growthbook_namespaces.md)	 - Namespaces partition your user population into buckets so that experiments using the same hash attribute do not overlap unintentionally
* [growthbook organizations](growthbook_organizations.md)	 - Organizations are used for multi-org deployments where different teams can run their own isolated feature flags and experiments
* [growthbook post-copy-transform](growthbook_post-copy-transform.md)	 - Post Copy Transform
* [growthbook post-visual-editor-AI-edit](growthbook_post-visual-editor-AI-edit.md)	 - Post Visual Editor AIEdit
* [growthbook post-visual-editor-AI-edit-resume](growthbook_post-visual-editor-AI-edit-resume.md)	 - Post Visual Editor AIEdit Resume
* [growthbook post-visual-editor-AI-image-gen](growthbook_post-visual-editor-AI-image-gen.md)	 - Post Visual Editor AIImage Gen
* [growthbook post-visual-editor-AI-promote-image](growthbook_post-visual-editor-AI-promote-image.md)	 - Post Visual Editor AIPromote Image
* [growthbook post-visual-editor-AI-suggestions](growthbook_post-visual-editor-AI-suggestions.md)	 - Post Visual Editor AISuggestions
* [growthbook post-visual-editor-AI-upload-signed-url](growthbook_post-visual-editor-AI-upload-signed-url.md)	 - Post Visual Editor AIUpload Signed Url
* [growthbook post-visual-editor-add-variant](growthbook_post-visual-editor-add-variant.md)	 - Post Visual Editor Add Variant
* [growthbook post-visual-editor-create-changeset](growthbook_post-visual-editor-create-changeset.md)	 - Post Visual Editor Create Changeset
* [growthbook post-visual-editor-create-experiment](growthbook_post-visual-editor-create-experiment.md)	 - Post Visual Editor Create Experiment
* [growthbook post-visual-editor-rename-experiment](growthbook_post-visual-editor-rename-experiment.md)	 - Post Visual Editor Rename Experiment
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
