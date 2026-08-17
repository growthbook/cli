## growthbook contextual-bandits update-linked-feature

Replace a Contextual Bandit's rule on a linked feature

### Synopsis

Replaces every `contextual-bandit-ref` rule pointing at this contextual bandit on the feature, keeping each rule's id and position in the rule list. Every field is replaced, so omitted optional fields revert to their defaults. Returns a 400 when the feature has no such rule on the target revision, or when it has several that are not identical to each other. The change lands in a draft revision that auto-publishes when the contextual bandit starts, unless `autoPublish` is set. Targeting (condition, Saved Groups, prerequisites, coverage) is inherited from the contextual bandit and cannot be set on the rule.

```
growthbook contextual-bandits update-linked-feature [flags]
```

### Examples

```
  growthbook contextual-bandits update-linked-feature --id <id> --feature-id <id> --variations '[{"variationId":"<id>","value":"<value>"}]'
```

### Options

```
      --all-environments           Apply the rule in every environment. Defaults to true.
      --all-projects               boolean flag
      --auto-publish               Publish the resulting revision immediately instead of leaving it as a draft.
      --body string                Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --description string         string value
      --draft-version int          Update the rule on this existing draft revision instead of starting a new one.
      --enabled                    boolean flag
      --environments stringArray   Environments to apply the rule in when not all.
  -f, --feature-id string          The linked feature id [required]
  -h, --help                       help for update-linked-feature
  -i, --id string                  The Contextual Bandit id [required]
  -p, --projects stringArray       list of values
  -v, --variations string          One entry per Contextual Bandit variation. Every variation must be covered exactly once. [required]
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
