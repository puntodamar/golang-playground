
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table publishers
(
    id int null auto_increment,
    name varchar(100) null,
    created_at datetime null,
    updated_at datetime null,

    constraint publishers_pk primary key (id)
);

create index publishers_name_index
    on publishers (name);



-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS publishers
