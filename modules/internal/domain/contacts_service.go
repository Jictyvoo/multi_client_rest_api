package domain

import (
	"database/sql"
	"errors"
	"github.com/jictyvoo/multi_client_rest_api/modules/internal/domain/entities"
	"github.com/jictyvoo/multi_client_rest_api/modules/internal/domain/interfaces"
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/corerrs"
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/dtos"
)

type ContactsService[T entities.ContactEntity] struct {
	repository interfaces.ContactsRepository
}

func NewContactsService[T entities.ContactEntity](repository interfaces.ContactsRepository) *ContactsService[T] {
	return &ContactsService[T]{
		repository: repository,
	}
}

func (service ContactsService[T]) Validate(dto interfaces.ContactDTO) (contact T, err error) {
	// check if the phone is valid
	contact.SetName(dto.Name())
	if err = contact.SetPhone(dto.Phone()); err != nil {
		return
	}

	// check if the contact already exists
	var tempContact interfaces.ContactDTO
	tempContact, err = service.repository.GetByPhone(contact.Phone())
	if err == nil && (tempContact != nil && len(tempContact.Phone()) > 0) {
		err = corerrs.ErrContactAlreadyExists
		return
	} else if errors.Is(sql.ErrNoRows, err) {
		err = nil
	}
	return
}

func (service ContactsService[T]) Add(contacts []dtos.ContactsDTO) error {
	for _, contactDTO := range contacts {
		contact, err := service.Validate(contactDTO)
		if err != nil {
			return err
		}

		// add the contact
		if err = service.repository.Add(contact); err != nil {
			return err
		}
	}
	return nil
}

func (service ContactsService[T]) AddAll(contacts []dtos.ContactsDTO) error {
	validatedContacts := make([]interfaces.ContactDTO, 0, len(contacts))
	for _, contactDTO := range contacts {
		contact, err := service.Validate(contactDTO)
		if err != nil {
			return err
		}
		validatedContacts = append(validatedContacts, contact)
	}

	// add the contact
	if err := service.repository.AddAll(validatedContacts); err != nil {
		return err
	}
	return nil
}

func (service ContactsService[T]) ListAll() ([]dtos.ContactsDTO, error) {
	contacts, err := service.repository.ListAll()
	if err != nil {
		return nil, err
	}

	var contactsList []dtos.ContactsDTO
	for _, contact := range contacts {
		contactsList = append(contactsList, contact)
	}

	return contactsList, nil
}
