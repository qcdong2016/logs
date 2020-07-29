# logs
🔥Simple, configurable Logging for Go.


```go
package main

import (
	"logs"
	"os"

	"github.com/gookit/color"
)

func ColorSelect(lv logs.Level) logs.ColorFunc {
	switch lv {
	case logs.LevelDebug:
		return color.Blue.Render
	case logs.LevelInfo:
		return color.Green.Render
	case logs.LevelWarn:
		return color.Yellow.Render
	case logs.LevelError:
		return color.Red.Render
	case logs.LevelFatal:
		return color.Red.Render
	case logs.LevelPanic:
		return color.Red.Render
	}
	return color.Yellow.Render
}

func main() {
	l := logs.NewLogger(
		logs.OptOutput(os.Stdout),
	)

	l.Use(
		logs.PartLevelColor(ColorSelect, logs.PartLayout("%-6s", logs.PartLevel())),
		logs.PartCaller(2, true),
		logs.PartTime("2006-01-02-15:04:05"),
		logs.PartColor(color.Blue.Render, logs.PartMessage()),
	)
	l.Log(logs.LevelDebug, "hello%d,%d", 123, 456)
	l.Log(logs.LevelInfo, "", 123, 456, "aaa")

	logs.Debug("234234")
	logs.Info("hello world")
}

```
