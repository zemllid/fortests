package main

// Frog описывает сущность "жаба" в системе.
type Frog struct {
	ID      int    `json:"id"`      // Уникальный идентификатор
	Name    string `json:"name"`    // Имя жабы
	Species string `json:"species"` // Вид жабы
	Habitat string `json:"habitat"` // Ареал обитания
	Age     int    `json:"age"`     // Возраст жабы
}
