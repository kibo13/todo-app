CREATE TABLE
    todo_items (
        id serial primary key,
        title varchar(255) not null,
        done boolean not null default false
    );