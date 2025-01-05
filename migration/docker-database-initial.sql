create table todo(
    id serial primary key,
    name varchar(50) not null,
    checked boolean default false not null,
    deliver_in date not null,
    createt_at timestamp default now() not null 
);


INSERT INTO todo(name, deliver_in) VALUES
('create todolist', '2025-01-09'),
('finish reading clean architecture', '2025-01-12');
