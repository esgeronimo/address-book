package service_test

import (
	"esgeronimo/address-book/internal/core/service"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// given
	repo := new(mockAddressBookRepository)
	repo.On("AddContact", "address-book-0", "contact-0-0").Return(nil)

	service := service.NewAddressBookContactService("address-book-0", repo)

	// when
	err := service.Add("contact-0-0")

	// then
	repo.AssertExpectations(t)
	assert.Nil(t, err, fmt.Sprintf("unexpected error encountered: %v", err))
}
