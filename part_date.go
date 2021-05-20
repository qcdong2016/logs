package logs

type partDate struct {
	layout string
}

func PartDate(layout string) *partDate {
	return &partDate{layout: layout}
}

func (p *partDate) Output(l *Logger, args *Args) string {
	return args.Time.Format(p.layout)
}
