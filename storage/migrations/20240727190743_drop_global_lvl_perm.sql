-- +goose Up
-- +goose StatementBegin
alter table roles_permissions
    drop column if exists target_level cascade;
alter table users_permissions
    drop column if exists target_level cascade;
alter table roles_permissions
    rename column global_level to level;
alter table users_permissions
    rename column global_level to level;
alter table permissions
    drop column if exists default_level_id cascade;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table roles_permissions
    add column if not exists target_level smallint;
alter table users_permissions
    add column if not exists target_level smallint;
alter table roles_permissions
    rename column level to global_level;
alter table users_permissions
    rename column level to global_level;
alter table permissions
    add column if not exists default_level_id smallint references permission_levels (id);
-- +goose StatementEnd
