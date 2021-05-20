package logs

var DefaultLogger = NewLogger().Use(
	PartLevel(),
	PartDate("2006-01-02"),
	PartTime("15:04:05"),
	PartCaller(4, true),
	PartMessage(),
)

func Debug(i ...interface{}) {
	DefaultLogger.Debug(i...)
}

func Debugf(format string, i ...interface{}) {
	DefaultLogger.Debugf(format, i...)
}

func Info(i ...interface{}) {
	DefaultLogger.Info(i...)
}

func Infof(format string, i ...interface{}) {
	DefaultLogger.Infof(format, i...)
}

func Warn(i ...interface{}) {
	DefaultLogger.Warn(i...)
}

func Warnf(format string, i ...interface{}) {
	DefaultLogger.Warnf(format, i...)
}

func Error(i ...interface{}) {
	DefaultLogger.Error(i...)
}

func Errorf(format string, i ...interface{}) {
	DefaultLogger.Errorf(format, i...)
}

func Fatal(i ...interface{}) {
	DefaultLogger.Fatal(i...)
}

func Fatalf(format string, i ...interface{}) {
	DefaultLogger.Fatalf(format, i...)
}

func Panic(i ...interface{}) {
	DefaultLogger.Panic(i...)
}

func Panicf(format string, i ...interface{}) {
	DefaultLogger.Panicf(format, i...)
}
