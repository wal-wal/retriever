package mantra_use_case

import (
	mantra_response "retriever-core/mantra/dto/response"
	"retriever-core/mantra/secondary"
	"sort"
)

type ReadAllMantrasUseCase struct {
	repo mantra_secondary.MantraSecondaryPort
}

func NewReadAllMantrasUseCase(repo mantra_secondary.MantraSecondaryPort) ReadAllMantrasUseCase {
	return ReadAllMantrasUseCase{
		repo: repo,
	}
}

func (r *ReadAllMantrasUseCase) Execute() (list []mantra_response.ReadMantraResDTO, err error) {
	mantras, err := r.repo.ReadAllMantras()
	if err != nil {
		return nil, err
	}
	if mantras == nil {
		return []mantra_response.ReadMantraResDTO{}, err
	}
	for _, v := range mantras {
		mantra := mantra_response.NewReadMantraResDTO(v)
		list = append(list, *mantra)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].CreatedAt.After(list[j].CreatedAt)
	})
	return list, nil
}
