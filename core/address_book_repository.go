package core

import "esgeronimo/address-book/core/model"

// AddressBookRepository is a datastore port for aggregate model.AddressBook
type AddressBookRepository interface {
	Get(ID string) (*model.AddressBook, error)
	Add(addressBook model.AddressBook) error
	AddContact(addressBookID string, contactName string) error
}
