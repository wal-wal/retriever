package presentation

import (
	"application/domain/mantra/dto/request"
	"application/domain/mantra/dto/response"
	"application/domain/mantra/use_case"
)

type MantraPortImpl struct {
	readAllMantrasUseCase use_case.ReadAllMantrasUseCase
	createMantraUseCase   use_case.CreateMantraUseCase
}

func New(ReadAllMantrasUseCase use_case.ReadAllMantrasUseCase,
	CreateMantraUseCase use_case.CreateMantraUseCase) *MantraPortImpl {
	return &MantraPortImpl{
		readAllMantrasUseCase: ReadAllMantrasUseCase,
		createMantraUseCase:   CreateMantraUseCase,
	}
}
func (r *MantraPortImpl) ReadAllMantras() []response.ReadAllMantraResDTO {
	return r.readAllMantrasUseCase.Execute()
}

func (r *MantraPortImpl) CreateMantra(dto request.CreateMantraReqDTO) error {
	return r.createMantraUseCase.Execute(dto)
}
