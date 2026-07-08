## growthbook code-references post-code-refs

Submit list of code references

### Synopsis

Submit list of code references

```
growthbook code-references post-code-refs [flags]
```

### Examples

```
  growthbook code-references post-code-refs --branch <value> --repo-name <value> --refs '[{"filePath":"/bin/shameful_sans_till.xlw","startingLineNumber":873333,"lines":"<value>","flagKey":"<value>","contentHash":"<value>"}]'
```

### Options

```
      --body string             Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -b, --branch string           [required]
      --delete-missing string   Whether to delete code references that are no longer present in the submitted data (options: true, false) (default "false")
  -h, --help                    help for post-code-refs
      --refs string             [required]
      --repo-name string        [required]
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
      --password string               HTTP Basic auth: your Secret Key as the username, with an empty password. password
      --profile growthbook profiles   Use a named credential/server profile (manage with growthbook profiles)
      --server string                 Select a server by index (for indexed servers) or name (for named servers)
      --server-url string             Override the default server URL
      --timeout string                HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                         Print the CLI Usage schema in KDL format
      --username string               HTTP Basic auth: your Secret Key as the username, with an empty password. username
```

### SEE ALSO

* [growthbook code-references](growthbook_code-references.md)	 - Intended for use with our code reference CI utility, [`gb-find-code-refs`](https://github
