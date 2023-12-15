package primary_port

import "retriever/application_regercy/domain/mantra/dto/response"

type MantraPort interface {
	ReadAllMantras() []response.ReadAllMantraResDTO
}
