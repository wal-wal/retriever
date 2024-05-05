package mantra_use_case

import (
	"github.com/google/uuid"
	"retriever-core/mantra/dto/response"
	"retriever-core/mantra/secondary"
)

type ReadMantraUseCase struct {
	repo mantra_secondary.MantraSecondaryPort
}

func NewReadMantraUseCase(repo mantra_secondary.MantraSecondaryPort) ReadMantraUseCase {
	return ReadMantraUseCase{
		repo: repo,
	}
}

func (r *ReadMantraUseCase) Execute(mantraId uuid.UUID) (mantraResponse mantra_response.ReadMantraResDTO, err error) {
	mantra, err := r.repo.FindMantraById(mantraId)
	if err != nil {
		return mantra_response.ReadMantraResDTO{}, err
	}
	return mantraResponse.Of(mantra), nil
}
