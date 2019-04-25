
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table users
(
    id         int auto_increment primary key,
    name       varchar(150) not null,
    username   varchar(20)  not null,
    email      varchar(100) not null,
    updated_at datetime     null,
    deleted_at datetime     null,
    created_at datetime     null,

    constraint users_email_uindex
        unique (email),

    constraint users_username_uindex
        unique (username)
);

create index users_email_index
    on users (email);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS users
