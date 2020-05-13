package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ganitzsh/write-testable-code-go/handlers"
	"github.com/Ganitzsh/write-testable-code-go/services"
	"github.com/Ganitzsh/write-testable-code-go/tests"
)

func servicesFuncMockGen(contactsToReturn []*services.Contact, err error) func() ([]*services.Contact, error) {
	return func() ([]*services.Contact, error) {
		return contactsToReturn, err
	}
}

func TestListContactHandlerSwap(t *testing.T) {
	testCases := []struct {
		name            string
		mockFunc        func() ([]*services.Contact, error)
		expectedPayload string
		expectedStatus  int
	}{
		{
			"service error, expected status 500 with error message",
			servicesFuncMockGen(nil, tests.ErrRandomError),
			`{"status":500,"message":"random error"}`,
			http.StatusInternalServerError,
		},
		{
			"no errors, contact list returned with status 200",
			servicesFuncMockGen(tests.Contacts, nil),
			`[{"name":"John Hammond"},{"name":"Dwayne Johnson"}]`,
			http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Log(testCase.name)

		// Swap function used by handler with the mock function
		// we created
		services.ListContacts = testCase.mockFunc

		req, err := http.NewRequest("GET", "/contacts", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.ListContactsHandlerSwap)

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
