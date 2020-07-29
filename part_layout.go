package logs

import "fmt"

type partLayout struct {
	layout string
	part   IPart
}

func PartLayout(layout string, part IPart) IPart {
	return &partLayout{layout: layout, part: part}
}

func (p *partLayout) Output(l *Logger, args *Args) string {
	return fmt.Sprintf(p.layout, p.part.Output(l, args))
}
