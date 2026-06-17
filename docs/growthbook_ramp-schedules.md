## growthbook ramp-schedules

Multi-step rollout schedules that gradually increase feature rule traffic over time, with optional real-time monitoring

### Synopsis

Multi-step rollout schedules that gradually increase feature rule traffic over time, with optional real-time monitoring. Each step supports interval timers, approval gates, and hold conditions. Monitored steps are backed by a live analysis experiment that can automatically hold, roll back, or advance the ramp based on guardrail and signal metric health.

```
growthbook ramp-schedules [flags]
```

### Options

```
  -h, --help   help for ramp-schedules
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
* [growthbook ramp-schedules add-target](growthbook_ramp-schedules_add-target.md)	 - Add a target rule to a ramp schedule
* [growthbook ramp-schedules api-advance](growthbook_ramp-schedules_api-advance.md)	 - Advance to the next step, overriding any holds
* [growthbook ramp-schedules approve-step](growthbook_ramp-schedules_approve-step.md)	 - Approve the current step
* [growthbook ramp-schedules complete](growthbook_ramp-schedules_complete.md)	 - Complete a ramp schedule immediately
* [growthbook ramp-schedules delete](growthbook_ramp-schedules_delete.md)	 - Delete a single rampSchedule
* [growthbook ramp-schedules eject-target](growthbook_ramp-schedules_eject-target.md)	 - Remove a target rule from a ramp schedule
* [growthbook ramp-schedules get](growthbook_ramp-schedules_get.md)	 - Get a single rampSchedule
* [growthbook ramp-schedules get-ramp-schedule-status](growthbook_ramp-schedules_get-ramp-schedule-status.md)	 - Get ramp schedule status summary
* [growthbook ramp-schedules jump](growthbook_ramp-schedules_jump.md)	 - Jump to a specific step
* [growthbook ramp-schedules list](growthbook_ramp-schedules_list.md)	 - Get all rampSchedules
* [growthbook ramp-schedules pause](growthbook_ramp-schedules_pause.md)	 - Pause a ramp schedule
* [growthbook ramp-schedules post](growthbook_ramp-schedules_post.md)	 - Create a ramp schedule
* [growthbook ramp-schedules refresh-monitoring](growthbook_ramp-schedules_refresh-monitoring.md)	 - Trigger a manual monitoring update
* [growthbook ramp-schedules restart](growthbook_ramp-schedules_restart.md)	 - Restart a terminal ramp schedule
* [growthbook ramp-schedules resume](growthbook_ramp-schedules_resume.md)	 - Resume a paused ramp schedule
* [growthbook ramp-schedules rollback](growthbook_ramp-schedules_rollback.md)	 - Roll back a ramp schedule
* [growthbook ramp-schedules set-auto-update](growthbook_ramp-schedules_set-auto-update.md)	 - Toggle automatic monitoring updates
* [growthbook ramp-schedules set-monitoring-mode](growthbook_ramp-schedules_set-monitoring-mode.md)	 - Set ramp monitoring mode
* [growthbook ramp-schedules start](growthbook_ramp-schedules_start.md)	 - Start a ramp schedule
* [growthbook ramp-schedules update](growthbook_ramp-schedules_update.md)	 - Update a single rampSchedule
* [growthbook ramp-schedules update-ramp-schedule-lockdown](growthbook_ramp-schedules_update-ramp-schedule-lockdown.md)	 - Update ramp lockdown configuration
* [growthbook ramp-schedules update-ramp-schedule-monitoring](growthbook_ramp-schedules_update-ramp-schedule-monitoring.md)	 - Update ramp monitoring configuration
* [growthbook ramp-schedules update-ramp-schedule-steps](growthbook_ramp-schedules_update-ramp-schedule-steps.md)	 - Update ramp schedule steps
