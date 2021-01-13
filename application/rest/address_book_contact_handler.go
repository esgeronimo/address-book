package rest

import (
	"esgeronimo/address-book/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddressBookContactHandler interface {
	AddContact(c *gin.Context)
}

type addressBookContactHandler struct {
	repo core.AddressBookRepository
}

func NewAddressBookContactHandler(repo core.AddressBookRepository) AddressBookContactHandler {
	return &addressBookContactHandler{
		repo: repo,
	}
}

type addContactReq struct {
	ContactName string `json:"contactName" binding:"required"`
}

func (a *addressBookContactHandler) AddContact(c *gin.Context) {
	addressBookID := c.Params.ByName("addressBookID")
	if addressBookID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "addressBookID is missing"})
		return
	}

	var req addContactReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := core.NewAddressBookContactService(addressBookID, a.repo)
	if err := service.Add(req.ContactName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
