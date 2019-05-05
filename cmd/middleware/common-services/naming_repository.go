package services

import (
	"github.com/joaoaneto/pd-middleware/cmd/middleware/app"
	"github.com/joaoaneto/pd-middleware/cmd/middleware/distribution"
)

// NamingRepository represents the object name storage.
type NamingRepository struct {
	names []NamingRecord
}

// NamingRecord represents the object name itself.
type NamingRecord struct {
	name  string
	proxy interface{}
}

// NewNamingRepository creates a brand NewNamingRepository instance.
func NewNamingRepository() NamingRepository {
	client := distribution.NewClientProxy("localhost", 4000, 1)
	proxy := app.NewCalculatorProxy(client)

	record := NamingRecord{"Calculator", proxy}
	return NamingRepository{[]NamingRecord{record}}
}

// List will return all the records.
func (repository *NamingRepository) List() []NamingRecord {
	return repository.names
}
