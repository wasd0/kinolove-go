-- +goose Up
-- +goose StatementBegin
insert into permissions (name, description, default_level_id)
values ('user', 'Взаимодействие с пользователем', 1),
       ('role', 'Взаимодействие с ролями', 0),
       ('permission', 'Взаимодействие с разрешениями', 0),
       ('movie', 'Взаимодействие с контентом', 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete
from permissions
where name in ('user', 'role', 'permission', 'movie');
-- +goose StatementEnd
