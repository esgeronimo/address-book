package core

import (
	"esgeronimo/address-book/core/model"
	"esgeronimo/address-book/infrastructure/db"
	"testing"
)

func TestGetAll(t *testing.T) {
	// given
	repo := db.NewMockAddressBookRepository(map[string]model.AddressBook{
		"address-book-0": model.NewAddressBook("address-book-0", []model.Contact{
			model.NewContact("contact-0-0", "eugene"),
			model.NewContact("contact-0-1", "karl"),
		}),
		"address-book-1": model.NewAddressBook("address-book-1", []model.Contact{
			model.NewContact("contact-1-0", "geronimo"),
		}),
	})

	service := NewAddressBookContactService("address-book-1", repo)

	// when
	contacts, err := service.GetAll()

	// then
	if err != nil {
		t.Errorf("unexpected error encountered: %v", err)
	}

	actualSize := len(contacts)
	const expectedSize int = 1
	if actualSize != expectedSize {
		t.Errorf("expected size is %d, actual size is %d", expectedSize, actualSize)
	}
}

func TestGetAll_error(t *testing.T) {
	// given
	repo := db.NewMockAddressBookRepository(map[string]model.AddressBook{
		"address-book-0": model.NewAddressBook("address-book-0", []model.Contact{
			model.NewContact("contact-0-0", "eugene"),
			model.NewContact("contact-0-1", "karl"),
		}),
	})

	service := NewAddressBookContactService("address-book-1", repo)

	// when
	_, err := service.GetAll()

	// then
	if err == nil {
		t.Errorf("error is expected but err is nil")
	}
}
