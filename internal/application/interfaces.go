package application

type Logger interface {
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)

	DebugErr(err error)
	InfoErr(err error)
	WarnErr(err error)
	ErrorErr(err error)
}
