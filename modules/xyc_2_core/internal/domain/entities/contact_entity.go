package entities

import (
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/domain/utils"
	"github.com/nyaruka/phonenumbers"
	"strings"
)

type ContactEntity struct {
	name      string
	cellphone *phonenumbers.PhoneNumber
}

func (c ContactEntity) Name() string {
	return c.name
}

func (c ContactEntity) Phone() string {
	formatted := phonenumbers.Format(c.cellphone, phonenumbers.E164)

	// Remove the plus sign from the beginning of the phone number
	return strings.TrimPrefix(formatted, "+")
}

func (c *ContactEntity) SetName(name string) {
	c.name = name
}

func (c *ContactEntity) SetPhone(phone string) error {
	parsedPhone, err := phonenumbers.Parse(phone, "BR")
	if !phonenumbers.IsValidNumber(parsedPhone) {
		return utils.ErrInvalidPhone
	}

	c.cellphone = parsedPhone
	return err
}
