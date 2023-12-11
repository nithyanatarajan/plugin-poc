package registry

import (
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/nithyanatarajan/plugin-poc/pkg/shared"
	"github.com/nithyanatarajan/plugin-poc/providers/logging"
)

type ServiceProviderTwo struct {
	logger hclog.Logger
}

func (*ServiceProviderTwo) GetTransactionID() uuid.UUID {
	return uuid.New()
}

func (*ServiceProviderTwo) GetName() string {
	return "ServiceProviderTwo"
}

func init() {
	plugin := shared.ServiceProviderPlugin{
		Impl: &ServiceProviderTwo{
			logger: logging.GetLogger(),
		},
	}
	RegisterServiceProvider(&plugin)
}
