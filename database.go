package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	i *sqlx.DB
}

func CreateDatabase(c *Config) *Database {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	database := Database{
		i: db,
	}
	return &database
}

func (v Database) AddUser(id int64, name string) error {
	tx := v.i.MustBegin()
	_, err := tx.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", id, name)
	if err != nil {
		return err
	}
	err = tx.Commit()
	return err
}
