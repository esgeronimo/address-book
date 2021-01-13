package core

import (
	"esgeronimo/address-book/core/model"
)

type AddressBookService interface {
	Get(addressBookID string) (model.AddressBook, error)
}

type addressBookService struct {
	repo AddressBookRepository
}

func NewAddressBookService(repo AddressBookRepository) AddressBookService {
	return &addressBookService{
		repo: repo,
	}
}

func (a addressBookService) Get(addressBookID string) (addressBook model.AddressBook, err error) {
	addressBook, err = a.repo.Get(addressBookID)
	if err != nil {
		return (model.AddressBook)(nil), err
	}
	return addressBook, nil
}
