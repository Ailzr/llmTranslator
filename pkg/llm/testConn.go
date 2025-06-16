package llm

import (
	"llmTranslator/configs"
	"llmTranslator/logHelper"
)

func TestConn() bool {
	llm, err := NewLLMTool(configs.Setting.LLM.Provider)
	if err != nil {
		logHelper.Error("%v", err)
		return false
	}
	return llm.TestLLM()
}
