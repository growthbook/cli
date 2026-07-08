# Live CLI e2e suite

End-to-end tests that drive the **compiled `growthbook` binary** as a subprocess
against a **real GrowthBook server**, asserting on parsed output and exit codes.
Unlike the offline tests in [`internal/cli/custom_e2e_test.go`](../../internal/cli/custom_e2e_test.go)
(which run in-process against a mock), this suite validates the actual shipped
artifact against a real API — flag→body serialization, path params, auth, output
formatting, and error surfacing — across the command surface.

It exists for the **Stainless→Speakeasy cutover**: higher-confidence, one-command
validation so we don't rely on manual testing before promoting the CLI. It is
**gated by the `live` build tag**, so the default offline CI (`go test ./...`)
never runs it.

## Run it locally

```bash
./scripts/e2e-setup.sh                       # docker up + bootstrap creds into .env.local
go build -o growthbook ./cmd/growthbook       # build the binary the suite drives
set -a && source .env.local && set +a         # export GB_SERVER_URL + GBCLI_BEARER_AUTH
go test -tags live ./test/live/ -v            # run the suite (writes coverage.txt)
./scripts/e2e-teardown.sh                     # docker down -v (removes creds + volumes)
```

`scripts/e2e-setup.sh` brings up [`docker-compose.e2e.yml`](docker-compose.e2e.yml)
(mongo + the published `growthbook/growthbook` image), waits for the API, registers
the first user (creating the org + a default project), mints an admin secret key,
and creates the built-in demo datasource. It's idempotent. `TestMain` auto-loads
`.env.local`, so the `source` step is only needed if you invoke `go test` in a
shell that didn't already export those vars.

## The coverage report

After each run, `test/live/coverage.txt` (also printed to stdout) diffs the
commands the suite exercised against the generated command-surface golden
([`internal/cli/testdata/command-surface.txt`](../../internal/cli/testdata/command-surface.txt)).
It lists exactly what was and wasn't hit — the reviewer-facing confidence
artifact. "Tests pass" doesn't imply full coverage; the report makes the gaps
explicit.

## What's covered vs. not

Covered: full create→list→get→update→delete lifecycles for the core createable
resources (projects, environments, attributes, saved-groups, SDK-connections,
features, metrics, dimensions, segments, fact-tables, experiments), read-only
`list` smoke across the remaining ungated groups, the output-format matrix
(pretty/json/yaml/table/toon), and error paths (4xx → non-zero exit).

Deliberately **not** covered (shown as uncovered in the report):

- **Stateful workflows** — feature-revisions, ramp-schedules, saved-group /
  constant revisions, experiment start/stop. They need multi-step setup and are
  out of scope for the core suite.
- **Premium/enterprise-gated** resources (e.g. archetypes) — a free self-hosted
  test instance returns 403; those tests `t.Skip` automatically.
- **Not-in-this-API-version / multi-org-only** endpoints (visual-changesets,
  `organizations list`).
- The deprecated **`-v1`** command mirrors.

## Notes / constraints

- **Rate limit**: the compose file sets `API_RATE_LIMIT_MAX` high so a full run
  doesn't trip the default 60-req/min-per-key limit.
- **Datasource**: the API can't create data sources, so Phase 3/4 rely on the
  demo datasource created during setup; without it they skip.
- **Delete-after-archive**: features and saved-groups can't be REST-deleted while
  live, so the suite archives them first (also exercising the archive path).

## CI

[`.github/workflows/e2e-live.yaml`](../../.github/workflows/e2e-live.yaml) runs
this on-demand (`workflow_dispatch`): it builds, sets up, runs the suite, and
uploads `coverage.txt` as an artifact — so reviewers can trigger a green run
without local docker.
