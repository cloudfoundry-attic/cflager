package main

import (
	"errors"

	"github.com/cloudfoundry-incubator/cf-lager"
)

func main() {
	logger := cf_lager.New("cf-lager-integration")

	logger.Debug("some", "debug", "for you")
	logger.Info("some", "info", "for you")
	logger.Error("some", "error", "for you", errors.New("error"))
	logger.Fatal("some", "fatal", "for you", errors.New("fatal"))
}
