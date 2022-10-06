package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

// Transaction type from the assignment
type Transaction struct {
	// a UUID of transaction
	ID string
	// in USD, typically a value between 0.01 and 1000 USD.
	Amount float64
	// a 2-letter country code of where the bank is located
	BankCountryCode string
}

// loads latencies map from the given JSON
func loadLatencies(path string) map[string]int {
	f, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Could not open " + path)
	}
	var latenciesData map[string]int
	err = json.Unmarshal(f, &latenciesData)
	if err != nil {
		log.Fatal("Could not unmarshal latencies.")
	}
	return latenciesData
}

// loads transactions from the given CSV
func loadTransactions(path string) []Transaction {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("Could not open " + path)
	}
	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+path, err)
	}
	// get transactions and skip the csv header
	var transactions = make([]Transaction, len(records))
	for i, line := range records[1:] {
		amount, err := strconv.ParseFloat(line[1], 32)
		if err != nil {
			log.Fatal("Could not parse float value from " + line[1])
		}
		transactions[i] = Transaction{ID: line[0], Amount: amount, BankCountryCode: line[2]}
	}

	return transactions
}
