-- +goose Up
-- +goose StatementBegin
create table if not exists users_permissions
(
    user_id       uuid references users (id),
    permission_id bigint references permissions (id),
    target_level  smallint not null references permission_levels (id),
    global_level  smallint not null references permission_levels (id),
    cause         varchar,
    date_expire   timestamptz,
    primary key (user_id, permission_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users_permissions;
-- +goose StatementEnd
