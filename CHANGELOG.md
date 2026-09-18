# Changelog

## [Unreleased]

### Changed

- Requests send `User-Agent: growthbook-cli/<version> (<go version>; <os>/<arch>)` instead of the
  Speakeasy generator's default, so GrowthBook can attribute API usage to the CLI and to a real
  release rather than the generator's version. A `User-Agent` supplied with `-H`/`--header` is kept
  as a second product token behind the CLI's own instead of replacing it, so a proxy or wrapper can
  still identify itself.
- The once-a-day compatibility check sends the same `User-Agent`. It bypasses the SDK client, so it
  previously went out as Go's default and was counted as generic Go traffic rather than as the CLI.

## [2.6.0] - 2026-09-15

Additive, plus one flag fix. Twenty-four new commands and twenty-two new flags; nothing was
removed, and `--dry-run` behaves exactly as before everywhere, so no action is needed on upgrade.

### Added

- **`growthbook holdouts` — a new command group.** Manage holdouts end to end: `create`, `get`,
  `list`, `update`, `start` (begin the active period), `start-analysis` (begin the analysis
  period), and `stop`. `list` filters on `--stage`, `--datasource-id`, `--project-id`, and
  `--archived`.

- **`growthbook auto-runs` — a new command group.** `create`, `get`, `list`, `update`, and
  `append-artifact` for recording an agent's run and the artifacts it produced.

- **Product-analytics commands on `growthbook analytics-explorations`.** `get` returns a saved
  exploration by id; `search` finds explorable resources; `get-columns` and `get-column-values`
  enumerate a source's columns and their values; `run-sql` runs a SQL exploration alongside the
  existing `run-metric` / `run-funnel` / `run-fact-table` / `run-data-source`.

- **Warehouse SQL commands on `growthbook data-sources`.** `search-warehouse-tables`,
  `get-warehouse-table-schema`, `preview-warehouse-column-values`, and `run-sql-query` for
  exploring a data source's warehouse directly.

- **`growthbook experiments post-comment`.** Add a comment to an experiment.

- **`--restrict-access` on `growthbook projects create` / `update`.** When true, only members with
  an explicit role on the project — directly or via a team — can access it. Members with
  `manageTeam` keep access. Requires a Pro or Enterprise plan.

- **`--include-referenced-prerequisites` on `growthbook SDK-connections create` / `update`.**
  Carry prerequisite feature flags into the payload even when they target other projects. Defaults
  to true for new connections.

- **`--include-archived` on `growthbook metrics list`.** Defaults to `true`; pass `false` to return
  only non-archived metrics.

- **`--replaces` on `growthbook fact-metrics create` / `update`.** Ids of older metrics this one
  supersedes. Informational only — it links the old and new definitions in the UI and keeps
  showing results from snapshots taken before an experiment switched to this metric. Settable only
  through the API.

- **`--creatable` on `growthbook custom-fields create` / `update`.** For `enum` and `multiselect`
  fields, let users enter values beyond the predefined list.

- **`--timestamp-column` and `--user-id-columns` on `growthbook fact-tables create` / `update`.**
  Name the event-timestamp column, and map identifier types to columns for SQL that does not alias
  its columns to the identifier-type names.

- **`--default-managed-by` on `growthbook fact-tables bulk-import`.** Fallback `managedBy` for fact
  tables and fact metrics that omit it. Defaults to `api`.

- **`--comparison` on `growthbook dashboards create` / `update`.** Dashboard-wide
  compare-to-previous-period; takes precedence over any per-block comparison.

- **`--chart-settings` on the four `growthbook analytics-explorations run-*` commands.**

- **`sql-exploration` dashboard blocks.** A new block type returned by every `growthbook
  dashboards` command and accepted inside `blocks[]` — a nested body field, so pass it in `--body`
  JSON.

- **Dimension breakdowns on results.** `dimension.dimensions` is now returned under
  `growthbook experiments results`, `growthbook experiments list-results`, and every `growthbook
  reports` command that returns results, with `date`, `dynamic`, `static`, and `slice` variants.
  Response-only.

