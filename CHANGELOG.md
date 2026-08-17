# Changelog

## [Unreleased]

## [2.0.0] - 2026-08-17

Upgrading from 1.x requires two changes: drop `--merge-now`, and update any renamed
`contextual-bandits` / `contextual-bandit-queries` commands. Both fail loudly (unknown
flag / unknown command) rather than silently misbehaving, so a dry run of your scripts
will surface everything that needs attention.

### Removed

- **`--merge-now` is gone from all publish commands** (`feature-revisions publish`,
  `feature-revisions-v1 post-feature-revision-publish`, `config-revisions publish`,
  `constant-revisions publish`, `saved-group-revisions publish`). The GrowthBook API no longer
  supports it as part of publish hardening. Scripts passing it will fail with an unknown-flag
  error rather than silently not merging.

### Changed

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
