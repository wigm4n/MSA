drop table users cascade ;
drop table tasks cascade ;
drop table messages cascade ;
drop table sessions cascade ;

create table users (
  id          serial not null primary key,
  email       varchar(64),
  firstname   varchar(255),
  lastname    varchar(255),
  password    text
);

create table tasks (
  id                serial not null primary key ,
  creator_user_id   integer references users(id),
  name              varchar(255),
  task_type         int,
  date              timestamp
);

create table messages (
  id                serial not null primary key,
  task_id           integer references tasks(id),
  user_id           varchar(255),
  text              text,
  date              timestamp
);

create table sessions (
  id                serial not null primary key,
  user_id           integer references users(id),
  token             varchar(255)
);

insert into users (email, firstname, lastname, password) values ('chyps97@gmail.com', 'Илья', 'Лобанов', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220');


insert into users (email, firstname, lastname, password) values ('test3@hse.ru', 'prep', 'prepov', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220');
insert into users (email, firstname, lastname, password) values ('hse2@mail.ru', 'Кирилл', 'Фурманов', '1234');

insert into tasks (creator_user_id, name, task_type, date) values (1, 'dz 1', 2, TIMESTAMP '2019-02-16 15:36:38');
insert into tasks (creator_user_id, name, task_type, date) values (1, 'dz 2', 4, TIMESTAMP '2019-02-16 15:36:38');
insert into tasks (creator_user_id, name, task_type, date) values (2, 'dz 3', 5, TIMESTAMP '2019-02-16 15:36:38');
insert into tasks (creator_user_id, name, task_type, date) values (1, 'dz 4', 6, TIMESTAMP '2019-02-16 15:36:38');

insert into messages (task_id, user_id, text, date) values (1, 'anonim', 'привет', TIMESTAMP '2019-02-16 15:36:38');
insert into messages (task_id, user_id, text, date) values (1, 'преп', 'и тебе привет', TIMESTAMP '2019-02-16 15:36:38');
insert into messages (task_id, user_id, text, date) values (1, 'анон', 'как дела', TIMESTAMP '2019-02-16 15:36:38');
insert into messages (task_id, user_id, text, date) values (1, 'крокодил', 'ух ты', TIMESTAMP '2019-02-16 15:36:38');
insert into messages (task_id, user_id, text, date) values (1, 'Вася', 'как это решать', TIMESTAMP '2019-02-16 15:36:38');
insert into messages (task_id, user_id, text, date) values (2, 'anonim', 'привет', TIMESTAMP '2019-02-16 15:36:38');
insert into messages (task_id, user_id, text, date) values (2, 'преп', 'и тебе привет', TIMESTAMP '2019-02-20 15:36:38');
insert into messages (task_id, user_id, text, date) values (2, 'анон', 'как дела', TIMESTAMP '2019-02-18 15:36:38');
insert into messages (task_id, user_id, text, date) values (3, 'крокодил', 'ух ты', TIMESTAMP '2019-02-16 15:36:38');
insert into messages (task_id, user_id, text, date) values (3, 'Вася', 'как это решать', TIMESTAMP '2019-02-16 15:36:38');

insert into messages (task_id, user_id, text, date) values (4, 'девочка', 'привет 2', TIMESTAMP '2019-02-16 15:36:38');
insert into messages (task_id, user_id, text, date) values (4, 'мальчик', 'и тебе привет 2', TIMESTAMP '2019-02-16 15:36:38');
insert into messages (task_id, user_id, text, date) values (4, 'анон', 'как дела 2', TIMESTAMP '2019-02-16 15:36:38');
insert into messages (task_id, user_id, text, date) values (4, 'крокодил', 'ух ты 2', TIMESTAMP '2019-02-16 15:36:38');
insert into messages (task_id, user_id, text, date) values (4, 'препод', 'как это решать 2', TIMESTAMP '2019-02-16 15:36:38');

--7110eda4d09e062aa5e4a390b0a572ac0d2c0220
-- insert into task (name, task_type) values ('dz 1', 2);
-- insert into task (name, task_type) values ('dz 2', 4);
--
-- insert into messages (task_id, user_id, text, date) values (1, 'anonim', 'привет', TIMESTAMP '2019-02-16 15:36:38');
-- insert into messages (task_id, user_id, text, date) values (1, 'преп', 'и тебе привет', TIMESTAMP '2019-02-16 15:36:38');
-- insert into messages (task_id, user_id, text, date) values (1, 'анон', 'как дела', TIMESTAMP '2019-02-16 15:36:38');
-- insert into messages (task_id, user_id, text, date) values (1, 'крокодил', 'ух ты', TIMESTAMP '2019-02-16 15:36:38');
-- insert into messages (task_id, user_id, text, date) values (1, 'Вася', 'как это решать', TIMESTAMP '2019-02-16 15:36:38');
--
-- insert into messages (task_id, user_id, text, date) values (2, 'девочка', 'привет 2', TIMESTAMP '2019-02-16 15:36:38');
-- insert into messages (task_id, user_id, text, date) values (2, 'мальчик', 'и тебе привет 2', TIMESTAMP '2019-02-16 15:36:38');
-- insert into messages (task_id, user_id, text, date) values (2, 'анон', 'как дела 2', TIMESTAMP '2019-02-16 15:36:38');
-- insert into messages (task_id, user_id, text, date) values (2, 'крокодил', 'ух ты 2', TIMESTAMP '2019-02-16 15:36:38');
-- insert into messages (task_id, user_id, text, date) values (2, 'препод', 'как это решать 2', TIMESTAMP '2019-02-16 15:36:38');
--
-- select * from messages;
-- select * from task;
--
-- select * from messages where task_id=1;
--
-- delete from messages where task_id=2;
-- delete from task where id=2;