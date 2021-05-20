package main

import (
	"os"

	"github.com/gookit/color"
	"github.com/qcdong2016/logs"
)

func main() {
	l := logs.NewLogger(
		logs.OptOutput(os.Stdout),
	)

	l.Use(
		logs.PartLevel(),
		logs.PartCaller(2, true),
		logs.PartDate("2006-01-02"),
		logs.PartTime("15:04:05"),
		logs.PartColor(color.Blue.Render, logs.PartMessage()),
	)
	l.Log(logs.LevelDebug, "hello%d,%d", 123, 456)
	l.Log(logs.LevelInfo, "", 123, 456, "aaa")

	logs.Debug("234234")
	logs.Info("hello world")
}
