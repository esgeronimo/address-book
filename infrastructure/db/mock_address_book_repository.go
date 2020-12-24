package db

import (
	"esgeronimo/address-book/core/model"
	"esgeronimo/address-book/core/port"
	"fmt"
)

type mockAddressBookRepository struct {
	addressBookMap map[string]model.AddressBook
}

// NewMockAddressBookRepository returns instance of db.mockAddressBookRepository
func NewMockAddressBookRepository(addressBookMap map[string]model.AddressBook) port.AddressBookRepository {
	return &mockAddressBookRepository{
		addressBookMap: addressBookMap,
	}
}

func (m mockAddressBookRepository) Get(addressBookID string) (model.AddressBook, error) {
	addressBook := m.addressBookMap[addressBookID]
	if addressBook == nil {
		return nil, fmt.Errorf("address book with ID = %s not found", addressBookID)
	}

	return addressBook, nil
}

func (m *mockAddressBookRepository) Add(model.AddressBook) error {
	return nil
}

func (m *mockAddressBookRepository) AddContact(addressBookID string, contact model.Contact) error {
	return nil
}
