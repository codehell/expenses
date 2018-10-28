CREATE table users (
  id serial primary key,
  username varchar(64),
  email varchar(128),
  password varchar(64),
  created_at
)