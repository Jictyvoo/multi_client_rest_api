package dtos

type (
	Contacts struct {
		FullName  string `json:"name"`
		Cellphone string `json:"cellphone"`
	}
	ContactsListDTO struct {
		Contacts []Contacts `json:"contacts"`
	}
)

func (c Contacts) Name() string {
	return c.FullName
}

func (c Contacts) Phone() string {
	return c.Cellphone
}
