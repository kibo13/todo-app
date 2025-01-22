CREATE TABLE
    todo_lists (
        id serial primary key,
        title varchar(255) not null,
        description varchar(255)
    );