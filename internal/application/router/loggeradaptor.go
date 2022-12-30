package router

import (
	"fmt"

	"github.com/LakeevSergey/mailer/internal/application"
)

type LoggerAdaptor struct {
	logger application.Logger
}

func NewLoggerAdaptor(logger application.Logger) *LoggerAdaptor {
	return &LoggerAdaptor{
		logger: logger,
	}
}

func (a *LoggerAdaptor) Print(v ...interface{}) {
	a.logger.Info(fmt.Sprintln(v...))
}
