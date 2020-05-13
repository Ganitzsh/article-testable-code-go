package tests

import (
	"errors"

	"github.com/Ganitzsh/write-testable-code-go/services"
	"github.com/Ganitzsh/write-testable-code-go/storage"
)

var (
	ErrRandomError = errors.New("random error")
	JohnModel      = storage.ContactModel{ID: "contact1", Name: "John Hammond"}
	DwayneModel    = storage.ContactModel{ID: "contact2", Name: "Dwayne Johnson"}
	John           = services.Contact{Name: JohnModel.Name}
	Dwayne         = services.Contact{Name: DwayneModel.Name}
	Contacts       = []*services.Contact{&John, &Dwayne}
	ContactModels  = []*storage.ContactModel{&JohnModel, &DwayneModel}
)
