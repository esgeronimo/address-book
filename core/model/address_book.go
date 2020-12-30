package model

type AddressBook interface {
	ID() string
	Contacts() []Contact
	AddContact(contactID string, contactName string)
}

type addressBook struct {
	id       string    `json:"addressBookID"`
	contacts []Contact `json:"contacts"`
}

func NewAddressBook(id string, contacts []Contact) AddressBook {
	return &addressBook{
		id:       id,
		contacts: contacts,
	}
}

func (a addressBook) ID() string {
	return a.id
}

func (a addressBook) Contacts() []Contact {
	return a.contacts
}

func (a *addressBook) AddContact(contactID string, contactName string) {
	a.contacts = append(a.contacts, NewContact(contactID, contactName))
}
