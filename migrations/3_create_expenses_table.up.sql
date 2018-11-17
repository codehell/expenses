CREATE TABLE expenses (
  id serial primary key,
  user_id integer references users(id),
  description varchar(255),
  created_at timestamp,
  updated_at timestamp
)