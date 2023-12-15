package secondary_port

import "application/domain/mantra/entity"

type MantraRepository interface {
	ReadAllMantras() []entity.Mantra
}
