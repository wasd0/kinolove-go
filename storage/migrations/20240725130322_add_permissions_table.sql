-- +goose Up
-- +goose StatementBegin
create table if not exists permissions
(
    id bigserial primary key,
    name varchar(50) not null,
    description varchar(100),
    default_level_id smallint references permission_levels(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists permissions;
-- +goose StatementEnd
