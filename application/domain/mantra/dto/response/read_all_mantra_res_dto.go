package response

import (
	"application/domain/mantra/entity"
	"github.com/google/uuid"
	"time"
)

type ReadAllMantraResDTO struct {
	Id        uuid.UUID
	Content   string
	Writer    string
	CreatedAt time.Time
}

func NewReadAllMantraResDTO(e entity.Mantra) *ReadAllMantraResDTO {
	return &ReadAllMantraResDTO{
		Id:        e.Id,
		Content:   e.Content,
		Writer:    e.Writer,
		CreatedAt: e.CreatedAt,
	}
}
