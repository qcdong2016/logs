package logs

type ColorFunc func(...interface{}) string

type partColor struct {
	colorfn ColorFunc
	part    IPart
}

func PartColor(fn ColorFunc, part IPart) IPart {
	return &partColor{colorfn: fn, part: part}
}

func (p *partColor) Output(l *Logger, args *Args) string {
	return p.colorfn(p.part.Output(l, args))
}
