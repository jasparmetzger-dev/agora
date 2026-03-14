#!/usr/bin/env bash
set -euo pipefail

# ─── Config ────────────────────────────────────────────────────────────────────
DB_URL="${DATABASE_URL:-postgresql://postgres:2319@localhost:5432/agora?sslmode=disable}"
MIGRATIONS_DIR="${MIGRATIONS_DIR:-./cmd/database/migrations}"
SQLC_CONFIG="${SQLC_CONFIG:-./sqlc.yaml}"

# ─── Helpers ───────────────────────────────────────────────────────────────────
log()  { echo "[INFO]  $*"; }
err()  { echo "[ERROR] $*" >&2; }
die()  { err "$*"; exit 1; }

require() {
  command -v "$1" &>/dev/null || die "'$1' not found in PATH. Please install it."
}

# ─── Dependency checks ─────────────────────────────────────────────────────────
require migrate
require sqlc

# ─── Get current migration version before applying ────────────────────────────
get_version() {
  migrate -database "$DB_URL" -path "$MIGRATIONS_DIR" version 2>/dev/null \
    | grep -oE '[0-9]+' | head -1 || echo "0"
}

# ─── Step 1: Run migrations ────────────────────────────────────────────────────
run_migrations() {
  log "Reading current DB version..."
  local version_before
  version_before=$(get_version)
  log "Version before migration: ${version_before:-none}"

  log "Applying migrations from '$MIGRATIONS_DIR'..."
  if ! migrate -database "$DB_URL" -path "$MIGRATIONS_DIR" up; then
    err "Migration failed. Rolling back to version $version_before..."

    if [[ "$version_before" == "0" ]]; then
      # No previous version — drop everything
      migrate -database "$DB_URL" -path "$MIGRATIONS_DIR" drop -f \
        && log "Rolled back: database dropped (was empty before)." \
        || die "Rollback (drop) also failed. Manual intervention required."
    else
      migrate -database "$DB_URL" -path "$MIGRATIONS_DIR" goto "$version_before" \
        && log "Rolled back to version $version_before." \
        || die "Rollback to version $version_before failed. Manual intervention required."
    fi

    die "Aborted after rollback."
  fi

  local version_after
  version_after=$(get_version)
  log "Migration successful. Version now: $version_after"
}

# ─── Step 2: Generate Go code with sqlc ───────────────────────────────────────
run_sqlc() {
  log "Running sqlc generate (config: $SQLC_CONFIG)..."
  if ! sqlc generate --file "$SQLC_CONFIG"; then
    die "sqlc generate failed. Check your SQL queries and sqlc.yaml."
  fi
  log "sqlc generate completed successfully."
}

# ─── Main ──────────────────────────────────────────────────────────────────────
main() {
  log "=== DB Migration & Code Generation ==="

  [[ -d "$MIGRATIONS_DIR" ]] || die "Migrations directory '$MIGRATIONS_DIR' does not exist."
  [[ -f "$SQLC_CONFIG" ]]    || die "sqlc config '$SQLC_CONFIG' does not exist."

  run_migrations
  run_sqlc

  log "=== Done ==="
}

main "$@"