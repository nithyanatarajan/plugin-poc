package logging

import (
	"os"

	"github.com/hashicorp/go-hclog"
)

var instance hclog.Logger

func GetLogger() hclog.Logger {
	if instance == nil {
		instance = hclog.New(&hclog.LoggerOptions{
			Name:       "PROVIDERS.LOGGER",
			Level:      hclog.Trace,
			Output:     os.Stderr,
			JSONFormat: true,
		})
	}
	return instance
}
