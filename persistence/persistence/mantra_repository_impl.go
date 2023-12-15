package persistence

import (
	"retriever/application_regercy/domain/mantra/entity"
)

type MantraRepositoryImpl struct {
	m map[int]entity.Mantra
}

func New(m map[int]entity.Mantra) *MantraRepositoryImpl {
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
