-- +goose Up
-- +goose StatementBegin
create extension if not exists "uuid-ossp";
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop extension if exists "uuid-ossp";
-- +goose StatementEnd
