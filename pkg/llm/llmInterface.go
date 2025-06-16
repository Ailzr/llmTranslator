package llm

import "fmt"

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
