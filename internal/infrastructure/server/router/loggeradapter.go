package router

import (
	"fmt"

	"github.com/LakeevSergey/mailer/internal/infrastructure"
)

type LoggerAdapter struct {
	logger infrastructure.Logger
}

func NewLoggerAdapter(logger infrastructure.Logger) *LoggerAdapter {
	return &LoggerAdapter{
		logger: logger,
	}
}

func (a *LoggerAdapter) Print(v ...interface{}) {
	a.logger.Info(fmt.Sprintln(v...))
}
