package mantra_use_case

import (
	"github.com/google/uuid"
	"retriever-core/mantra/dto/request"
	"retriever-core/mantra/dto/response"
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

func (r *MantraUseCase) CreateMantra(dto mantra_request.CreateMantraReqDTO) error {
	return r.createMantraUC.Execute(dto)
}

func (r *MantraUseCase) ReadAllMantras() ([]mantra_response.ReadAllMantraResDTO, error) {
	return r.readAllMantraUC.Execute()
}

func (r *MantraUseCase) DeleteMantra(id uuid.UUID) error {
	return r.deleteMantraUC.Execute(id)
}
