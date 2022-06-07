package models

type ContactsModel struct {
	Id        uint64
	FullName  string
	Cellphone string
}

func (c ContactsModel) Name() string {
	return c.FullName
}

func (c ContactsModel) Phone() string {
	return c.Cellphone
}
