package model

type Contact struct {
	contactID string
	name      string
}

// GetContactID returns contacts unique identifier
func (c Contact) GetContactID() string {
	return c.contactID
}
