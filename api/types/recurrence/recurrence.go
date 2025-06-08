package recurrence

import (
	"database/sql"
	"time"
)

type Recurrence struct {
	Id             int           `json:"id"`
	Name           string        `json:"name"`
	House_User     int           `json:"house_userID"`
	PayerAccountID int           `json:"payerAccountID"`
	Category       sql.NullInt32 `json:"categoryID"`
	Amount         float64       `json:"amount"`
	Currency       string        `json:"currency"`
	ConversionRate float64       `json:"conversionRate"`
	PayDate        time.Time     `json:"payDate"`
	DayCycle       int           `json:"dayCycle"`
}

type RecurrenceForm struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	House_User     int       `json:"house_userID"`
	PayerAccountID int       `json:"payerAccountID"`
	Category       int       `json:"categoryID"`
	Amount         float64   `json:"amount"`
	Currency       string    `json:"currency"`
	ConversionRate float64   `json:"conversionRate"`
	PayDate        time.Time `json:"payDate"`
	DayCycle       int       `json:"dayCycle"`
	UserId         int       `json:"userID"`
}
