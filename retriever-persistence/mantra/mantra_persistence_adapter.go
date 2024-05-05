package mantra_persistence

import (
	"github.com/google/uuid"
	"retriever-core/mantra/model"
	"retriever-persistence/mantra/mapper"
	"retriever-persistence/mantra/repository"
)

type MantraPersistenceAdapter struct {
	repository mantra_repository.MantraRepository
	mapper     mantra_mapper.MantraMapper
}

func NewMantraPersistenceAdapter(repository mantra_repository.MantraRepository) *MantraPersistenceAdapter {
	return &MantraPersistenceAdapter{
		repository: repository,
	}
}

func (r *MantraPersistenceAdapter) ReadAllMantras() (mantraList []mantra_model.Mantra, err error) {
	mantraEntityList, err := r.repository.ReadAllMantras()
	for _, v := range mantraEntityList {
		mantraList = append(mantraList, r.mapper.ToDomain(v))
	}
	return mantraList, err
}

func (r *MantraPersistenceAdapter) CreateMantra(Mantra mantra_model.Mantra) error {
	entity := r.mapper.ToEntity(Mantra)
	return r.repository.CreateMantra(entity)
}

func (r *MantraPersistenceAdapter) DeleteMantra(MantraId uuid.UUID) error {
	return r.repository.DeleteMantra(MantraId)
}

func (r *MantraPersistenceAdapter) FindMantraById(mantraId uuid.UUID) (mantra_model.Mantra, error) {
	mantraEntity, err := r.repository.FindMantraById(mantraId)
	return r.mapper.ToDomain(mantraEntity), err
}
