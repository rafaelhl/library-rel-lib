create schema if not exists library collate utf8mb4_0900_ai_ci;

create table if not exists library.shelf
(
    id       int           not null
        primary key,
    capacity int           not null,
    amount   int default 0 not null
);

create table if not exists library.book
(
    id          int auto_increment
        primary key,
    title       varchar(60)  not null,
    description varchar(255) null,
    author      varchar(60)  not null,
    edition     int          not null,
    shelf_id    int          not null,
    created_at  datetime     not null,
    updated_at  datetime     not null,
    constraint book_shelf_id_fk
        foreign key (shelf_id) references library.shelf (id)
);

insert into library.shelf (id, capacity) VALUE (1, 5);
insert into library.shelf (id, capacity) VALUE (2, 3);