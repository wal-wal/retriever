package entity

import (
	"github.com/google/uuid"
	"time"
)

type Mantra struct {
	Content   string
	Writer    string
	Id        uuid.UUID
	CreatedAt time.Time
}

func NewMantra(Content string, Writer string, Id uuid.UUID, CreatedAt time.Time) *Mantra {
	return &Mantra{
		Content:   Content,
		Writer:    Writer,
		Id:        Id,
		CreatedAt: CreatedAt,
	}
}
