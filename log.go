package logs

var DefaultLogger = NewLogger().Use(
	PartLevel(),
	PartTime("2006-01-02-15:04:05"),
	PartCaller(3, true),
	PartMessage(),
)

type Level int8

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func Debug(i ...interface{}) {
	DefaultLogger.Log(LevelDebug, "", i...)
}

func Debugf(format string, i ...interface{}) {
	DefaultLogger.Log(LevelDebug, format, i...)
}

func Info(i ...interface{}) {
	DefaultLogger.Log(LevelInfo, "", i...)
}

func Infof(format string, i ...interface{}) {
	DefaultLogger.Log(LevelInfo, format, i...)
}

func Warn(i ...interface{}) {
	DefaultLogger.Log(LevelWarn, "", i...)
}

func Warnf(format string, i ...interface{}) {
	DefaultLogger.Log(LevelWarn, format, i...)
}

func Error(i ...interface{}) {
	DefaultLogger.Log(LevelError, "", i...)
}

func Errorf(format string, i ...interface{}) {
	DefaultLogger.Log(LevelError, format, i...)
}

func Fatal(i ...interface{}) {
	DefaultLogger.Log(LevelFatal, "", i...)
}

func Fatalf(format string, i ...interface{}) {
	DefaultLogger.Log(LevelFatal, format, i...)
}

func Panic(i ...interface{}) {
	DefaultLogger.Log(LevelPanic, "", i...)
}

func Panicf(format string, i ...interface{}) {
	DefaultLogger.Log(LevelPanic, format, i...)
}
