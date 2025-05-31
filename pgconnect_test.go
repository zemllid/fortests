package main

import (
	"database/sql"
	"testing"
)

// Тест подключения к базе (мок-драйвер)
func TestPgConnect(t *testing.T) {
	mockDB, _ := sql.Open("postgres", "user=mock dbname=testdb sslmode=disable")

	if err := mockDB.Ping(); err != nil {
		t.Errorf("Ошибка подключения к БД: %v", err)
	}
}
