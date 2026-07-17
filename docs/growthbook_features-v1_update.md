## growthbook features-v1 update

Partially update a feature

### Synopsis

DEPRECATED: This will be removed in a future release, please migrate away from it as soon as possible

**Deprecated.** Use [POST /v2/features/:id](#operation/updateFeatureV2) instead.

Updates any combination of a feature's metadata (description, owner, tags, project), default value, environment settings (rules, kill switches, enabled state), prerequisites, holdout assignment, or JSON schema validation. All provided fields are merged into the existing feature and the result is immediately published as a new revision.

Returns 403 if the API key lacks permission or if approval rules are enabled for an affected environment and the org setting "REST API always bypasses approval requirements" is off.

```
growthbook features-v1 update [flags]
```

### Examples

```
  growthbook features-v1 update --id <id>
```

### Options

```
  -a, --archived                    boolean flag
  -b, --base-config string          The config backing this flag ("Config mode"), fixed at creation. Cannot be changed by an update — resend the current value or omit it; a different value (or null to detach) is rejected.
      --body string                 Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --custom-fields string        value
      --default-value string        string value
      --description string          Description of the feature
  -e, --environments string         value
  -h, --help                        help for update
      --holdout string              Holdout to assign this feature to. Pass null to remove the feature from its current holdout. Omit the field entirely to leave the holdout unchanged.
                                    
  -i, --id string                   The id of the requested resource [required]
  -j, --json-schema string          Use JSON schema to validate the payload of a JSON-type feature value (enterprise only).
      --owner string                The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.
      --prerequisites stringArray   Feature IDs. Each feature must evaluate to true
      --project string              An associated project ID
  -t, --tags stringArray            List of associated tags. Will override tags completely with submitted list
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

* [growthbook features-v1](growthbook_features-v1.md)	 - Control your feature flags programatically
