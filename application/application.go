package application

import (
	"esgeronimo/address-book/application/rest"
	"esgeronimo/address-book/core"
	"esgeronimo/address-book/core/model"
	"esgeronimo/address-book/core/port"
	"esgeronimo/address-book/infrastructure/db"
)

func Run() {
	repo := addressBookRepository()

	app := rest.NewRestApplication(
		repo,
		core.NewAddressBookService(repo),
	)

	app.Init()
}

type application interface {
	Init()
}

func addressBookRepository() port.AddressBookRepository {
	return db.NewMockAddressBookRepository(map[string]model.AddressBook{
		"test": model.NewAddressBook("address-book-1", []model.Contact{
			model.NewContact("contact-id-0", "eugene karl geronimo"),
			model.NewContact("contact-id-1", "ma. ciela salazar"),
		}),
	})
}
