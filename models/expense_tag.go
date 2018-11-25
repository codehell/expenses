package models

import "github.com/go-pg/pg/orm"

type ExpenseTag struct {
	Id        int      `json:"id"`
	ExpenseId int      `json:"expense_id"`
	Expense   *Expense `json:"expense"`
	TagId     int      `json:"tag_id"`
	Tag       *Tag     `json:"tag"`
	Model
}

func init() {
	orm.RegisterTable((*ExpenseTag)(nil))
}
