package service_test

import (
	"esgeronimo/address-book/core/model"

	"github.com/stretchr/testify/mock"
)

type mockAddressBookRepository struct {
	mock.Mock
}

func (m *mockAddressBookRepository) Get(ID string) (*model.AddressBook, error) {
	returnArgs := m.Called(ID)
	if returnArgs.Get(0) == nil {
		return nil, returnArgs.Error(1)
	}
	return returnArgs.Get(0).(*model.AddressBook), returnArgs.Error(1)
}

func (m *mockAddressBookRepository) Add(addressBook model.AddressBook) error {
	returnArgs := m.Called(addressBook)
	return returnArgs.Error(0)
}

func (m *mockAddressBookRepository) AddContact(addressBookID string, contactName string) error {
	returnArgs := m.Called(addressBookID, contactName)
	return returnArgs.Error(0)
}
