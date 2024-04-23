package mapper

import (
	"application/domain/mantra/model"
	"github.com/google/uuid"
	"persistence/mantra/entity"
	"time"
)

type MantraMapper struct{}

func (r *MantraMapper) ToDomain(entity entity.MantraEntity) *model.Mantra {
	createdAt, _ := time.Parse(time.DateTime, entity.CreatedAt)
	return &model.Mantra{
		MantraId:  uuid.MustParse(entity.MantraId),
		Content:   entity.Content,
		Writer:    entity.Writer,
		CreatedAt: createdAt,
	}
}

func (r *MantraMapper) ToEntity(model model.Mantra) *entity.MantraEntity {
	return &entity.MantraEntity{
		MantraId:  model.MantraId.String(),
		Content:   model.Content,
		Writer:    model.Writer,
		CreatedAt: model.CreatedAt.Format(time.DateTime),
	}
}
