package main

import (
	_ "llmTranslator/configs"
	"llmTranslator/ui"
)

func main() {
	ui.Init()
	ui.ShowAndRun()
}
