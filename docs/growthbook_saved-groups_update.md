## growthbook saved-groups update

Partially update a single saved group

### Synopsis

Partially update a single saved group

```
growthbook saved-groups update [flags]
```

### Examples

```
  growthbook saved-groups update --id <id>
```

### Options

```
      --body string                            Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -b, --bypass-approval bypassApprovalChecks   Set to true to skip the approval flow when the org requires approvals on saved groups. Requires the bypassApprovalChecks permission on the saved group's existing projects. When the org does not require approvals, this flag has no effect.
  -c, --condition string                       When type = 'condition', this is the JSON-encoded condition for the group
  -h, --help                                   help for update
  -i, --id string                              The id of the requested resource [required]
  -n, --name string                            The display name of the Saved Group
      --owner string                           The userId or email address of the owner. If an email address is provided, it will be used to look up the userId of the matching organization member. If an ID is provided, it will be validated as existing in the organization.
  -p, --projects stringArray                   list of values
  -v, --values stringArray                     When type = 'list', this is the list of values for the attribute key
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
  -o, --output-format string          Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --password string               HTTP Basic auth: use your GrowthBook Secret Key as the username and leave the password empty. password
      --profile growthbook profiles   Use a named credential/server profile (manage with growthbook profiles)
      --server string                 Select a server by index (for indexed servers) or name (for named servers)
      --server-url string             Override the default server URL
      --timeout string                HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                         Print the CLI Usage schema in KDL format
      --username string               HTTP Basic auth: use your GrowthBook Secret Key as the username and leave the password empty. username
```

### SEE ALSO

* [growthbook saved-groups](growthbook_saved-groups.md)	 - Defined sets of attribute values which can be used with feature rules for targeting features at particular users
