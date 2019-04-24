
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- auto-generated definition
create table book_tags
(
    book_id int null,
    tag_id  int null,

    constraint book_tags_books__fk
        foreign key (book_id) references books (id),

    constraint book_tags_tags__fk
        foreign key (tag_id) references tags (id)
);



-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS book_tags
