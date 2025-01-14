create table todos(
    id serial primary key,
    name varchar(50) not null,
    description varchar(50),
    checked boolean default false not null,
    deliver_in date,
    createt_at timestamp default now() not null 
);

create table users(
    id serial primary key,
    name varchar(50) not null,
    email varchar(60) not null unique,
    password varchar(70) not null
);

create table user_todos (
    user_id int references users(id),
    todo_id int references todos(id),
    constraint user_todos_pk primary key (user_id, todo_id)
);

INSERT INTO users(name, email, password) VALUES
('Luan', 'lluantenorio7@gmail.com', '$2a$12$7LVHCrakFO6wFtM23uvxQOj6pVsEm7bPPAwL2kn2rFuvgvgVDFSzG'),
('Cleiton', 'cleiton@gmail.com', '$2a$12$DRjW1RnB0z4Rt5lF2oJPFeAkgxrw9mZ5YuctvMYyLGE5/YcLVUvsC');

INSERT INTO todos(name, deliver_in) VALUES
('create todolist', '2025-01-09'),
('finish reading clean architecture', '2025-01-12'),
('Carry out analysis modeling', '2025-01-15');

INSERT INTO user_todos(user_id, todo_id) VALUES 
(1, 1), (2, 2), (1, 3), (2, 3);
