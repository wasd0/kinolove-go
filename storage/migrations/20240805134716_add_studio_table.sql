-- +goose Up
-- +goose StatementBegin
create table if not exists studio
(
    id bigserial primary key,
    name varchar not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists studio;
-- +goose StatementEnd
