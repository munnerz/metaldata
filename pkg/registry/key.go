package registry

import "strings"

const (
	keySeperator = "/"
)

// Key is a metadata key, forward-slash (/) separated.
type Key string

// Split splits a Key into it's parts by the key separator
func (k Key) Split() []string {
	return strings.Split(string(k), keySeperator)
}

// NewKey takes a list of strings representing a metadata key and returns a Key
// resource for it
func NewKey(s ...string) Key {
	var strs []string
	for _, str := range s {
		substrs := strings.Split(str, keySeperator)
		strs = append(strs, substrs...)
	}
	return Key(strings.Join(strs, keySeperator))
}
