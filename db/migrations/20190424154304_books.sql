
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table books
(
    id            int auto_increment,
    publisher_id  int null,
    name          varchar(200) CHARACTER SET utf8 not null,
    published_at  date null ,
    created_at    datetime null,
    updated_at    datetime null,
    deleted_at    datetime null,

    constraint books_pk primary key (id),

    constraint books_publishers__fk foreign key (publisher_id)
        references publishers (id)
            on delete set null
);

create index books_name_index
    on books (name);

create index books_published_at_index
    on books (published_at);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS books