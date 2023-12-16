package primary_port

import (
	"application/domain/mantra/dto/request"
	"application/domain/mantra/dto/response"
	"github.com/google/uuid"
)

type MantraPort interface {
	ReadAllMantras() []response.ReadAllMantraResDTO
	CreateMantra(dto request.CreateMantraReqDTO) error
	DeleteMantra(id uuid.UUID) error
}
