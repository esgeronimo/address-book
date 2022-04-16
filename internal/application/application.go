package application

import (
	"database/sql"
	"esgeronimo/address-book/internal/application/rest"
	"esgeronimo/address-book/internal/core/model"
	"esgeronimo/address-book/internal/core/repo"
	"esgeronimo/address-book/internal/core/service"
	dbimpl "esgeronimo/address-book/internal/infrastructure/db"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Application interface {
	Run()
}

var app Application

func Run() {
	repo := addressBookRepositoryMySQL()

	app = rest.NewRestApplication(
		repo,
		service.NewAddressBookService(repo),
	)
	app.Run()
}

func addressBookRepository() repo.AddressBookRepository {
	return dbimpl.NewMockAddressBookRepository(map[string]*model.AddressBook{
		"address-book-1": {
			ID: "address-book-1",
			Contacts: []model.Contact{
				{
					ContactID: "contact-id-0",
					Name:      "eugene karl geronimo",
				},
				{
					ContactID: "cotnact-id-1",
					Name:      "ma. ciela salazar",
				},
			},
		},
	})
}

func addressBookRepositoryMySQL() repo.AddressBookRepository {
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_ADDRESS"),
		DBName: os.Getenv("DB_DATABASE"),
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	return dbimpl.NewMySQLAddressBookRepository(db)
}
