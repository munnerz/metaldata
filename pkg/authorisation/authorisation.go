package authorisation

import (
	"net/http"

	"github.com/munnerz/metaldata/pkg/registry"
)

type Interface interface {
	Authorize(*http.Request) (registry.SourceRef, error)
}
