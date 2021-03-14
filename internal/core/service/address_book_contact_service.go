package service

import "esgeronimo/address-book/internal/core/repo"

type AddressBookContactService interface {
	Add(contactName string) error
}

type addressBookContactService struct {
	addressBookID string
	repo          repo.AddressBookRepository
}

func NewAddressBookContactService(addressBookID string, repo repo.AddressBookRepository) *addressBookContactService {
	return &addressBookContactService{
		addressBookID: addressBookID,
		repo:          repo,
	}
}

func (a *addressBookContactService) Add(contactName string) error {
	return a.repo.AddContact(a.addressBookID, contactName)
}
