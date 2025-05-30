package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"strconv"

	_ "github.com/lib/pq"
)

func main() {
	dbHost := getEnv("DB_HOST", "postgresf")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "frogs_db")
	
	delaySec, errD:= strconv.ParseInt(getEnv("DB_DELAY", "5"), 10, 64)
if errD != nil {
    panic(errD)
}
	log.Printf("wait for %v sec", delaySec)
	time.Sleep(time.Duration(delaySec) * time.Second)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
		
    log.Printf("test connect %v",psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

    log.Printf("success connect %v",psqlInfo)

	// Проверка соединения с БД
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		log.Fatalf("Ошибка ping БД: %v ", err)
	}
	log.Println("Подключение к БД успешно установлено")

	frogHandler := NewFrogHandler(db)

	http.HandleFunc("/frogs", frogHandler.FrogsHandler)
	http.HandleFunc("/frogs/", frogHandler.FrogHandler)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	serverAddr := ":8080"
	log.Printf("Сервер запущен на %s", serverAddr)
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Fatalf("Ошибка HTTP-сервера: %v", err)
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
