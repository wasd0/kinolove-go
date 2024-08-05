-- +goose Up
-- +goose StatementBegin
create table if not exists studios
(
    id bigserial primary key,
    name varchar not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists studios;
-- +goose StatementEnd
