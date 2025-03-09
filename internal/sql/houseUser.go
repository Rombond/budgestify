package db_sql

import (
	"database/sql"
	"fmt"
	"log/slog"
)

func CreateHouseUser(db *sql.DB, houseID int, userID int) (int, error) {
	var id int64 = -1
	query := fmt.Sprintf("INSERT INTO `%s` (`house`, `user`, `admin`) VALUES (?, ?, ?);", Tables[House_UserIdx].Key)
	results, err := db.Exec(query, houseID, userID, true)
	if err != nil {
		slog.Error("[CreateHouseUser] Inserting new house: " + err.Error())
		id = -1
		return int(id), err
	}
	id, err = results.LastInsertId()
	if err != nil {
		slog.Error("[CreateHouseUser] Error while getting lastInsertId: " + err.Error())
		id = -1
		return int(id), err
	}
	return int(id), nil
}

func InviteUserToHouse(db *sql.DB, houseID int, userID int) (int, error) {
	var id int64 = -1
	query := fmt.Sprintf("INSERT INTO `%s` (`house`, `user`, `admin`) VALUES (?, ?, ?);", Tables[House_UserIdx].Key)
	results, err := db.Exec(query, houseID, userID, false)
	if err != nil {
		slog.Error("[InviteUserToHouse] Inserting new house: " + err.Error())
		id = -1
		return int(id), err
	}
	id, err = results.LastInsertId()
	if err != nil {
		slog.Error("[InviteUserToHouse] Error while getting lastInsertId: " + err.Error())
		id = -1
		return int(id), err
	}
	return int(id), nil
}

func IsUserInThisHouse(db *sql.DB, houseID int, userID int) (bool, error) {
	isInside := false
	query := fmt.Sprintf("SELECT (Count(*) > 0) FROM `%s` WHERE user = ? AND house = ?;", Tables[House_UserIdx].Key)
	err := db.QueryRow(query, userID, houseID).Scan(&isInside)
	if err != nil {
		slog.Error("[IsUserInThisHouse] Error querying house id from user: " + err.Error())
		isInside = false
		return isInside, err
	}
	return isInside, nil
}

func IsHouseEmpty(db *sql.DB, houseID int) (bool, error) {
	isEmpty := false
	query := fmt.Sprintf("SELECT (Count(*) <= 0) FROM `%s` WHERE house = ?;", Tables[House_UserIdx].Key)
	err := db.QueryRow(query, houseID).Scan(&isEmpty)
	if err != nil {
		slog.Error("[IsUserInThisHouse] Error querying house id from user: " + err.Error())
		isEmpty = false
		return isEmpty, err
	}
	return isEmpty, nil
}

func DoesHouseGotAdmin(db *sql.DB, houseID int) (int, error) {
	adminCount := 0
	query := fmt.Sprintf("SELECT Count(*) FROM `%s` WHERE house = ? AND admin = 1;", Tables[House_UserIdx].Key)
	err := db.QueryRow(query, houseID).Scan(&adminCount)
	if err != nil {
		slog.Error("[IsUserInThisHouse] Error querying house id from user: " + err.Error())
		adminCount = 0
		return adminCount, err
	}
	return adminCount, nil
}

func GetHouseIDFromUser(db *sql.DB, userID int) ([]int, error) {
	var ids []int
	id := -1
	query := fmt.Sprintf("SELECT house FROM `%s` WHERE user = ?;", Tables[House_UserIdx].Key)
	rows, err := db.Query(query, userID)
	if err != nil {
		slog.Error("[GetHouseIDFromUser] Error querying house id from user: " + err.Error())
		return ids, err
	}
	for rows.Next() {
		rows.Scan(&id)
		ids = append(ids, id)
	}
	if err = rows.Err(); err != nil {
		return ids, err
	}
	return ids, nil
}

func LeaveHouse(db *sql.DB, houseID int, userID int) (bool, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE house = ? AND user = ?", Tables[House_UserIdx].Key)
	_, err := db.Exec(query, houseID, userID)
	if err != nil {
		slog.Error("[LeaveHouse] Error while leaving house: " + err.Error())
		return false, err
	}
	return true, nil
}

func IsUserAdmin(db *sql.DB, houseID int, userID int) (bool, error) {
	admin := false
	query := fmt.Sprintf("SELECT (Count(*) > 0) FROM `%s` WHERE house = ? AND admin = 1 AND user = ?;", Tables[House_UserIdx].Key)
	err := db.QueryRow(query, houseID, userID).Scan(&admin)
	if err != nil {
		slog.Error("[IsUserAdmin] Error querying admin: " + err.Error())
		admin = false
		return admin, err
	}
	return admin, nil
}
