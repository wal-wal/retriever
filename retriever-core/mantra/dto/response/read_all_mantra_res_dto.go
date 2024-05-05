package mantra_response

import (
	"retriever-core/mantra/model"
	"time"
)

type ReadAllMantraResDTO struct {
	MantraId  string    `json:"mantra_id"`
	Content   string    `json:"content"`
	Writer    string    `json:"writer"`
	CreatedAt time.Time `json:"created_at"`
}

func NewReadAllMantraResDTO(e mantra_model.Mantra) *ReadAllMantraResDTO {
	return &ReadAllMantraResDTO{
		MantraId:  e.MantraId.String(),
		Content:   e.Content,
		Writer:    e.Writer,
		CreatedAt: e.CreatedAt,
	}
}
