package use_case

import (
	"application/domain/mantra/dto/request"
	"application/domain/mantra/entity"
	"application/port/secondary_port"
	"github.com/google/uuid"
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
	mantra := entity.NewMantra(dto.Content, dto.Writer, uuid.New(), time.Now())
	err := r.repo.CreateMantra(*mantra)
	if err != nil {
		return err
	}
	return nil
}
