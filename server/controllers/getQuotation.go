package controllers

import (
	"encoding/json"
	"net/http"
	"server/services"
)

func GetQuotationController(w http.ResponseWriter, r *http.Request) {
	quotation, err := services.GetQuotation()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(quotation.USDBRL.Bid)
}
