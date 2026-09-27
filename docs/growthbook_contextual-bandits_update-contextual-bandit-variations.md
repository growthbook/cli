## growthbook contextual-bandits update-contextual-bandit-variations

Add or remove Contextual Bandit variations

### Synopsis

Adds and/or removes variations on a Contextual Bandit. Send `addVariations` and `removeVariationIds` independently; both are optional. New arms must carry a `values` entry for each linked feature. Running CBs publish the linked-feature updates; draft CBs stage them until start. Under an approval flow, unapproved drafts leave the added arm `pending` (zero weight, filtered from the SDK) until every linked feature's draft is live. Removed arms are tombstoned; their ids can never be re-added. Weights are reconciled server-side.

```
growthbook contextual-bandits update-contextual-bandit-variations [flags]
```

### Examples

```
  growthbook contextual-bandits update-contextual-bandit-variations --id <id>
```

### Options

```
  -a, --add-variations id                  New arms to add. Omit id to have the server generate one and `key` to have the server assign the next integer.
      --body string                        Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                               help for update-contextual-bandit-variations
  -i, --id string                          [required]
  -r, --remove-variation-ids stringArray   Ids of active arms to remove. Removed arms are tombstoned in place and their ids can never be re-added.
  -u, --update-variations name             Metadata edits to existing active arms. name, `description`, and `key` may be changed; values, weights, screenshots, and status are preserved.
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

* [growthbook contextual-bandits](growthbook_contextual-bandits.md)	 - Operations for contextual-bandits
