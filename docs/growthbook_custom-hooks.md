## growthbook custom-hooks

Sandboxed JavaScript validation hooks that run when features, configs, or their revisions are saved or published

### Synopsis

Sandboxed JavaScript validation hooks that run when features, configs, or their revisions are saved or published. Throwing an Error blocks the save; `addWarning(msg)` raises a soft warning. Hooks are scoped by projects, or pinned to a single feature/config via `entityType`/`entityId`; a config-scoped hook also runs for every config inheriting from it (its whole descendant lineage). Scope can be retargeted on update (or cleared with nulls). Requires an enterprise plan; not available on GrowthBook Cloud.

```
growthbook custom-hooks [flags]
```

### Options

```
  -h, --help   help for custom-hooks
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
* [growthbook custom-hooks delete](growthbook_custom-hooks_delete.md)	 - Delete a single custom hook
* [growthbook custom-hooks get](growthbook_custom-hooks_get.md)	 - Get a single custom hook
* [growthbook custom-hooks list](growthbook_custom-hooks_list.md)	 - Get all custom hooks
* [growthbook custom-hooks list-custom-hook-history](growthbook_custom-hooks_list-custom-hook-history.md)	 - List a custom hook's version history
* [growthbook custom-hooks post](growthbook_custom-hooks_post.md)	 - Create a single custom hook
* [growthbook custom-hooks revert](growthbook_custom-hooks_revert.md)	 - Revert a custom hook to a previous version
* [growthbook custom-hooks test](growthbook_custom-hooks_test.md)	 - Dry-run hook code in the sandbox
* [growthbook custom-hooks update](growthbook_custom-hooks_update.md)	 - Partially update a single custom hook
