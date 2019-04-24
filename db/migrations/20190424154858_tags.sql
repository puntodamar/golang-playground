
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table tags
(
    id      int auto_increment,
    name    varchar(20) not null,

    constraint tags_pk primary key (id)
);

create unique index tags_name_uindex
    on tags (name);



-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS  tags
