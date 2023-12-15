package secondary_port

import "retriever/application_regercy/domain/mantra/entity"

type MantraRepository interface {
	ReadAllMantras() []entity.Mantra
}
