#!/usr/bin/env bash
# Regenerate the CLI against an UNMERGED GrowthBook OpenAPI spec, without
# touching the committed .speakeasy/workflow.yaml (whose source must always
# point at `main` for CI). Use it to verify a spec PR generates cleanly before
# it lands.
#
#   ./scripts/gen-from-branch.sh                       # defaults to main
#   ./scripts/gen-from-branch.sh kc/fix-invalid-openapi-spec   # a branch/tag/SHA
#   ./scripts/gen-from-branch.sh ../growthbook/packages/back-end/generated/spec.yaml  # a local file
#   ./scripts/gen-from-branch.sh kc/my-branch --skip-compile   # extra flags pass through to `speakeasy run`
#
# The committed workflow.yaml is restored on exit (even on failure/Ctrl-C).
# Generated output is left in place so you can inspect it; `git checkout` to revert.
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
WORKFLOW="$ROOT/.speakeasy/workflow.yaml"

ARG="${1:-main}"
if [ $# -gt 0 ]; then shift; fi

if [ -f "$ARG" ]; then
  SPEC_LOC="$ARG"
  echo "Spec source: local file $SPEC_LOC"
else
  SPEC_LOC="https://raw.githubusercontent.com/growthbook/growthbook/${ARG}/packages/back-end/generated/spec.yaml"
  echo "Spec source: ref '$ARG' -> $SPEC_LOC"
fi

# Restore the committed workflow no matter how we exit.
BACKUP="$(mktemp)"
cp "$WORKFLOW" "$BACKUP"
trap 'cp "$BACKUP" "$WORKFLOW"; rm -f "$BACKUP"; echo "Restored committed .speakeasy/workflow.yaml"' EXIT

# Swap only the source input line (ends in spec.yaml); leaves the overlay line
# (overlay.yaml) untouched. Rebuild the whole value so refs containing slashes
# (e.g. kc/foo) work.
tmp="$(mktemp)"
sed -E "s#^( *- location: ).*spec\.yaml.*#\1${SPEC_LOC}#" "$WORKFLOW" > "$tmp"
mv "$tmp" "$WORKFLOW"

# Skip the registry upload: a pre-merge branch spec shouldn't be published to
# the shared Speakeasy registry.
echo "Running: speakeasy run --skip-upload-spec $*"
speakeasy run --skip-upload-spec "$@"
