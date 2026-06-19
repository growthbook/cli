#!/usr/bin/env bash
# Smoke-test the npm shim end to end: pack npm/, install it into a throwaway
# project, and confirm the postinstall downloads + checksum-verifies the release
# binary and the launcher actually runs it.
#
# Requires a GitHub Release to already exist for the version (the shim's
# postinstall pulls assets from releases/download/v<version>). Run it after a
# release tag has built, as a final check before trusting an npm publish.
#
#   ./scripts/smoke-npm.sh 1.0.0-next.0
set -euo pipefail

VERSION="${1:-}"
if [ -z "$VERSION" ]; then
  echo "usage: $0 <released-version>   e.g. $0 1.0.0-next.0" >&2
  exit 2
fi
VERSION="${VERSION#v}"

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
NPM_DIR="$ROOT/npm"
ASSET_URL="https://github.com/growthbook/cli/releases/download/v${VERSION}/checksums.txt"

# Friendly preflight: the postinstall will 404 if the release isn't published.
if command -v curl >/dev/null 2>&1; then
  if ! curl -sfL -o /dev/null "$ASSET_URL"; then
    echo "No published release assets found at v${VERSION}." >&2
    echo "Tag + push v${VERSION} (so release.yaml builds the GitHub Release) before smoke-testing." >&2
    exit 1
  fi
fi

WORK="$(mktemp -d)"
PKG_BACKUP="$(mktemp)"
cp "$NPM_DIR/package.json" "$PKG_BACKUP"
cleanup() {
  cp "$PKG_BACKUP" "$NPM_DIR/package.json"  # restore committed version
  rm -f "$PKG_BACKUP" "$NPM_DIR"/growthbook-*.tgz
  rm -rf "$WORK"
}
trap cleanup EXIT

echo "==> Packing npm shim at version $VERSION"
( cd "$NPM_DIR" && npm version "$VERSION" --no-git-tag-version --allow-same-version >/dev/null )
TARBALL="$(cd "$NPM_DIR" && npm pack --silent)"

echo "==> Installing the tarball into a throwaway project"
( cd "$WORK" && npm init -y >/dev/null && npm install "$NPM_DIR/$TARBALL" )

echo "==> Running the installed binary"
BIN="$WORK/node_modules/.bin/growthbook"
"$BIN" version || true
"$BIN" --help >/dev/null

echo "==> OK: growthbook@$VERSION packs, installs (downloaded + checksum-verified), and runs"
