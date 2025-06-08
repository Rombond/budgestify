package account

import "database/sql"

type Account struct {
	Id              int             `json:"id"`
	Name            string          `json:"name"`
	House_User      int             `json:"house_userID"`
	Amount          float64         `json:"amount"`
	Currency        string          `json:"currency"`
	TheoricalAmount sql.NullFloat64 `json:"theoricalAmount"`
}

type AccountForm struct {
	Id              int     `json:"id"`
	Name            string  `json:"name"`
	House_User      int     `json:"house_userID"`
	Amount          float64 `json:"amount"`
	Currency        string  `json:"currency"`
	TheoricalAmount float64 `json:"theoricalAmount"`
	UserId          int     `json:"userID"`
}
