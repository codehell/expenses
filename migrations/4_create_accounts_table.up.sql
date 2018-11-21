CREATE table accounts (
  id serial primary key,
  user_id integer references users(id),
  amount decimal(14,2),
  created_at timestamp,
  updated_at timestamp
);