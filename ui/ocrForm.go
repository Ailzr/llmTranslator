package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/viper"
	"llmTranslator/logHelper"
	"llmTranslator/utils"
)

// 创建OCR表单
func createOCRForm() *widget.Form {

	//创建表单
	form := &widget.Form{}

	//设置语言下拉框
	langCombo := widget.NewSelect(
		[]string{"日语", "英语"}, func(s string) {
			viper.Set("ocr.lang", utils.LangUnMap[s])
		})
	//获取当前设置的语言
	langSet := viper.GetString("ocr.lang")
	//根据当前设置的语言设置下拉框的选中项
	langCombo.SetSelected(utils.LangMap[langSet])
	//将语言下拉框添加到表单中
	form.AppendItem(widget.NewFormItem("需要翻译的语言", langCombo))

	//创建API Key输入框
	apiKeyEntry := widget.NewEntry()
	//获取当前设置的API Key
	apiKeySet := viper.GetString("ocr.api_key")
	//设置API Key输入框的占位符
	apiKeyEntry.SetPlaceHolder("请输入API Key")
	//设置API Key输入框的文本
	apiKeyEntry.SetText(apiKeySet)
	//禁用API Key输入框
	apiKeyEntry.Disabled()

	//设置ocr提供者下拉框
	ocrProviderCombo := widget.NewSelect(
		[]string{"paddle-ocr", "baidu-ocr"}, func(s string) {
			viper.Set("ocr.provider", s)
			//刷新表单
			fyne.Do(func() {
				form.Refresh()
			})
		})
	//设置ocr提供者下拉框的选中项
	ocrProviderCombo.SetSelected(viper.GetString("ocr.engine"))
	//将ocr提供者下拉框添加到表单中
	form.AppendItem(widget.NewFormItem("OCR提供者", ocrProviderCombo))
	//将API Key输入框添加到表单中
	form.AppendItem(widget.NewFormItem("API KEY", apiKeyEntry))

	//设置表单的提交按钮文本
	form.SubmitText = "保存设置"
	//设置表单的提交事件
	form.OnSubmit = func() {
		//保存API Key
		viper.Set("ocr.api_key", apiKeyEntry.Text)
		//写入配置文件
		_ = viper.WriteConfig()
	}
	//设置表单的取消按钮文本
	form.CancelText = "取消"
	//设置表单的取消事件
	form.OnCancel = func() {
		//重置表单
		resetOCRForm(langCombo, ocrProviderCombo, apiKeyEntry)
	}

	//返回表单
	return form
}

// 新增：重置控件的函数
func resetOCRForm(langCombo *widget.Select, ocrProviderCombo *widget.Select, apiKeyEntry *widget.Entry) {
	// 1. 重新加载配置文件（清除内存中的未保存修改）
	if err := viper.ReadInConfig(); err != nil {
		logHelper.Error("读取配置文件失败: %v", err)
		logHelper.WriteLog("读取配置文件失败: %v", err)
		return
	}

	// 2. 恢复语言选择框
	langSet := viper.GetString("ocr.lang")
	fyne.Do(
		func() {
			langCombo.SetSelected(utils.LangMap[langSet])
		})

	// 3. 恢复OCR提供者选择框
	providerSet := viper.GetString("ocr.provider")
	fyne.Do(func() {
		ocrProviderCombo.SetSelected(providerSet)
	})

	if providerSet == "paddle-ocr" {
		apiKeyEntry.Disable()
	} else {
		apiKeyEntry.Enable()
	}

	// 4. 恢复API Key输入框
	apiKeySet := viper.GetString("ocr.api_key")
	fyne.Do(func() {
		apiKeyEntry.SetText(apiKeySet)
	})
}
