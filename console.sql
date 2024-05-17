create table todo_items(
    id int primary key ,
    Name varchar(50),
    create_at timestamp,
    update_at timestamp
);

INSERT INTO todo_items (id, Name, create_at, update_at) VALUES
 (1, 'Sản phẩm 1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
 (2, 'Sản phẩm 2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'Sản phẩm 3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 'Sản phẩm 4', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 'Sản phẩm 5', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
