package mantra_repository

import (
	"database/sql"
	"github.com/google/uuid"
	"retriever-core/mantra/model"
	"retriever-persistence/mantra/entity"
	"retriever-persistence/mantra/mapper"
)

type MantraPersistenceAdapter struct {
	db     *sql.DB
	mapper mantra_mapper.MantraMapper
}

func NewMantraPersistenceAdapter(db *sql.DB) *MantraPersistenceAdapter {
	return &MantraPersistenceAdapter{
		db: db,
	}
}

func (r *MantraPersistenceAdapter) ReadAllMantras() (list []mantra_model.Mantra, err error) {
	rows, err := r.db.Query("SELECT * FROM tbl_mantra")
	if err != nil {
		return list, err
	}
	mantraEntity := mantra_entity.MantraEntity{}
	for rows.Next() {
		err = rows.Scan(&mantraEntity.MantraId, &mantraEntity.Writer, &mantraEntity.Content, &mantraEntity.CreatedAt)
		list = append(list, r.mapper.ToDomain(mantraEntity))
	}
	return list, nil
}

func (r *MantraPersistenceAdapter) CreateMantra(Mantra mantra_model.Mantra) error {
	mantraEntity := r.mapper.ToEntity(Mantra)
	_, err := r.db.Exec("insert into tbl_mantra values (?, ?, ?, ?)", mantraEntity.MantraId, mantraEntity.Writer, mantraEntity.Content, mantraEntity.CreatedAt)
	return err
}

func (r *MantraPersistenceAdapter) DeleteMantra(MantraId uuid.UUID) error {
	_, err := r.db.Exec("delete from tbl_mantra where mantraId = ?", MantraId.String())
	return err
}
