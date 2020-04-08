package gobitlaunch

import (
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
