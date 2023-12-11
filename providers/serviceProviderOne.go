package main

import (
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
)

type ServiceProviderOne struct {
	logger hclog.Logger
}

func (s *ServiceProviderOne) GetTransactionID() uuid.UUID {
	s.logger.Debug("\n<<<<<<<<<<<<<<<<<<<<Log from ServiceProviderOne.GetTransactionID>>>>>>>>>>>>>>>>>>>>")
	return uuid.New()
}
