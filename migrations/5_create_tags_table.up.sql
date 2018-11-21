CREATE TABLE tags (
  id serial primary key,
  user_id integer references users(id),
  name varchar(100),
  created_at timestamp,
  updated_at timestamp
);