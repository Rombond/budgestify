package db_sql

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/Rombond/budgestify/api/types/accountHouse"
)

func GetAccountHouse(db *sql.DB, id int) (accountHouse.AccountHouse, error) {
	var accountHouse accountHouse.AccountHouse
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE id = ?;", Tables[AccountHouseIdx].Key)
	err := db.QueryRow(query, id).Scan(&accountHouse.Id, &accountHouse.Name, &accountHouse.Amount)
	if err != nil {
		slog.Error("[GetAccountHouse] Error querying accountHouse: " + err.Error())
	}
	return accountHouse, err
}

func CreateAccountHouse(db *sql.DB, name string, amount int) (int, error) {
	var id int64 = -1
	query := fmt.Sprintf("INSERT INTO `%s` (`name`, `amount`) VALUES (?, ?);", Tables[AccountHouseIdx].Key)
	results, err := db.Exec(query, name, amount)
	if err != nil {
		slog.Error("[CreateAccountHouse] Inserting new accountHouse: " + err.Error())
		id = -1
		return int(id), err
	}
	id, err = results.LastInsertId()
	if err != nil {
		slog.Error("[CreateAccountHouse] Error while getting lastInsertId: " + err.Error())
		id = -1
		return int(id), err
	}
	return int(id), nil
}

func ChangeAccountHouseName(db *sql.DB, accountHouseID int, newName string) (bool, error) {
	query := fmt.Sprintf("UPDATE %s SET `name` = ? WHERE id = ?", Tables[AccountHouseIdx].Key)
	_, err := db.Exec(query, newName, accountHouseID)
	if err != nil {
		slog.Error("[ChangeAccountHouseName] Error while updating accountHouse name: " + err.Error())
		return false, err
	}
	return true, nil
}

func ChangeAccountHouseAmount(db *sql.DB, accountHouseID int, newAmount int) (bool, error) {
	query := fmt.Sprintf("UPDATE %s SET `amount` = ? WHERE id = ?", Tables[AccountHouseIdx].Key)
	_, err := db.Exec(query, newAmount, accountHouseID)
	if err != nil {
		slog.Error("[ChangeAccountHouseAmount] Error while updating accountHouse amount: " + err.Error())
		return false, err
	}
	return true, nil
}

func TransactionAccountHouse(db *sql.DB, accountHouseID int, amount int) (bool, error) {
	query := fmt.Sprintf("UPDATE %s SET `amount` = `amount` + (?) WHERE id = ?", Tables[AccountHouseIdx].Key)
	_, err := db.Exec(query, amount, accountHouseID)
	if err != nil {
		slog.Error("[ChangeAccountHouseAmount] Error while updating accountHouse amount: " + err.Error())
		return false, err
	}
	return true, nil
}

func DeleteAccountHouse(db *sql.DB, accountHouseID int) (bool, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", Tables[AccountHouseIdx].Key)
	_, err := db.Exec(query, accountHouseID)
	if err != nil {
		slog.Error("[DeleteAccountHouse] Error while deleting accountHouse: " + err.Error())
		return false, err
	}
	return true, nil
}

func GetHouseIDFromAccount(db *sql.DB, accountHouseID int) (int, error) {
	id := -1
	query := fmt.Sprintf("SELECT id FROM `%s` WHERE account = ?;", Tables[HouseIdx].Key)
	err := db.QueryRow(query, accountHouseID).Scan(&id)
	if err != nil {
		slog.Error("[GetHouseIDFromAccount] Error querying house: " + err.Error())
	}
	return id, err
}

func DoesUserCanEditAccount(db *sql.DB, accountHouseID int, userID int) (bool, error) {
	canEdit := false
	query := fmt.Sprintf("SELECT hu.admin FROM `%s` h JOIN `%s` ah ON h.account = ah.id JOIN `%s` hu ON hu.house = h.id WHERE ah.id = ? AND hu.user = ?;", Tables[HouseIdx].Key, Tables[AccountHouseIdx].Key, Tables[House_UserIdx].Key)
	err := db.QueryRow(query, accountHouseID, userID).Scan(&canEdit)
	if err != nil {
		slog.Error("[GetHouseIDFromAccount] Error querying house: " + err.Error())
	}
	return canEdit, nil
}
