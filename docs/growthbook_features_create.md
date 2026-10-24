## growthbook features create

Create a single feature

### Synopsis

Creates a new Feature Flag. The caller needs Create access in its Project, plus Publish access for any environment the Feature Flag starts enabled in — one that starts disabled everywhere needs Create alone. Rules are supplied as a top-level `rules` array; each rule includes `allEnvironments` / `environments` scope fields.

### Naming for human review

The `id` is the Feature Flag's permanent name in the UI and code. Keep it short (`checkout-express-pay`) and put context in `description`. To create it disabled, list every environment in `environments` with `enabled: false`.

### Config-backed features (Config mode)

A JSON feature can be backed by a shared **config** — the config supplies the base JSON value and schema, and the feature's *rule* values become override *patches* merged on top (nested objects deep-merge; arrays and scalars replace). The default value is exactly a config with no overrides (see below). Config backing is set exclusively through dedicated fields — never a raw `$extends: ["@config:…"]` inside a value string (that is rejected). `@const:` references inside values still work.

- **Top-level (`baseConfig`):** set `valueType: "json"` and `baseConfig: "<configKey>"` to put the Feature Flag in Config mode. The config must be live. This is the family root and the base the default value patches.
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
      --allow-unscoped-rules               When the organization requires rules to target specific environments or Projects, set to true to save a rule that applies to All Environments and All Projects. Only set this when a person has explicitly asked you to. If a check blocks you, stop and ask a person instead of working around it.
      --archived                           boolean flag
  -b, --base-config valueType: "json"      Key of the config backing this flag ("Config mode"). Requires valueType: "json" and a live config. The config supplies the base JSON and schema; `defaultValue` and rule values are override patches on top. null or omitted for a plain flag.
      --body string                        Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --comment string                     Comment to record on the feature's initial revision. Defaults to an empty comment.
      --custom-fields string               value
      --default-value valueType            Default value when feature is enabled. Type must match valueType. In Config mode (`baseConfig` set) the default must be exactly a config with no overrides: send `"{}"` to use `baseConfig`, or set `defaultValueConfig` to point at a descendant. [required]
      --default-value-config baseConfig    Optional. A config within baseConfig's family that the default value resolves to instead of `baseConfig` itself. null or omitted means the default is `baseConfig`. The default is exactly this config and carries no overrides of its own.
      --description string                 Markdown. What the Feature Flag gates and why, rollout plan, and links. Put context here, not in the key.
  -e, --environments enabled: false        Per-environment enabled state, keyed by environment ID. An environment you leave out uses its "Default state for new features" setting, which may be enabled. To create the Feature Flag disabled, list every environment with enabled: false. Rules go on the top-level `rules` field.
  -h, --help                               help for create
      --id featureRegexValidator           Unique, permanent key, shown as the Feature Flag's name in the UI and used in code. Keep it short. If featureRegexValidator in `GET /api/v1/settings` is set, follow it (`featureKeyExample` shows the format); otherwise use kebab-case (e.g. `checkout-express-pay`). Put context in `description`. [required]
      --ignore-warnings                    Set to true to acknowledge the warnings listed in a blocked response and continue. This covers experiment guards, locked dependents, and references affected by an archive. When the organization treats schema failures as warnings, it also covers schema and invariant warnings. It never bypasses a rejected Custom Hook. On revision publish endpoints, it can also force-publish an out-of-date draft when the caller has Bypass draft approvals access. Only set this when a person has explicitly asked you to. If a check blocks you, stop and ask a person instead of working around it.
  -j, --json-schema string                 Use JSON schema to validate the payload of a JSON-type feature value (enterprise only).
      --owner string                       The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization. When omitted, it defaults to the requesting member: a personal access token's owner, or the member an organization API key assumes. An organization API key acting as itself must send it, or the request fails with a 400.
      --prerequisites true                 Feature IDs. Each feature must evaluate to true
      --project string                     An associated project ID
  -r, --rules allEnvironments              Feature rules. Each rule carries its own environment scope via allEnvironments / `environments`.
      --skip-hooks skipSchemaValidation    Set to true to publish despite a Custom Hook rejection. This does not bypass schema validation; use skipSchemaValidation for that. The caller must have Bypass draft approvals access for Feature Flags, Configs, and Constants in every Project. Otherwise, this field is ignored. Only set this when a person has explicitly asked you to. If a check blocks you, stop and ask a person instead of working around it.
      --skip-schema-validation skipHooks   Set to true to publish despite schema validation errors, failed invariants, or schema changes that invalidate dependent resources. This does not bypass a rejected Custom Hook; use skipHooks for that. The caller must have Bypass draft approvals access for Feature Flags, Configs, and Constants in every Project. Otherwise, this field is ignored. Only set this when a person has explicitly asked you to. If a check blocks you, stop and ask a person instead of working around it.
      --tags stringArray                   Tags for filtering and grouping. Use these instead of encoding groups in names.
      --targeting-all-projects project     Make this feature discoverable in — and served to — every project, beyond its primary project. Requires the `targetFeatures` permission (FlagsTarget policy) unscoped to any project. Governance stays with `project`.
      --targeting-projects project         Secondary project IDs this feature is targeted in and served to, beyond its primary project. Adding a project requires the `targetFeatures` permission (FlagsTarget policy) in that project. Governance stays with `project`.
  -v, --value-type string                  The data type of the feature payload. Boolean by default. (options: boolean, string, number, json) [required]
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

* [growthbook features](growthbook_features.md)	 - Control your feature flags programmatically
