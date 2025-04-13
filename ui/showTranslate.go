package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var contentShow *widget.Label

func (mw *MainWindow) CreateShowWindow() {
	if mw.TranslatorWindow != nil {
		mw.TranslatorWindow.Close()
		mw.TranslatorWindow = nil
	}
	mw.TranslatorWindow = mw.App.NewWindow("Translator")

	mw.TranslatorWindow.SetIcon(resourceLlmTranslatorPng)

	contentShow = widget.NewLabel("")
	contentShow.Wrapping = fyne.TextWrapWord

	mw.TranslatorWindow.SetContent(contentShow)

	mw.TranslatorWindow.Resize(fyne.NewSize(800, 300))
	mw.TranslatorWindow.Show()
}

func (mw *MainWindow) ShowTranslate(text string) {
	if contentShow != nil {
		contentShow.SetText(text)
	}
}
