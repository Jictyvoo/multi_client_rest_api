package interfaces

import "github.com/jictyvoo/multi_client_rest_api/services/apicontracts/dtos"

type (
	ContactDTO         = dtos.ContactsDTO
	ContactsRepository interface {
		ListAll() (contactsList []ContactDTO, err error)
		Add(ContactDTO) error
		AddAll([]ContactDTO) error
		GetByPhone(string) (ContactDTO, error)
	}
)
