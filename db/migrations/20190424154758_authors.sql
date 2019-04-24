
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table authors
(
    id      int auto_increment,
    name    varchar(150) CHARACTER SET utf8 null,

    constraint authors_pk primary key (id)
);



-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE if exists authors
