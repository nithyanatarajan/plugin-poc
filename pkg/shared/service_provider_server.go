package shared

import (
	"github.com/google/uuid"
)

type ServiceProviderRPCServer struct {
	Impl ServiceProvider
}

func (s *ServiceProviderRPCServer) GetTransactionID(_ any, resp *uuid.UUID,
) error {
	*resp = s.Impl.GetTransactionID()
	return nil
}

func (s *ServiceProviderRPCServer) GetName(_ any, resp *string) error {
	*resp = s.Impl.GetName()
	return nil
}
