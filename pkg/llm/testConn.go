package llm

import (
	"github.com/spf13/viper"
	"llmTranslator/logHelper"
)

func TestConn() bool {
	provider := viper.GetString("llm.provider")
	switch provider {
	case "ollama":
		return ollamaTest()
	default:
		logHelper.Error("不支持的provider: ", provider)
		logHelper.WriteLog("不支持的provider: " + provider)
		return false
	}
}
