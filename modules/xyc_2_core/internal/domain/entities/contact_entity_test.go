package entities

import (
	"errors"
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/corerrs"
	"testing"
)

func TestContactEntity_Name(t *testing.T) {
	expectedNames := []string{
		"lowercase name only",
		"Homer Simpson",
		"Harrison Ford",
		"ALL UPPERCASE NAME",
		"Mark Hamill",
	}

	for _, name := range expectedNames {
		var contact ContactEntity
		contact.SetName(name)
		if contact.Name() != name {
			t.Errorf("Name was not parsed correctly.\nReceived:`%s`\nExpected:`%s`", contact.Name(), name)
		}
	}
}

func TestContactEntity_SetPhone__Valid(t *testing.T) {
	contact := ContactEntity{}
	err := contact.SetPhone("5541930306905")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// Verify if the parsed phone is correct
	const expectedPhoneFormat = "5541930306905"
	resultPhone := contact.Phone()
	if contact.Phone() != expectedPhoneFormat {
		t.Errorf("Phone was not parsed correctly.\nReceived:`%s`\nExpected:`%s`", resultPhone, expectedPhoneFormat)
	}
}

func TestContactEntity_SetPhone__InValid(t *testing.T) {
	contact := ContactEntity{}
	invalidPhones := []string{
		"",
		"123456789",
		"0012930306905",
		"+55 () 930306905",
		"+55 (41 9AS12121-30-30-6905",
	}

	for _, phone := range invalidPhones {
		err := contact.SetPhone(phone)
		if err == nil && !errors.Is(err, corerrs.ErrInvalidPhone) {
			t.Errorf("The phone number `%s` provided should be invalid", phone)
		}
	}
}
