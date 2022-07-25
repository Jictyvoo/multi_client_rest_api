package entities

type ContactEntity interface {
	SetName(name string)
	SetPhone(phone string) error
	Phone() string
	Name() string
}
