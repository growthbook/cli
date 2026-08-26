# Changelog

## [Unreleased]

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
