package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/viper"
	"llmTranslator/langMap"
)

var ocrProvider string

// 创建OCR表单
func createOCRForm() *widget.Form {

	//创建表单
	form := &widget.Form{}

	//读取配置
	ocrLang := viper.GetString("ocr.lang")
	ocrProvider = viper.GetString("ocr.provider")
	ocrAPPID := viper.GetStringMapString("ocr.appid")
	ocrAPIKey := viper.GetStringMapString("ocr.api_key")
	ocrAPISecret := viper.GetStringMapString("ocr.api_secret")

	//设置语言下拉框
	langCombo := widget.NewSelect([]string{"日语", "英语"}, nil)
	//根据当前设置的语言设置下拉框的选中项
	langCombo.SetSelected(langMap.LangMap[ocrLang])
	//将语言下拉框添加到表单中
	form.AppendItem(widget.NewFormItem("需要翻译的语言", langCombo))

	//创建APP ID输入框
	appIdEntry := widget.NewEntry()
	//设置APP ID输入框的文本
	appIdEntry.SetText(ocrAPPID[ocrProvider])
	//创建API Key输入框
	apiKeyEntry := widget.NewEntry()
	//设置API Key输入框的文本
	apiKeyEntry.SetText(ocrAPIKey[ocrProvider])
	//创建API Secret输入框
	apiSecretEntry := widget.NewEntry()
	//设置API Secret输入框的文本
	apiSecretEntry.SetText(ocrAPISecret[ocrProvider])

	openInput := widget.NewCheck("开启或者关闭输入框", func(b bool) {
		if b {
			fyne.Do(func() {
				appIdEntry.Enable()
				apiKeyEntry.Enable()
				apiSecretEntry.Enable()
				appIdEntry.SetText(ocrAPPID[ocrProvider])
				apiKeyEntry.SetText(ocrAPIKey[ocrProvider])
				apiSecretEntry.SetText(ocrAPISecret[ocrProvider])
				form.Refresh()
			})
		} else {
			fyne.Do(func() {
				appIdEntry.Disable()
				apiKeyEntry.Disable()
				apiSecretEntry.Disable()
				form.Refresh()
			})
		}
	})

	if ocrProvider == "paddle" || ocrProvider == "dango" {
		openInput.Checked = false
	}

	//设置ocr提供者下拉框
	ocrProviderCombo := widget.NewSelect(
		[]string{"paddle", "dango", "baidu"}, func(s string) {
			ocrProvider = s
			fyne.Do(func() {
				if ocrProvider == "paddle" || ocrProvider == "dango" {
					openInput.Checked = false
				} else {
					openInput.Checked = true
				}
			})
		})
	//设置ocr提供者下拉框的选中项
	ocrProviderCombo.SetSelected(ocrProvider)
	//将ocr提供者下拉框添加到表单中
	form.AppendItem(widget.NewFormItem("OCR提供者", ocrProviderCombo))
	//将APP ID输入框添加到表单中
	form.AppendItem(widget.NewFormItem("APP ID", appIdEntry))
	//将API Key输入框添加到表单中
	form.AppendItem(widget.NewFormItem("API KEY", apiKeyEntry))
	//将API Secret输入框添加到表单中
	form.AppendItem(widget.NewFormItem("API Secret", apiSecretEntry))

	//设置表单的提交按钮文本
	form.SubmitText = "保存设置"
	//设置表单的提交事件
	form.OnSubmit = func() {
		ocrLang = langMap.LangUnMap[langCombo.Selected]

		ocrAPPID[ocrProvider] = appIdEntry.Text
		ocrAPIKey[ocrProvider] = apiKeyEntry.Text
		ocrAPISecret[ocrProvider] = apiSecretEntry.Text
		//保存
		viper.Set("ocr.lang", ocrLang)
		viper.Set("ocr.provider", ocrProvider)
		viper.Set("ocr.appid", ocrAPPID)
		viper.Set("ocr.api_key", ocrAPIKey)
		viper.Set("ocr.api_secret", ocrAPISecret)
		//写入配置文件
		_ = viper.WriteConfig()

		dialog.ShowInformation("提示", "保存成功", mw.Window)
	}
	//设置表单的取消按钮文本
	form.CancelText = "取消"
	//设置表单的取消事件
	form.OnCancel = func() {
		//重置表单
		langCombo.SetSelected(langMap.LangUnMap[ocrLang])
		ocrProviderCombo.SetSelected(ocrProvider)
		appIdEntry.SetText(ocrAPPID[ocrProvider])
		apiKeyEntry.SetText(ocrAPIKey[ocrProvider])
		apiSecretEntry.SetText(ocrAPISecret[ocrProvider])
		if ocrProvider == "paddle" || ocrProvider == "dango" {
			openInput.Checked = false
		} else {
			openInput.Checked = true
		}
	}

	//返回表单
	return form
}
