CREATE TABLE users(
    id serial primary key,
    login varchar(30) not null unique check(login != ''),
    full_name varchar(255),
    hash_pwd varchar(255) not null check(hash_pwd != '')
);

CREATE TABLE info (
    id serial primary key,
    owner integer not null references users (id) on delete cascade,
    login bytea not null,
    pwd bytea not null,
    project varchar(255),
    source varchar(255) not null,
    type_source varchar(255) not null,
    type_access varchar(255) not null default 'basic' check (type_access='basic' or type_access='secondary'),
    comment text
);

CREATE TABLE groups (
    id serial primary key,
    name varchar(255) not null,
    comment text
);

CREATE TABLE access (
    id serial primary key,
    owner integer references users (id) on delete cascade,
    owner_group integer references groups (id) on delete cascade,
    final timestamp,
    count_request integer check (count_request >= 0),
    type_access varchar(255) not null default 'none' check (type_access='show' or type_access='edit' or type_access='full'),
    check (owner > 0 or owner_group > 0)
);

CREATE TABLE groups_to_users (
    id serial primary key,
    user_id integer not null references users (id) on delete cascade,
    group_id integer not null references groups (id) on delete cascade
);