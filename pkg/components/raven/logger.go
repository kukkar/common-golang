package raven

import (
	"github.com/kukkar/common-golang/pkg/logger"
)

var _ Logger = (*DefaultLogger)(nil)

// This logger uses default getsetgo logger.
type DefaultLogger struct {
}

func (this *DefaultLogger) Debug(a ...interface{}) {
	logger.Info(a)
}

func (this *DefaultLogger) Info(a ...interface{}) {
	logger.Info(a)
}
func (this *DefaultLogger) Warning(a ...interface{}) {
	logger.Warning(a)
}
func (this *DefaultLogger) Error(a ...interface{}) {
	logger.Error(a)
}

func (this *DefaultLogger) Fatal(a ...interface{}) {
	logger.Error(a)
}
