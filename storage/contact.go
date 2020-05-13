package storage

import "github.com/Ganitzsh/write-testable-code-go/database"

// ContactModel is a specific type created
// to comply with our storage and the possible
// feature we need to build our application
type ContactModel struct {
	ID   string `storage:primary`
	Name string `storage:unique`
}

// ContactStorage defines the behaviour
// of the sotrage interface and its
// set of functions
type ContactStorage interface {
	GetContactsFromStorage() ([]*ContactModel, error)
}

type contactStorageImplemetation struct{}

func NewContactStorageImplementation() *contactStorageImplemetation {
	return &contactStorageImplemetation{}
}

// GetContactsFromStorage will retrieve the Contacts
// from the database and return them
func (*contactStorageImplemetation) GetContactsFromStorage() ([]*ContactModel, error) {
	contacts := []*ContactModel{}

	if err := database.FindAll("contacts", &contacts); err != nil {
		return nil, err
	}

	return contacts, nil
}
