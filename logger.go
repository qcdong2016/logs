package logs

import (
	"bytes"
	"io"
	"os"
)

type Args struct {
	Level  Level
	Format string
	Args   []interface{}
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
	}

	buf := &bytes.Buffer{}

	for _, part := range l.parts {
		buf.WriteString(part.Output(l, msg))
		buf.WriteByte(' ')
	}
	buf.WriteByte('\n')

	l.output.Write(buf.Bytes())
}
