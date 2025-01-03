package models

import "github.com/Danielopes7/api-go/db"

func Delete(id int64) (int64, error) {
	db, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer db.Close()

	result, err := db.Exec(`DELETE from todos WHERE id = $1`, id)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
