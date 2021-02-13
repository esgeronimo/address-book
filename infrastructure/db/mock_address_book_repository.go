package db

import (
	"esgeronimo/address-book/core"
	"esgeronimo/address-book/core/model"
	"fmt"

	"github.com/google/uuid"
)

type mockAddressBookRepository struct {
	addressBookMap map[string]*model.AddressBook
}

// NewMockAddressBookRepository returns instance of db.mockAddressBookRepository
func NewMockAddressBookRepository(addressBookMap map[string]*model.AddressBook) core.AddressBookRepository {
	return &mockAddressBookRepository{
		addressBookMap: addressBookMap,
	}
}

func (m mockAddressBookRepository) Get(addressBookID string) (*model.AddressBook, error) {
	addressBook := m.addressBookMap[addressBookID]
	if addressBook.ID == "" {
		return nil, fmt.Errorf("address book with ID = %s not found", addressBookID)
	}

	return addressBook, nil
}

func (m *mockAddressBookRepository) Add(model.AddressBook) error {
	return nil
}

func (m *mockAddressBookRepository) AddContact(addressBookID string, contactName string) error {
	if addressBook := m.addressBookMap[addressBookID]; addressBook.ID != "" {
		addressBook.Contacts = append(addressBook.Contacts, model.Contact{
			ContactID: uuid.New().String(),
			Name:      contactName,
		})
		return nil
	}

	return fmt.Errorf("address book with ID = %s not found - contact will not be added", addressBookID)
}
