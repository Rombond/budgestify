package db_sql

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/Rombond/budgestify/api/types/category"
)

func GetCategory(db *sql.DB, id int) (category.Category, error) {
	var category category.Category
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE id = ?;", Tables[CategoryIdx].Key)
	err := db.QueryRow(query, id).Scan(&category.Id, &category.Name, &category.Icon, &category.Parent, &category.House)
	if err != nil {
		slog.Error("[GetCategory] Error querying category: " + err.Error())
	}
	return category, err
}

func GetCategories(db *sql.DB, houseID int) ([]category.Category, error) {
	var categories []category.Category
	var category category.Category
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE house = ?;", Tables[CategoryIdx].Key)
	rows, err := db.Query(query, houseID)
	if err != nil {
		slog.Error("[GetCategories] Error querying categories: " + err.Error())
		return categories, err
	}
	for rows.Next() {
		rows.Scan(&category.Id, &category.Name, &category.Icon, &category.Parent, &category.House)
		categories = append(categories, category)
	}
	if err = rows.Err(); err != nil {
		return categories, err
	}
	return categories, err
}

func CreateCategory(db *sql.DB, name string, icon string, parent int, houseID int) (int, error) {
	properties := "(`name`"
	placeholders := "?"
	args := []interface{}{name}

	if icon != "" {
		properties += ", `icon`"
		placeholders += ", ?"
		args = append(args, icon)
	}
	if parent > 0 {
		properties += ", `parent`"
		placeholders += ", ?"
		args = append(args, parent)
	}
	properties += ", `house`)"
	placeholders += ", ?"
	args = append(args, houseID)

	query := fmt.Sprintf("INSERT INTO `%s` %s VALUES (%s);", Tables[CategoryIdx].Key, properties, placeholders)

	results, err := db.Exec(query, args...)
	if err != nil {
		slog.Error("[CreateCategory] Error while inserting new category: " + err.Error())
		return -1, err
	}

	id, err := results.LastInsertId()
	if err != nil {
		slog.Error("[CreateCategory] Error while getting lastInsertId: " + err.Error())
		return -1, err
	}
	return int(id), nil
}

func ChangeCategory(db *sql.DB, id int, name string, icon string, parent int) (bool, error) {
	params := ""
	args := []interface{}{id}

	if name != "" {
		params += "`name` = ? "
		args = append(args, name)
	}
	if icon != "" {
		params += "`icon` = ? "
		args = append(args, icon)
	}
	if parent > 0 {
		params += "`parent` = ? "
		args = append(args, parent)
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", Tables[CategoryIdx].Key, params)
	_, err := db.Exec(query, args...)
	if err != nil {
		slog.Error("[ChangeCategory] Error while updating category: " + err.Error())
		return false, err
	}
	return true, nil
}

func DeleteCategory(db *sql.DB, id int) (bool, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", Tables[CategoryIdx].Key)
	_, err := db.Exec(query, id)
	if err != nil {
		slog.Error("[DeleteCategory] Error while deleting category: " + err.Error())
		return false, err
	}
	return true, nil
}
