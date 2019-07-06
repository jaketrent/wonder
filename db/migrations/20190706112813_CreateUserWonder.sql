-- +goose Up
begin;
  create sequence user_wonders_id_seq
  start with 1 
  increment by 1
  no minvalue
  no maxvalue
  cache 1;

  create table user_wonders (
    id int primary key default nextval('user_wonders_id_seq'),
    user_id int not null references users(id),
    description varchar(255) not null,
    created date default now()
  );
end;


-- +goose Down
begin;
  drop sequence user_wonders_id_seq;
  drop table user_wonders;
end;


