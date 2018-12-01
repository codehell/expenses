package models

import "github.com/go-pg/pg/orm"

type ExpenseTag struct {
	ExpenseId int      `json:"expense_id" sql:",pk"`
	TagId     int      `json:"tag_id" sql:",pk"`
}

func init() {
	orm.RegisterTable((*ExpenseTag)(nil))
}

func (et *ExpenseTag) StoreExpenseTag() error {
	err := db.Insert(et)
	if err != nil {
		return err
	}
	return nil
}

func (et *ExpenseTag) DeleteExpenseTag() error {
	err := db.Delete(et)
	if err != nil {
		return err
	}
	return nil
}