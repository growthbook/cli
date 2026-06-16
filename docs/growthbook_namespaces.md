## growthbook namespaces

Namespaces partition your user population into buckets so that experiments using the same hash attribute do not overlap unintentionally

### Synopsis

Namespaces partition your user population into buckets so that experiments using the same hash attribute do not overlap unintentionally. Each namespace defines a 0–1 range and individual experiments claim sub-ranges within it.

```
growthbook namespaces [flags]
```

### Options

```
  -h, --help   help for namespaces
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

* [growthbook](growthbook.md)	 - GrowthBook REST API: GrowthBook offers a full REST API for interacting with the application
* [growthbook namespaces delete](growthbook_namespaces_delete.md)	 - Delete a namespace
* [growthbook namespaces get](growthbook_namespaces_get.md)	 - Get a single namespace
* [growthbook namespaces get-namespace-memberships](growthbook_namespaces_get-namespace-memberships.md)	 - Get namespace membership
* [growthbook namespaces list](growthbook_namespaces_list.md)	 - Get all namespaces
* [growthbook namespaces post](growthbook_namespaces_post.md)	 - Create a namespace
* [growthbook namespaces post-namespace-rotate-seed](growthbook_namespaces_post-namespace-rotate-seed.md)	 - Rotate namespace seed
* [growthbook namespaces put](growthbook_namespaces_put.md)	 - Update a namespace
