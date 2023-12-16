package secondary_port

import (
	"application/domain/mantra/entity"
	"github.com/google/uuid"
)

type MantraRepository interface {
	ReadAllMantras() []entity.Mantra
	CreateMantra(mantra entity.Mantra) error
	DeleteMantra(id uuid.UUID) error
}
