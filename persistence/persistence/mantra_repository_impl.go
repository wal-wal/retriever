package persistence

import (
	"application/domain/mantra/entity"
	"github.com/google/uuid"
)

type MantraRepositoryImpl struct {
	m map[uuid.UUID]entity.Mantra
}

func New(m map[uuid.UUID]entity.Mantra) *MantraRepositoryImpl {
	return &MantraRepositoryImpl{
		m: m,
	}
}

func (r *MantraRepositoryImpl) ReadAllMantras() (list []entity.Mantra) {
	for _, v := range r.m {
		list = append(list, v)
	}
	return list
}

func (r *MantraRepositoryImpl) CreateMantra(mantra entity.Mantra) error {
	r.m[mantra.Id] = mantra
	return nil
}
