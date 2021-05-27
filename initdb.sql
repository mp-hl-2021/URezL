DROP TABLE IF EXISTS accounts CASCADE;
CREATE TABLE accounts
(
    id        serial primary key,
    login     varchar(255) not null,
    password  varchar(255) not null,

    createdAt timestamp without time zone default now(),
    updatedAt timestamp without time zone default now(),

    unique (login)
);

DROP TABLE IF EXISTS oldLinkByNewLink CASCADE;
CREATE TABLE oldLinkByNewLink
(
    newLink   varchar(255) primary key,
    oldLink   varchar(255) not null,
    lifetime  int,
    accountId int,
    isBad     boolean not null,

    createdAt timestamp without time zone default now(),
    updatedAt timestamp without time zone default now()
);

DROP TABLE IF EXISTS linksByAccountId CASCADE;
CREATE TABLE linksByAccountId
(
    accountId int not null,
    newLink varchar(255) not null,
    oldLink varchar(255) not null,

    createdAt timestamp without time zone default now(),
    updatedAt timestamp without time zone default now()
);
