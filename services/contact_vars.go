package services

import "github.com/Ganitzsh/write-testable-code-go/storage"

// ListContacts will retrieve the data from the storage and
// transform it to a usable entity our application
// can work with
var ListContacts = func() ([]*Contact, error) {
	// 1. Retrieve contacts from storage
	contactsFromStorage, err := storage.GetContactsFromStorage()
	if err != nil {
		return nil, err
	}

	// 2. Transform them to strip storage-specific properties
	contacts := []*Contact{}

	for _, model := range contactsFromStorage {
		contacts = append(contacts, &Contact{
			Name: model.Name,
		})
	}

	// 3. Return their value
	return contacts, nil
}
