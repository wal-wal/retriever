package use_case

import (
	"application/domain/mantra/dto/response"
	"application/port/secondary_port"
	"sort"
)

type ReadAllMantrasUseCase struct {
	repo secondary_port.MantraRepository
}

func NewReadAllMantrasUseCase(repo secondary_port.MantraRepository) *ReadAllMantrasUseCase {
	return &ReadAllMantrasUseCase{
		repo: repo,
	}
}

func (r *ReadAllMantrasUseCase) Execute() (list []response.ReadAllMantraResDTO) {
	mantras := r.repo.ReadAllMantras()
	for _, v := range mantras {
		mantra := response.NewReadAllMantraResDTO(v)
		list = append(list, *mantra)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].CreatedAt.After(list[j].CreatedAt)
	})
	return list
}
