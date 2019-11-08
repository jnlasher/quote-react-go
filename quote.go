package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// QuoteStruct struct for containing a quote
type QuoteStruct struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

// FetchRandomQuote Fetches a random quote from the database
func FetchRandomQuote() (*QuoteStruct, error) {
	query := "SELECT * FROM quotes ORDER BY RANDOM() LIMIT 1"
	row, err := QueryDB(query)
	if err != nil {
		return nil, err
	}

	if !row.Next() {
		return nil, errors.New("No quote found in database")
	}

	var (
		id     int64
		quote  string
		author string
	)
	if err := row.Scan(&id, &quote, &author); err != nil {
		log.Fatalf("Failed Scanning - %s", err)
	}

	quoteStruct := &QuoteStruct{
		Quote:  quote,
		Author: author,
	}

	return quoteStruct, nil
}

// WriteNewQuote writes a new quote to the database
func WriteNewQuote(request *http.Request) error {
	var newQuote QuoteStruct
	err := json.NewDecoder(request.Body).Decode(&newQuote)
	if err != nil {
		return err
	}

	if newQuote.Quote == "" || newQuote.Author == "" {
		return errors.New("A quote or author cannot be empty")
	}

	query := "INSERT INTO quotes (quote, author) VALUES (?, ?)"
	_, err = ExecDB(query, newQuote.Quote, newQuote.Author)
	if err != nil {
		return err
	}

	return nil
}
