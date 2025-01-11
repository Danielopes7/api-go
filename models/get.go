package models

import (
	"github.com/Danielopes7/api-go/db"
)

func Get(id int64) (todo Todo, err error) {
	db, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer db.Close()

	row := db.QueryRow(`SELECT * FROM todos WHERE id = $1`, id)
	if err != nil {
		return
	}
	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}
