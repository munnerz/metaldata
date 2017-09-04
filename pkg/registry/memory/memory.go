package memory

import (
	"fmt"
	"log"

	"github.com/munnerz/metaldata/pkg/registry"
	"github.com/munnerz/metaldata/pkg/util/errors"
)

type memory struct {
	meta map[string]string
}

var _ registry.Interface = &memory{}

func NewMemory(meta map[string]string) registry.Interface {
	if meta == nil {
		meta = make(map[string]string)
	}
	return &memory{meta}
}

func (m *memory) Get(src registry.SourceRef, key registry.Key) (string, error) {
	mapKey := fmt.Sprintf("%s:%s", src, key)
	if val, ok := m.meta[mapKey]; ok {
		return val, nil
	}
	return "", errors.NewNotFound(fmt.Sprintf("could not find key '%v' for '%v'", key, src))
}

func (m *memory) Set(src registry.SourceRef, key registry.Key, val string) error {
	mapKey := fmt.Sprintf("%s:%s", src, key)
	m.meta[mapKey] = val
	log.Printf("setting key '%s' for ref '%s'", key, src)
	return nil
}
