package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Тест получения списка жаб (GET /frogs)
func TestGetFrogs(t *testing.T) {
	req := httptest.NewRequest("GET", "/frogs", nil)
	w := httptest.NewRecorder()

	frogHandler := &FrogHandler{DB: nil} // Создаём мок-обработчик без БД
	frogHandler.FrogsHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Ожидался %d, но получен %d", http.StatusOK, w.Code)
	}
}

// Тест обработки некорректного запроса (DELETE /frogs)
func TestDeleteFrogInvalid(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/frogs/abc", nil) // Некорректный ID
	w := httptest.NewRecorder()

	frogHandler := &FrogHandler{DB: nil}
	frogHandler.FrogHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Ожидался %d, но получен %d", http.StatusBadRequest, w.Code)
	}
}
