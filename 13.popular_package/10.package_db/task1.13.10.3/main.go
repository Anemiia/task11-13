package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	UserID int    `json:"user_id"`
}

// Создание таблицы пользователей
func CreateUserTable() error {
	db, _ := sql.Open("sqlite3", "users.db")
	defer db.Close()

	query := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL
	);`
	db.Exec(query)

	query = `CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		text TEXT NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`
	db.Exec(query)

	return nil
}

// Вставка пользователя в таблицу
func InsertUser(user User) error {
	db, _ := sql.Open("sqlite3", "users.db")
	defer db.Close()

	query := `INSERT INTO users (name, age) VALUES (?, ?);`
	_, err := db.Exec(query, user.Name, user.Age)
	if err != nil {
		return fmt.Errorf("ошибка вставки пользователя: %w", err)
	}

	return nil
}

// Выборка пользователя из таблицы
func SelectUser(userID int) (User, error) {
	db, _ := sql.Open("sqlite3", "users.db")
	defer db.Close()

	query := `SELECT id, name, age FROM users WHERE id = ?;`
	row := db.QueryRow(query, userID)

	var user User
	row.Scan(&user.ID, &user.Name, &user.Age)


	query = `SELECT id, text, user_id FROM comments WHERE user_id = ?;`
	rows, _ := db.Query(query, userID)

	defer rows.Close()

	for rows.Next() {
		var comment Comment
		rows.Scan(&comment.ID, &comment.Text, &comment.UserID)

		user.Comments = append(user.Comments, comment)
	}

	return user, nil
}

// Обновление информации о пользователе
func UpdateUser(user User) error {
	db, _ := sql.Open("sqlite3", "users.db")
	defer db.Close()

	query := `UPDATE users SET name = ?, age = ? WHERE id = ?;`
	db.Exec(query, user.Name, user.Age, user.ID)

	return nil
}

// Удаление пользователя из таблицы
func DeleteUser(userID int) error {
	db, _ := sql.Open("sqlite3", "users.db")
	defer db.Close()

	query := `DELETE FROM users WHERE id = ?;`
	db.Exec(query, userID)

	query = `DELETE FROM comments WHERE user_id = ?;`
	db.Exec(query, userID)

	return nil
}
