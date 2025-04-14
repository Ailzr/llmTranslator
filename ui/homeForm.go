package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"llmTranslator/pkg/llm"
	"llmTranslator/pkg/ocr"
)

func createHomeForm(mw *MainWindow) *widget.Form {
	var form *widget.Form

	ocrConnLabel := widget.NewLabel("")
	llmConnLabel := widget.NewLabel("")

	testBtn := widget.NewButton("测试", func() {
		ocrConnLabel.SetText("测试中...")
		llmConnLabel.SetText("测试中...")

		go func() {
			if ocr.OCRTest() {
				fyne.Do(func() {
					ocrConnLabel.SetText("OCR测试成功")
				})
			} else {
				fyne.Do(func() {
					ocrConnLabel.SetText("OCR测试失败")
				})
			}
		}()

		go func() {
			if llm.TestConn() {
				fyne.Do(func() {
					llmConnLabel.SetText("LLM测试成功")
				})
			} else {
				fyne.Do(func() {
					llmConnLabel.SetText("LLM测试失败")
				})
			}
		}()
	})

	//TODO 添加输入文本框，将翻译按键修改为仅调用LLM翻译文本框内容

	translateText := widget.NewLabel("")

	translateBtn := widget.NewButton("翻译", func() {
		mw.Translate()
	})

	form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "测试链接", Widget: testBtn},
			{Text: "OCR链接状态", Widget: ocrConnLabel},
			{Text: "LLM链接状态", Widget: llmConnLabel},
			{Text: "翻译", Widget: translateBtn},
			{Text: "翻译结果", Widget: translateText},
		},
	}

	return form
}
