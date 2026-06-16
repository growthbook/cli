## growthbook namespaces post-namespace-rotate-seed

Rotate namespace seed

### Synopsis

⚠️ Dangerous: sets a new seed for a multiRange namespace. Every user's bucket position within the namespace is re-computed immediately, which re-randomizes traffic eligibility for **all** experiments currently using this namespace. Only do this if you intentionally want to reshuffle all allocations across experiments. This could be useful when re-using a namespace for a new set of experiments.

```
growthbook namespaces post-namespace-rotate-seed [flags]
```

### Examples

```
  growthbook namespaces post-namespace-rotate-seed --id <id>
```

### Options

```
      --body string   Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help          help for post-namespace-rotate-seed
  -i, --id string     The unique id of the namespace [required]
  -s, --seed string   A specific value to use as the new seed. If omitted, a random value is generated.
```

### Options inherited from parent commands

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth            If using Bearer auth, pass the Secret Key as the token:
                               `bash
                               curl https://api.growthbook.io/api/v1/features   -H "Authorization: Bearer secret_abc123DEF456"
                               ```
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --domain string          Server template variable: domain
      --dry-run                Preview the request that would be sent without executing it (output to stderr)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
      --include-headers        Include HTTP response headers in the output
  -q, --jq string              Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive         Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --password               If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                               `bash
                               curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                               # The ":" at the end stops curl from asking for a password
                               ```
                                password
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
      --username               If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                               `bash
                               curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                               # The ":" at the end stops curl from asking for a password
                               ```
                                username
```

### SEE ALSO

* [growthbook namespaces](growthbook_namespaces.md)	 - Namespaces partition your user population into buckets so that experiments using the same hash attribute do not overlap unintentionally
