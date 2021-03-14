package service

import (
	"esgeronimo/address-book/internal/core/model"
	"esgeronimo/address-book/internal/core/repo"
)

type AddressBookService interface {
	Get(addressBookID string) (*model.AddressBook, error)
}

type addressBookService struct {
	repo repo.AddressBookRepository
}

func NewAddressBookService(repo repo.AddressBookRepository) *addressBookService {
	return &addressBookService{
		repo: repo,
	}
}

func (a addressBookService) Get(addressBookID string) (addressBook *model.AddressBook, err error) {
	return a.repo.Get(addressBookID)
}
