#!/bin/bash

source ".env"

export GOOSE=~/go/bin/goose
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="$DB_URL"
export GOOSE_MIGRATION_DIR="$MIGRATIONS_PATH"

exec "$GOOSE" up