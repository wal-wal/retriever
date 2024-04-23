package secondary

import (
	"application/domain/mantra/model"
	"github.com/google/uuid"
)

type MantraSecondaryPort interface {
	ReadAllMantras() ([]model.Mantra, error)
	CreateMantra(mantra model.Mantra) error
	DeleteMantra(MantraId uuid.UUID) error
}
