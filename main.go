package main

import (
	_ "llmTranslator/configs"
	"llmTranslator/hotkey"
	"llmTranslator/ui"
)

func main() {
	hotkey.AddTranslateHotKey(ui.GetMainWindow())
	hotkey.AddCaptureRectangleHotKey(ui.GetMainWindow())
	ui.ShowAndRun()
}
