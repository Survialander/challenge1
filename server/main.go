package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"server/controllers"
)

func main() {
	err := setupDatabase()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/cotacao", controllers.GetQuotationController)
	mux.HandleFunc("/database", controllers.CheckDatabaseController)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func setupDatabase() error {
	file, err := os.Create("database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()

	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return err
	}
	defer db.Close()

	createQuotationTableSQL := `CREATE TABLE quotation (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"quotation" VARCHAR(50)		
	);`

	_, err = db.Exec(createQuotationTableSQL)

	if err != nil {
		return err
	}

	return nil
}
