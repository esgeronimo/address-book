package db

import (
	"database/sql"
	"esgeronimo/address-book/internal/core/model"
	"log"
)

type mysqlAddressBookRepository struct {
	db *sql.DB
}

func NewMySQLAddressBookRepository(db *sql.DB) *mysqlAddressBookRepository {
	return &mysqlAddressBookRepository{
		db: db,
	}
}

func (m *mysqlAddressBookRepository) Get(addressBookID string) (*model.AddressBook, error) {
	// Yeah this is stupid. This query will make sense if address book already has other info outside of its identifier
	var addressBook model.AddressBook
	m.db.QueryRow("SELECT UUID FROM ADDRESS_BOOK WHERE UUID=?", addressBookID).Scan(&addressBook.ID)

	contactRS, err := m.db.Query("SELECT UUID, NAME FROM CONTACT WHERE ADDRESS_BOOK_ID=?", addressBookID)
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

func (m *mysqlAddressBookRepository) Add(addressBook model.AddressBook) error {
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO ADDRESS_BOOK (UUID) VALUES (?)", addressBook.ID)
	if err != nil {
		return err
	}

	for _, contact := range addressBook.Contacts {
		_, err = tx.Exec("INSERT INTO CONTACT (UUID, NAME, ADDRESS_BOOK_ID) VALUES (?, ?, ?)", contact.ContactID, contact.Name, addressBook.ID)
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (m *mysqlAddressBookRepository) AddContact(addressBookID string, contactName string) error {
	return nil
}
