package mantra_primary

import (
	"github.com/google/uuid"
	"retriever-core/mantra/dto/request"
	"retriever-core/mantra/dto/response"
)

type MantraPrimaryPort interface {
	ReadAllMantras() ([]mantra_response.ReadAllMantraResDTO, error)
	CreateMantra(dto mantra_request.CreateMantraReqDTO) error
	DeleteMantra(id uuid.UUID) error
}
