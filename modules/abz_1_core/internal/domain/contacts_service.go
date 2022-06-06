package domain

import (
	"github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core/internal/core"
	"github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core/internal/domain/entities"
	"github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core/internal/domain/interfaces"
	"github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core/internal/domain/utils"
	"github.com/wrapped-owls/goremy-di/remy"
)

type ContactsService struct {
	repository interfaces.ContactsRepository
}

func NewContactsService() *ContactsService {
	return &ContactsService{
		repository: remy.Get[interfaces.ContactsRepository](core.Injector),
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
	if tempContact, err = service.repository.GetByPhone(contact.Phone()); err == nil && len(tempContact.Phone()) > 0 {
		err = utils.ErrContactAlreadyExists
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
