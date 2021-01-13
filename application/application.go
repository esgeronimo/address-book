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
	return db.NewMockAddressBookRepository(map[string]model.AddressBook{
		"address-book-1": model.NewAddressBook("address-book-1", []model.Contact{
			model.NewContact("contact-id-0", "eugene karl geronimo"),
			model.NewContact("contact-id-1", "ma. ciela salazar"),
		}),
	})
}
