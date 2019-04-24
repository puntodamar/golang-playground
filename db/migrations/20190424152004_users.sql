
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table users
(
    id            int auto_increment,
    name          varchar(150) not null,
    username      varchar(20) not null,
    created_at    datetime null,
    updated_at    datetime null,
    deleted_at    datetime null,

    constraint users_pk primary key (id)
);

create unique index users_username_uindex
    on users (username);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS users
