package models

type Tag struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	User   *User
	Name   string `json:"name"`
	Expenses []Expense `json:"expenses" pg:"many2many:expense_tags"`
	Model
}

func (t *Tag) StoreTag() error {
	err := db.Insert(t)
	if err != nil {
		return err
	}
	return nil
}