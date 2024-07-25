-- +goose Up
-- +goose StatementBegin
insert into roles (name, description)
values ('user', 'Рядовой пользователь'),
       ('system', 'Владелец продукта'),
       ('admin', 'Администратор'),
       ('moderator', 'Модератор'),
       ('blocked', 'Заблокированный на платформе');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete
from roles
where name in ('user', 'admin', 'system', 'moderator', 'blocked')
-- +goose StatementEnd
