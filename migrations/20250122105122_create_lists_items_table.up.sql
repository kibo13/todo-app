CREATE TABLE
    lists_items (
        id serial primary key,
        item_id int references todo_items (id) on delete cascade not null,
        list_id int references todo_lists (id) on delete cascade not null
    );