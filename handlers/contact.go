package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ganitzsh/write-testable-code-go/services"
)

// HttpError is a custom type to render errors
// for HTTP client properly
type HttpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func handleHTTPError(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(HttpError{
		Status:  status,
		Message: err.Error(),
	})
}

// ListContactsHandlerInterface will be served over HTTP
// to allow the communication of the contacts list using
// the service implementation of ContactService
func ListContactsHandlerInterface(w http.ResponseWriter, r *http.Request) {
	contacts, err := services.Contacts.ListContacts()

	if err != nil {
		handleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contacts)
}
