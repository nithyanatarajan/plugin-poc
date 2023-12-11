package shared

import (
	"github.com/google/uuid"
)

type ServiceProviderRPCServer struct {
	Impl ServiceProvider
}

func (s *ServiceProviderRPCServer) GetTransactionID(args any, resp *uuid.UUID) error {
	*resp = s.Impl.GetTransactionID()
	return nil
}
