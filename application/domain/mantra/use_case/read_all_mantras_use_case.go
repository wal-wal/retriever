package use_case

import (
	"application/domain/mantra/dto/response"
	"application/port/secondary_port"
	"sort"
)

type ReadAllMantrasUseCase struct {
	repo secondary_port.MantraRepository
}

func New(repo secondary_port.MantraRepository) *ReadAllMantrasUseCase {
	return &ReadAllMantrasUseCase{
		repo: repo,
	}
}

func (r *ReadAllMantrasUseCase) Execute() (list []response.ReadAllMantraResDTO) {
	mantras := r.repo.ReadAllMantras()
	for _, v := range mantras {
		mantra := response.Init(v)
		list = append(list, *mantra)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Id < list[j].Id
	})
	return list
}
