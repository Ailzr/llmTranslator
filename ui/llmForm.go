package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"llmTranslator/configs"
	"slices"
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

	form.AppendItem(widget.NewFormItem("LLM提供者", providerSelector))
	form.AppendItem(widget.NewFormItem("模型", modelSelector))
	form.AppendItem(widget.NewFormItem("自定义模型", modelInputOpen))
	form.AppendItem(widget.NewFormItem("模型名称", modelInput))

	form.SubmitText = "保存"
	form.OnSubmit = func() {
		configs.Setting.LLM.Provider = providerSelector.Selected
		configs.Setting.LLM.Model = modelSelector.Selected
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
