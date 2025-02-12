package houseUser

type HouseUser struct {
	Id              int `json:"id"`
	HouseID         int `json:"houseID"`
	UserID          int `json:"userID"`
	TheoricalAmount int `json:"theoricalAmount"`
}
