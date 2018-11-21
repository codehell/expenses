CREATE TABLE expense_tag (
  id serial primary key,
  expense_id integer references expenses(id),
  tag_id integer references tags(id),
  created_at timestamp,
  updated_at timestamp
)