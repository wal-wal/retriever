package mantra_repository

import (
	"database/sql"
	"github.com/google/uuid"
	"retriever-persistence/mantra/entity"
)

type MantraRepository struct {
	db *sql.DB
}

func NewMantraRepository(db *sql.DB) MantraRepository {
	return MantraRepository{
		db: db,
	}
}

func (r *MantraRepository) ReadAllMantras() (list []mantra_entity.MantraEntity, err error) {
	rows, err := r.db.Query("SELECT * FROM tbl_mantra")
	if err != nil {
		return list, err
	}
	mantraEntity := mantra_entity.MantraEntity{}
	for rows.Next() {
		err = rows.Scan(&mantraEntity.MantraId, &mantraEntity.Writer, &mantraEntity.Content, &mantraEntity.CreatedAt)
		list = append(list, mantraEntity)
	}
	return list, nil
}

func (r *MantraRepository) CreateMantra(mantraEntity mantra_entity.MantraEntity) error {
	_, err := r.db.Exec("insert into tbl_mantra values (?, ?, ?, ?)", mantraEntity.MantraId, mantraEntity.Writer, mantraEntity.Content, mantraEntity.CreatedAt)
	return err
}

func (r *MantraRepository) DeleteMantra(mantraId uuid.UUID) error {
	_, err := r.db.Exec("delete from tbl_mantra where mantraId = ?", mantraId)
	return err
}

func (r *MantraRepository) FindMantraById(mantraId uuid.UUID) (entity mantra_entity.MantraEntity, err error) {
	err = r.db.QueryRow("select * from tbl_mantra where mantraId = ?", mantraId).Scan(&entity.MantraId, &entity.Content, &entity.Writer, &entity.CreatedAt)
	return entity, err
}
