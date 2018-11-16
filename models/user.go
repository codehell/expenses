package models

import (
	"github.com/go-pg/pg/orm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id int64 `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	Model
}

func (u *User) BeforeInsert(db orm.DB) error {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	if u.UpdatedAt.IsZero() {
		u.UpdatedAt = time.Now()
	}
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
