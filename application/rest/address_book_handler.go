package rest

import (
	"esgeronimo/address-book/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddressBookHandler interface {
	GetAddressBook(c *gin.Context)
}

type addressBookHandler struct {
	service core.AddressBookService
}

func NewAddressBookHandler(service core.AddressBookService) AddressBookHandler {
	return &addressBookHandler{
		service: service,
	}
}

func (a addressBookHandler) GetAddressBook(c *gin.Context) {
	addressBook, err := a.service.Get(c.Params.ByName("addressBookID"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": addressBook})
}
