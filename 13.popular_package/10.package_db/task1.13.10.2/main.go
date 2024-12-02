package main

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int
	Username string
	Email    string
}

// Создание таблицы пользователей
func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL
	);`
	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("ошибка создания таблицы пользователей: %w", err)
	}

	return nil
}

// Вставка пользователя в таблицу
func InsertUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query, args, err := PrepareQuery("insert", "users", user)
	if err != nil {
		return err
	}

	_, err = db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

// Выборка пользователя из таблицы
func SelectUser(userID int) (User, error) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	query, args, err := PrepareQuery("select", "users", User{ID: userID})
	if err != nil {
		return User{}, err
	}

	row := db.QueryRow(query, args...)

	var user User
	err = row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return user, fmt.Errorf("ошибка выборки пользователя: %w", err)
	}

	return user, nil
}

// Обновление информации о пользователе
func UpdateUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query, args, err := PrepareQuery("update", "users", user)
	if err != nil {
		return err
	}

	_, err = db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

// Удаление пользователя из таблицы
func DeleteUser(userID int) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query, args, err := PrepareQuery("delete", "users", User{ID: userID})
	if err != nil {
		return err
	}

	_, err = db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

// Функция для подготовки запроса
func PrepareQuery(operation string, table string, user User) (string, []interface{}, error) {
    // Создаем новый билдер запросов с использованием squirrel и устанавливаем формат placeholder'ов
    psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question)

    switch operation {
    case "insert":
        // Создаем запрос на вставку данных в таблицу
        query, args, err := psql.Insert(table).
            Columns("username", "email"). // Указываем столбцы для вставки
            Values(user.Username, user.Email). // Указываем значения для вставки
            ToSql() // Преобразуем запрос в SQL строку и аргументы
        return query, args, err
    case "select":
        // Создаем запрос на выборку данных из таблицы
        query, args, err := psql.Select("*"). // Указываем, что выбираем все столбцы
            From(table). // Указываем таблицу
            Where(squirrel.Eq{"id": user.ID}). // Указываем условие выборки
            ToSql() // Преобразуем запрос в SQL строку и аргументы
        return query, args, err
    case "update":
        // Создаем запрос на обновление данных в таблице
        query, args, err := psql.Update(table).
            Set("username", user.Username). // Устанавливаем новое значение для столбца username
            Set("email", user.Email). // Устанавливаем новое значение для столбца email
            Where(squirrel.Eq{"id": user.ID}). // Указываем условие обновления
            ToSql() // Преобразуем запрос в SQL строку и аргументы
        return query, args, err
    case "delete":
        // Создаем запрос на удаление данных из таблицы
        query, args, err := psql.Delete(table).
            Where(squirrel.Eq{"id": user.ID}). // Указываем условие удаления
            ToSql() // Преобразуем запрос в SQL строку и аргументы
        return query, args, err
    default:
        return "", nil, fmt.Errorf("неверная операция")
    }

}