package gobitlaunch

import (
	"encoding/json"
	"fmt"
	"time"
)

// Transaction represents a transaction
type Transaction struct {
	ID           string    `json:"id"`
	TID          string    `json:"transactionId"`
	Date         time.Time `json:"date"`
	Address      string    `json:"address"`
	Symbol       string    `json:"cryptoSymbol"`
	AmountUSD    float64   `json:"amountUsd"`
	AmountCrypto string    `json:"amountCrypto"`
	Status       string    `json:"status"`
	StatusURL    string    `json:"statusUrl"`
	QrCodeURL    string    `json:"qrCodeUrl"`
}

// CreateTransactionOptions represents options for create a new transaction
type CreateTransactionOptions struct {
	AmountUSD        int    `json:"amountUsd"`
	CryptoSymbol     string `json:"cryptoSymbol"`
	LightningNetwork bool   `json:"lightningNetwork"`
}

// TransactionService manages account API actions
type TransactionService struct {
	client *Client
}

// Create transaction
func (ss *TransactionService) Create(opts *CreateTransactionOptions) (*Transaction, error) {
	b, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}
	req, err := ss.client.NewRequest("POST", "/transactions", b)
	if err != nil {
		return nil, err
	}

	s := Transaction{}
	if err := ss.client.DoRequest(req, &s); err != nil {
		return nil, err
	}

	return &s, nil
}

// Show transaction
func (ss *TransactionService) Show(id string) (*Transaction, error) {
	req, err := ss.client.NewRequest("GET", "/transactions/"+id, nil)
	if err != nil {
		return nil, err
	}

	t := Transaction{}
	if err := ss.client.DoRequest(req, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

// List transactions
func (ss *TransactionService) List(page, perPage int) ([]Transaction, error) {
	q := fmt.Sprintf("?page=%d&items=%d", page, perPage)
	req, err := ss.client.NewRequest("GET", "/transactions"+q, nil)
	if err != nil {
		return nil, err
	}

	a := struct {
		History []Transaction
	}{}
	if err := ss.client.DoRequest(req, &a); err != nil {
		return nil, err
	}

	return a.History, nil
}
