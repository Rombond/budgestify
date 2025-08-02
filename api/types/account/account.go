package account

import "database/sql"

type Account struct {
	Id                int             `json:"id"`
	Name              string          `json:"name"`
	House_User        int             `json:"house_userID"`
	Amount            float64         `json:"amount"`
	Currency          string          `json:"currency"`
	TheoreticalAmount sql.NullFloat64 `json:"theoreticalAmount"`
}

type AccountForm struct {
	Id                int     `json:"id"`
	Name              string  `json:"name"`
	House_User        int     `json:"house_userID"`
	Amount            float64 `json:"amount"`
	Currency          string  `json:"currency"`
	TheoreticalAmount float64 `json:"theoreticalAmount"`
	UserId            int     `json:"userID"`
}
