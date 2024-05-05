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
	readMantraUC    ReadMantraUseCase
}

func NewMantraUseCase(readAllMantraUseCase ReadAllMantrasUseCase,
	createMantraUseCase CreateMantraUseCase,
	deleteMantraUseCase DeleteMantraUseCase,
	readMantraUseCase ReadMantraUseCase) *MantraUseCase {
	return &MantraUseCase{
		readAllMantraUC: readAllMantraUseCase,
		createMantraUC:  createMantraUseCase,
		deleteMantraUC:  deleteMantraUseCase,
		readMantraUC:    readMantraUseCase,
	}
}

func (r *MantraUseCase) CreateMantra(dto mantra_request.CreateMantraReqDTO) error {
	return r.createMantraUC.Execute(dto)
}

func (r *MantraUseCase) ReadAllMantras() ([]mantra_response.ReadMantraResDTO, error) {
	return r.readAllMantraUC.Execute()
}

func (r *MantraUseCase) DeleteMantra(id uuid.UUID) error {
	return r.deleteMantraUC.Execute(id)
}

func (r *MantraUseCase) ReadMantra(mantraId uuid.UUID) (mantra_response.ReadMantraResDTO, error) {
	return r.readMantraUC.Execute(mantraId)
}
