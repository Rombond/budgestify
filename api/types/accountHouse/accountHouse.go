package accountHouse

type AccountHouse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	UserId int    `json:"userID"`
}

type AccountHouseForm struct {
	Name    string `json:"name"`
	Amount  int    `json:"amount"`
	UserId  int    `json:"userID"`
	HouseID int    `json:"houseID"`
}
