package db

import (
	"database/sql"
	"esgeronimo/address-book/core/model"
	"fmt"
	"log"
)

type mysqlAddressBookRepository struct {
	IPAddress    string
	Port         int
	DatabaseName string
	DriverName   string
	Account      string
	Password     string
}

func (m *mysqlAddressBookRepository) getConnString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.Account, m.Password, m.IPAddress, m.Port, m.DatabaseName)
}

func (m *mysqlAddressBookRepository) Get(addressBookID string) (*model.AddressBook, error) {
	db, err := sql.Open(m.DriverName, m.getConnString())

	// Yeah this is stupid. This query will make sense if address book already has other info outside of its identifier
	var addressBook model.AddressBook
	db.QueryRow("SELECT UUID FROM ADDRESS_BOOK WHERE UUID=?", addressBookID).Scan(&addressBook.ID)

	contactRS, err := db.Query("SELECT UUID, NAME FROM CONTACT WHERE ADDRESS_BOOK_ID=?", addressBookID)
	if err != nil {
		return nil, err
	}

	var contacts = []model.Contact{}
	for contactRS.Next() {
		var contact model.Contact
		err = contactRS.Scan(&contact.ContactID, &contact.Name)
		if err != nil {
			log.Print(err)
		}
		contacts = append(contacts, contact)
	}
	addressBook.Contacts = contacts

	return &addressBook, nil
}
