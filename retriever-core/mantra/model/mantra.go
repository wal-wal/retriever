package mantra_model

import (
	"github.com/google/uuid"
	"time"
)

type Mantra struct {
	MantraId  uuid.UUID
	Speaker   string
	Content   string
	CreatedAt time.Time
	Writer    string
}

func NewMantra(mantraId uuid.UUID, speaker string, content string, createdAt time.Time, writer string) *Mantra {
	return &Mantra{
		MantraId:  mantraId,
		Speaker:   speaker,
		Content:   content,
		CreatedAt: createdAt,
		Writer:    writer,
	}
}
