package domain

import (
	"database/sql"
	"errors"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/domain/entities"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/domain/interfaces"
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/corerrs"
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/dtos"
)

type ContactsService struct {
	repository interfaces.ContactsRepository
}

func NewContactsService(repository interfaces.ContactsRepository) *ContactsService {
	return &ContactsService{
		repository: repository,
	}
}

func (service ContactsService) Validate(dto interfaces.ContactDTO) (contact entities.ContactEntity, err error) {
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

func (service ContactsService) Add(contacts []dtos.ContactsDTO) error {
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

func (service ContactsService) AddAll(contacts []dtos.ContactsDTO) error {
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
