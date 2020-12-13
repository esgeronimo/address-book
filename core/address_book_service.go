package core

import (
	"esgeronimo/address-book/core/model"
	"esgeronimo/address-book/core/port"
	"fmt"
)

type AddressBookService struct {
	repo port.AddressBookRepository
}

// NewAddressBookService creates a core.AddressBookService instance
func NewAddressBookService(repo port.AddressBookRepository) *AddressBookService {
	return &AddressBookService{
		repo: repo,
	}
}

// AddContact adds a contact to specified address book
func (a *AddressBookService) AddContact(addressBookID string, contactName string) (string, error) {
	return "", fmt.Errorf("not implemented yet")
}

// GetContacts returns list of contacts for a specified address book
func (a *AddressBookService) GetContacts(addressBookID string) ([]model.AddressBook, error) {
	return nil, fmt.Errorf("not implemented yet")
}

// SearchContacts returns list of contacts for a specified address book and search token
func (a *AddressBookService) SearchContacts(addressBookID string, searchToken string) ([]model.AddressBook, error) {
	return nil, fmt.Errorf("not implemented yet")
}
