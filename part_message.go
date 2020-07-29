package logs

import (
	"bytes"
	"fmt"
)

type partMessage struct {
}

func PartMessage() IPart {
	return &partMessage{}
}

func (p *partMessage) Output(l *Logger, args *Args) string {
	if args.Format == "" {

		buf := bytes.Buffer{}

		for i, v := range args.Args {
			buf.WriteString(fmt.Sprint(v))
			if i < len(args.Args)-1 {
				buf.WriteByte(' ')
			}
		}

		return buf.String()
	}

	return fmt.Sprintf(args.Format, args.Args...)
}
