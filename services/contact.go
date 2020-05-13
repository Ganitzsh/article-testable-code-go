package services

import "github.com/Ganitzsh/write-testable-code-go/storage"

// Contacts is a global instance to be reachable by other packages to use
var Contacts ContactService = NewContactServiceImplementation()

// Contact is the final representation of
// a contact in the context of our application
// designed for the end user
type Contact struct {
	Name string `json:"name"`
}

// ContactService defines the behaviour of the
// application layer of the contact entity
// within the application
type ContactService interface {
	ListContacts() ([]*Contact, error)
}

// ContactServiceImplementation is the real life implementation
// of our service for the application
type ContactServiceImplementation struct {
	ContactStorage storage.ContactStorage
}

// NewContactServiceImplementation is a simple comstructor returning
// a ContactServiceImplementation initialized with a storage
func NewContactServiceImplementation() *ContactServiceImplementation {
	return &ContactServiceImplementation{
		ContactStorage: storage.NewContactStorageImplementation(),
	}
}

// ListContacts will retrieve the data from the storage and
// transform it to a usable entity our application
// can work with
func (s *ContactServiceImplementation) ListContacts() ([]*Contact, error) {
	// 1. Retrieve contacts from storage
	contactsFromStorage, err := s.ContactStorage.GetContactsFromStorage()
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
