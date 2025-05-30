package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/lib/pq" // подключаем PostgreSQL-драйвер
)

// PgConnect устанавливает соединение с базой данных PostgreSQL и возвращает указатель на sql.DB.
// Если происходит ошибка на любом этапе, приложение завершает работу с логированием фатальной ошибки.
func PgConnect() *sql.DB {
	// Формирование строки подключения.
	// Здесь параметры (пользователь, пароль, имя БД) заданы явно и предполагается, что
	// приложение подключается к локальной базе без ssl (sslmode=disable).
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		"postgres", // имя пользователя (желательно получать из переменной окружения)
		"postgres", // пароль (также лучше передавать через переменную окружения или секреты)
		"frogs_db", // имя базы данных
	)

	// Открываем соединение с базой данных.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		// Если не удалось открыть подключение – логируем ошибку и завершаем работу.
		log.Fatalf("Ошибка при открытии БД: %v", err)
	}

	// Проверка соединения с БД посредством вызова Ping().
	if err = db.Ping(); err != nil {
		// Если не удалось установить реальное соединение – логируем и завершаем работу.
		log.Fatalf("Невозможно подключиться к базе данных: %v", err)
	}
	tableInit(db)
	return db
}

// Создаём таблицу
func tableInit(db *sql.DB) {
	file, err := os.Open("./db-init/init.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	b, err := io.ReadAll(file)
	bstring := string(b) //Преобразуем []byte в string
	_, err = db.Exec(bstring)
	if err != nil {
		log.Fatal(err)
	}
}
