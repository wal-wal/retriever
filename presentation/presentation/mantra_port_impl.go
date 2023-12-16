package presentation

import (
	"application/domain/mantra/dto/request"
	"application/domain/mantra/dto/response"
	"application/domain/mantra/use_case"
	"github.com/google/uuid"
)

type MantraPortImpl struct {
	readAllMantrasUseCase use_case.ReadAllMantrasUseCase
	createMantraUseCase   use_case.CreateMantraUseCase
	deleteMantraUseCase   use_case.DeleteMantraUseCase
}

func New(readAllMantrasUseCase use_case.ReadAllMantrasUseCase,
	createMantraUseCase use_case.CreateMantraUseCase,
	deleteMantraUseCase use_case.DeleteMantraUseCase) *MantraPortImpl {
	return &MantraPortImpl{
		readAllMantrasUseCase: readAllMantrasUseCase,
		createMantraUseCase:   createMantraUseCase,
		deleteMantraUseCase:   deleteMantraUseCase,
	}
}
func (r *MantraPortImpl) ReadAllMantras() []response.ReadAllMantraResDTO {
	return r.readAllMantrasUseCase.Execute()
}

func (r *MantraPortImpl) CreateMantra(dto request.CreateMantraReqDTO) error {
	return r.createMantraUseCase.Execute(dto)
}

func (r *MantraPortImpl) DeleteMantra(id uuid.UUID) error {
	return r.deleteMantraUseCase.Execute(id)
}
