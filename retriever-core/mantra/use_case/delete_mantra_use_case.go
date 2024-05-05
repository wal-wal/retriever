package mantra_use_case

import (
	"github.com/google/uuid"
	"retriever-core/mantra/secondary"
)

type DeleteMantraUseCase struct {
	repo mantra_secondary.MantraSecondaryPort
}

func NewDeleteMantraUseCase(repo mantra_secondary.MantraSecondaryPort) DeleteMantraUseCase {
	return DeleteMantraUseCase{
		repo: repo,
	}
}

func (r *DeleteMantraUseCase) Execute(id uuid.UUID) error {
	err := r.repo.DeleteMantra(id)
	if err != nil {
		return err
	}
	return nil
}
