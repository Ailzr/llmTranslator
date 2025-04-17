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
	prompt := fmt.Sprintf("你是一个翻译助手，将以下文本翻译成{%s}，保持专业术语准确，保留数字和专有名词，不要回复其他内容，仅回复翻译出来的文本！需要翻译的内容：\n%s", targetLang, text)

	respText := ""
	var err error

	// 调用翻译函数
	provider := configs.Setting.LLM.Provider

	switch provider {
	case "ollama":
		respText, err = ollamaTranslate(prompt)
		if err != nil {
			logHelper.Error(err.Error())
			logHelper.WriteLog("LLM翻译时错误:" + err.Error())
		}
	default:
		logHelper.Error("未知的翻译提供者")
		logHelper.WriteLog("未知的翻译提供者")
	}

	return respText
}
