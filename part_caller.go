package logs

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type partCaller struct {
	skip      int
	shortFile bool
}

func PartCaller(skip int, shortFile bool) IPart {
	return &partCaller{skip: skip, shortFile: shortFile}
}

func (p *partCaller) Output(l *Logger, args *Args) string {

	_, file, line, ok := runtime.Caller(p.skip)
	if !ok {
		file = "???"
		line = 0
	}

	if p.shortFile {
		file = filepath.Base(file)
	}

	return fmt.Sprintf("%s:%d", file, line)
}
