CREATE table users (
  id serial primary key,
  email varchar(128),
  password varchar(64),
  created_at timestamp,
  updated_at timestamp
)