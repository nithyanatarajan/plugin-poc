package shared

import (
	"net/rpc"

	"github.com/google/uuid"
	"github.com/hashicorp/go-plugin"
)

type ServiceProvider interface {
	GetTransactionID() uuid.UUID
	GetName() string
}

type ServiceProviderPlugin struct {
	Impl ServiceProvider
}

func (p *ServiceProviderPlugin) Server(*plugin.MuxBroker) (any, error) {
	return &ServiceProviderRPCServer{Impl: p.Impl}, nil
}

func (ServiceProviderPlugin) Client(_ *plugin.MuxBroker, c *rpc.Client) (
	any, error) {
	return &ServiceProviderRPCClient{client: c}, nil
}
