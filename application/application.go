package application

import (
	"esgeronimo/address-book/application/rest"
	"esgeronimo/address-book/core"
	"esgeronimo/address-book/core/model"
	"esgeronimo/address-book/infrastructure/db"
)

type Application interface {
	Run()
}

var app Application

func init() {
	repo := addressBookRepository()

	app = rest.NewRestApplication(
		repo,
		core.NewAddressBookService(repo),
	)
}

func Run() {
	app.Run()
}

func addressBookRepository() core.AddressBookRepository {
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
