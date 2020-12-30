package rest

import (
	"esgeronimo/address-book/core"
	"esgeronimo/address-book/core/port"

	"github.com/gin-gonic/gin"
)

type restApplication struct {
	repo               port.AddressBookRepository
	addressBookService core.AddressBookService
}

func NewRestApplication(repo port.AddressBookRepository, addressBookService core.AddressBookService) *restApplication {
	return &restApplication{
		repo:               repo,
		addressBookService: addressBookService,
	}
}

func (p restApplication) Init() {
	r := gin.Default()

	addressBookHandler := NewAddressBookHandler(p.addressBookService)
	r.GET("/address-book/:addressBookID", addressBookHandler.GetAddressBook)

	addressBookContactHandler := NewAddressBookContactHandler()
	r.POST("/address-book/:addressBookID/contacts", addressBookContactHandler.AddContact)
}
