package db_sql

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/Rombond/budgestify/internal/sql/tables/house"
	"github.com/Rombond/budgestify/internal/sql/tables/user"
)

type StateSetup struct {
	IsDbInitialized   bool
	IsUserCreated     bool
	IsHouseCreated    bool
	IsCategoryCreated bool
}

var state StateSetup

func GetStateSetup(db *sql.DB, login string) StateSetup {
	state.IsDbInitialized = setDbInitialized(db)
	state.IsUserCreated = setUserCreated(db)
	state.IsHouseCreated = false
	state.IsCategoryCreated = false
	if login == "" {
		return state
	}
	userID := user.GetUserID(login)
	if userID == -1 {
		return state
	}
	state.IsHouseCreated = setHouseCreated(db, userID)
	houseID := house.GetHouseID(userID)
	if houseID == -1 {
		return state
	}
	state.IsCategoryCreated = setCategoryCreated(db, houseID)
	return state
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
	query := fmt.Sprintf("SELECT (Count(*) > 0) FROM `%s`;", Tables[User].Key)
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
	query := fmt.Sprintf("SELECT (Count(*) > 0) FROM `%s`;", Tables[House].Key)
	err := db.QueryRow(query).Scan(&isThereHouses)
	if err != nil {
		slog.Error("[setHouseCreated] Error querying number of houses: " + err.Error())
		isCreated = false
	}
	if isThereHouses {
		query = fmt.Sprintf("SELECT (Count(*) > 0) FROM `%s` WHERE user = %d;", Tables[House_User].Key, userID)
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
	query := fmt.Sprintf("SELECT (Count(*) > 0) FROM `%s`;", Tables[Category].Key)
	err := db.QueryRow(query).Scan(&isThereCategories)
	if err != nil {
		slog.Error("[setCategoryCreated] Error querying number of categories: " + err.Error())
		isCreated = false
	}
	if isThereCategories {
		query = fmt.Sprintf("SELECT (Count(*) > 0) FROM `%s` WHERE user = %d;", Tables[Category].Key, houseID)
		err = db.QueryRow(query).Scan(&isCreated)
		if err != nil {
			slog.Error("[setCategoryCreated] Error querying house's categories: " + err.Error())
			isCreated = false
		}
	}
	return isCreated
}
