package domain

import (
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/core"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/domain/entities"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/domain/interfaces"
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/corerrs"
	"github.com/wrapped-owls/goremy-di/remy"
)

type ContactsService struct {
	repository interfaces.ContactsRepository
}

func NewContactsService(repository interfaces.ContactsRepository) *ContactsService {
	if repository == nil {
		repository = remy.Get[interfaces.ContactsRepository](core.Injector)
	}
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
	if err == nil && (len(tempContact.Phone()) > 0) {
		err = corerrs.ErrContactAlreadyExists
		return
	}
	return
}

func (service ContactsService) Add(contacts []interfaces.ContactDTO) error {
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
