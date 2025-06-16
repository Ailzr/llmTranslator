package llm

import (
	"fmt"
	"llmTranslator/configs"
	"llmTranslator/langMap"
	"llmTranslator/logHelper"
)

// 翻译函数
func Translate(text string) string {

	targetLang := langMap.LangMap[configs.Setting.AppSetting.TargetLang]
	// 构造提示词
	prompt := fmt.Sprintf(configs.Setting.LLM.Prompt, targetLang, text)

	var err error

	// 调用翻译函数
	llm, err := NewLLMTool(configs.Setting.LLM.Provider)
	if err != nil {
		logHelper.Error("%v", err)
		return ""
	}
	respText, err := llm.Translate(prompt)
	if err != nil {
		logHelper.Error("%v", err)
		return ""
	}
	return respText
}
