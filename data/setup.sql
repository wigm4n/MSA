DROP TABLE users CASCADE ;

create table users
(
  id        serial not null
    constraint users_pkey
    primary key,
  email     varchar(64),
  firstname varchar(255),
  lastname  varchar(255),
  password  text
);

alter table users
  owner to ilobanov;

create unique index users_email_uindex
  on users (email);