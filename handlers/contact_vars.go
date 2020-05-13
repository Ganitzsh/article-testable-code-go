package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ganitzsh/write-testable-code-go/services"
)

// ListContactsHandlerSwap will be served over HTTP
// to allow the communication of the contacts list using
// the global variable
func ListContactsHandlerSwap(w http.ResponseWriter, r *http.Request) {
	contacts, err := services.ListContacts()

	if err != nil {
		handleHTTPError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contacts)
}
