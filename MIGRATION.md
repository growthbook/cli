# Migrating from the legacy GrowthBook CLI

The GrowthBook CLI was rewritten on a new code generator. Most workflows are unchanged; this
guide covers what moved.

## Install — upgrade intentionally

This is a new generation with breaking changes, so upgrade on purpose:

- `npm i -g growthbook` installs the latest published version; `npm update -g growthbook`
  moves an existing install onto it.
- Not ready yet? Pin the legacy line: `npm i -g growthbook@0.2`.
- Also available via Homebrew, the install script, or `go install` (see the README).

## Auth & config — migrates automatically

On first run the CLI detects a legacy `~/.growthbook/config.toml`, imports every profile into
the OS keychain + new config, and prints what it did. Nothing to do by hand.

| Legacy                      | New                                                                         |
| --------------------------- | --------------------------------------------------------------------------- |
| `~/.growthbook/config.toml` | `~/.config/growthbook/config.yaml` + `cli.yaml`; secrets in the OS keychain |
| `auth login` / `logout`     | `auth login` / `logout` (plus `auth whoami`)                                |
| `-p, --profile <name>`      | `--profile <name>` / `GBCLI_PROFILE`                                        |
| `-u, --apiBaseUrl <url>`    | `--server-url <url>` (or `--server` / `--domain`)                           |
| key per profile             | `--bearer-auth` / `GBCLI_BEARER_AUTH` / keychain / active profile           |

Profiles are managed with `growthbook profiles list|use|set|remove`. Credential precedence is
**flag > env > active profile** (the server URL has no env tier, so it's flag > profile).

## Command map

Most commands are unchanged. The feature commands now target the **latest (v2)** API by default;
the previous v1 surface is still available under `features-v1` (deprecated).

| Legacy                                         | New                                                               |
| ---------------------------------------------- | ----------------------------------------------------------------- |
| `features list` / `get` / `toggle`             | **`features list` / `get` / `toggle`** (now the latest, v2, API)  |
| `features generate-types`                      | `generate-types` (top-level)                                      |
| (old v1 feature surface, if you still need it) | `features-v1 …` (deprecated; prints a one-line warning to stderr) |
| `experiments results`                          | `experiments results` (one) / `experiments list-results` (many)   |
| `datasources …`                                | `data-sources …`                                                  |
| `sdkconnections …`                             | `SDK-connections …`                                               |
| `savedgroups …`                                | `saved-groups …`                                                  |
| `vcs list` / `get`                             | `visual-changesets list` / `get`                                  |

> The new CLI exposes the **full** GrowthBook REST API — run `growthbook --help` to explore the
> complete command tree.

### `features-v1` deprecation

`features-v1 …` reaches the legacy v1 feature endpoints and prints a deprecation notice to
**stderr** (never stdout, so piped JSON stays clean). Prefer the plain `features` group, which
targets the v2 API. When a future API version lands, `features` advances to it and the prior
version becomes the next `features-vN`. Re-pointing the base command that way is a **breaking
change**: it ships as a **major version** with the change called out in the changelog, so you can
pin the prior major (or switch to the explicit `features-vN` command) on your own schedule.

## `generate-types`

```
growthbook generate-types --project prj_123
```

Flags: `--output` (default `./growthbook-types`), `-f/--filename` (default `app-features.ts`),
`-p/--project`.

## New conveniences

- `--output-format pretty|json|yaml|toon`
- `--jq '<filter>'` to post-process JSON output
- `--all` to auto-paginate list commands (streams NDJSON for JSON output)
- `--body` / stdin (including `--body -`) for request payloads
- `--dry-run` to print the request without sending it
- `--agent-mode` structured errors + TOON output (auto-enabled under AI coding agents)
- `--debug` for verbose diagnostics
