package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/viper"
	"llmTranslator/logHelper"
	"slices"
)

// 创建LLM表单
func createLLMForm() *widget.Form {
	form := &widget.Form{}

	// LLM提供者选择器
	providerSelector := widget.NewSelect([]string{"ollama"}, func(s string) {
		viper.Set("llm.provider", s)
	})
	providerSelector.SetSelected(viper.GetString("llm.provider"))
	form.AppendItem(widget.NewFormItem("LLM提供者", providerSelector))

	// 模型输入框
	modelInput := widget.NewEntry()
	modelInput.SetText(viper.GetString("llm.model"))
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

	choiceList := viper.GetStringSlice("llm.support")
	// 模型选择
	modelSelector := widget.NewSelect(choiceList, func(s string) {
		viper.Set("llm.model", s)
		modelInput.SetText(s)
	})
	modelSelector.SetSelected(viper.GetString("llm.model"))

	form.AppendItem(widget.NewFormItem("模型", modelSelector))
	form.AppendItem(widget.NewFormItem("自定义模型", modelInputOpen))
	form.AppendItem(widget.NewFormItem("模型名称", modelInput))

	form.SubmitText = "保存设置"
	form.OnSubmit = func() {
		if modelInputOpen.Checked && modelInput.Text != "" {
			useCustomModel(modelInput, modelSelector)
		}
		err := viper.WriteConfig()
		if err != nil {
			logHelper.Error("保存配置失败: " + err.Error())
			logHelper.WriteLog("保存配置失败: " + err.Error())
			return
		}
	}
	form.CancelText = "取消"
	form.OnCancel = func() {
		resetLLMForm(providerSelector, modelSelector, modelInput, modelInputOpen)
	}

	return form
}

func useCustomModel(entry *widget.Entry, selector *widget.Select) {
	modelList := viper.GetStringSlice("llm.support")
	if !slices.Contains(modelList, entry.Text) {
		modelList = append(modelList, entry.Text)
		selector.Options = modelList
		viper.Set("llm.support", modelList)
	}
	selector.SetSelected(entry.Text)
	viper.Set("llm.model", entry.Text)
}

func resetLLMForm(provider, model *widget.Select, modelInput *widget.Entry, modelInputOpen *widget.Check) {
	err := viper.ReadInConfig()
	if err != nil {
		logHelper.Error("读取配置失败: " + err.Error())
		logHelper.WriteLog("读取配置失败: " + err.Error())
		return
	}
	provider.SetSelected(viper.GetString("llm.provider"))
	model.SetSelected(viper.GetString("llm.model"))
	modelInput.SetText(viper.GetString("llm.model"))
	modelInput.Disable()
	modelInputOpen.Checked = false
}
