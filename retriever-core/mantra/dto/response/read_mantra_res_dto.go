package mantra_response

import (
	"retriever-core/mantra/model"
	"time"
)

type ReadMantraResDTO struct {
	MantraId  string    `json:"mantra_id"`
	Content   string    `json:"content"`
	Writer    string    `json:"writer"`
	CreatedAt time.Time `json:"created_at"`
}

func NewReadMantraResDTO(e mantra_model.Mantra) *ReadMantraResDTO {
	return &ReadMantraResDTO{
		MantraId:  e.MantraId.String(),
		Content:   e.Content,
		Writer:    e.Writer,
		CreatedAt: e.CreatedAt,
	}
}

func (r *ReadMantraResDTO) Of(mantra mantra_model.Mantra) ReadMantraResDTO {
	return ReadMantraResDTO{
		MantraId:  mantra.MantraId.String(),
		Content:   mantra.Content,
		Writer:    mantra.Writer,
		CreatedAt: mantra.CreatedAt,
	}
}
