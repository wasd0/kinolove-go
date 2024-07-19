-- +goose Up
-- +goose StatementBegin
alter table users
    alter column password type bytea using password::bytea;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table users
    alter column password type varchar(512) using password::varchar(512);

-- +goose StatementEnd
