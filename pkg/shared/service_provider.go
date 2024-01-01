package shared

import (
	"fmt"
	"net/rpc"

	"github.com/google/uuid"
	"github.com/hashicorp/go-plugin"
)

type ServiceProvider interface {
	GetTransactionID() uuid.UUID
	GetName() string
	HandleCallback([]byte) SomeStruct
}

type SomeStruct struct {
	Data string
}

func (s SomeStruct) String() string {
	return fmt.Sprintf("Data: %v", s.Data)
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