- **`force` on ramp-schedule template patch actions**, under `steps[].actions[].patch` on every
  `growthbook ramp-schedule-templates` command, and **`actions` on**
  `growthbook ramp-schedules update-ramp-schedule-steps` steps.

### Fixed

- **`--dry-run` on `growthbook fact-tables bulk-import` and `growthbook releases
  publish-revisions` never reached the server; the API's own dry run is now `--validate-only`.**
  Both endpoints take a `dryRun` body field meaning "validate and report, write nothing." It
  generated a command-local `--dry-run` flag that shadowed the CLI's global `--dry-run` ("preview
  the request, skip the network call"), and because the global is read off the command, passing it
  both set `dryRun: true` in the payload *and* short-circuited to `Network call skipped` — so the
  server-side validation could never actually run, and `--dry-run` vanished from the command's
  `Diagnostics` help section.

  The body field's flag is now `--validate-only` (`-v`); the wire field is still `dryRun`, so the
  request is unchanged. `--dry-run` goes back to meaning what it means on every other command, and
  does exactly what it did before on these two — it always skipped the network call, so no script
  changes behavior. To get the server-side gate report, switch `--dry-run` to `--validate-only`.

### Changed

- **`--number-format` accepts `time:milliseconds`.** A new option on `growthbook fact-tables
  create-virtual-column` and `update-virtual-column`, alongside `currency`, `time:seconds`,
  `memory:bytes`, and `memory:kilobytes`. Purely additive — existing values are unaffected.

## [2.5.0] - 2026-08-31

Additive only — two new flags, no commands or flags were removed or renamed, so no action is
needed on upgrade.

### Added

- **`--attribute-scope-all-projects` on `growthbook experiments create` and `growthbook
  experiments update`.** A picker preference: show attributes from all projects in the
  experiment's targeting UI instead of only those in scope for its project and linked features.
  It does not loosen enforcement — when the organization requires registered attributes with
  project scoping, out-of-scope attributes are still rejected.

- **`effectStandardError` on analysis results.** The standard error of the estimated effect,
  alongside the `percentChange` it belongs to. Returned by `growthbook experiments results`,
  `growthbook experiments list-results`, and every `growthbook reports` command that returns
  results, under `results[].metrics[].variations[].analyses[]`. Response-only, always present.

- **`effectStandardError` on experiment result variations.** The standard error of `lift`, under
  `experimentResults[].variations[]` from `growthbook metrics list-experiments`. Response-only
  and optional — returned only when available.

### Changed

- **`scheduledStopPlan.fallback` is now optional.** It was required on every scheduled stop plan;
  it is now required only when `mode` is `"auto-ship"`, and ignored for other modes. A relaxation,
  so nothing to migrate — plans that already send `fallback` keep working, and `notify`,
  `force-ship`, and `stop` plans can now omit it. Affects `growthbook experiments set-schedule
  --scheduled-stop-plan` and the `statusUpdateSchedule.scheduledStopPlan` body field on
  `growthbook experiments create` / `update`.

## [2.4.0] - 2026-08-26

Additive only — no commands or flags were added, removed, or renamed, so no action is needed on
upgrade. The new fields are all nested, so there are no new flags.

### Added

- **`errorMessage` on analysis results.** Each `analyses[]` entry under
  `results[].metrics[].variations[]` can now carry the reason that analysis failed, instead of
  reporting absent numbers with no explanation. Returned by `growthbook experiments results`,
  `growthbook experiments list-results`, and every `growthbook reports` command that returns
  results. Response-only.

- **`linkedFunnelMetricId` on funnel-exploration dashboard blocks.** Returned by every `growthbook
  dashboards` command that returns blocks, and accepted by `growthbook dashboards update` inside
  `blocks[]` — a nested body field, so pass it in `--body` JSON. Nullable. `dashboards create`
  does not take funnel-exploration blocks, so it is update-only on the request side.

### Changed

- **`seed` is now always returned on contextual bandits.** It was previously omitted when unset;
  every `growthbook contextual-bandits` command now includes it. Nothing to migrate — a script
  reading `.seed` gets a string where it could previously have got nothing.

- **Funnel-exploration blocks cap `config.dataset.steps` at 20 entries.** The API enforces the
  limit, so a longer list is rejected server-side; the CLI does not check it before sending.

