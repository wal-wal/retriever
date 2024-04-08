package model

import (
	"github.com/google/uuid"
	"time"
)

type Mantra struct {
	MantraId  uuid.UUID
	Content   string
	Writer    string
	CreatedAt time.Time
}

func NewMantra(MantraId uuid.UUID, Content string, Writer string, CreatedAt time.Time) *Mantra {
	return &Mantra{
		MantraId:  MantraId,
		Content:   Content,
		Writer:    Writer,
		CreatedAt: CreatedAt,
	}
}
