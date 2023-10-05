package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", getDailyQuotation)

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		panic(err)
	}
}

func getDailyQuotation(w http.ResponseWriter, r *http.Request) {
	client := http.Client{}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Printf("Error: %v \n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %v \n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error: %v \n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	err = saveQuotationIntoFile(string(data))
	if err != nil {
		fmt.Printf("Error: %v \n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}

func saveQuotationIntoFile(quotation string) error {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	quotationInfo := []byte("Dolar:" + quotation)
	_, err = file.Write(quotationInfo)
	if err != nil {
		return err
	}

	return nil
}
