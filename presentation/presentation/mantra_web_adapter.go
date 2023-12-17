package presentation

import (
	"application/domain/mantra/dto/request"
	"application/domain/mantra/dto/response"
	"application/domain/mantra/use_case"
	"github.com/google/uuid"
)

type MantraWebAdapter struct {
	readAllMantrasUseCase use_case.ReadAllMantrasUseCase
	createMantraUseCase   use_case.CreateMantraUseCase
	deleteMantraUseCase   use_case.DeleteMantraUseCase
}

func New(readAllMantrasUseCase use_case.ReadAllMantrasUseCase,
	createMantraUseCase use_case.CreateMantraUseCase,
	deleteMantraUseCase use_case.DeleteMantraUseCase) *MantraWebAdapter {
	return &MantraWebAdapter{
		readAllMantrasUseCase: readAllMantrasUseCase,
		createMantraUseCase:   createMantraUseCase,
		deleteMantraUseCase:   deleteMantraUseCase,
	}
}
func (r *MantraWebAdapter) ReadAllMantras() []response.ReadAllMantraResDTO {
	return r.readAllMantrasUseCase.Execute()
}

func (r *MantraWebAdapter) CreateMantra(dto request.CreateMantraReqDTO) error {
	return r.createMantraUseCase.Execute(dto)
}

func (r *MantraWebAdapter) DeleteMantra(id uuid.UUID) error {
	return r.deleteMantraUseCase.Execute(id)
}
