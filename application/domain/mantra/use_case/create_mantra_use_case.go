package use_case

import (
	"application/domain/mantra/dto/request"
	"application/port/secondary_port"
	"math/rand"
	"time"
)

type CreateMantraUseCase struct {
	repo secondary_port.MantraRepository
}

func NewCreateMantraUseCase(repo secondary_port.MantraRepository) *CreateMantraUseCase {
	return &CreateMantraUseCase{
		repo: repo,
	}
}

func (r *CreateMantraUseCase) Execute(dto request.CreateMantraReqDTO) error {
	err := r.repo.CreateMantra(dto.ToEntity(rand.Int(), time.Now()))
	if err != nil {
		return err
	}
	return nil
}
