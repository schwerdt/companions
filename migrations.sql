\c postgres;
drop database animals;
create database animals;
\c animals;

create table pet_types(
	id bigserial primary key,
	pet_type varchar(255)
);

create table pets(
	id bigserial primary key,
	name varchar(255),
	pet_type_id int references pet_types(id),
	created_at timestamptz default now()
);

create table guardians(
	id bigserial primary key,
	name varchar(255),
	created_At timestamptz default now()
);

create table pet_guardians(
	pet_id bigserial,
	guardian_id bigserial,
	primary_guardian boolean,
	created_at timestamptz default now()
);

create table animals(
	id bigserial primary key,
	dog_id bigint references pets (id),
	cat_id bigint references pets (id),
	created_at timestamptz default now()
);

insert into pet_types (pet_type) values ('Cat'), ('Dog');
