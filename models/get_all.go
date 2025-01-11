package models

import (
	"github.com/Danielopes7/api-go/db"
)

func GetAll() (todos []Todo, err error) {
	db, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer db.Close()

	rows, err := db.Query(`SELECT * FROM todo`)

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}

	return
}
