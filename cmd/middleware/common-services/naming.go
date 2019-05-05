package services

import (
	"fmt"

	"github.com/joaoaneto/pd-middleware/cmd/middleware/distribution"
)

// Naming is an interface to implement naming service capabilities.
type Naming interface {
	Lookup(name string) distribution.ClientProxy
}

// NamingService represents the object naming service.
type NamingService struct {
	repository NamingRepository
}

// NamingRecordNotFoundErr represents a error raised when a record isn't found.
type NamingRecordNotFoundErr struct {
	name string
}

func (err *NamingRecordNotFoundErr) Error() string {
	return fmt.Sprintf("%s - %s", err.name, "naming record not found")
}

// NewNamingService creates a brand NamingService instance.
func NewNamingService() NamingService {
	repository := NewNamingRepository()
	return NamingService{repository}
}

// Lookup is method that given a name will search for a proxy assoaciated to it.
func (namingService *NamingService) Lookup(name string) (interface{}, error) {
	names := namingService.repository.List()
	for _, value := range names {
		if value.name == name {
			return value.proxy, nil
		}
	}

	return nil, &NamingRecordNotFoundErr{name}
}
