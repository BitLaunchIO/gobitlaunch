package gobitlaunch

import (
	"fmt"
	"time"
)

// Account represents a BitLaunch account
type Account struct {
	ID                  string    `json:"id"`
	Email               string    `json:"email"`
	EmailConfirmed      bool      `json:"emailConfirmed"`
	Created             time.Time `json:"created"`
	Used                int       `json:"used"`
	Limit               int       `json:"limit"`
	Twofa               bool      `json:"twofa"`
	Balance             int       `json:"balance"`
	CostPerHr           int       `json:"costPerHr"`
	LowBalanceAlertDays int       `json:"billingAlert"`
	NegativeAllowance   int       `json:"negativeAllowance"`
}

type usageData struct {
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Cost        int       `json:"cost"`
	Hours       int       `json:"hours"`
	Amount      int       `json:"amount"`
	Type        string    `json:"type"`
}

// AccountUsage represents the usage data of a BitLaunch account
type AccountUsage struct {
	Server     []usageData `json:"serverUsage"`
	Backup     []usageData `json:"backupUsage"`
	Bandwidth  []usageData `json:"bandwidthUsage"`
	Protection []usageData `json:"protectionUsage"`
	TotalUSD   int         `json:"totalUsd"`
	PrevMonth  string      `json:"prevMonth"`
	ThisMonth  string      `json:"thisMonth"`
	NextMonth  string      `json:"nextMonth"`
}

// AccountHistory represents the usage data of a BitLaunch account
type AccountHistory struct {
	History []struct {
		ID          string    `json:"id"`
		Time        time.Time `json:"time"`
		Description string    `json:"description"`
	} `json:"history"`
	Total int `json:"total"`
}

// AccountService manages account API actions
type AccountService struct {
	client *Client
}

// Show the account
func (as *AccountService) Show() (*Account, error) {
	req, err := as.client.NewRequest("GET", "/user", nil)
	if err != nil {
		return nil, err
	}

	account := &Account{}
	if err := as.client.DoRequest(req, account); err != nil {
		return nil, err
	}

	return account, nil
}

// Usage shows the account usage
func (as *AccountService) Usage(filter ...string) (*AccountUsage, error) {
	var search string

	if len(filter) > 1 {
		return nil, fmt.Errorf("Too many arguments")
	} else if len(filter) == 0 {
		search = "latest"
	} else {
		search = filter[0]
	}

	req, err := as.client.NewRequest("GET", "/usage?period="+search, nil)
	if err != nil {
		return nil, err
	}

	usage := &AccountUsage{}
	if err := as.client.DoRequest(req, usage); err != nil {
		return nil, err
	}

	return usage, nil
}

// History shows the account history/activity
func (as *AccountService) History(page, perPage int) (*AccountHistory, error) {
	q := fmt.Sprintf("?page=%d&items=%d", page, perPage)

	req, err := as.client.NewRequest("GET", "/security/history"+q, nil)
	if err != nil {
		return nil, err
	}

	history := &AccountHistory{}
	if err := as.client.DoRequest(req, history); err != nil {
		return nil, err
	}

	return history, nil
}
