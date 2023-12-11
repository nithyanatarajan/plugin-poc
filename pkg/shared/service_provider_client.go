package shared

import (
	"net/rpc"

	"github.com/google/uuid"
)

type ServiceProviderRPCClient struct{ client *rpc.Client }

func (c *ServiceProviderRPCClient) GetTransactionID() uuid.UUID {
	var resp uuid.UUID
	err := c.client.Call("Plugin.GetTransactionID", new(any), &resp)
	if err != nil {
		panic(err)
	}

	return resp
}

func (c *ServiceProviderRPCClient) GetName() string {
	var resp string
	err := c.client.Call("Plugin.GetName", new(any), &resp)
	if err != nil {
		panic(err)
	}

	return resp
}
