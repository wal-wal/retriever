package use_case

import (
	"application/port/secondary"
	"github.com/google/uuid"
)

type DeleteMantraUseCase struct {
	repo secondary.MantraSecondaryPort
}

func NewDeleteMantraUseCase(repo secondary.MantraSecondaryPort) *DeleteMantraUseCase {
	return &DeleteMantraUseCase{
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
