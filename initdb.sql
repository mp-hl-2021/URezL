CREATE TABLE accounts
(
    id        serial primary key,
    login     varchar(255) not null,
    password  varchar(255) not null,

    createdAt timestamp without time zone default now(),
    updatedAt timestamp without time zone default now(),
    unique (login)
);
