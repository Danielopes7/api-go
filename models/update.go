package models

import "github.com/Danielopes7/api-go/db"

func Update(id int64, todo Todo) (int64, error) {
	db, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer db.Close()

	result, err := db.Exec(`UPDATE todos SET title = $1, description = $2, done = $3 WHERE id = $4`, todo.Title, todo.Description, todo.Done, id)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
