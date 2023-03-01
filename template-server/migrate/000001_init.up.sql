CREATE TABLE templates
(
    id serial not null unique,
    owner_id int not null,
    name varchar(255) not null,
    filename varchar(255) not null unique
);