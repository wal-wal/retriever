package repository

import (
	"application/domain/mantra/model"
	"database/sql"
	"github.com/google/uuid"
	entity2 "persistence/mantra/entity"
	"persistence/mantra/mapper"
)

type MantraPersistenceAdapter struct {
	db     *sql.DB
	mapper mapper.MantraMapper
}

func NewMantraPersistenceAdapter(db *sql.DB) *MantraPersistenceAdapter {
	return &MantraPersistenceAdapter{
		db: db,
	}
}

func (r *MantraPersistenceAdapter) ReadAllMantras() (list []model.Mantra, err error) {
	rows, err := r.db.Query("SELECT * FROM tbl_mantra")
	if err != nil {
		return list, err
	}
	mantraEntity := entity2.MantraEntity{}
	for rows.Next() {
		err = rows.Scan(&mantraEntity.MantraId, &mantraEntity.Writer, &mantraEntity.Content, &mantraEntity.CreatedAt)
		list = append(list, *r.mapper.ToDomain(mantraEntity))
	}
	return list, nil
}

func (r *MantraPersistenceAdapter) CreateMantra(Mantra model.Mantra) error {
	mantraEntity := r.mapper.ToEntity(Mantra)
	_, err := r.db.Exec("insert into tbl_mantra values (?, ?, ?, ?)", mantraEntity.MantraId, mantraEntity.Writer, mantraEntity.Content, mantraEntity.CreatedAt)
	return err
}

func (r *MantraPersistenceAdapter) DeleteMantra(MantraId uuid.UUID) error {
	_, err := r.db.Exec("delete from tbl_mantra where mantraId = ?", MantraId.String())
	return err
}
