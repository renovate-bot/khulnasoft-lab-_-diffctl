package enumerator

import (
	"github.com/khulnasoft-lab/driftctl/pkg/iac/config"
	"github.com/khulnasoft-lab/driftctl/pkg/iac/terraform/state/backend"
	"github.com/sirupsen/logrus"
)

type StateEnumerator interface {
	Origin() string
	Enumerate() ([]string, error)
}

func GetEnumerator(config config.SupplierConfig, opts *backend.Options) (StateEnumerator, error) {

	switch config.Backend {
	case backend.BackendKeyFile:
		return NewFileEnumerator(config), nil
	case backend.BackendKeyS3:
		return NewS3Enumerator(config), nil
	case backend.BackendKeyGS:
		return NewGSEnumerator(config)
	case backend.BackendKeyAzureRM:
		return NewAzureRMEnumerator(config, opts.AzureRMBackendOptions)
	}

	logrus.WithFields(logrus.Fields{
		"backend": config.Backend,
	}).Debug("No enumerator for backend")

	return nil, nil
}
