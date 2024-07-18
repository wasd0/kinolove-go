-- +goose Up
-- +goose StatementBegin
create table if not exists users
(
    id            uuid primary key default uuid_generate_v4(),
    username      varchar(50)  not null,
    password      varchar(510) not null,
    is_active     boolean          default true,
    date_reg      timestamptz      default now(),
    date_pass_upd timestamptz,
    constraint password_length_check check ( length(password) > 8 ),
    constraint username_length_check check ( length(username) > 5 )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
