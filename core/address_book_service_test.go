package core

import (
	"esgeronimo/address-book/core/model"
	"esgeronimo/address-book/infrastructure/db"
	"testing"
)

func TestGet(t *testing.T) {
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

	service := NewAddressBookService(repo)

	// when
	targetAddressBookID := "address-book-0"
	addressBook, err := service.Get(targetAddressBookID)

	// then - service call should not return an error
	if err != nil {
		t.Errorf("unexpected error encountered: %v", err)
	}

	// then - number of contacts for targeted address book ID is equal to expected size
	expectedSize := 2
	actualSize := len(addressBook.Contacts())
	if expectedSize != actualSize {
		t.Errorf("expected size is %d, actual size is %d", expectedSize, actualSize)
	}

	// then - retrieved contacts equal to expected contact names
	contains := func(contacts []model.Contact, contactName string) bool {
		for _, contact := range contacts {
			if contact.Name() == contactName {
				return true
			}
		}

		return false
	}
	retrievedContacts := addressBook.Contacts()
	expectedContactNames := []string{"eugene", "karl"}
	for _, contactName := range expectedContactNames {
		if !contains(retrievedContacts, contactName) {
			t.Errorf("address book %s is expected to contain contact with name %s", targetAddressBookID, contactName)
		}
	}
}
func TestGet_error(t *testing.T) {
	// given
	repo := db.NewMockAddressBookRepository(map[string]model.AddressBook{
		"address-book-0": model.NewAddressBook("address-book-0", []model.Contact{
			model.NewContact("contact-0-0", "eugene"),
			model.NewContact("contact-0-1", "karl"),
		}),
	})

	service := NewAddressBookService(repo)

	// when
	targetAddressBookID := "address-book-1"
	_, err := service.Get(targetAddressBookID)

	// then
	if err == nil {
		t.Error("error is expected but err is nil")
	}
}
