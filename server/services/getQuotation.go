package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type EconomyApiResponse struct {
	USDBRL struct {
		Bid string
	}
}

type Quotation struct {
	price string
}

func GetQuotation() (*EconomyApiResponse, error) {
	client := http.Client{}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var QuotationInfo EconomyApiResponse
	err = json.Unmarshal(data, &QuotationInfo)
	if err != nil {
		return nil, err
	}

	err = saveQuotation(QuotationInfo.USDBRL.Bid)
	if err != nil {
		return nil, err
	}

	return &QuotationInfo, nil
}

func saveQuotation(quotation string) error {
	db, err := getDbInstance()
	if err != nil {
		return err
	}
	defer db.Close()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	stmt, err := db.Prepare("INSERT INTO Quotation(quotation) VALUES(?)")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, quotation)
	if err != nil {
		return err
	}

	return nil
}

func getDbInstance() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		return nil, err
	}

	return db, err
}
