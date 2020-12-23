package db

import (
	"esgeronimo/address-book/core/model"
	"esgeronimo/address-book/core/port"
	"fmt"

	"github.com/google/uuid"
)

type mockAddressBookRepository struct {
	addressBooks []model.AddressBook
}

// NewMockAddressBookRepository returns instance of db.mockAddressBookRepository
func NewMockAddressBookRepository() port.AddressBookRepository {
	contacts := []model.Contact{}

	contact1 := model.NewContact(uuid.New().String(), "Eugene")
	contacts = append(contacts, contact1)

	contact2 := model.NewContact(uuid.New().String(), "Karl")
	contacts = append(contacts, contact2)

	addressBooks := []model.AddressBook{}
	addressBooks = append(addressBooks, model.NewAddressBook("address-book-1234", contacts))

	return &mockAddressBookRepository{
		addressBooks: addressBooks,
	}
}

func (m mockAddressBookRepository) Get(addressBookID string) (model.AddressBook, error) {
	for _, addressBook := range m.addressBooks {
		if addressBook.ID() == addressBookID {
			return addressBook, nil
		}
	}

	return nil, fmt.Errorf("address book with ID = %s not found", addressBookID)
}

func (m *mockAddressBookRepository) Add(model.AddressBook) error {
	return nil
}

func (m *mockAddressBookRepository) AddContact(addressBookID string, contact model.Contact) error {
	return nil
}
