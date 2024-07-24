#!/bin/bash

source ".env"

export GOOSE=~/go/bin/goose
export JET=~/go/bin/jet
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="$DB_URL"
export GOOSE_MIGRATION_DIR="$MIGRATIONS_PATH"

"$GOOSE" up
exec "$JET" -dsn="${DB_URL}"?sslmode=disable -schema=public -path=./internal/entity/.gen