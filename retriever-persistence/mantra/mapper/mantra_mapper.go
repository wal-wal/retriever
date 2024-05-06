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
		Speaker:   entity.Speaker,
		Content:   entity.Content,
		CreatedAt: createdAt,
		Writer:    entity.Writer,
	}
}

func (r *MantraMapper) ToEntity(model mantra_model.Mantra) mantra_entity.MantraEntity {
	return mantra_entity.MantraEntity{
		MantraId:  model.MantraId.String(),
		Speaker:   model.Speaker,
		Content:   model.Content,
		CreatedAt: model.CreatedAt.Format(time.DateTime),
		Writer:    model.Writer,
	}
}
