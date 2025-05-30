package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq" // драйвер PostgreSQL
)

func main() {
	// Чтение настроек подключения из переменных окружения
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "frogs_db")

	// Формирование строки подключения
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Подключение к базе данных PostgreSQL
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	// Проверка соединения с БД
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		log.Fatalf("Ошибка ping БД: %v", err)
	}
	log.Println("Подключение к БД успешно установлено")

	// Создаём экземпляр обработчика напрямую
	frogHandler := NewFrogHandler(db)

	// Регистрируем обработчики
	http.HandleFunc("/frogs", frogHandler.FrogsHandler)
	http.HandleFunc("/frogs/", frogHandler.FrogHandler)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	serverAddr := ":8080"
	log.Printf("Сервер запущен на %s", serverAddr)
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Fatalf("Ошибка HTTP-сервера: %v", err)
	}
}

// getEnv возвращает значение переменной окружения или значение по умолчанию
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
