## growthbook features-v1 create

Create a single feature

### Synopsis

DEPRECATED: This will be removed in a future release, please migrate away from it as soon as possible

**Deprecated.** Use [POST /v2/features](#operation/postFeatureV2) instead.

```
growthbook features-v1 create [flags]
```

### Examples

```
  growthbook features-v1 create --id <id> --value-type number --default-value <value>
```

### Options

```
  -a, --archived                    boolean flag
      --body string                 Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --custom-fields string        value
      --default-value string        Default value when feature is enabled. Type must match valueType. [required]
      --description string          Description of the feature
  -e, --environments string         A dictionary of environments that are enabled for this feature. Keys supply the names of environments. Environments belong to organization and are not specified will be disabled by default.
  -h, --help                        help for create
  -i, --id string                   A unique key name for the feature. Feature keys can only include letters, numbers, hyphens, and underscores. [required]
  -j, --json-schema string          Use JSON schema to validate the payload of a JSON-type feature value (enterprise only).
      --owner string                The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization. When omitted, it defaults to the user associated with the request's Personal Access Token (PAT), if one is being used.
      --prerequisites stringArray   Feature IDs. Each feature must evaluate to true
      --project string              An associated project ID
  -t, --tags stringArray            List of associated tags
  -v, --value-type string           The data type of the feature payload. Boolean by default. (options: boolean, string, number, json) [required]
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
