package core

type AddressBookContactService interface {
	Add(contactName string) error
}

type addressBookContactService struct {
	addressBookID string
	repo          AddressBookRepository
}

func NewAddressBookContactService(addressBookID string, repo AddressBookRepository) AddressBookContactService {
	return &addressBookContactService{
		addressBookID: addressBookID,
		repo:          repo,
	}
}

func (a *addressBookContactService) Add(contactName string) error {
	return a.repo.AddContact(a.addressBookID, contactName)
}
