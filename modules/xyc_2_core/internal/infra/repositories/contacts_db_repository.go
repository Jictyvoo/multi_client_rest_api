package repositories

import (
	"database/sql"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/domain/interfaces"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/infra/models"
)

type ContactsDbRepository struct {
	db *sql.DB
}

func NewContactsDbRepository(db *sql.DB) *ContactsDbRepository {
	return &ContactsDbRepository{db: db}
}

func (repo ContactsDbRepository) ListAll() (contactsList []interfaces.ContactDTO, err error) {
	const sqlCmd = `SELECT id, nome, celular FROM contacts`
	var rows *sql.Rows

	rows, err = repo.db.Query(sqlCmd)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if err == nil {
			err = closeErr
		}
	}(rows)

	contactsList = make([]interfaces.ContactDTO, 0, 11)
	for rows.Next() {
		var contact models.ContactsModel
		err = rows.Scan(&contact.Id, &contact.FullName, &contact.Cellphone)
		if err != nil {
			return nil, err
		}
		contactsList = append(contactsList, interfaces.ContactDTO{
			FullName:  contact.Name(),
			Cellphone: contact.Phone(),
		})
	}
	return
}

func (repo ContactsDbRepository) Add(dto interfaces.ContactDTO) error {
	const sqlCmd = `INSERT INTO contacts (nome, celular) VALUES ($1, $2)`
	_, err := repo.db.Exec(sqlCmd, dto.Name(), dto.Phone())
	return err
}

func (repo ContactsDbRepository) AddAll(dto []interfaces.ContactDTO) (err error) {
	// Create a transaction to ensure that it will be rollback if an error occurs
	var tx *sql.Tx
	tx, err = repo.db.Begin()
	if err != nil {
		return err
	}
	defer func(tx *sql.Tx) {
		rollbackErr := tx.Rollback()
		if err == nil {
			err = rollbackErr
		}
	}(tx)

	// Insert the list of contacts
	for _, contact := range dto {
		_, err = tx.Exec(`INSERT INTO contacts (nome, celular) VALUES ($1, $2)`, contact.Name(), contact.Phone())
		if err != nil {
			return err
		}
	}
	return err
}

func (repo ContactsDbRepository) GetByPhone(s string) (interfaces.ContactDTO, error) {
	const sqlCmd = `SELECT nome, celular FROM contacts WHERE celular = $1`
	var contact models.ContactsModel
	err := repo.db.QueryRow(sqlCmd, s).Scan(&contact.FullName, &contact.Cellphone)
	return interfaces.ContactDTO{
		FullName:  contact.Name(),
		Cellphone: contact.Phone(),
	}, err
}
