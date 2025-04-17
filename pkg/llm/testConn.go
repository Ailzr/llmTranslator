package llm

import (
	"llmTranslator/configs"
	"llmTranslator/logHelper"
)

func TestConn() bool {
	provider := configs.Setting.LLM.Provider
	switch provider {
	case "ollama":
		return ollamaTest()
	default:
		logHelper.Error("不支持的provider: ", provider)
		return false
	}
}
