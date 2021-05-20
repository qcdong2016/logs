package logs

type partTime struct {
	layout string
}

func PartTime(layout string) *partTime {
	return &partTime{layout: layout}
}

func (p *partTime) Output(l *Logger, args *Args) string {
	return args.Time.Format(p.layout)
}
