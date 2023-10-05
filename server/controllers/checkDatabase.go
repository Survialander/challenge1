package controllers

import (
	"net/http"
	"server/services"
)

func CheckDatabaseController(w http.ResponseWriter, r *http.Request) {
	err := services.CheckDatabase()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	w.WriteHeader(http.StatusNoContent)
	w.Header().Add("Content-type", "application/json")
}
