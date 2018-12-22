package models

type Expense struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Amount      string `json:"amount" binding:"required,numeric,max=16"`
	Description string `json:"description" binding:"required,max=255"`
	Tags        []Tag  `json:"tags" pg:"many2many:expense_tags"`
	Model
}

func (e *Expense) GetExpense() error {
	err := db.Select(e)
	if err != nil {
		return err
	}
	return nil
}

func (e *Expense) StoreExpense() error {
	err := db.Insert(e)
	if err != nil {
		return err
	}
	var et ExpenseTag
	for _, t := range e.Tags {
		et = ExpenseTag{ExpenseId: e.Id, TagId: t.Id}
		err = et.StoreExpenseTag()
		if err != nil {
			return err
		}
	}
	return nil
}

//DeleteExpense Delete an expense and his related tags
func (e *Expense) DeleteExpense() error {
	var et ExpenseTag
	_, err := db.Model(&et).Where("expense_id = ?", e.Id).Delete()
	if err != nil {
		return err
	}
	err = db.Delete(e)
	if err != nil {
		return err
	}
	return nil
}
