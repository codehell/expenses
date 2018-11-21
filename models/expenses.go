package models

import (
	"github.com/shopspring/decimal"
)

type Expense struct {
	Id int64 `json:"id"`
	UserId int64 `json:"user_id"`
	Amount decimal.Decimal `json:"amount"`
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

func (e *Expense) StoreExpense() error {
	err := db.Insert(e)
	if err != nil {
		return err
	}
	return nil
}
