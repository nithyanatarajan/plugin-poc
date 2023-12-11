package main

import (
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
)

type ServiceProviderOne struct {
	logger hclog.Logger
}

func (s *ServiceProviderOne) GetTransactionID() uuid.UUID {
	msg := "\n<<<<<<<<<<<<<<<<<<<<" +
		"Log from ServiceProviderOne.GetTransactionID" +
		">>>>>>>>>>>>>>>>>>>>"
	s.logger.Debug(msg)
	return uuid.New()
}
