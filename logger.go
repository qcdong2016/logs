package logs

import (
	"bytes"
	"io"
	"os"
	"time"
)

type Args struct {
	Level  Level
	Format string
	Args   []interface{}
	Time   time.Time
}

type IPart interface {
	Output(l *Logger, args *Args) string
}

type Logger struct {
	parts  []IPart
	output io.Writer
}

func NewLogger(options ...OptionFunc) *Logger {
	l := &Logger{}
	l.parts = []IPart{}
	l.output = os.Stdout

	for _, o := range options {
		o(l)
	}

	return l
}

func (l *Logger) Use(parts ...IPart) *Logger {
	for _, v := range parts {
		l.parts = append(l.parts, v)
	}
	return l
}

func (l *Logger) Log(level Level, format string, args ...interface{}) {

	msg := &Args{
		Level:  level,
		Format: format,
		Args:   args,
		Time:   time.Now(),
	}

	buf := &bytes.Buffer{}

	for _, part := range l.parts {
		buf.WriteString(part.Output(l, msg))
		buf.WriteByte(' ')
	}
	buf.WriteByte('\n')

	l.output.Write(buf.Bytes())
}

func (l *Logger) Debug(i ...interface{}) {
	l.Log(LevelDebug, "", i...)
}

func (l *Logger) Debugf(format string, i ...interface{}) {
	l.Log(LevelDebug, format, i...)
}

func (l *Logger) Info(i ...interface{}) {
	l.Log(LevelInfo, "", i...)
}

func (l *Logger) Infof(format string, i ...interface{}) {
	l.Log(LevelInfo, format, i...)
}

func (l *Logger) Warn(i ...interface{}) {
	l.Log(LevelWarn, "", i...)
}

func (l *Logger) Warnf(format string, i ...interface{}) {
	l.Log(LevelWarn, format, i...)
}

func (l *Logger) Error(i ...interface{}) {
	l.Log(LevelError, "", i...)
}

func (l *Logger) Errorf(format string, i ...interface{}) {
	l.Log(LevelError, format, i...)
}

func (l *Logger) Fatal(i ...interface{}) {
	l.Log(LevelFatal, "", i...)
}

func (l *Logger) Fatalf(format string, i ...interface{}) {
	l.Log(LevelFatal, format, i...)
}

func (l *Logger) Panic(i ...interface{}) {
	l.Log(LevelPanic, "", i...)
}

func (l *Logger) Panicf(format string, i ...interface{}) {
	l.Log(LevelPanic, format, i...)
}
