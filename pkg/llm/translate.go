package llm

import (
	"fmt"
	"github.com/spf13/viper"
	"llmTranslator/logHelper"
)

// 翻译函数
func Translate(text, targetLang string) string {

	sourceLang := viper.GetString("ocr.lang")
	// 构造提示词
	prompt := fmt.Sprintf(
		`你是一个翻译助手，将以下文本:{%s}翻译成{%s}，保持专业术语准确，保留数字和专有名词，注意：因为文本为OCR识别而来，所以可能会有部分错误，如果你可以猜测出原文，可以根据你的猜测将其修改的更通畅，不要回复其他内容，仅回复翻译出来的文本！：
%s`, sourceLang, targetLang, text,
	)

	respText := ""
	var err error

	// 调用翻译函数
	provider := viper.GetString("llm.provider")

	switch provider {
	case "ollama":
		respText, err = ollamaTranslate(prompt, text, sourceLang, targetLang)
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
