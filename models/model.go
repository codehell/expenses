package models

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var db *pg.DB

type dbConfig struct {
	Addr     string
	User     string
	Database string
	Password string
}

func init() {
	var config dbConfig
	var confFile string
	env := os.Getenv("GCP_ENVIRONMENT")
	log.Println("environment: ", env)
	if env == "develop" {
		confFile = "./database.yaml"
	} else {
		confFile = "./database.pro.yaml"
	}
	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatal(err)
	}
	db = pg.Connect(&pg.Options{
		User:     config.User,
		Database: config.Database,
		Addr:     config.Addr,
		Password: config.Password,
	})
}

type Model struct {
	CreatedAt time.Time `json:"created_at" sql:"default:now()"`
	UpdatedAt time.Time `json:"updated_at" sql:"default:now()"`
}

func (m *Model) BeforeInsert(db orm.DB) error {
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now()
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = time.Now()
	}
	return nil
}
