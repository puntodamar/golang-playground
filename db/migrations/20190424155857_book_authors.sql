
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- auto-generated definition
create table book_authors
(
    book_id   int null,
    author_id int auto_increment,

    constraint book_authors_authors__fk
        foreign key (author_id) references authors (id),

    constraint book_authors_books__fk
        foreign key (book_id) references books (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS  book_authors
