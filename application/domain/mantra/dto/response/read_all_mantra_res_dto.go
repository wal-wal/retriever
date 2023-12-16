package response

import (
	"application/domain/mantra/entity"
	"time"
)

type ReadAllMantraResDTO struct {
	Id        int
	Content   string
	Writer    string
	CreatedAt time.Time
}

func Init(e entity.Mantra) *ReadAllMantraResDTO {
	return &ReadAllMantraResDTO{
		Id:        e.Id,
		Content:   e.Content,
		Writer:    e.Writer,
		CreatedAt: e.CreatedAt,
	}
}
