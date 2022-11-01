CREATE TABLE users
(
    id serial not null unique primary key,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);