package application

import (
	"esgeronimo/address-book/internal/application/rest"
	"esgeronimo/address-book/internal/core/model"
	"esgeronimo/address-book/internal/core/repo"
	"esgeronimo/address-book/internal/core/service"
	"esgeronimo/address-book/internal/infrastructure/db"
)

type Application interface {
	Run()
}

var app Application

func init() {
	repo := addressBookRepository()

	app = rest.NewRestApplication(
		repo,
		service.NewAddressBookService(repo),
	)
}

func Run() {
	app.Run()
}

func addressBookRepository() repo.AddressBookRepository {
	return db.NewMockAddressBookRepository(map[string]*model.AddressBook{
		"address-book-1": &model.AddressBook{
			ID: "address-book-1",
			Contacts: []model.Contact{
				model.Contact{
					ContactID: "contact-id-0",
					Name:      "eugene karl geronimo",
				},
				model.Contact{
					ContactID: "cotnact-id-1",
					Name:      "ma. ciela salazar",
				},
			},
		},
	})
}
