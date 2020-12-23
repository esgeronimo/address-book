package db

import "esgeronimo/address-book/core/model"

type mockContactRepository struct{}

func NewMockContactRepository() *mockContactRepository {
	return &mockContactRepository{}
}

func (m *mockContactRepository) Save(contact model.Contact) error {
	return nil
}
