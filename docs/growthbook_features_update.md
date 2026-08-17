## growthbook features update

Partially update a feature

### Synopsis

Updates the Feature Flag and immediately publishes a new revision. The caller needs Edit access in the Feature Flag's Project and Publish access for every affected environment. When approval is required, use the revision endpoints instead, unless the caller can bypass draft approvals.

Other top-level fields are patch-merged: omit a field to leave it unchanged. The `rules` field, when supplied, replaces the entire `rules` array in one operation. To preserve existing rules, fetch the Feature Flag, update the returned `rules` array, and send the complete array back. Safe-rollout rules round-trip through `safeRolloutId`; use `POST /v2/features/:id/revisions/:version/rules` to create new ones.

```
growthbook features update [flags]
```

### Examples

```
  growthbook features update --id <id>
```

### Options

```
      --body string                                   Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --body-param.archived                           boolean flag
      --body-param.base-config string                 The config backing this flag, fixed at creation. Cannot be changed by an update — resend the current value or omit it; a different value is rejected.
      --body-param.custom-fields string               value
      --body-param.default-value string               string value
      --body-param.default-value-config baseConfig    Optional. A config within baseConfig's family that the default value resolves to instead of `baseConfig` itself. null or omitted means the default is `baseConfig`. The default is exactly this config and carries no overrides of its own.
      --body-param.description string                 Description of the feature
      --body-param.environments rules                 Per-environment enabled state. V2 rules are specified on the top-level rules field.
      --body-param.holdout null                       Holdout to assign this feature to. Pass null to remove the feature from its current holdout. Omit the field entirely to leave the holdout unchanged.
                                                      
      --body-param.ignore-warnings                    Set to true to acknowledge the warnings listed in a blocked response and continue. This covers experiment guards, locked dependents, and references affected by an archive. When the organization treats schema failures as warnings, it also covers schema and invariant warnings. It never bypasses a rejected Custom Hook. On revision publish endpoints, it can also force-publish an out-of-date draft when the caller has Bypass draft approvals access.
      --body-param.json-schema string                 Use JSON schema to validate the payload of a JSON-type feature value (enterprise only).
      --body-param.owner string                       The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.
      --body-param.prerequisites true                 Feature IDs. Each feature must evaluate to true
      --body-param.project string                     An associated project ID
      --body-param.rules rules                        Replaces all feature rules atomically. Behavior differs from v1: v1 PUT applies per-environment patches, v2 PUT swaps the entire rules array in one revision. To preserve existing rules during a partial edit, GET the feature first, mutate the returned `rules` array, and PUT the full array back. Safe-rollout rules round-trip via their `safeRolloutId` (creation requires `POST /v2/features/:id/revisions/:version/rules`).
      --body-param.skip-hooks skipSchemaValidation    Set to true to publish despite a Custom Hook rejection. This does not bypass schema validation; use skipSchemaValidation for that. The caller must have Bypass draft approvals access for Feature Flags, Configs, and Constants in every Project. Otherwise, this field is ignored.
      --body-param.skip-schema-validation skipHooks   Set to true to publish despite schema validation errors, failed invariants, or schema changes that invalidate dependent resources. This does not bypass a rejected Custom Hook; use skipHooks for that. The caller must have Bypass draft approvals access for Feature Flags, Configs, and Constants in every Project. Otherwise, this field is ignored.
      --body-param.tags stringArray                   List of associated tags. Will override tags completely with submitted list
      --body-param.targeting-all-projects project     Make this feature discoverable in — and served to — every project, beyond its primary project. Governance/approvals stay with `project`.
      --body-param.targeting-projects project         Secondary project IDs this feature is targeted in and served to, beyond its primary project. Governance/approvals stay with `project`.
  -h, --help                                          help for update
      --id string                                     The id of the requested resource [required]
      --ignore-warnings ignoreWarnings                Deprecated — pass ignoreWarnings in the request body instead.
  -s, --skip-schema-validation skipSchemaValidation   Deprecated — pass skipSchemaValidation in the request body instead.
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
