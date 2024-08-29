-- +goose Up
-- +goose StatementBegin
alter table movies
    add image_url text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table movies
    drop column image_url;
-- +goose StatementEnd
