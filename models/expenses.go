package models

type Expense struct {
	Id int64 `json:"id"`
	UserId []User `json:"user"`
	Model
}