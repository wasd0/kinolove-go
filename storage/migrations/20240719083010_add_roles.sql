-- +goose Up
-- +goose StatementBegin
create table if not exists roles
(
    id bigint primary key,
    name varchar(20),
    description varchar(50)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists roles;
-- +goose StatementEnd
