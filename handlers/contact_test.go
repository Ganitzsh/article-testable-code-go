package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ganitzsh/write-testable-code-go/handlers"
	"github.com/Ganitzsh/write-testable-code-go/services"
	"github.com/Ganitzsh/write-testable-code-go/tests"
)

// We can now implement a mock ContactService to
// be replaced in the handlers used in our tests
type contactServiceMock struct {
	contactsToReturn []*services.Contact
	err              error
}

func newContactServiceMock(contactsToReturn []*services.Contact, err error) *contactServiceMock {
	return &contactServiceMock{contactsToReturn, err}
}

func (m *contactServiceMock) ListContacts() ([]*services.Contact, error) {
	return m.contactsToReturn, m.err
}

func TestListContactHandlerInterface(t *testing.T) {
	testCases := []struct {
		name            string
		mockInstance    *contactServiceMock
		expectedPayload string
		expectedStatus  int
	}{
		{
			"service error, expected status 500 with error message",
			newContactServiceMock(nil, tests.ErrRandomError),
			`{"status":500,"message":"random error"}`,
			http.StatusInternalServerError,
		},
		{
			"no errors, contact list returned with status 200",
			newContactServiceMock(tests.Contacts, nil),
			`[{"name":"John Hammond"},{"name":"Dwayne Johnson"}]`,
			http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Log(testCase.name)

		// Replace the instance of ContactService used by the handler
		// with our mock
		services.Contacts = testCase.mockInstance

		req, err := http.NewRequest("GET", "/contacts", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.ListContactsHandlerInterface)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != testCase.expectedStatus {
			t.Errorf("wrong status code: got %v want %v",
				status, testCase.expectedStatus)
		}

		if rr.Body.String() != testCase.expectedPayload+"\n" {
			t.Errorf("unexpected payload: got %v want %v",
				rr.Body.String(), testCase.expectedPayload)
		}
	}
}
