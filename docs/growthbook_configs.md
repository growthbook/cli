## growthbook configs

Reusable, typed, inheritable JSON objects referenced from feature flag values as `@config:key`

### Synopsis

Reusable, typed, inheritable JSON objects referenced from feature flag values as `@config:key`. A config carries a field `schema` (with TypeScript/JSON Schema import-export) and a lineage `parent`; it resolves like a `json` constant, composed via `$extends`. Inheritance is expressed via `parent`, never an in-value `@config:` entry. Schema fields colliding with a published ancestor's key follow 'base wins': identical re-declarations are stripped with a warning, differing ones are rejected.

```
growthbook configs [flags]
```

### Options

```
  -h, --help   help for configs
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

* [growthbook](growthbook.md)	 - GrowthBook REST API: A command-line interface for GrowthBook — manage feature flags, experiments, metrics, and more from your terminal
* [growthbook configs archive](growthbook_configs_archive.md)	 - Archive a single config
* [growthbook configs create](growthbook_configs_create.md)	 - Create a single config
* [growthbook configs delete](growthbook_configs_delete.md)	 - Delete a single config
* [growthbook configs get](growthbook_configs_get.md)	 - Get a single config
* [growthbook configs get-key-usage](growthbook_configs_get-key-usage.md)	 - Get the feature rules and default values implementing each key
* [growthbook configs get-lineage](growthbook_configs_get-lineage.md)	 - Get the full lineage (family tree) for a config
* [growthbook configs get-references](growthbook_configs_get-references.md)	 - Get features and configs that reference this config
* [growthbook configs get-schema](growthbook_configs_get-schema.md)	 - Export a config's schema
* [growthbook configs list](growthbook_configs_list.md)	 - Get all configs
* [growthbook configs lock](growthbook_configs_lock.md)	 - Lock a config at its current published revision
* [growthbook configs unarchive](growthbook_configs_unarchive.md)	 - Unarchive a single config
* [growthbook configs unlock](growthbook_configs_unlock.md)	 - Unlock a config
* [growthbook configs update](growthbook_configs_update.md)	 - Partially update a single config
* [growthbook configs verify-schema](growthbook_configs_verify-schema.md)	 - Verify a config's schema against a source (drift check)
