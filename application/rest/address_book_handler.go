package rest

import (
	"esgeronimo/address-book/core"
	"esgeronimo/address-book/core/model"
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

type getAddressBookResp struct {
	ID       string   `json:"id"`
	Contacts []string `json:"contacts"`
}

func (a addressBookHandler) mapResponse(addressBook model.AddressBook) getAddressBookResp {
	var contacts []string
	for _, contact := range addressBook.Contacts() {
		contacts = append(contacts, contact.Name())
	}
	return getAddressBookResp{
		ID:       addressBook.ID(),
		Contacts: contacts,
	}
}

func (a addressBookHandler) GetAddressBook(c *gin.Context) {
	addressBook, err := a.service.Get(c.Params.ByName("addressBookID"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": a.mapResponse(addressBook)})
}
