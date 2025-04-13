package ui

import (
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/viper"
	"llmTranslator/logHelper"
	"slices"
)

func createLLMForm() *widget.Form {
	var form *widget.Form

	// LLM提供者选择器
	providerSelector := widget.NewSelect([]string{"ollama"}, func(s string) {
		viper.Set("llm.provider", s)
	})
	providerSelector.SetSelected(viper.GetString("llm.provider"))

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
	})
	modelInputOpen.Checked = false

	choiceList := viper.GetStringSlice("llm.support")
	// 模型选择
	modelSelector := widget.NewSelect(choiceList, func(s string) {
		viper.Set("llm.model", s)
		modelInput.SetText(s)
	})
	modelSelector.SetSelected(viper.GetString("llm.model"))

	form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "LLM提供者", Widget: providerSelector},
			{Text: "模型", Widget: modelSelector},
			{Text: "自定义模型", Widget: modelInputOpen},
			{Text: "模型名称", Widget: modelInput},
		},
		SubmitText: "保存设置",
		OnSubmit: func() {
			if modelInputOpen.Checked && modelInput.Text != "" {
				useCustomModel(modelInput, modelSelector)
			}
			err := viper.WriteConfig()
			if err != nil {
				logHelper.Error("保存配置失败: " + err.Error())
				logHelper.WriteLog("保存配置失败: " + err.Error())
				return
			}
		},
		CancelText: "取消",
		OnCancel: func() {
			resetLLMForm(providerSelector, modelSelector, modelInput, modelInputOpen)
		},
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
