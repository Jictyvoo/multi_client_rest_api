package services

import (
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/dtos"
)

// ContactsServiceFacade defines how the contacts service facade should be implemented and methods that it have to
// manipulate contacts
type ContactsServiceFacade interface {
	Add(contacts []dtos.ContactsDTO) error
	ListAll() ([]dtos.ContactsDTO, error)
}
