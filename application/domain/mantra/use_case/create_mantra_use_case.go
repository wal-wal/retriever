package use_case

import (
	"application/domain/mantra/dto/request"
	"application/domain/mantra/model"
	"application/port/secondary"
	"github.com/google/uuid"
	"time"
)

type CreateMantraUseCase struct {
	repo secondary.MantraSecondaryPort
}

func NewCreateMantraUseCase(repo secondary.MantraSecondaryPort) *CreateMantraUseCase {
	return &CreateMantraUseCase{
		repo: repo,
	}
}

func (r *CreateMantraUseCase) Execute(dto request.CreateMantraReqDTO) error {
	mantra := model.NewMantra(uuid.New(), dto.Content, dto.Writer, time.Now())
	err := r.repo.CreateMantra(*mantra)
	if err != nil {
		return err
	}
	return nil
}
