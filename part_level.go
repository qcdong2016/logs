package logs

type Level int8

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

type partLevel struct {
}

func PartLevel() *partLevel {
	return &partLevel{}
}

func (p *partLevel) Output(l *Logger, args *Args) string {

	switch args.Level {
	case LevelDebug:
		return "DBG"
	case LevelInfo:
		return "INF"
	case LevelWarn:
		return "WAN"
	case LevelError:
		return "ERR"
	case LevelFatal:
		return "FTL"
	case LevelPanic:
		return "PNC"
	}

	return ""
}
