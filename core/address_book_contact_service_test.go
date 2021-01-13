package core

import (
	"esgeronimo/address-book/core/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockAddressBookRepository struct {
	mock.Mock
}

func (m *mockAddressBookRepository) Get(ID string) (model.AddressBook, error) {
	returnArgs := m.Called(ID)
	return returnArgs.Get(0).(model.AddressBook), returnArgs.Error(1)
}

func (m *mockAddressBookRepository) Add(addressBook model.AddressBook) error {
	returnArgs := m.Called(addressBook)
	return returnArgs.Error(0)
}

func (m *mockAddressBookRepository) AddContact(addressBookID string, contactName string) error {
	returnArgs := m.Called(addressBookID, contactName)
	return returnArgs.Error(0)
}

func TestAdd(t *testing.T) {
	// given
	repo := new(mockAddressBookRepository)
	repo.On("AddContact", "address-book-0", "contact-0-0").Return(nil)

	service := NewAddressBookContactService("address-book-0", repo)

	// when
	err := service.Add("contact-0-0")

	// then
	repo.AssertExpectations(t)
	assert.Nil(t, err, fmt.Sprintf("unexpected error encountered: %v", err))
}
