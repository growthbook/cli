# Changelog

## [Unreleased]

Additive only — no commands or flags were removed or renamed, so no action is needed on upgrade.

### Added

- **`growthbook custom-hooks test --original-function-args`** — the entity state *before* the
  change. When supplied, the hook also runs against that state and Incremental Changes Only
  suppression is applied, and the response gains a `suppressed` object listing the error and
  warnings the previous state produced too (the ones a real save would hide).

- **`savedGroups` targeting on feature rules and experiment phases.** `growthbook features
  create` / `update` accept `savedGroups` inside `rules[]`, and `growthbook experiments create` /
  `update-experiment` accept it inside `phases[]`. Both are nested body fields, so pass them in
  `--body` JSON — there are no new flags.

- **`create-snapshot` now reports 409 conflicts.** `growthbook snapshots create-snapshot` and
  `growthbook experiments create-snapshot` surface a typed conflict when the dimension is already
  up to date (with `overallResultsAsOf`) or when a full refresh is required, rather than an
  unmodelled error status.

- Dashboard responses carry additional global-control settings variants. No flag changes.

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
