with system_id as (select id as id
                   from users
                   where users.username = 'system'
                   limit 1),
     role_id as (select id as id
                 from roles
                 where roles.name = 'system'
                 limit 1)

insert
into users_roles (user_id, role_id, date_active)
values ((select id from system_id), (select id from role_id), now());