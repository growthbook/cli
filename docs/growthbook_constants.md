## growthbook constants

Reusable named values referenced from feature flag values as `@const:key` and resolved into the SDK payload at build time

### Synopsis

Reusable named values referenced from feature flag values as `@const:key` and resolved into the SDK payload at build time. String constants are interpolated via `{ { @const:key } }`; JSON (object) constants are composed via an `$extends` array.

```
growthbook constants [flags]
```

### Options

```
  -h, --help   help for constants
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
  -o, --output-format string          Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --password string               HTTP Basic auth: use your GrowthBook Secret Key as the username and leave the password empty. password
      --profile growthbook profiles   Use a named credential/server profile (manage with growthbook profiles)
      --server string                 Select a server by index (for indexed servers) or name (for named servers)
      --server-url string             Override the default server URL
      --timeout string                HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                         Print the CLI Usage schema in KDL format
      --username string               HTTP Basic auth: use your GrowthBook Secret Key as the username and leave the password empty. username
```

### SEE ALSO

* [growthbook](growthbook.md)	 - GrowthBook REST API: A command-line interface for GrowthBook — manage feature flags, experiments, metrics, and more from your terminal
* [growthbook constants archive](growthbook_constants_archive.md)	 - Archive a single constant
* [growthbook constants create](growthbook_constants_create.md)	 - Create a single constant
* [growthbook constants delete](growthbook_constants_delete.md)	 - Delete a single constant
* [growthbook constants get](growthbook_constants_get.md)	 - Get a single constant
* [growthbook constants get-constant-references](growthbook_constants_get-constant-references.md)	 - Get features and constants that reference this constant
* [growthbook constants list](growthbook_constants_list.md)	 - Get all constants
* [growthbook constants unarchive](growthbook_constants_unarchive.md)	 - Unarchive a single constant
* [growthbook constants update](growthbook_constants_update.md)	 - Partially update a single constant
