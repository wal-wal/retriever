package secondary_port

import "retriever/application/domain/mantra/entity"

type MantraRepository interface {
	ReadAllMantras() []entity.Mantra
}
