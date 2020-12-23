package rest

import (
	"esgeronimo/address-book/application"

	"github.com/gin-gonic/gin"
)

type httpApplication struct{}

func NewHttpApplication() application.Application {
	return &ginApplication{}
}

func (g ginApplication) Init() {
	r := gin.Default()

	r.GET("/address-book/:addressBookID", GetAddressBookHandler())
	r.GET("/address-book/:addressBookID/contacts", GetAddressBookContactsHandler())
	r.POST("/address-book/:addressBookID/contacts", AddAddressBookContactHandler())
}
