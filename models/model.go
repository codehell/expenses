package models

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"time"
)

var db *pg.DB

func init()  {
	db = pg.Connect(&pg.Options{
		User:     "codehell",
		Database: "awesome",
	})
}

type Model struct {
	CreatedAt time.Time `json:"created_at" sql:"default:now()"`
	UpdatedAt time.Time `json:"updated_at" sql:"default:now()"`
}

func (m *Model) BeforeInsert (db orm.DB) error {
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now()
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = time.Now()
	}
	return nil
}