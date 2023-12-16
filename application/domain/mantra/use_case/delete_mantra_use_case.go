package use_case

import (
	"application/port/secondary_port"
	"github.com/google/uuid"
)

type DeleteMantraUseCase struct {
	repo secondary_port.MantraRepository
}

func NewDeleteMantraUseCase(repo secondary_port.MantraRepository) *DeleteMantraUseCase {
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
