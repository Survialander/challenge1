package services

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func CheckDatabase() error {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT quotation FROM quotation")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var q Quotation

		err := rows.Scan(&q.price)
		if err != nil {
			return err
		}

		fmt.Printf("%v \n", q)
	}

	return nil
}
