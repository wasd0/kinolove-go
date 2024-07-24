-- +goose Up
-- +goose StatementBegin
insert into users (username, password, is_active)
values ('system', '$2a$10$1Y30U4hEoy6Ppz0zm0r8B.vxGHCP.limDMk5FWIdg//PXFb0aWPO.', true)
on conflict do nothing;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from users where id = '$2a$10$1Y30U4hEoy6Ppz0zm0r8B.vxGHCP.limDMk5FWIdg//PXFb0aWPO.'
-- +goose StatementEnd
