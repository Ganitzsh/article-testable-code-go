package storage

import "github.com/Ganitzsh/write-testable-code-go/database"

// GetContactsFromStorage will retrieve the Contacts
// from the database and return them
var GetContactsFromStorage = func() ([]*ContactModel, error) {
	contacts := []*ContactModel{}

	if err := database.FindAll("contacts", &contacts); err != nil {
		return nil, err
	}

	return contacts, nil
}
