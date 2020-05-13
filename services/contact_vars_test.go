package services_test

import (
	"testing"

	"github.com/Ganitzsh/write-testable-code-go/services"
	"github.com/Ganitzsh/write-testable-code-go/storage"
	"github.com/Ganitzsh/write-testable-code-go/tests"
)

func storageFuncMockGen(contactsToReturn []*storage.ContactModel, err error) func() ([]*storage.ContactModel, error) {
	return func() ([]*storage.ContactModel, error) {
		return contactsToReturn, err
	}
}

func TestListContactsSwap(t *testing.T) {
	testCases := []struct {
		name             string
		mockFunc         func() ([]*storage.ContactModel, error)
		expectedError    error
		expectedContacts []*services.Contact
	}{
		{
			"error with database",
			storageFuncMockGen(nil, tests.ErrRandomError),
			tests.ErrRandomError,
			nil,
		},
		{
			"no errors, contacts retrieved successfully",
			storageFuncMockGen(tests.ContactModels, nil),
			nil,
			tests.Contacts,
		},
	}

	for _, testCase := range testCases {
		t.Log(testCase.name)

		storage.GetContactsFromStorage = testCase.mockFunc

		contacts, err := services.ListContacts()
		if err != testCase.expectedError {
			t.Fatalf("Unexpected error, got %v want %v", err, testCase.expectedError)
		}
		if contacts != nil && testCase.expectedContacts == nil {
			t.Fatalf("Invalid contacts, got %v want %v", contacts, testCase.expectedContacts)
		}
	}
}
