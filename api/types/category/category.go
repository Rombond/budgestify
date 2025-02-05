package category

import "database/sql"

type Category struct {
	Id     int            `json:"id"`
	Name   string         `json:"name"`
	Icon   sql.NullString `json:"icon"`
	Parent sql.NullInt32  `json:"parent"`
	House  int            `json:"houseID"`
}

type CategoryForm struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Icon    string `json:"icon"`
	Parent  int    `json:"parent"`
	HouseId int    `json:"houseID"`
	UserId  int    `json:"userID"`
}
