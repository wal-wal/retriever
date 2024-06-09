package mantra_secondary

import (
	"github.com/google/uuid"
	"retriever-core/mantra/model"
)

type MantraSecondaryPort interface {
	ReadAllMantras() ([]mantra_model.Mantra, error)

	CreateMantra(mantra mantra_model.Mantra) error

	DeleteMantra(MantraId uuid.UUID) error

	FindMantraById(mantraId uuid.UUID) (mantra_model.Mantra, error)

	FindMantrasBySpeaker(speakerName string) (list []mantra_model.Mantra, err error)
	FindMantraByLimitAndOffset(limit int) (mantra_model.Mantra, error)
	GetTableSize() (size int, err error)
}
