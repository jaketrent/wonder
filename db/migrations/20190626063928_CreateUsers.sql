-- +goose Up
begin;
  create sequence users_id_seq
  start with 1 
  increment by 1
  no minvalue
  no maxvalue
  cache 1;

  create table users (
    id int primary key default nextval('users_id_seq'),
    name varchar(255) not null
  );
end;


-- +goose Down
begin;
  drop sequence users_id_seq;
  drop table users;
end;

