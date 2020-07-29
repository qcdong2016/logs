package logs

import (
	"time"
)

type partTime struct {
	layout string
}

func PartTime(layout string) IPart {
	return &partTime{layout: layout}
}

func (p *partTime) Output(l *Logger, args *Args) string {
	return time.Now().Format(p.layout)
}
