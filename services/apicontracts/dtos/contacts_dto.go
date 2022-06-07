package dtos

type (
	ContactsDTO struct {
		FullName  string `json:"name"`
		Cellphone string `json:"cellphone"`
	}
	ContactsListDTO struct {
		Contacts []ContactsDTO `json:"contacts"`
	}
)

func (c ContactsDTO) Name() string {
	return c.FullName
}

func (c ContactsDTO) Phone() string {
	return c.Cellphone
}
