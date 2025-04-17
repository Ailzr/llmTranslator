package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"llmTranslator/pkg/llm"
	"llmTranslator/pkg/ocr"
)

func createHomeForm() *widget.Form {
	var form *widget.Form

	ocrConnLabel := canvas.NewText("", color.White)
	llmConnLabel := canvas.NewText("", color.White)

	testOCR := widget.NewButton("OCR测试", func() {
		ocrConnLabel.Text = "测试中..."
		ocrConnLabel.Color = color.White

		go func() {
			if ocr.OCRTest() {
				fyne.Do(func() {
					ocrConnLabel.Text = "OCR测试成功"
					ocrConnLabel.Color = color.NRGBA{G: 0x80, A: 0xff}
				})
			} else {
				fyne.Do(func() {
					ocrConnLabel.Text = "OCR测试失败"
					ocrConnLabel.Color = color.NRGBA{R: 0x80, A: 0xff}
				})
			}
		}()

	})

	testLLM := widget.NewButton("LLM测试", func() {
		llmConnLabel.Text = "测试中..."
		llmConnLabel.Color = color.White
		go func() {
			if llm.TestConn() {
				fyne.Do(func() {
					llmConnLabel.Text = "LLM测试成功"
					llmConnLabel.Color = color.NRGBA{G: 0x80, A: 0xff}
				})
			} else {
				fyne.Do(func() {
					llmConnLabel.Text = "LLM测试失败"
					llmConnLabel.Color = color.NRGBA{R: 0x80, A: 0xff}
				})
			}
		}()
	})

	translateInput := widget.NewMultiLineEntry()
	translateInput.SetPlaceHolder("请输入需要翻译的文本")
	translateText := widget.NewLabel("")
	translateText.Wrapping = fyne.TextWrapBreak

	translateBtn := widget.NewButton("翻译", func() {
		if translateInput.Text == "" {
			return
		}
		translateText.SetText("翻译中...")
		go func() {
			text := llm.Translate(translateInput.Text)
			fyne.Do(func() {
				translateText.SetText(text)
			})
		}()
	})
	form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "测试链接", Widget: container.NewGridWithColumns(2, testOCR, testLLM)},
			{Text: "链接状态", Widget: container.NewGridWithColumns(2, ocrConnLabel, llmConnLabel)},
			{Text: "输入文本", Widget: translateInput},
			{Text: "翻译", Widget: translateBtn},
			{Text: "翻译结果", Widget: translateText},
		},
	}

	return form
}
