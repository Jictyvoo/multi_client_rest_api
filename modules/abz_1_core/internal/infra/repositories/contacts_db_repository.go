package repositories

import (
	"database/sql"
	"github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core/internal/domain/interfaces"
)

type ContactsDbRepository struct {
	db *sql.DB
}

func NewContactsDbRepository(db *sql.DB) *ContactsDbRepository {
	return &ContactsDbRepository{db: db}
}

func (repo ContactsDbRepository) Add(dto interfaces.ContactDTO) error {
	//TODO implement me
	panic("implement me")
}

func (repo ContactsDbRepository) AddAll(dto interfaces.ContactDTO) error {
	//TODO implement me
	panic("implement me")
}

func (repo ContactsDbRepository) GetByPhone(s string) (interfaces.ContactDTO, error) {
	//TODO implement me
	panic("implement me")
}
