-- +goose Up
-- +goose StatementBegin
create table if not exists movies
(
    id               bigserial primary key,
    title            varchar(512) not null,
    episode_duration int,
    episode_count    smallint,
    alter_titles     text,
    description      text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists movies;
-- +goose StatementEnd
