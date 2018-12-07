package models

import (
	"github.com/go-pg/pg/orm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int        `json:"id"`
	Email    string     `json:"email"`
	Password string     `json:"password,omitempty"`
	Expenses []*Expense `json:"expenses"`
	Tags     []*Tag     `json:"tags"`
	Model
}

func (u *User) BeforeInsert(db orm.DB) error {
	_ = u.Model.BeforeInsert(db)
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) StoreUser() error {
	err := db.Insert(u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetUserByEmail() error {
	err := db.Model(u).Where("email = ?", u.Email).Select()
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetUserExpenses() error {
	err := db.Model(u).
		Where("id = ?", u.Id).
		Column("user.*", "Expenses", "Tags", "Expenses.Tags").
		First()
	if err != nil {
		return err
	}
	return nil
}
