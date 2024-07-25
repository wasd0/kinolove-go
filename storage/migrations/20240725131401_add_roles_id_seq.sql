-- +goose Up
-- +goose StatementBegin
create sequence if not exists roles_id_seq start 1;

alter sequence roles_id_seq owned by roles.id;

alter table roles
    alter column id set default nextval('roles_id_seq');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop sequence if exists roles_id_seq;

alter table roles
    alter column id type bigint using id::bigint;
-- +goose StatementEnd
