package llm

import (
	"fmt"
	"llmTranslator/configs"
	"llmTranslator/logHelper"
)

type InterfaceLLM interface {
	Translate(text string) (string, error)
	TestLLM() bool
}

func NewLLMTool(provider string) (InterfaceLLM, error) {
	switch provider {
	case "ollama":
		return &OllamaLLM{}, nil
	default:
		return nil, fmt.Errorf("不支持的LLM供应商")
	}
}

func TestConn() bool {
	llm, err := NewLLMTool(configs.Setting.LLM.Provider)
	if err != nil {
		logHelper.Error("%v", err)
		return false
	}
	return llm.TestLLM()
}

func Translate(text string) string {
	// 调用翻译函数
	llm, err := NewLLMTool(configs.Setting.LLM.Provider)
	if err != nil {
		logHelper.Error("%v", err)
		return ""
	}
	respText, err := llm.Translate(text)
	if err != nil {
		logHelper.Error("%v", err)
		return ""
	}
	return respText
}
