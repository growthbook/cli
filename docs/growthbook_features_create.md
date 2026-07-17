## growthbook features create

Create a single feature

### Synopsis

Creates a new feature. Rules are supplied as a top-level `rules` array; each rule includes `allEnvironments` / `environments` scope fields.

### Config-backed features (Config mode)

A JSON feature can be backed by a shared **config** — the config supplies the base JSON value and schema, and the feature's *rule* values become override *patches* merged on top (nested objects deep-merge; arrays and scalars replace). The default value is exactly a config with no overrides (see below). Config backing is set exclusively through dedicated fields — never a raw `$extends: ["@config:…"]` inside a value string (that is rejected). `@const:` references inside values still work.

- **Top-level (`baseConfig`):** set `valueType: "json"` and `baseConfig: "<configKey>"` to put the flag in Config mode. The config must be live. This is the family root and the base the default value patches.
- **Default value:** unlike rules, the default is exactly a config with no overrides of its own — send `defaultValue: "{}"` to use `baseConfig`. To resolve the default to a *descendant* of `baseConfig` instead, set `defaultValueConfig` to that descendant's key (it must be within `baseConfig`'s family); omit/null to use `baseConfig` directly.
- **Rules & experiment variations:** each carries its own `config` field naming the family config that value patches (omit/null to patch the base). `value` is the override patch.

Example:

```json
{



  "id": "checkout-config",
  "valueType": "json",
  "baseConfig": "purchase-flow",
  "defaultValue": "{}",
  "rules": [
    { "type": "force", "config": "purchase-flow-vip", "value": "{\"maxItems\": 20}", "allEnvironments": true }
  ]
}
```

```
growthbook features create [flags]
```

### Examples

```
  growthbook features create --id <id> --value-type number --default-value <value>
```

### Options

```
  -a, --archived                                      boolean flag
  -b, --base-config valueType: "json"                 Key of the config backing this flag ("Config mode"). Requires valueType: "json" and a live config. The config supplies the base JSON and schema; `defaultValue` and rule values are override patches on top. null or omitted for a plain flag.
      --body string                                   Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --custom-fields string                          value
      --default-value valueType                       Default value when feature is enabled. Type must match valueType. In Config mode (`baseConfig` set) the default must be exactly a config with no overrides: send `"{}"` to use `baseConfig`, or set `defaultValueConfig` to point at a descendant. [required]
      --default-value-config baseConfig               Optional. A config within baseConfig's family that the default value resolves to instead of `baseConfig` itself. null or omitted means the default is `baseConfig`. The default is exactly this config and carries no overrides of its own.
      --description string                            Description of the feature
  -e, --environments rules                            Per-environment enabled state. V2 rules are specified on the top-level rules field.
  -h, --help                                          help for create
      --id string                                     A unique key name for the feature. Feature keys can only include letters, numbers, hyphens, and underscores. [required]
      --ignore-warnings blockPublishOnSchemaError     Proceed despite soft validation warnings — e.g. publishing values that don't match the schema when the org has blockPublishOnSchemaError disabled (warn mode).
  -j, --json-schema string                            Use JSON schema to validate the payload of a JSON-type feature value (enterprise only).
      --owner string                                  The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization. Optional when authenticating with a Personal Access Token (PAT): when omitted, the owner defaults to the PAT's user. Required when authenticating with an organization secret API key (which has no associated user): omitting it fails with a 400.
      --prerequisites true                            Feature IDs. Each feature must evaluate to true
      --project string                                An associated project ID
  -r, --rules allEnvironments                         Feature rules. Each rule carries its own environment scope via allEnvironments / `environments`.
  -s, --skip-schema-validation bypassApprovalChecks   Skip JSON-schema validation of the value(s) being written. Only honored for callers with org-wide bypass authority (the bypassApprovalChecks permission on all projects); ignored otherwise. Validation is enforced by default.
  -t, --tags stringArray                              List of associated tags
  -v, --value-type string                             The data type of the feature payload. Boolean by default. (options: boolean, string, number, json) [required]
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

* [growthbook features](growthbook_features.md)	 - Control your feature flags programatically
