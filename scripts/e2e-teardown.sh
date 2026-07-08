#!/usr/bin/env bash
# Tear down the live CLI e2e stack started by scripts/e2e-setup.sh, removing
# volumes so the next setup starts from a clean, unprovisioned instance.
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
COMPOSE="$ROOT/test/live/docker-compose.e2e.yml"

echo "==> Tearing down the GrowthBook stack (removing volumes)"
docker compose -f "$COMPOSE" down -v

# Drop the generated creds so a stale key can't linger past the instance.
rm -f "$ROOT/.env.local"
echo "==> Done."
