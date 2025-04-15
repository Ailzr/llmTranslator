package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"llmTranslator/configs"
	"llmTranslator/langMap"
	"llmTranslator/pkg/ocr"
	"slices"
)

var apiKeyEntry, apiSecretEntry *widget.Entry
var ocrProvider string
var ocrForm *widget.Form

// 创建OCR表单
func createOCRForm() *widget.Form {

	//创建表单
	ocrForm = &widget.Form{}

	//读取配置
	ocrLang := configs.Setting.OCR.Lang
	ocrProvider = configs.Setting.OCR.Provider

	//设置语言下拉框
	langCombo := widget.NewSelect([]string{"日语", "英语"}, nil)
	//根据当前设置的语言设置下拉框的选中项
	langCombo.SetSelected(langMap.LangMap[ocrLang])
	//将语言下拉框添加到表单中
	ocrForm.AppendItem(widget.NewFormItem("需要翻译的语言", langCombo))

	//创建API Key输入框
	apiKeyEntry = widget.NewEntry()
	//创建API Secret输入框
	apiSecretEntry = widget.NewEntry()

	apiKeyEntry.Password = true
	apiSecretEntry.Password = true

	//设置ocr提供者下拉框
	ocrProviderCombo := widget.NewSelect(
		[]string{"paddle", "dango", "baidu"}, func(s string) {
			ocrProvider = s
			go setApiInfo()
		})
	//设置ocr提供者下拉框的选中项
	ocrProviderCombo.SetSelected(ocrProvider)
	//将ocr提供者下拉框添加到表单中
	ocrForm.AppendItem(widget.NewFormItem("OCR提供者", ocrProviderCombo))
	//将API Key输入框添加到表单中
	ocrForm.AppendItem(widget.NewFormItem("API KEY", apiKeyEntry))
	//将API Secret输入框添加到表单中
	ocrForm.AppendItem(widget.NewFormItem("API Secret", apiSecretEntry))

	//设置表单的提交按钮文本
	ocrForm.SubmitText = "保存"
	//设置表单的提交事件
	ocrForm.OnSubmit = func() {
		ocrLang = langMap.LangUnMap[langCombo.Selected]

		configs.Setting.OCR.Lang = ocrLang
		configs.Setting.OCR.Provider = ocrProvider
		switch ocrProvider {
		case "baidu":
			configs.Setting.OCR.Baidu.APIKey = apiKeyEntry.Text
			configs.Setting.OCR.Baidu.APISecret = apiSecretEntry.Text
		}
		configs.WriteSettingToFile()

		dialog.ShowInformation("提示", "保存成功", mw.Window)
	}
	//设置表单的取消按钮文本
	ocrForm.CancelText = "取消"
	//设置表单的取消事件
	ocrForm.OnCancel = func() {
		//重置表单
		langCombo.SetSelected(langMap.LangUnMap[ocrLang])
		ocrProvider = configs.Setting.OCR.Provider
		ocrProviderCombo.SetSelected(ocrProvider)
		go setApiInfo()
	}

	//返回表单
	return ocrForm
}

func setApiInfo() {
	if slices.Contains(ocr.LocalOCR, ocrProvider) {
		fyne.Do(func() {
			apiKeyEntry.Disable()
			apiSecretEntry.Disable()
			apiKeyEntry.SetText("")
			apiSecretEntry.SetText("")
			ocrForm.Refresh()
		})
	} else {
		fyne.Do(func() {
			apiKeyEntry.Enable()
			apiSecretEntry.Enable()
			switch ocrProvider {
			case "baidu":
				apiKeyEntry.SetText(configs.Setting.OCR.Baidu.APIKey)
				apiSecretEntry.SetText(configs.Setting.OCR.Baidu.APISecret)
			}
			ocrForm.Refresh()
		})
	}
}
