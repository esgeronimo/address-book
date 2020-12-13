package model

import (
	"github.com/google/uuid"
)

type addressBook struct {
	id       string
	contacts []Contact
}

type AddressBook interface {
	GetContacts() []Contact
	AddContact(contactName string)
}

func NewAddressBook(contactNames []string, generator IDGenerator) *addressBook {
	return newAddressBook(contactNames, generator.Generate())
}

func newAddressBook(contactNames []string, generatedAddressBookID string) *addressBook {
	var contacts = []Contact{}
	for _, contactName := range contactNames {
		contacts = append(contacts, Contact{
			contactID: uuid.New().String(),
			name:      contactName,
		})
	}
	return &addressBook{
		id:       generatedAddressBookID,
		contacts: contacts,
	}
}

// GetContacts returns a list of contacts in the address book
func (a *addressBook) GetContacts() []Contact {
	return a.contacts
}

// AddContact adds a new contact in the address book
func (a *addressBook) AddContact(contactName string) {
	a.contacts = append(a.contacts, Contact{
		contactID: uuid.New().String(),
		name:      contactName,
	})
}
