package request

import (
	"application/domain/mantra/entity"
	"time"
)

type CreateMantraReqDTO struct {
	Content string `json:"content"`
	Writer  string `json:"writer"`
}

func (c CreateMantraReqDTO) ToEntity(Id int, CreatedAt time.Time) entity.Mantra {
	return entity.Mantra{
		Id:        Id,
		Content:   c.Content,
		Writer:    c.Writer,
		CreatedAt: CreatedAt,
	}
}
