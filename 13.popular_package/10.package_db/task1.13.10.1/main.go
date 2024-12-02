package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // SQLite драйвер
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL
	);`
	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("ошибка создания таблицы: %w", err)
	}
	return nil
}

func InsertUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query := `INSERT INTO users (name, age) VALUES (?, ?)`
	_, err = db.Exec(query, user.Name, user.Age)
	if err != nil {
		return fmt.Errorf("ошибка добавления пользователя: %w", err)
	}
	return nil
}

func SelectUser(id int) (User, error) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	var user User
	query := `SELECT id, name, age FROM users WHERE id = ?`
	row := db.QueryRow(query, id)
	err = row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("пользователь с ID %d не найден", id)
		}
		return User{}, fmt.Errorf("ошибка выборки пользователя: %w", err)
	}
	return user, nil
}

func UpdateUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query := `UPDATE users SET name = ?, age = ? WHERE id = ?`
	result, err := db.Exec(query, user.Name, user.Age, user.ID)
	if err != nil {
		return fmt.Errorf("ошибка обновления пользователя: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("пользователь с ID %d не найден", user.ID)
	}
	return nil
}

func DeleteUser(id int) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query := `DELETE FROM users WHERE id = ?`
	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка удаления пользователя: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("пользователь с ID %d не найден", id)
	}
	return nil
}

