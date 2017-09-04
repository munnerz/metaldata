package registry

type Interface interface {
	Get(SourceRef, Key) (string, error)
	Set(SourceRef, Key, string) error
}

// SourceRef identifies something to get a metadata key for. This could be a
// source IP address or an instance name.
type SourceRef string
