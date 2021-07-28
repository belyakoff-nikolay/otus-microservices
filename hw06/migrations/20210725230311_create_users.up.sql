CREATE SEQUENCE IF NOT EXISTS user_id_seq;

CREATE TABLE IF NOT EXISTS users(
      ID bigint not null primary key default NEXTVAL('user_id_seq'),
      FirstName varchar(50) not null,
      LastName varchar(50) not null,
      Email varchar(150) not null
);