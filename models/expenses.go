package models

type Expense struct {
	Id int64 `json:"id"`
	UserId int `json:"user_id"`
	Description string `json:"description"`
	Model
}

func (e *Expense) GetExpense() error {
	err := db.Select(e)
	if err != nil {
		return err
	}
	return nil
}
