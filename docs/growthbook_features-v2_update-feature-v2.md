## growthbook features-v2 update-feature-v2

Partially update a feature

### Synopsis

Updates any combination of a feature's metadata, default value, environment state, and rules. Other top-level fields are patch-merged: omit a field to leave it unchanged. The `rules` field, when supplied, replaces the entire `rules` array atomically in a single revision (v1 PUT applied per-environment patches; v2 swaps the full flat array). To preserve existing rules during a partial edit, GET the feature first, mutate the returned `rules` array, and PUT the full array back. Safe-rollout rules round-trip via their `safeRolloutId`; use `POST /v2/features/:id/revisions/:version/rules` to create new ones. Returns 403 if approval rules are enabled for an affected environment and the bypass setting is off.

```
growthbook features-v2 update-feature-v2 [flags]
```

### Examples

```
  growthbook features-v2 update-feature-v2 --id <id>
```

### Options

```
  -a, --archived               boolean flag
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --custom-fields string   value
      --default-value string   string value
      --description string     Description of the feature
  -e, --environments rules     Per-environment enabled state. V2 rules are specified on the top-level rules field.
  -h, --help                   help for update-feature-v2
      --holdout null           Holdout to assign this feature to. Pass null to remove the feature from its current holdout. Omit the field entirely to leave the holdout unchanged.
                               
  -i, --id string              The id of the requested resource [required]
  -j, --json-schema string     Use JSON schema to validate the payload of a JSON-type feature value (enterprise only).
      --owner string           The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.
      --prerequisites true     Feature IDs. Each feature must evaluate to true
      --project string         An associated project ID
  -r, --rules rules            Replaces all feature rules atomically. Behavior differs from v1: v1 PUT applies per-environment patches, v2 PUT swaps the entire rules array in one revision. To preserve existing rules during a partial edit, GET the feature first, mutate the returned `rules` array, and PUT the full array back. Safe-rollout rules round-trip via their `safeRolloutId` (creation requires `POST /v2/features/:id/revisions/:version/rules`).
  -t, --tags stringArray       List of associated tags. Will override tags completely with submitted list
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

* [growthbook features-v2](growthbook_features-v2.md)	 - Control your feature flags programatically
