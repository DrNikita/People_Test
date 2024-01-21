CREATE TABLE people.public.persons (
    id SERIAL NOT NULL,
    name varchar(255) NOT NULL ,
    surname varchar(255) NOT NULL ,
    patronymic varchar(255),
    age integer,
    gender varchar(10),
    country_id varchar(5),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    primary key (id)
);