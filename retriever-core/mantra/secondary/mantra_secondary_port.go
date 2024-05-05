package mantra_secondary

import (
	"github.com/google/uuid"
	"retriever-core/mantra/model"
)

type MantraSecondaryPort interface {
	ReadAllMantras() ([]mantra_model.Mantra, error)
	CreateMantra(mantra mantra_model.Mantra) error
	DeleteMantra(MantraId uuid.UUID) error
}
