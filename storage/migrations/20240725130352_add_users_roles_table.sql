-- +goose Up
-- +goose StatementBegin
create table if not exists users_roles
(
    user_id     uuid references users (id),
    role_id     bigint references roles (id),
    date_active timestamptz not null,
    date_expire timestamptz,
    primary key (user_id, role_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users_roles;
-- +goose StatementEnd
