package rest

import (
	"esgeronimo/address-book/core"

	"github.com/gin-gonic/gin"
)

type restApplication struct {
	repo               core.AddressBookRepository
	addressBookService core.AddressBookService
}

func NewRestApplication(repo core.AddressBookRepository, addressBookService core.AddressBookService) *restApplication {
	return &restApplication{
		repo:               repo,
		addressBookService: addressBookService,
	}
}

func (p restApplication) Run() {
	r := p.setup()
	r.Run(":8080")
}

func (p restApplication) setup() *gin.Engine {
	r := gin.Default()

	addressBookHandler := NewAddressBookHandler(p.addressBookService)
	r.GET("/address-book/:addressBookID", addressBookHandler.GetAddressBook)

	addressBookContactHandler := NewAddressBookContactHandler(p.repo)
	r.POST("/address-book/:addressBookID/contacts", addressBookContactHandler.AddContact)

	return r
}
