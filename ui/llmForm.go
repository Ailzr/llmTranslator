package ui

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"llmTranslator/configs"
	"llmTranslator/langMap"
	"slices"
	"strconv"
)

// 创建LLM表单
func createLLMForm() *widget.Form {
	form := &widget.Form{}

	// LLM提供者选择器
	providerSelector := widget.NewSelect([]string{"ollama"}, nil)
	providerSelector.SetSelected(configs.Setting.LLM.Provider)

	// 模型输入框
	modelInput := widget.NewEntry()
	modelInput.SetText(configs.Setting.LLM.Model)
	modelInput.Disable()

	modelInputOpen := widget.NewCheck("添加自定义模型", func(b bool) {
		if b {
			modelInput.Enable()
		} else {
			modelInput.Disable()
		}
		fyne.Do(func() {
			form.Refresh()
		})
	})
	modelInputOpen.Checked = false

	choiceList := configs.Setting.LLM.Support
	// 模型选择
	modelSelector := widget.NewSelect(choiceList, func(s string) {
		modelInput.SetText(s)
	})
	modelSelector.SetSelected(configs.Setting.LLM.Model)

	promptInput := widget.NewMultiLineEntry()
	promptInput.Wrapping = fyne.TextWrapWord
	promptInput.SetText(fmt.Sprintf("注: 提示词不提供修改，如果需要修改，自行在configs/setting.json中修改\n"+configs.Setting.LLM.Prompt, langMap.LangMap[configs.Setting.AppSetting.TargetLang], "{需要翻译的文本}"))
	promptInput.SetPlaceHolder("请输入prompt")
	promptInput.Disable()

	temperatureInput := widget.NewEntry()
	temperatureInput.SetPlaceHolder("请输入温度，如果不知道怎么设置请不要随意改动")
	temperatureInput.SetText(fmt.Sprintf("%.1f", configs.Setting.LLM.Temperature))

	form.AppendItem(widget.NewFormItem("LLM提供者", providerSelector))
	form.AppendItem(widget.NewFormItem("模型", modelSelector))
	form.AppendItem(widget.NewFormItem("温度", temperatureInput))
	form.AppendItem(widget.NewFormItem("自定义模型", modelInputOpen))
	form.AppendItem(widget.NewFormItem("模型名称", modelInput))
	form.AppendItem(widget.NewFormItem("提示词", promptInput))
	form.SubmitText = "保存"
	form.OnSubmit = func() {
		temp, err := strconv.ParseFloat(temperatureInput.Text, 32)
		if err != nil {
			dialog.ShowError(errors.New("温度值转换错误，保存失败"), mw.Window)
			return
		}
		configs.Setting.LLM.Provider = providerSelector.Selected
		configs.Setting.LLM.Model = modelSelector.Selected
		configs.Setting.LLM.Temperature = float32(temp)
		if modelInputOpen.Checked && modelInput.Text != "" {
			useCustomModel(modelInput, modelSelector)
		}
		configs.WriteSettingToFile()
		dialog.ShowInformation("提示", "保存成功", mw.Window)
	}
	form.CancelText = "取消"
	form.OnCancel = func() {
		providerSelector.SetSelected(configs.Setting.LLM.Provider)
		modelSelector.SetSelected(configs.Setting.LLM.Model)
		modelInput.SetText(configs.Setting.LLM.Model)
		modelInput.Disable()
	}

	return form
}

func useCustomModel(entry *widget.Entry, selector *widget.Select) {
	modelList := configs.Setting.LLM.Support
	if !slices.Contains(modelList, entry.Text) {
		configs.Setting.LLM.Support = append(configs.Setting.LLM.Support, entry.Text)
	}
	selector.SetSelected(entry.Text)
	configs.Setting.LLM.Model = entry.Text
}
