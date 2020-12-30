package model

type Contact interface {
	Name() string
}

type contact struct {
	contactID string `json:"contactID"`
	name      string `json:"name"`
}

func NewContact(contactID string, name string) Contact {
	return &contact{
		contactID: contactID,
		name:      name,
	}
}

func (c contact) Name() string {
	return c.name
}