## [2.3.0] - 2026-08-25

Additive only — no commands or flags were removed or renamed, so no action is needed on upgrade.

### Added

- **`growthbook config-revisions set-property`** — stage a single property of a config draft
  revision's own value, leaving every other property untouched. Prefer it over
  `config-revisions set-value` when you are changing one field: a whole-value write built from a
  stale read silently drops properties someone else added in the meantime.

  `--value null` sets the property to null rather than removing it — use `delete-property` to
  remove one. Pass `--version-param new` to auto-create a draft.

  ```bash
  growthbook config-revisions set-property --key my-config --version-param new \
    --property timeout --value 30
  ```

- **`growthbook config-revisions delete-property`** — stage removal of a single property from a
  config draft revision's own value. The config then inherits that property from its parent, if it
  has one, and every other property is untouched. Also accepts `--version-param new`.

  ```bash
  growthbook config-revisions delete-property --key my-config --version-param new \
    --property timeout
  ```

## [2.2.0] - 2026-08-24

Additive only — no commands or flags were removed or renamed, so no action is needed on upgrade.

### Added

- **`growthbook settings set-approvals`** — replace the organization's approval requirements.
  `PUT /v1/settings/approvals` covers two independent families: `--require-reviews` for feature
  flags, configs, and constants, and `--approval-flows` for saved groups. Each family is replaced
  wholesale when supplied, and omitting one leaves it untouched, so read the current values from
  `growthbook settings get` before sending a partial update.

  ```bash
  growthbook settings set-approvals --body '{
    "requireReviews": [{"requireReviewOn": true, "blockSelfApproval": true}],
    "approvalFlows": {"savedGroups": [{"required": true}]}
  }'
  ```

- **Approval settings gained `blockSelfApproval` and `requiredApproverTeams`.** Both appear on
  every entry of `requireReviews` in `growthbook settings get`, alongside a new top-level
  `approvalFlows.savedGroups` describing the saved-group approval rules.

- **`additionalRoles` on members and team project roles.** Responses from `growthbook members
  list`, `growthbook members update-member-role`, and the `growthbook teams` commands now carry
  `additionalRoles` on the member and inside each `projectRoles[]` entry. It is also accepted in
  the request bodies of `members update-member-role`, `teams create`, and `teams update` — a
  nested body field, so pass it in `--member` / `--project-roles` / `--body` JSON. There are no
  new flags.

## [2.1.1] - 2026-08-21

One flag is gone. It only ever set a server-generated diagnostic, so nothing that worked before
stops working — but a script that passes it will now fail with `unknown flag`, so it is called out
below.

### Removed

