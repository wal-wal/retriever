package mantra_response

import (
	"retriever-core/mantra/model"
	"time"
)

type ReadMantraResDTO struct {
	MantraId  string    `json:"mantra_id"`
	Speaker   string    `json:"speaker"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func NewReadMantraResDTO(model mantra_model.Mantra) *ReadMantraResDTO {
	return &ReadMantraResDTO{
		MantraId:  model.MantraId.String(),
		Speaker:   model.Speaker,
		Content:   model.Content,
		CreatedAt: model.CreatedAt,
	}
}

func (r *ReadMantraResDTO) Of(model mantra_model.Mantra) ReadMantraResDTO {
	return ReadMantraResDTO{
		MantraId:  model.MantraId.String(),
		Speaker:   model.Speaker,
		Content:   model.Content,
		CreatedAt: model.CreatedAt,
	}
}
