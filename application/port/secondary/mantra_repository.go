package secondary

import (
	"application/domain/mantra/entity"
	"github.com/google/uuid"
)

type MantraSecondaryPort interface {
	ReadAllMantras() []entity.Mantra
	CreateMantra(mantra entity.Mantra) error
	DeleteMantra(id uuid.UUID) error
}
