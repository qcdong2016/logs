package logs

type ColorLevelSelector func(lv Level) ColorFunc

type partLevelColor struct {
	colorfn ColorLevelSelector
	part    IPart
}

func PartLevelColor(colorfn ColorLevelSelector, part IPart) IPart {
	return &partLevelColor{colorfn: colorfn, part: part}
}

func (p *partLevelColor) Output(l *Logger, args *Args) string {
	return p.colorfn(args.Level)(p.part.Output(l, args))
}
