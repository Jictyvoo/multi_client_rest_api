package entities

import (
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/corerrs"
	"github.com/nyaruka/phonenumbers"
	"strings"
)

type _contactData struct {
	name      string
	cellphone *phonenumbers.PhoneNumber
}

type ContactEntity struct {
	_data *_contactData
}

func (c *ContactEntity) data() *_contactData {
	if c._data == nil {
		c._data = &_contactData{}
	}
	return c._data
}

func (c ContactEntity) Name() string {
	return c.data().name
}

func (c ContactEntity) Phone() string {
	formatted := phonenumbers.Format(c.data().cellphone, phonenumbers.E164)

	// Remove the plus sign from the beginning of the phone number
	return strings.TrimPrefix(formatted, "+")
}

func (c *ContactEntity) SetName(name string) {
	c.data().name = name
}

func (c *ContactEntity) SetPhone(phone string) error {
	parsedPhone, err := phonenumbers.Parse(phone, "BR")
	if !phonenumbers.IsValidNumber(parsedPhone) {
		return corerrs.ErrInvalidPhone
	}

	c.data().cellphone = parsedPhone
	return err
}
