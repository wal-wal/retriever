package use_case

import (
	"application/domain/mantra/dto/request"
	"application/domain/mantra/dto/response"
	"github.com/google/uuid"
)

type MantraUseCase struct {
	readAllMantraUC ReadAllMantrasUseCase
	createMantraUC  CreateMantraUseCase
	deleteMantraUC  DeleteMantraUseCase
}

func NewMantraUseCase(readAllMantraUseCase ReadAllMantrasUseCase, createMantraUseCase CreateMantraUseCase, deleteMantraUseCase DeleteMantraUseCase) *MantraUseCase {
	return &MantraUseCase{
		readAllMantraUC: readAllMantraUseCase,
		createMantraUC:  createMantraUseCase,
		deleteMantraUC:  deleteMantraUseCase,
	}
}

func (r *MantraUseCase) CreateMantra(dto request.CreateMantraReqDTO) error {
	return r.createMantraUC.Execute(dto)
}

func (r *MantraUseCase) ReadAllMantras() []response.ReadAllMantraResDTO {
	return r.readAllMantraUC.Execute()
}

func (r *MantraUseCase) DeleteMantra(id uuid.UUID) error {
	return r.deleteMantraUC.Execute(id)
}
