CREATE TABLE
    users_lists (
        id serial primary key,
        user_id int references users (id) on delete cascade not null,
        list_id int references todo_lists (id) on delete cascade not null
    );