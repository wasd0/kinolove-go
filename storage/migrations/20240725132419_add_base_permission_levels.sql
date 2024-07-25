-- +goose Up
-- +goose StatementBegin
insert into permission_levels (id, name)
values (0, 'denied'),
       (1, 'read'),
       (2, 'create'),
       (4, 'edit'),
       (8, 'delete');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete
from permission_levels
where id in (0, 1, 2, 4, 8)
-- +goose StatementEnd
