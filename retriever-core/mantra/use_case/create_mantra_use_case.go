package mantra_use_case

import (
	"github.com/google/uuid"
	"retriever-core/mantra/dto/request"
	"retriever-core/mantra/model"
	"retriever-core/mantra/secondary"
	"time"
)

type CreateMantraUseCase struct {
	repo mantra_secondary.MantraSecondaryPort
}

func NewCreateMantraUseCase(repo mantra_secondary.MantraSecondaryPort) CreateMantraUseCase {
	return CreateMantraUseCase{
		repo: repo,
	}
}

func (r *CreateMantraUseCase) Execute(dto mantra_request.CreateMantraReqDTO) error {
	mantra := mantra_model.NewMantra(uuid.New(), dto.Content, dto.Writer, time.Now())
	err := r.repo.CreateMantra(*mantra)
	if err != nil {
		return err
	}
	return nil
}
