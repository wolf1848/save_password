CREATE TABLE users(
    id serial primary key,
    login varchar(30) not null unique,
    full_name varchar(255) null,
    hash_pwd varchar(255) not null
);

CREATE TABLE info (
    id serial primary key,
    hash_login varchar(255) not null,
    hash_pwd varchar(255) not null,
);


//login
//password
//project Проект где используется
//source Источник доступа
//type_source Тип ресурса
//type_access тип доступа (основной, вторичный)
// comment комментарий
