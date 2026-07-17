## growthbook config-revisions set-metadata

Update config metadata in a draft revision

### Synopsis

Stages metadata changes (name, owner, description, project, lineage parent, extensibility) on the draft. Pass `version: "new"` to auto-create a draft. The change is only applied to the live config when the revision is merged.

```
growthbook config-revisions set-metadata [flags]
```

### Examples

```
  growthbook config-revisions set-metadata --key <key> --version-param <value>
```

### Options

```
      --body string                                   Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --description string                            string value
      --extends parent                                Replace the composition mixins layered on top of parent, in precedence order (later overrides earlier; all override `parent`; own keys win last). Send the complete set; an empty array clears all mixins.
      --extensible                                    boolean flag
  -h, --help                                          help for set-metadata
  -i, --ignore-warnings blockPublishOnSchemaError     Proceed despite soft validation warnings — e.g. publishing values that don't match the schema when the org has blockPublishOnSchemaError disabled (warn mode).
  -k, --key string                                    [required]
  -n, --name string                                   string value
      --owner string                                  The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.
      --parent key                                    Change the lineage parent (the key to inherit from). Empty string detaches from the parent.
      --project string                                string value
      --revision-comment string                       string value
      --revision-title string                         string value
  -s, --skip-schema-validation bypassApprovalChecks   Skip JSON-schema validation of the value(s) being written. Only honored for callers with org-wide bypass authority (the bypassApprovalChecks permission on all projects); ignored otherwise. Validation is enforced by default.
  -v, --version-param string                          [required]
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

* [growthbook config-revisions](growthbook_config-revisions.md)	 - Draft revisions for configs, including value and schema edits, schema import (JSON Schema / TypeScript / inferred), approvals, and lifecycle (publish, discard, revert)
