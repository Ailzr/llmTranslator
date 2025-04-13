package ui

import (
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/viper"
)

func createOCRForm() *widget.Form {

	var form *widget.Form

	//设置语言下拉框
	langCombo := widget.NewSelect(
		[]string{"日语", "英语"}, func(s string) {
			switch s {
			case "日语":
				viper.Set("ocr.lang", "japan")
			case "英语":
				viper.Set("ocr.lang", "en")
			default:
				viper.Set("ocr.lang", "japan")
			}
		})
	langSet := viper.GetString("ocr.lang")
	switch langSet {
	case "japan":
		langCombo.SetSelected("日语")
	case "en":
		langCombo.SetSelected("英语")
	}

	apiKeyEntry := widget.NewEntry()
	apiKeySet := viper.GetString("ocr.api_key")
	apiKeyEntry.SetPlaceHolder("请输入API Key")
	apiKeyEntry.SetText(apiKeySet)
	apiKeyEntry.Disabled()

	//设置ocr引擎下拉框
	ocrEngineCombo := widget.NewSelect(
		[]string{"paddle-ocr", "baidu-ocr"}, func(s string) {
			switch s {
			case "paddle-ocr":
				viper.Set("ocr.engine", "paddle-ocr")
				apiKeyEntry.Disable()
			case "baidu-ocr":
				viper.Set("ocr.engine", "baidu-ocr")
				apiKeyEntry.Enable()
			default:
				viper.Set("ocr.engine", "paddle-ocr")
			}
		})
	ocrEngineCombo.SetSelected(viper.GetString("ocr.engine"))

	form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "需要翻译的语言", Widget: langCombo},
			{Text: "OCR引擎", Widget: ocrEngineCombo},
			{Text: "API KEY", Widget: apiKeyEntry},
		},
		SubmitText: "保存设置",
		OnSubmit: func() {
			viper.Set("ocr.api_key", apiKeyEntry.Text)
			_ = viper.WriteConfig()
		},
		CancelText: "取消",
		OnCancel: func() {
			reserOCRForm(langCombo, ocrEngineCombo, apiKeyEntry)
		},
	}

	return form
}

// 新增：重置控件的函数
func reserOCRForm(langCombo *widget.Select, ocrEngineCombo *widget.Select, apiKeyEntry *widget.Entry) {
	// 1. 重新加载配置文件（清除内存中的未保存修改）
	if err := viper.ReadInConfig(); err != nil {
		// 处理错误（可选）
		return
	}

	// 2. 恢复语言选择框
	langSet := viper.GetString("ocr.lang")
	switch langSet {
	case "japan":
		langCombo.SetSelected("日语")
	case "en":
		langCombo.SetSelected("英语")
	default:
		langCombo.ClearSelected()
	}

	// 3. 恢复OCR引擎选择框
	engineSet := viper.GetString("ocr.engine")
	switch engineSet {
	case "paddle-ocr":
		ocrEngineCombo.SetSelected("paddle-ocr")
		apiKeyEntry.Disable()
	case "baidu-ocr":
		ocrEngineCombo.SetSelected("baidu-ocr")
		apiKeyEntry.Enable()
	default:
		ocrEngineCombo.SetSelected("paddle-ocr")
	}

	// 4. 恢复API Key输入框
	apiKeySet := viper.GetString("ocr.api_key")
	apiKeyEntry.SetText(apiKeySet)
}
