package logger

type Common interface {
	Info(msg string)
	Infof(format string, args ...interface{})
	Fatal(err error, msg string)
	Fatalf(err error, format string, args ...interface{})
	Error(err error, msg string)
	Errorf(err error, format string, args ...interface{})
}
