package mantra_use_case

import (
	"retriever-core/mantra/dto/response"
	"retriever-core/mantra/secondary"
)

type ReadMantrasBySpeakerUseCase struct {
	repo mantra_secondary.MantraSecondaryPort
}

func NewReadMantrasBySpeakerUseCase(repo mantra_secondary.MantraSecondaryPort) ReadMantrasBySpeakerUseCase {
	return ReadMantrasBySpeakerUseCase{
		repo: repo,
	}
}

func (r *ReadMantrasBySpeakerUseCase) Execute(speakerName string) (list []mantra_response.ReadMantraResDTO, err error) {
	mantras, err := r.repo.FindMantrasBySpeaker(speakerName)
	if err != nil {
		return list, err
	}
	for _, v := range mantras {
		list = append(list, *mantra_response.NewReadMantraResDTO(v))
	}
	return list, nil
}
