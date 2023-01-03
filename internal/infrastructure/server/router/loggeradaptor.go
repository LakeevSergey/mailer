package router

import (
	"fmt"

	"github.com/LakeevSergey/mailer/internal/infrastructure"
)

type LoggerAdaptor struct {
	logger infrastructure.Logger
}

func NewLoggerAdaptor(logger infrastructure.Logger) *LoggerAdaptor {
	return &LoggerAdaptor{
		logger: logger,
	}
}

func (a *LoggerAdaptor) Print(v ...interface{}) {
	a.logger.Info(fmt.Sprintln(v...))
}
