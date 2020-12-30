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

	// when
	err := service.Add("contact-0-0")

	// then
	if err != nil {
		t.Errorf("unexpected error encountered: %v", err)
	}
}
