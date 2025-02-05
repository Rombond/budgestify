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
		slog.Error("[GetHouse] Error querying house: " + err.Error())
	}
	return house, err
}

func GetHouses(db *sql.DB, userID int) ([]house.House, error) {
	var houses []house.House
	var house house.House
	query := fmt.Sprintf("SELECT h.* FROM %s h JOIN %s hu ON h.id = hu.house WHERE hu.user = ?;", Tables[HouseIdx].Key, Tables[House_UserIdx].Key)
	rows, err := db.Query(query, userID)
	if err != nil {
		slog.Error("[GetCategories] Error querying house id from user: " + err.Error())
		return houses, err
	}
	house.UserId = userID
	for rows.Next() {
		rows.Scan(&house.Id, &house.Name, &house.AccountId)
		houses = append(houses, house)
	}
	if err = rows.Err(); err != nil {
		return houses, err
	}
	return houses, err
}

func CreateHouse(db *sql.DB, name string) (int, error) {
	var id int64 = -1
	query := fmt.Sprintf("INSERT INTO `%s` (`name`) VALUES (?);", Tables[HouseIdx].Key)
	results, err := db.Exec(query, name)
	if err != nil {
		slog.Error("[CreateHouse] Inserting new house: " + err.Error())
		id = -1
		return int(id), err
	}
	id, err = results.LastInsertId()
	if err != nil {
		slog.Error("[CreateHouse] Error while getting lastInsertId: " + err.Error())
		id = -1
		return int(id), err
	}
	return int(id), nil
}

func ChangeHouseName(db *sql.DB, houseID int, newName string) (bool, error) {
	query := fmt.Sprintf("UPDATE %s SET `name` = ? WHERE id = ?", Tables[HouseIdx].Key)
	_, err := db.Exec(query, newName, houseID)
	if err != nil {
		slog.Error("[ChangeHouseName] Error while updating house: " + err.Error())
		return false, err
	}
	return true, nil
}

func DeleteHouse(db *sql.DB, houseID int) (bool, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", Tables[HouseIdx].Key)
	_, err := db.Exec(query, houseID)
	if err != nil {
		slog.Error("[DeleteHouse] Error while deleting house: " + err.Error())
		return false, err
	}
	return true, nil
}

func AddAccount(db *sql.DB, houseID int, accountHouseID int) (bool, error) {
	query := fmt.Sprintf("UPDATE %s SET `account` = ? WHERE id = ?", Tables[HouseIdx].Key)
	_, err := db.Exec(query, accountHouseID, houseID)
	if err != nil {
		slog.Error("[AddAccount] Error while updating house: " + err.Error())
		return false, err
	}
	return true, nil
}
