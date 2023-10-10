package models

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Todo struct {
	Id      int    `json:"id" db:"id"`
	Subject string `json:"subject" db:"subject"`
}

func Create(subject string, db *sqlx.DB) {
	query := "INSERT INTO todos (subject) VALUES ($1)"
	result, err := db.Exec(query, subject)
	if err != nil {
		log.Println("error while creating todo ", err)
	} else {
		log.Println(result)
	}
}

func Update(newSubject string, id int, db *sqlx.DB) {
	query := "UPDATE todos SET subject = $1 WHERE id=$2"
	result, err := db.Exec(query, newSubject, id)
	if err != nil {
		log.Println("error while updating todo ", err)
	} else {
		log.Println(result)
	}
}

func Delete(id int, db *sqlx.DB) {
	query := "DELETE FROM todos WHERE id=$1"
	result, err := db.Exec(query, id)
	if err != nil {
		log.Println("error while deleting todo ", err)
	} else {
		log.Println(result)
	}
}

func ReadAll(db *sqlx.DB) []Todo {
	query := "SELECT * FROM todos"
	todos := []Todo{}
	err := db.Select(&todos, query)
	if err != nil {
		log.Println("error while selecting todos ", err)
	}
	return todos
}
