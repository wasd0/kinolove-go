-- +goose Up
-- +goose StatementBegin
create table if not exists titles_studios
(
    title_id bigint references movies(id),
    studio_id bigint references studios(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists titles_studios;
-- +goose StatementEnd
