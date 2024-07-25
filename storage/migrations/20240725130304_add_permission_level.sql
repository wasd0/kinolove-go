-- +goose Up
-- +goose StatementBegin
create table if not exists permission_levels (
    id smallint primary key,
    name varchar(50) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists permission_levels;
-- +goose StatementEnd
