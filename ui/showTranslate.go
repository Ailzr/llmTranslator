package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"llmTranslator/logHelper"
	"llmTranslator/pkg/llm"
	"llmTranslator/pkg/ocr"
	"llmTranslator/utils"
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

	//TODO 翻译框大小从配置中读取，动态修改写入到配置中保存

	mw.TranslatorWindow.Resize(fyne.NewSize(800, 300))
	mw.TranslatorWindow.Show()
}

func (mw *MainWindow) ShowTranslate(text string) {
	if contentShow != nil {
		contentShow.SetText(text)
	} else {
		dialog.ShowInformation("错误", "未选择需要翻译的部分", mw.Window)
	}
}

func (mw *MainWindow) Translate() {
	//TODO 将主窗口也隐藏，避免点击翻译按钮时截屏被遮挡

	mw.ShowTranslate("翻译中...")
	mw.TranslatorWindow.Hide()

	img, err := utils.CaptureImg()
	utils.SaveImgToPng(img, "tmp")
	if err != nil {
		logHelper.Error(err.Error())
		logHelper.WriteLog(err.Error())
		mw.ShowTranslate("截图失败")
	}

	ocrResult := ocr.GetOCRResult()
	if ocrResult == "" {
		mw.ShowTranslate("未识别到文字")
	}
	mw.TranslatorWindow.Show()
	result := llm.Translate(ocrResult, "简体中文")
	mw.ShowTranslate(result)
}
