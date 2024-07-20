@echo off

set GOOSE=%USERPROFILE%\go\bin\goose.exe
set JET=%USERPROFILE%\go\bin\jet.exe
set GOOSE_DRIVER=postgres
set GOOSE_DBSTRING=%DB_URL%
set GOOSE_MIGRATION_DIR=%MIGRATIONS_PATH%

"%GOOSE%" up
"%JET%" -dsn="%DB_URL%"?sslmode=disable -schema=public -path=.\internal\entity\.gen