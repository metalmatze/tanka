package client

import (
	"github.com/Masterminds/semver"
)

type Client interface {
	Get(namespace, kind, name string) (map[string]interface{}, error)
	GetByLabels(labels map[string]interface{}) ([]map[string]interface{}, error)

	// Apply the configuration to the cluster. `data` must contain a plaintext
	// format that is `kubectl-apply(1)` compatible
	Apply(data []byte) error

	// Delete the specified object from the cluster
	Delete(namespace, kind, name string) error
	DeleteByLabels(labels map[string]interface{}) error

	// Info returns known informational data about the client. Best effort based,
	// fields of `ClientInfo` that cannot be stocked with valuable data, e.g.
	// due to an error shall be left nil.
	Info() *Info
}

type Info struct {
	ClientVersion    *semver.Version
	ServerVersion    *semver.Version
	Context, Cluster map[string]interface{}
}
