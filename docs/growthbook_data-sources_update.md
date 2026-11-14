## growthbook data-sources update

Update a data source

### Synopsis

Fields left out keep their current values. To clear one, send an explicit empty value (`[]`, `""` or `false`). Connection settings are changed in the app.

```
growthbook data-sources update [flags]
```

### Examples

```
  growthbook data-sources update --id <id>
```

### Options

```
  -a, --assignment-queries id            Replaces the list. An entry with an existing id keeps `description`, `identifierTypes`, `includesNameColumns`, `dimensionColumns`, `projects` and `columnMapping` if left out; clear the others with `""`, `false` or `[]`. Each query is validated against the warehouse.
      --body string                      Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --description string               string value
  -e, --event-tracker eventForwarder     eventForwarder is only accepted unchanged (options: segment, snowplow, jitsu, freshpaint, ga4, fullstory, matomo, heap, rudderstack, amplitude, mparticle, mixpanel, firebase, keen, clevertap, langfuse, phoenix, eventForwarder, custom)
  -h, --help                             help for update
      --id string                        The id of the requested resource [required]
      --identifier-join-queries string   Replaces the list
      --identifier-types ""              Replaces the list. An entry matching an existing id keeps its description if left out; send "" to clear it.
      --max-concurrent-queries int       integer value
      --mixpanel-settings string         JSON object
  -n, --name string                      string value
      --pipeline-settings allowWriting   Merged into the stored settings: fields left out keep their values, but any change must leave allowWriting, `mode`, `writeDataset` and `unitsTableRetentionHours` set, with a non-empty `writeDataset` when `allowWriting` is true. Send `[]` to clear `excludedExperimentIds` or `incrementalOptInExperimentIds`, and `includedExperimentIds: null` to apply the mode to every experiment again (`[]` applies it to none).
      --project-ids stringArray          list of values
      --query-cache-ttl-mins int         integer value
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

* [growthbook data-sources](growthbook_data-sources.md)	 - How GrowthBook connects and queries your data, including cached database schema metadata (information schemas) for tables and columns
