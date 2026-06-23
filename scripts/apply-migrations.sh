#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

: "${POSTGRES_HOST:?POSTGRES_HOST is required}"
: "${POSTGRES_PORT:=5432}"
: "${POSTGRES_USER:?POSTGRES_USER is required}"
: "${POSTGRES_PASS:?POSTGRES_PASS is required}"
: "${POSTGRES_DB:?POSTGRES_DB is required}"

export PGPASSWORD="$POSTGRES_PASS"

for migration in migrations/20*.sql; do
  [ -e "$migration" ] || continue
  echo "Applying $migration"
  psql \
    -v ON_ERROR_STOP=1 \
    -h "$POSTGRES_HOST" \
    -p "$POSTGRES_PORT" \
    -U "$POSTGRES_USER" \
    -d "$POSTGRES_DB" \
    -f "$migration"
done
