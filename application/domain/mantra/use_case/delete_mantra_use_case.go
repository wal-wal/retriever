package use_case

import (
	"application/port/secondary"
	"github.com/google/uuid"
)

type DeleteMantraUseCase struct {
	repo secondary.MantraRepository
}

func NewDeleteMantraUseCase(repo secondary.MantraRepository) *DeleteMantraUseCase {
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
