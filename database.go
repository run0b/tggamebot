package main

import (
	"errors"
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

func (v Database) SetUserClass(id int64, class int) error {
	_, err := v.i.Exec("UPDATE users SET class = $1 WHERE id = $2", class, id)

	return err
}

func (v Database) GetUserClass(id int64) (int, error) {
	rows, err := v.i.Query("SELECT class FROM users WHERE id = $1", id)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	if err != nil {
		return -1, err
	}
	if !rows.Next() {
		return -1, errors.New("No user")
	}
	var class int
	if err := rows.Scan(&class); err != nil {
		return -1, nil
	}

	return class, err
}
