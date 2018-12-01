CREATE TABLE expense_tags (
  expense_id integer not null references expenses(id),
  tag_id integer not null references tags(id),
  constraint pk_expense_tags primary key (expense_id, tag_id)
)