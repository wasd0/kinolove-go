-- +goose Up
-- +goose StatementBegin
create table if not exists title_studio
(
    title_id bigint references movies(id),
    studio_id bigint references studio(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists title_studio;
-- +goose StatementEnd
