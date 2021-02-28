package rest

import (
	"esgeronimo/address-book/core/repo"
	"esgeronimo/address-book/core/service"

	"github.com/gin-gonic/gin"
)

type restApplication struct {
	repo               repo.AddressBookRepository
	addressBookService service.AddressBookService
}

func NewRestApplication(repo repo.AddressBookRepository, addressBookService service.AddressBookService) *restApplication {
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
