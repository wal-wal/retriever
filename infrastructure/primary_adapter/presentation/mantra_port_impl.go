package presentation

import (
	"retriever/application/domain/mantra/dto/response"
	"retriever/application/domain/mantra/use_case"
)

type MantraPortImpl struct {
	readAllMantrasUseCase use_case.ReadAllMantrasUseCase
}

func New(useCase use_case.ReadAllMantrasUseCase) *MantraPortImpl {
	return &MantraPortImpl{
		readAllMantrasUseCase: useCase,
	}
}
func (r *MantraPortImpl) ReadAllMantras() []response.ReadAllMantraResDTO {
	return r.readAllMantrasUseCase.Execute()
}