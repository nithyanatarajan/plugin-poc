package registry

import (
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/nithyanatarajan/plugin-poc/pkg/shared"
	"github.com/nithyanatarajan/plugin-poc/providers/logging"
)

type ServiceProviderOne struct {
	logger hclog.Logger
}

func (*ServiceProviderOne) GetTransactionID() uuid.UUID {
	return uuid.New()
}

func (*ServiceProviderOne) GetName() string {
	return "ServiceProviderOne"
}

func init() {
	plugin := shared.ServiceProviderPlugin{
		Impl: &ServiceProviderOne{
			logger: logging.GetLogger(),
		},
	}
	RegisterServiceProvider(&plugin)
}
