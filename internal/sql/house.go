package db_sql

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/Rombond/budgestify/api/types/house"
)

func GetHouse(db *sql.DB, id int) (house.House, error) {
	var house house.House
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE id = ?;", Tables[HouseIdx].Key)
	err := db.QueryRow(query, id).Scan(&house.Id, &house.Name, &house.AccountId)
	if err != nil {
		slog.Error("[GetUserID] Error querying user id: " + err.Error())
	}
	return house, err
}

func GetHouseIDFromUser(db *sql.DB, userID int) (int, error) {
	id := -1
	query := fmt.Sprintf("SELECT house FROM `%s` WHERE user = ?;", Tables[House_UserIdx].Key)
	err := db.QueryRow(query, userID).Scan(&id)
	if err != nil {
		slog.Error("[GetHouseIDFromUser] Error querying house id from user: " + err.Error())
		id = -1
		return id, err
	}
	return id, nil
}

func CreateHouse(db *sql.DB, name string) (int, error) {
	id := -1
	query := fmt.Sprintf("INSERT INTO `%s` (`name`) VALUES (?);", Tables[HouseIdx].Key)
	err := db.QueryRow(query, name).Scan(&id)
	if err != nil {
		slog.Error("[CreateHouse] Inserting new house: " + err.Error())
		id = -1
		return id, err
	}
	return id, nil
}

func ChangeHouseName(db *sql.DB, houseID int, newName string) (bool, error) {
	query := fmt.Sprintf("UPDATE %s SET `name` = ? WHERE id = ?", Tables[houseID].Key)
	_, err := db.Exec(query, newName, houseID)
	if err != nil {
		slog.Error("[ChangeHouseName] Error while updating house: " + err.Error())
		return false, err
	}
	return true, nil
}
