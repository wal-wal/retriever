package mantra_mapper

import (
	"github.com/google/uuid"
	"retriever-core/mantra/model"
	"retriever-persistence/mantra/entity"
	"time"
)

type MantraMapper struct{}

func (r *MantraMapper) ToDomain(entity mantra_entity.MantraEntity) mantra_model.Mantra {
	createdAt, _ := time.Parse(time.DateTime, entity.CreatedAt)
	return mantra_model.Mantra{
		MantraId:  uuid.MustParse(entity.MantraId),
		Content:   entity.Content,
		Writer:    entity.Writer,
		CreatedAt: createdAt,
	}
}

func (r *MantraMapper) ToEntity(model mantra_model.Mantra) mantra_entity.MantraEntity {
	return mantra_entity.MantraEntity{
		MantraId:  model.MantraId.String(),
		Content:   model.Content,
		Writer:    model.Writer,
		CreatedAt: model.CreatedAt.Format(time.DateTime),
	}
}
