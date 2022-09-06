CREATE TABLE users(
    id serial primary key,
    name varchar(255) not null,
    login varchar(255) not null,
    age integer not null
);

insert into users (name,login,age) values ('ivan','livan',17),('katarsis','ekt',152);