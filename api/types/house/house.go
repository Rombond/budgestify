package house

import "database/sql"

type House struct {
	Id        int           `json:"id"`
	Name      string        `json:"name"`
	AccountId sql.NullInt32 `json:"accountId"`
	UserId    int           `json:"userID"`
}

type HouseForm struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	HouseID int    `json:"houseID"`
	UserId  int    `json:"userID"`
}
