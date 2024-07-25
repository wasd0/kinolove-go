-- +goose Up
-- +goose StatementBegin
create table if not exists roles_permissions
(
    role_id       bigint references roles (id),
    permission_id bigint references permissions (id),
    target_level  smallint references permission_levels (id) not null,
    global_level  smallint references permission_levels (id) not null,
    primary key (role_id, permission_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists roles_permissions;
-- +goose StatementEnd
