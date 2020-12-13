package db

import "esgeronimo/address-book/core/model"

type mockAddressBookRepository struct{}

// NewMockAddressBookRepository returns instance of db.mockAddressBookRepository
func NewMockAddressBookRepository() *mockAddressBookRepository {
	return &mockAddressBookRepository{}
}

func (m *mockAddressBookRepository) Get(ID string) (model.AddressBook, error) {
	addressBook := model.AddressBook{}
	addressBook.AddContact("Eugene Geronimo")
	return addressBook, nil
}
