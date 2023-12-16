package primary_port

import (
	"application/domain/mantra/dto/request"
	"application/domain/mantra/dto/response"
)

type MantraPort interface {
	ReadAllMantras() []response.ReadAllMantraResDTO
	CreateMantra(dto request.CreateMantraReqDTO) error
}
