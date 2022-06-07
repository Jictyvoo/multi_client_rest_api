package interfaces

type (
	ContactDTO interface {
		Name() string
		Phone() string
	}
	ContactsRepository interface {
		ListAll() (contactsList []ContactDTO, err error)
		Add(ContactDTO) error
		AddAll([]ContactDTO) error
		GetByPhone(string) (ContactDTO, error)
	}
)
