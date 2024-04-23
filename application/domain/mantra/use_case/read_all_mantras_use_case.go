package use_case

import (
	"application/domain/mantra/dto/response"
	"application/port/secondary"
	"sort"
)

type ReadAllMantrasUseCase struct {
	repo secondary.MantraSecondaryPort
}

func NewReadAllMantrasUseCase(repo secondary.MantraSecondaryPort) *ReadAllMantrasUseCase {
	return &ReadAllMantrasUseCase{
		repo: repo,
	}
}

func (r *ReadAllMantrasUseCase) Execute() (list []response.ReadAllMantraResDTO, err error) {
	mantras, err := r.repo.ReadAllMantras()
	if err != nil {
		return nil, err
	}
	if mantras == nil {
		return []response.ReadAllMantraResDTO{}, err
	}
	for _, v := range mantras {
		mantra := response.NewReadAllMantraResDTO(v)
		list = append(list, *mantra)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].CreatedAt.After(list[j].CreatedAt)
	})
	return list, nil
}
