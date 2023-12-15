package primary_port

import "retriever/application/domain/mantra/dto/response"

type MantraPort interface {
	ReadAllMantras() []response.ReadAllMantraResDTO
}
