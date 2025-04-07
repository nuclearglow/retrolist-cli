package retrolist

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	UUID      string    `json:"uuid"`
	Title     string    `json:"title"`
	Quantity  uint      `json:"Quantity"`
	Done      bool      `json:"Done"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func NewItem(title string, quantity uint) *Item {
	if quantity < 1 {
		quantity = 1
	}

	return &Item{
		UUID:      uuid.New().String(),
		Title:     title,
		Quantity:  quantity,
		Done:      false,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
