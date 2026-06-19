# growthbook (npm)

npm wrapper for the [GrowthBook CLI](https://github.com/growthbook/cli). Preserves
the familiar global install:

```bash
npm i -g growthbook
growthbook --help
```

## How it works

This package ships no binary itself. On install, `install.js` (postinstall):

1. Resolves the asset for your platform/arch (`growthbook_<Os>_<arch>.<ext>`),
   matching the GoReleaser archive naming.
2. Downloads it from the matching GitHub release (`v<package version>`).
3. Verifies the archive's SHA-256 against the release `checksums.txt`.
4. Extracts the binary into `vendor/` using the system `tar`.

`bin/growthbook.js` then execs that vendored binary, forwarding args and exit code.

The npm `version` must match a published GitHub release tag — release automation
should set them together.

### Environment

- `GROWTHBOOK_CLI_SKIP_DOWNLOAD=1` — skip the postinstall download (air-gapped/CI
  setups that vendor the binary another way).

### Notes / future hardening

Prebuilt binaries are provided for linux, darwin, and windows on x64/arm64. Other
platforms should `go install github.com/growthbook/cli/cmd/growthbook@latest`.

A future option is the per-platform `optionalDependencies` pattern (à la esbuild):
publish `@growthbook/cli-<os>-<arch>` packages carrying the raw binary so npm
selects the right one with no postinstall download (works with `--ignore-scripts`).
The postinstall-download approach here is simpler to ship and is the first cut.
