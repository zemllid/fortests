package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// FrogHandler содержит ссылку на базу данных
type FrogHandler struct {
	DB *sql.DB
}

// NewFrogHandler возвращает новый экземпляр FrogHandler
func NewFrogHandler(db *sql.DB) *FrogHandler {
	return &FrogHandler{DB: db}
}

// FrogsHandler обрабатывает запросы по маршруту "/frogs"
// GET: получение списка всех жаб
// POST: создание новой жабы
func (h *FrogHandler) FrogsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getFrogs(w, r)
	case http.MethodPost:
		h.createFrog(w, r)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

// FrogHandler обрабатывает запросы по маршруту "/frogs/{id}"
// GET: получение информации о жабе по id
// PUT: обновление информации о жабе по id
// DELETE: удаление жабы по id
func (h *FrogHandler) FrogHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем id, удаляя префикс "/frogs/"
	idStr := strings.TrimPrefix(r.URL.Path, "/frogs/")
	if idStr == "" {
		http.Error(w, "Отсутствует ID жабы", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID жабы", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getFrogByID(w, r, id)
	case http.MethodPut:
		h.updateFrog(w, r, id)
	case http.MethodDelete:
		h.deleteFrog(w, r, id)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

// getFrogs возвращает список всех жаб
func (h *FrogHandler) getFrogs(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT id, name, species, habitat, age FROM frogs")
	if err != nil {
		http.Error(w, "Ошибка запроса к базе данных всех жаб", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var frogs []Frog
	for rows.Next() {
		var frog Frog
		if err := rows.Scan(&frog.ID, &frog.Name, &frog.Species, &frog.Habitat, &frog.Age); err != nil {
			http.Error(w, "Ошибка чтения данных из БД", http.StatusInternalServerError)
			return
		}
		frogs = append(frogs, frog)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(frogs)
}

// getFrogByID возвращает детали конкретной жабы
func (h *FrogHandler) getFrogByID(w http.ResponseWriter, r *http.Request, id int) {
	var frog Frog
	err := h.DB.QueryRow("SELECT id, name, species, habitat, age FROM frogs WHERE id = $1", id).
		Scan(&frog.ID, &frog.Name, &frog.Species, &frog.Habitat, &frog.Age)
	if err == sql.ErrNoRows {
		http.Error(w, "Жаба не найдена", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Ошибка запроса к базе данных", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(frog)
}

// createFrog создаёт новую запись о жабе
func (h *FrogHandler) createFrog(w http.ResponseWriter, r *http.Request) {
	var frog Frog
	if err := json.NewDecoder(r.Body).Decode(&frog); err != nil {
		http.Error(w, "Неверные входные данные", http.StatusBadRequest)
		return
	}

	err := h.DB.QueryRow(
		"INSERT INTO frogs (name, species, habitat, age) VALUES ($1, $2, $3, $4) RETURNING id",
		frog.Name, frog.Species, frog.Habitat, frog.Age,
	).Scan(&frog.ID)

	if err != nil {
		http.Error(w, "Ошибка записи в базу данных", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(frog)
}

// updateFrog обновляет данные о жабе по её ID
func (h *FrogHandler) updateFrog(w http.ResponseWriter, r *http.Request, id int) {
	var frog Frog
	if err := json.NewDecoder(r.Body).Decode(&frog); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}
	frog.ID = id

	result, err := h.DB.Exec(
		"UPDATE frogs SET name = $1, species = $2, habitat = $3, age = $4 WHERE id = $5",
		frog.Name, frog.Species, frog.Habitat, frog.Age, frog.ID,
	)
	if err != nil {
		http.Error(w, "Ошибка обновления записи в базе данных", http.StatusInternalServerError)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Ошибка обработки запроса", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Жаба не найдена", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(frog)
}

// deleteFrog удаляет жабу по её ID
func (h *FrogHandler) deleteFrog(w http.ResponseWriter, r *http.Request, id int) {
	result, err := h.DB.Exec("DELETE FROM frogs WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Ошибка удаления записи из базы данных", http.StatusInternalServerError)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Ошибка обработки запроса", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Жаба не найдена", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
