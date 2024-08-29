-- +goose Up
-- +goose StatementBegin
create table if not exists genres_titles
(
    genre_id bigint references genres(id),
    title_id bigint references movies(id),
    primary key (genre_id, title_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists genres_titles;
-- +goose StatementEnd
