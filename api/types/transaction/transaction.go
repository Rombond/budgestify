package transaction

import (
	"database/sql"
	"time"
)

type Transaction struct {
	Id             int           `json:"id"`
	Name           string        `json:"name"`
	CategoryID     sql.NullInt32 `json:"categoryID"`
	Amount         float64       `json:"amount"`
	PayerID        int           `json:"payerID"`
	PayerAccountID sql.NullInt32 `json:"payerAccountID"`
	PayDate        time.Time     `json:"payDate"`
	Currency       string        `json:"currency"`
	ConversionRate float64       `json:"conversionRate"`
}

type TransactionForm struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	CategoryID     int       `json:"categoryID"`
	Amount         float64   `json:"amount"`
	PayerID        int       `json:"payerID"`
	PayerAccountID int       `json:"payerAccountID"`
	PayDate        time.Time `json:"payDate"`
	Currency       string    `json:"currency"`
	ConversionRate float64   `json:"conversionRate"`
	UserId         int       `json:"userID"`
	House_User     int       `json:"house_userID"`
}
