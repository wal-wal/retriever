package response

import "retriever/application/domain/mantra/entity"

type ReadAllMantraResDTO struct {
	Id      int
	Content string
	Writer  string
}

func Init(e entity.Mantra) *ReadAllMantraResDTO {
	return &ReadAllMantraResDTO{
		Id:      e.Id,
		Content: e.Content,
		Writer:  e.Writer,
	}
}
