package interfaces

type (
	ContactDTO interface {
		Name() string
		Phone() string
	}
	ContactsRepository interface {
		Add(ContactDTO) error
		AddAll(ContactDTO) error
		GetByPhone(string) error
	}
)
