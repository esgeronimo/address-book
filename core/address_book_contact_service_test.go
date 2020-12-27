package core

import (
	"esgeronimo/address-book/core/model"
	"esgeronimo/address-book/infrastructure/db"
	"testing"
)

func TestAdd(t *testing.T) {
	// given
	repo := db.NewMockAddressBookRepository(map[string]model.AddressBook{
		"address-book-0": model.NewAddressBook("address-book-0", []model.Contact{}),
	})

	service := NewAddressBookContactService("address-book-0", repo)

	newContact := model.NewContact("contact-0-0", "eugene")

	// when
	err := service.Add(newContact)

	// then
	if err != nil {
		t.Errorf("unexpected error encountered: %v", err)
	}
}
