package mantra_use_case

import (
	"math/rand"
	"retriever-core/mantra/dto/response"
	"retriever-core/mantra/secondary"
	"time"
)

type ReadMantraByRandomUseCase struct {
	repo mantra_secondary.MantraSecondaryPort
}

func NewReadMantraByRandomUseCase(repo mantra_secondary.MantraSecondaryPort) ReadMantraByRandomUseCase {
	return ReadMantraByRandomUseCase{
		repo: repo,
	}
}

func (r *ReadMantraByRandomUseCase) Execute() (mantra_response.ReadMantraResDTO, error) {
	tblSize, err := r.repo.GetTableSize()
	if err != nil {
		return mantra_response.ReadMantraResDTO{}, err
	}
	rand.Seed(stringToInt64(time.Now().Format("20061206")))
	mantra, err := r.repo.FindMantraByLimitAndOffset(rand.Intn(tblSize))
	if err != nil {
		return mantra_response.ReadMantraResDTO{}, err
	}
	return *mantra_response.NewReadMantraResDTO(mantra), err
}

func stringToInt64(str string) int64 {
	var a int64 = 0
	for _, v := range str {
		a = a*10 + int64(v)
	}
	return a
}
