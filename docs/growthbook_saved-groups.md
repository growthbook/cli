## growthbook saved-groups

Defined sets of attribute values which can be used with feature rules for targeting features at particular users

### Synopsis

Defined sets of attribute values which can be used with feature rules for targeting features at particular users.

```
growthbook saved-groups [flags]
```

### Options

```
  -h, --help   help for saved-groups
```

### Options inherited from parent commands

```
      --agent-mode                    Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth                   If using Bearer auth, pass the Secret Key as the token:
                                      `bash
                                      curl https://api.growthbook.io/api/v1/features   -H "Authorization: Bearer secret_abc123DEF456"
                                      ```
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
      --password                      If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                                      `bash
                                      curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                                      # The ":" at the end stops curl from asking for a password
                                      ```
                                       password
      --profile growthbook profiles   Use a named credential/server profile (manage with growthbook profiles)
      --server string                 Select a server by index (for indexed servers) or name (for named servers)
      --server-url string             Override the default server URL
      --timeout string                HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                         Print the CLI Usage schema in KDL format
      --username                      If using HTTP Basic auth, pass the Secret Key as the username and leave the password blank:
                                      `bash
                                      curl https://api.growthbook.io/api/v1/features   -u secret_abc123DEF456:
                                      # The ":" at the end stops curl from asking for a password
                                      ```
                                       username
```

### SEE ALSO

* [growthbook](growthbook.md)	 - GrowthBook REST API: GrowthBook offers a full REST API for interacting with the application
* [growthbook saved-groups archive](growthbook_saved-groups_archive.md)	 - Archive a single saved group
* [growthbook saved-groups delete](growthbook_saved-groups_delete.md)	 - Deletes a single saved group
* [growthbook saved-groups get](growthbook_saved-groups_get.md)	 - Get a single saved group
* [growthbook saved-groups get-saved-group-references](growthbook_saved-groups_get-saved-group-references.md)	 - Get features, experiments, and saved groups that reference this saved group
* [growthbook saved-groups list](growthbook_saved-groups_list.md)	 - Get all saved group
* [growthbook saved-groups post](growthbook_saved-groups_post.md)	 - Create a single saved group
* [growthbook saved-groups unarchive](growthbook_saved-groups_unarchive.md)	 - Unarchive a single saved group
* [growthbook saved-groups update](growthbook_saved-groups_update.md)	 - Partially update a single saved group
