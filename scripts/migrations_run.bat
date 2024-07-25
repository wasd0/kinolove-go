@echo off
setlocal EnableDelayedExpansion

set "envFile=.env"

if not exist "%envFile%" (
    echo Файл .env не найден
    exit /b
)


for /f "tokens=1,* delims==" %%a in ('type "%envFile%"') do (
    set "var=%%a"
    set "val=%%b"
    set "var=!var:~0,-1!"
    set "val=!val:~1!"
    if "!var!"=="%desiredVar%" (
        echo %desiredVar% = !val!
    )
)


set GOOSE=%USERPROFILE%\go\bin\goose.exe
set GOOSE_DRIVER=postgres
set GOOSE_DBSTRING=%DB_DRIVER%://%DB_USR%:%DB_PWD%@%DB_HOST%:%DB_PORT%/%DB_NAME%
set GOOSE_MIGRATION_DIR=%MIGRATIONS_PATH%

set JET=%USERPROFILE%\go\bin\jet.exe

"%GOOSE%" up

"%JET%" -dsn="%GOOSE_DBSTRING%"?sslmode=disable -schema=public -path=.\internal\entity\.gen