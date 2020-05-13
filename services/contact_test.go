package services_test

import (
	"testing"

	"github.com/Ganitzsh/write-testable-code-go/services"
	"github.com/Ganitzsh/write-testable-code-go/storage"
	"github.com/Ganitzsh/write-testable-code-go/tests"
)

// We can now implement a mock ContactStorage to
// be replaced in the ContactService used in our tests
type contactStorageMock struct {
	contactsToReturn []*storage.ContactModel
	err              error
}

func newContactStorageMock(contactsToReturn []*storage.ContactModel, err error) *contactStorageMock {
	return &contactStorageMock{contactsToReturn, err}
}

func (m *contactStorageMock) GetContactsFromStorage() ([]*storage.ContactModel, error) {
	return m.contactsToReturn, m.err
}

func TestListContactsInterface(t *testing.T) {
	testCases := []struct {
		name             string
		mockInstance     storage.ContactStorage
		expectedError    error
		expectedContacts []*services.Contact
	}{
		{
			"error with database",
			newContactStorageMock(nil, tests.ErrRandomError),
			tests.ErrRandomError,
			nil,
		},
		{
			"no errors, contacts retrieved successfully",
			newContactStorageMock(tests.ContactModels, nil),
			nil,
			tests.Contacts,
		},
	}

	contactService := services.NewContactServiceImplementation()
	for _, testCase := range testCases {
		t.Log(testCase.name)

		contactService.ContactStorage = testCase.mockInstance

		contacts, err := contactService.ListContacts()
		if err != testCase.expectedError {
			t.Fatalf("Unexpected error, got %v want %v", err, testCase.expectedError)
		}
		if contacts != nil && testCase.expectedContacts == nil {
			t.Fatalf("Invalid contacts, got %v want %v", contacts, testCase.expectedContacts)
		}
	}
}
