\c postgres;
drop database companionships;
create database companionships;
\c companionships;

create table creature_types(
	id bigserial primary key,
	creature_type varchar(255)
);

create table creatures(
	id bigserial primary key,
	name varchar(255),
	creature_type_id int references creature_types(id),
	created_at timestamptz default now()
);

create table hobbies(
	id bigserial primary key,
	hobby varchar(255)
);

create table foods(
	id bigserial primary key,
	food varchar(255)
);

create table creature_hobbies(
	creature_id bigserial,
	hobby_id bigserial references hobbies (id),
	favorite_hobby boolean,
	created_at timestamptz default now()
);

create table creature_foods(
	creature_id bigserial,
	food_id bigserial references foods (id),
	favorite_food boolean,
	created_at timestamptz default now()
);

create table companionships(
	id bigserial primary key,
	pet_id bigint references creatures (id),
	guardian_id bigint references creatures (id),
	created_at timestamptz default now()
);

insert into creature_types (creature_type) values ('Cat'), ('Dog'), ('Human'), ('Gorilla');
