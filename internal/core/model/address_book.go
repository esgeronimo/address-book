package model

type AddressBook struct {
	ID       string    `json:"id"`
	Contacts []Contact `json:"contacts"`
}