- **`growthbook fact-tables update --columns-error`.** `columnsError` is the error message the
  server produces when it cannot parse a fact table's SQL schema. It was writable by accident and
  is now read-only, maintained server-side ([growthbook#6692]). Responses still include it.

  Scripts passing `--columns-error` now fail with `unknown flag` instead of overwriting the
  server's diagnostic. The fix is to delete the flag — there is no replacement, and nothing to
  migrate to.

### Added

- **`-c` shorthand for `--columns`** on `growthbook fact-tables update`.
- Fact table responses gained `columnRefreshPending`, `archived`, `autoSliceUpdatesEnabled`, and
  fact table columns gained `sql`, `topValues`, `topValuesDate`, `dateCreated`, `dateUpdated`. No
  flag changes.

[growthbook#6692]: https://github.com/growthbook/growthbook/pull/6692

## [2.1.0] - 2026-08-19

Additive only — no commands or flags were removed or renamed, so no action is needed on upgrade.
One request field is newly deprecated but still accepted; see "Deprecated" below.

### Added

- **`growthbook custom-hooks test --original-function-args`** — the entity state *before* the
  change. When supplied, the hook also runs against that state and Incremental Changes Only
  suppression is applied, and the response gains a `suppressed` object listing the error and
  warnings the previous state produced too (the ones a real save would hide).

- **`savedGroups` targeting on feature rules and experiment phases.** `growthbook features
  create` / `update` accept `savedGroups` inside `rules[]`, and `growthbook experiments create` /
  `update-experiment` accept it inside `phases[]`. Both are nested body fields, so pass them in
  `--body` JSON — there are no new flags. Each entry is `{"match": "all"|"any"|"none", "ids":
  [...]}`.

  Note that **responses still return the older `savedGroupTargeting` shape**
  (`{"matchType": ..., "savedGroups": [...]}`), so a value you send as `savedGroups` reads back
  under the other name. That asymmetry is deliberate: it keeps a GET response postable back
  unchanged.

- **`create-snapshot` models 409 conflicts.** `growthbook snapshots create-snapshot` and
  `growthbook experiments create-snapshot` now distinguish two Incremental Pipeline conflicts —
  the dimension was already computed from the latest Overall Results (carrying
  `overallResultsAsOf`), and a full refresh of Overall Results is required first — instead of one
  unmodelled error status. Both need `--output-format json` to see the `details` payload.

- Dashboard responses carry additional global-control settings variants. No flag changes.

### Deprecated

- **`savedGroupTargeting` — use `savedGroups` instead.** Deprecated on feature rules
  (`features create` / `update`) and experiment phases (`experiments create` /
  `update-experiment`). It is still accepted, so nothing breaks now, but it will be removed in a
  future release. If both are sent, `savedGroups` wins.

  Migrating is a rename plus a field rename: `{"matchType": "all", "savedGroups": ["grp_x"]}`
  becomes `{"match": "all", "ids": ["grp_x"]}`. Only *request* bodies are affected — responses
  continue to return `savedGroupTargeting`.

## [2.0.1] - 2026-08-17

A narrow fix: 2.0.0 unintentionally renamed the request-body flags, and this restores their
original names. **It does not undo anything else in 2.0.0** — the removals and command renames in
that release still apply, so read the 2.0.0 notes below when upgrading from 1.x.

If you have not already migrated scripts to `--body-param.*`, this release needs no action from
you.

### Fixed

- **`--body-param.` prefixes are gone; request-body flags use their plain names again**
  (`--body-param.description` is `--description` once more). 2.0.0 renamed 142 flags across 15
  commands as an unintended side effect: the API deprecated the `ignoreWarnings` and
  `skipSchemaValidation` *query* parameters in favour of body fields of the same name, and that
  name collision made the generator prefix the entire request body. Dropping the deprecated query
  parameters removes the collision.

  `--ignore-warnings` and `--skip-schema-validation` still work on the same commands — they now
  set the body field, which is where the API wants them.

  If you already migrated to `--body-param.*`, revert to the plain names.

## [2.0.0] - 2026-08-17

Upgrading from 1.x: drop `--merge-now`, and update any renamed `contextual-bandits` /
`contextual-bandit-queries` commands. Both fail loudly — unknown flag or unknown command —
rather than silently misbehaving, so a dry run of your scripts will surface what needs
attention.

> **Upgrading to 2.0.1 or later?** Skip the `--body-param.` rename described under "Changed"
> below — it was an unintended side effect and 2.0.1 reverts it. Keep using the plain flag
> names. The rest of this release still applies.

### Removed

- **`--merge-now` is gone from all publish commands** (`feature-revisions publish`,
  `feature-revisions-v1 post-feature-revision-publish`, `config-revisions publish`,
  `constant-revisions publish`, `saved-group-revisions publish`). The GrowthBook API no longer
  supports it as part of publish hardening. Scripts passing it will fail with an unknown-flag
  error rather than silently not merging.

- **`contextual-bandits start` / `stop` / `refresh` no longer accept a request body** — the
  `--body` and `--body-param` flags are gone. These commands now take only `--id`.

### Changed

- **Request-body flags are now prefixed `--body-param.`** on the commands that take a
  structured body. For example `growthbook configs update --description X` becomes
  `growthbook configs update --body-param.description X`.

  > **Reverted in 2.0.1** — this was an unintended side effect, not a deliberate change.
  > Upgrade to 2.0.1 and keep using the plain flag names.

  Affected commands:

  `features create`, `features update`, `configs create`, `configs update`,
  `configs archive`, `config-revisions publish`, `config-revisions revert`,
  `config-revisions set-metadata`, `config-revisions set-projection`,
  `config-revisions set-schema`, `config-revisions set-value`,
  `feature-revisions add-rule`, `feature-revisions update-rule`,
  `feature-revisions publish`, `feature-revisions set-default-value`.

  Passing raw JSON via `--body` is unaffected, and remains the most stable way to script
  these commands.

- **`contextual-bandits` and `contextual-bandit-queries` commands renamed** to drop the redundant
  group name from every subcommand. These groups were missed by the 1.0.0 naming cleanup; this
  finishes it.

  | Before                                                       | After                                      |
  | ------------------------------------------------------------ | ------------------------------------------ |
  | `contextual-bandits get-contextual-bandit-current-weights`   | `contextual-bandits get-current-weights`   |
  | `contextual-bandits get-contextual-bandit-event`             | `contextual-bandits get-event`             |
  | `contextual-bandits list-contextual-bandit-events`           | `contextual-bandits list-events`           |
  | `contextual-bandits get-contextual-bandit-snapshot`          | `contextual-bandits get-snapshot`          |
  | `contextual-bandits list-contextual-bandit-snapshots`        | `contextual-bandits list-snapshots`        |
  | `contextual-bandits get-contextual-bandit-results`           | `contextual-bandits get-results`           |
  | `contextual-bandits get-contextual-bandit-linked-features`   | `contextual-bandits list-linked-features`  |
  | `contextual-bandits delete-contextual-bandit-linked-feature` | `contextual-bandits delete-linked-feature` |
  | `contextual-bandit-queries create-contextual-bandit-query`   | `contextual-bandit-queries create`         |
  | `contextual-bandit-queries get-contextual-bandit-query`      | `contextual-bandit-queries get`            |
  | `contextual-bandit-queries update-contextual-bandit-query`   | `contextual-bandit-queries update`         |
  | `contextual-bandit-queries delete-contextual-bandit-query`   | `contextual-bandit-queries delete`         |

- **`go install` path now carries the major version.** Go requires a `/vN` module suffix at v2+:

  ```bash
  go install github.com/growthbook/cli/v2/cmd/growthbook@latest
  ```

  This affects `go install` only — npm, Homebrew, the install script, and the prebuilt binaries are
  unchanged.

### Added

- **`learnings`** — new command group: `create`, `get`, `list`, `search`, `update`, `delete`.
- **`releases publish-revisions`** — atomic multi-entity publish.
- **`analytics-explorations run-funnel`** — funnel exploration.
- **`metrics list-experiments`** — experiments for a metric.
- **`fact-tables create-virtual-column` / `update-virtual-column` / `delete-virtual-column`**.
- **`experiments set-schedule`**.
- **`contextual-bandits cancel`**, plus `add-linked-feature` and `update-linked-feature`.
- **Revision workflow parity** — `recall-review`, `reopen`, `schedule-publish` and `undo-review`
  are now available across the `config-revisions`, `constant-revisions` and
  `saved-group-revisions` groups (previously only on `feature-revisions`).
- `configs archive` / `unarchive`, `constants archive` / `unarchive` and `saved-groups archive`
  accept request-body options (e.g. `--ignore-warnings`, `--skip-hooks`).
- `fact-tables` gains `--aggregated-fact-table-settings.restate-chunk-days`;
  `ramp-schedule-templates` gains `--order`.

## [1.0.0] - 2026-07-17

Initial general availability of the rewritten CLI, replacing the legacy (0.x) TypeScript CLI.

- Full coverage of the GrowthBook REST API, with command groups targeting the newest version of
  each endpoint (superseded versions remain available under a `-vN` suffix).
- OS-keychain credential storage, named profiles, and automatic import of a legacy
  `~/.growthbook/config.toml`.
- `--output-format pretty|json|yaml|toon`, `--jq`, `--all` auto-pagination, `--body`/stdin input,
  `--dry-run`, `--debug`, and agent mode.
- Distributed via npm, Homebrew, an install script, `go install`, and prebuilt binaries.

Upgrading from the 0.x CLI is covered in [MIGRATION.md](MIGRATION.md).
