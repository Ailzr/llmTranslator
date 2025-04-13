package ui

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"llmTranslator/pkg/llm"
	"llmTranslator/pkg/ocr"
)

func createHomeForm(mw *MainWindow) *widget.Form {
	var form *widget.Form

	connLabel := widget.NewLabel("")

	testBtn := widget.NewButton("测试", func() {
		connLabel.SetText("测试中...")
		mw.App.Driver().DoFromGoroutine(func() {
			if llm.TestConn() {
				connLabel.SetText(fmt.Sprintf("连接成功"))
			} else {
				connLabel.SetText("连接失败")
			}
		}, false)
	})

	translateText := widget.NewLabel("")

	translateBtn := widget.NewButton("翻译", func() {
		translateText.SetText("翻译中...")
		mw.App.Driver().DoFromGoroutine(func() {
			text := ocr.GetOCRResult()
			if text == "" {
				translateText.SetText("未识别到文字")
				return
			}
			result := llm.Translate(text, "zh")
			translateText.SetText(result)
		}, false)
	})

	form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "测试链接", Widget: testBtn},
			{Text: "链接状态", Widget: connLabel},
			{Text: "翻译", Widget: translateBtn},
			{Text: "翻译结果", Widget: translateText},
		},
	}

	return form
}
