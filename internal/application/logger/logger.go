package logger

import "github.com/rs/zerolog"

type Logger struct {
	logger zerolog.Logger
}

func NewLogger(logger zerolog.Logger) *Logger {
	return &Logger{logger: logger}
}

func (l *Logger) Debug(message string) {
	l.logger.Debug().Msg(message)
}
func (l *Logger) Info(message string) {
	l.logger.Info().Msg(message)
}
func (l *Logger) Warn(message string) {
	l.logger.Warn().Msg(message)
}
func (l *Logger) Error(message string) {
	l.logger.Error().Msg(message)
}

func (l *Logger) DebugErr(err error) {
	l.logger.Debug().Err(err).Send()
}
func (l *Logger) InfoErr(err error) {
	l.logger.Info().Err(err).Send()
}
func (l *Logger) WarnErr(err error) {
	l.logger.Warn().Err(err).Send()
}
func (l *Logger) ErrorErr(err error) {
	l.logger.Error().Err(err).Send()
}
