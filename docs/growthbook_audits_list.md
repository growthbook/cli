## growthbook audits list

Get audit history for an entity or entity type

### Synopsis

Includes changes to child entities (e.g. a feature's revisions). Newest first. Requires permission to view audit logs.

```
growthbook audits list [flags]
```

### Examples

```
  growthbook audits list --entity-type rampScheduleTemplate
```

### Options

```
  -b, --before nextCursor    Only return entries older than this date-time. To page back, pass the previous page's nextCursor instead.
      --entity-id string     Omit for every entity of this type
      --entity-type string   options: agreement, approvalFlow, revision, aiPrompt, aiCredential, attribute, experiment, project, environment, feature, featureRevisionLog, urlRedirect, metric, metricAnalysis, metricGroup, populationData, population, datasource, comment, sdk-connection, user, organization, apiKey, oauthAuthCode, oauthGrant, oauthRefreshToken, installation, savedGroup, constant, config, segment, archetype, team, vercelNativeIntegration, factTable, customField, autoRun, experimentTemplate, safeRollout, decisionCriteria, execReport, holdout, savedQuery, dashboard, dashboardTemplate, incrementalRefresh, vector, customHook, ssoConnection, sqlResultChunk, rampSchedule, rampScheduleTemplate, learning, contextualBandit, interleaving, eventForwarderConfig [required]
  -h, --help                 help for list
  -l, --limit int            Defaults to 50
```

### Options inherited from parent commands

```
      --agent-mode                    Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth string            Bearer auth token: your Secret Key or Personal Access Token, sent as an Authorization Bearer header.
      --color string                  Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                         Log request and response diagnostics to stderr
      --domain string                 Server template variable: domain
      --dry-run                       Preview the request that would be sent without executing it (output to stderr)
  -H, --header stringArray            Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
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

* [growthbook audits](growthbook_audits.md)	 - Audit history of who changed what, and the event log that drives event webhooks
