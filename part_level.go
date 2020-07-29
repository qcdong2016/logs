package logs

type partLevel struct {
}

func PartLevel() IPart {
	return &partLevel{}
}

func (p *partLevel) Output(l *Logger, args *Args) string {

	switch args.Level {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	case LevelPanic:
		return "PANIC"
	}

	return ""
}
