package core

import (
	"esgeronimo/address-book/core/model"
	"esgeronimo/address-book/core/port"

	"github.com/google/uuid"
)

type AddressBookContactService interface {
	Add(contactName string) error
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

func (a *addressBookContactService) Add(contactName string) error {
	return a.repo.AddContact(a.addressBookID, model.NewContact(uuid.New().String(), contactName))
}
