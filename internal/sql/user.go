package db_sql

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Rombond/budgestify/api/types/user"
)

func GetUser(db *sql.DB, id int) (user.User, error) {
	var user user.User
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE id = ?;", Tables[UserIdx].Key)
	err := db.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Login, &user.Hash)
	if err != nil {
		slog.Error("[GetUser] Error querying user: " + err.Error())
	}
	return user, err
}

func GetUserID(db *sql.DB, login string) (int, error) {
	id := -1
	query := fmt.Sprintf("SELECT `id` FROM `%s` WHERE login = ?;", Tables[UserIdx].Key)
	err := db.QueryRow(query, login).Scan(&id)
	if err != nil {
		slog.Error("[GetUserID] Error querying user id: " + err.Error())
		id = -1
		return id, err
	}
	return id, nil
}

func CreateUser(db *sql.DB, name string, login string, hash []byte) (int, error) {
	var id int64 = -1
	idLogin, err := GetUserID(db, login)
	if idLogin != -1 {
		slog.Error("[CreateUser] Error user already exist")
		return int(id), errors.New("Login already taken")
	}
	query := fmt.Sprintf("INSERT INTO `%s` (`name`, `login`, `hash`) VALUES (?, ?, ?);", Tables[UserIdx].Key)
	results, err := db.Exec(query, name, login, hash)
	if err != nil {
		slog.Error("[CreateUser] Error while inserting new user: " + err.Error())
		return int(id), err
	}
	id, err = results.LastInsertId()
	if err != nil {
		slog.Error("[CreateUser] Error while getting lastInsertId: " + err.Error())
		id = -1
		return int(id), err
	}
	UpdateStateSetup(db, login)
	return int(id), nil
}

func ChangeUserHash(db *sql.DB, userID int, hash []byte) (bool, error) {
	query := fmt.Sprintf("UPDATE %s SET `hash` = ? WHERE id = ?", Tables[UserIdx].Key)
	_, err := db.Exec(query, hash, userID)
	if err != nil {
		slog.Error("[ChangeUserHash] Error while updating user: " + err.Error())
		return false, err
	}
	return true, nil
}

func ChangeUserName(db *sql.DB, userID int, name string) (bool, error) {
	query := fmt.Sprintf("UPDATE %s SET `name` = ? WHERE id = ?", Tables[UserIdx].Key)
	_, err := db.Exec(query, name, userID)
	if err != nil {
		slog.Error("[ChangeUserName] Error while updating user: " + err.Error())
		return false, err
	}
	return true, nil
}
