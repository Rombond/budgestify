package db_sql

import (
	"database/sql"
	"fmt"
	"log/slog"
)

type StateSetup struct {
	IsDbInitialized   bool `json:"isDatabaseInitialized"`
	IsUserCreated     bool `json:"isUserCreated"`
	IsHouseCreated    bool `json:"isHouseCreated"`
	IsCategoryCreated bool `json:"isCategoryCreated"`
}

var state StateSetup

func GetStateSetup() *StateSetup {
	return &state
}

func GetSetupDone() bool {
	return state.IsDbInitialized || state.IsUserCreated || state.IsHouseCreated || state.IsCategoryCreated
}

func UpdateStateSetup(db *sql.DB, login string) {
	if !state.IsDbInitialized {
		state.IsDbInitialized = setDbInitialized(db)
		if !state.IsDbInitialized {
			return
		}
	}
	if !state.IsUserCreated {
		state.IsUserCreated = setUserCreated(db)
		if !state.IsUserCreated {
			return
		}
	}
	if login == "" {
		return
	}
	userID, err := GetUserID(db, login)
	if !state.IsHouseCreated {
		if userID == -1 || err != nil {
			return
		}
		state.IsHouseCreated = setHouseCreated(db, userID)
		if !state.IsHouseCreated {
			return
		}
	}
	houseIDs, err := GetHouseIDFromUser(db, userID)
	if err != nil || len(houseIDs) > 1 {
		return
	}
	state.IsCategoryCreated = setCategoryCreated(db, houseIDs[0])
}

func setDbInitialized(db *sql.DB) bool {
	isInitialized := true
	for _, pair := range Tables {
		if IsTableMissing(db, pair.Key) {
			isInitialized = false
			break
		}
	}
	return isInitialized
}

func setUserCreated(db *sql.DB) bool {
	isCreated := true
	query := fmt.Sprintf("SELECT (Count(*) > 0) FROM `%s`;", Tables[UserIdx].Key)
	err := db.QueryRow(query).Scan(&isCreated)
	if err != nil {
		slog.Error("[setUserCreated] Error querying number of users: " + err.Error())
		isCreated = false
	}
	return isCreated
}

func setHouseCreated(db *sql.DB, userID int) bool {
	isCreated := true
	isThereHouses := true
	query := fmt.Sprintf("SELECT (Count(*) > 0) FROM `%s`;", Tables[HouseIdx].Key)
	err := db.QueryRow(query).Scan(&isThereHouses)
	if err != nil {
		slog.Error("[setHouseCreated] Error querying number of houses: " + err.Error())
		isCreated = false
	}
	if isThereHouses {
		query = fmt.Sprintf("SELECT (Count(*) > 0) FROM `%s` WHERE user = %d;", Tables[House_UserIdx].Key, userID)
		err = db.QueryRow(query).Scan(&isCreated)
		if err != nil {
			slog.Error("[setHouseCreated] Error querying user's house: " + err.Error())
			isCreated = false
		}
	}
	return isCreated
}

func setCategoryCreated(db *sql.DB, houseID int) bool {
	isCreated := true
	isThereCategories := true
	query := fmt.Sprintf("SELECT (Count(*) > 0) FROM `%s`;", Tables[CategoryIdx].Key)
	err := db.QueryRow(query).Scan(&isThereCategories)
	if err != nil {
		slog.Error("[setCategoryCreated] Error querying number of categories: " + err.Error())
		isCreated = false
	}
	if isThereCategories {
		query = fmt.Sprintf("SELECT (Count(*) > 0) FROM `%s` WHERE user = %d;", Tables[CategoryIdx].Key, houseID)
		err = db.QueryRow(query).Scan(&isCreated)
		if err != nil {
			slog.Error("[setCategoryCreated] Error querying house's categories: " + err.Error())
			isCreated = false
		}
	}
	return isCreated
}
