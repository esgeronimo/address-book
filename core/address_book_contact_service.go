package core

import (
	"esgeronimo/address-book/core/model"
	"esgeronimo/address-book/core/port"
)

type AddressBookContactService interface {
	GetAll() ([]model.Contact, error)
	Add(contact model.Contact) error
}

type addressBookContactService struct {
	addressBookID string
	repo          port.AddressBookRepository
}

func NewAddressBookContactService(addressBookID string, repo port.AddressBookRepository) AddressBookContactService {
	return &addressBookContactService{
		addressBookID: addressBookID,
		repo:          repo,
	}
}

func (a addressBookContactService) GetAll() ([]model.Contact, error) {
	addressBook, err := a.repo.Get(a.addressBookID)
	if err != nil {
		return nil, err
	}

	return addressBook.Contacts(), nil
}

func (a *addressBookContactService) Add(contact model.Contact) error {
	return a.repo.AddContact(a.addressBookID, contact)
}
