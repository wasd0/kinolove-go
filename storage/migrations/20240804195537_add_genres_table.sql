-- +goose Up
-- +goose StatementBegin
create table if not exists genres
(
    id      bigserial primary key,
    name   varchar(255) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists genres;
-- +goose StatementEnd
