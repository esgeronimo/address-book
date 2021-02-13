package core_test

import (
	"errors"
	"esgeronimo/address-book/core"
	"esgeronimo/address-book/core/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestGet(t *testing.T) {
	// given
	repo := new(mockAddressBookRepository)
	repo.On("Get", "address-book-0").Return(&model.AddressBook{
		ID: "address-book-0",
		Contacts: []model.Contact{
			model.Contact{
				ContactID: "contact-0-0",
				Name:      "eugene",
			},
			model.Contact{
				ContactID: "contact-0-1",
				Name:      "karl",
			},
		},
	}, nil)

	service := core.NewAddressBookService(repo)

	// when
	targetAddressBookID := "address-book-0"
	addressBook, err := service.Get(targetAddressBookID)

	// then
	repo.AssertExpectations(t)

	// then - service call should not return an error
	assert.Nil(t, err, fmt.Sprintf("enexpected error encountered: %v", err))

	// then - number of contacts for targeted address book ID is equal to expected size
	expectedSize := 2
	actualSize := len(addressBook.Contacts)
	assert.Equal(t, expectedSize, actualSize, fmt.Sprintf("expected size is %d, actual size is %d", expectedSize, actualSize))

	// then - retrieved contacts equal to expected contact names
	contains := func(contacts []model.Contact, contactName string) bool {
		for _, contact := range contacts {
			if contact.Name == contactName {
				return true
			}
		}

		return false
	}
	retrievedContacts := addressBook.Contacts
	expectedContactNames := []string{"eugene", "karl"}
	for _, contactName := range expectedContactNames {
		assert.True(t, contains(retrievedContacts, contactName),
			fmt.Sprintf("address book %s is expected to contain contact with name %s", targetAddressBookID, contactName))
	}
}
func TestGet_error(t *testing.T) {
	// given
	targetAddressBookID := "address-book-0"
	repo := new(mockAddressBookRepository)
	repo.On("Get", targetAddressBookID).Return(nil, errors.New(""))

	service := core.NewAddressBookService(repo)

	// when
	_, err := service.Get(targetAddressBookID)

	// then
	repo.AssertExpectations(t)
	assert.NotNil(t, err, "error is expected but err is nil")
}
