CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE relations
(
    id serial not null unique,
    first_user int references users(id) on delete cascade not null,
    second_user int references users(id) on delete cascade not null,
    relations_type int not null
);