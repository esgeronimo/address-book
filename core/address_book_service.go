package core

import (
	"esgeronimo/address-book/core/model"
	"esgeronimo/address-book/core/port"
)

type AddressBookService interface {
	Get(addressBookID string) (model.AddressBook, error)
}

type addressBookService struct {
	repo port.AddressBookRepository
}

func NewAddressBookService(repo port.AddressBookRepository) AddressBookService {
	return &addressBookService{
		repo: repo,
	}
}

func (a addressBookService) Get(addressBookID string) (model.AddressBook, error) {
	return a.repo.Get(addressBookID)
}
