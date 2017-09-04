package ip

import (
	"log"
	"net"
	"net/http"

	"github.com/munnerz/metaldata/pkg/authorisation"
	"github.com/munnerz/metaldata/pkg/registry"
)

type ipAuthorisation struct {
}

func NewIPAuthorisation() authorisation.Interface {
	return &ipAuthorisation{}
}

func (i *ipAuthorisation) Authorize(req *http.Request) (registry.SourceRef, error) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)

	if err != nil {
		return "", err
	}

	log.Printf("authorised remote address '%s'", ip)
	return registry.SourceRef(ip), nil
}
