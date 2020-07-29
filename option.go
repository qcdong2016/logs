package logs

import "io"

type OptionFunc func(l *Logger)

func OptOutput(o io.Writer) OptionFunc {
	return func(l *Logger) {
		l.output = o
	}
}
