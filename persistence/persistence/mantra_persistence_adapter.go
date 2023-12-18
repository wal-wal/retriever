package persistence

import (
	"application/domain/mantra/entity"
	"github.com/google/uuid"
)

type MantraPersistenceAdapter struct {
	m map[uuid.UUID]entity.Mantra
}

func NewMantraPersistenceAdapter(m map[uuid.UUID]entity.Mantra) *MantraPersistenceAdapter {
	return &MantraPersistenceAdapter{
		m: m,
	}
}

func (r *MantraPersistenceAdapter) ReadAllMantras() (list []entity.Mantra, err error) {
	for _, v := range r.m {
		list = append(list, v)
	}
	return list, nil
}

func (r *MantraPersistenceAdapter) CreateMantra(mantra entity.Mantra) error {
	r.m[mantra.Id] = mantra
	return nil
}

func (r *MantraPersistenceAdapter) DeleteMantra(id uuid.UUID) error {
	delete(r.m, id)
	return nil
}
