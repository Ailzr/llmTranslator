package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"llmTranslator/configs"
	"llmTranslator/logHelper"
	"net/http"
	"time"
)

const ollamaUrl = "http://localhost:11434/api/generate"

// Ollama 请求结构
type OllamaRequest struct {
	Model   string  `json:"model"`
	Prompt  string  `json:"prompt"`
	Stream  bool    `json:"stream"`
	Options Options `json:"options,omitempty"`
}

type Options struct {
	Temperature float32 `json:"temperature,omitempty"`
	MaxTokens   int     `json:"num_predict,omitempty"`
}

// Ollama 响应结构
type OllamaResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
	Error    string `json:"error,omitempty"`
}

func ollamaTest() bool {
	client := &http.Client{Timeout: time.Duration(configs.Setting.LLM.MaxResponseTime) * time.Second}

	jsonBody, err := json.Marshal(OllamaRequest{Model: configs.Setting.LLM.Model})
	resp, err := client.Post(
		fmt.Sprintf("%s/api/generate", configs.Setting.LLM.BaseUrl[configs.Setting.LLM.Provider]),
		"application/json",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		logHelper.Error("API请求失败: ", err)
		logHelper.WriteLog("API请求失败: " + err.Error())
		return false
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		logHelper.Error("API返回错误: ", resp.Status)
		logHelper.WriteLog("API返回错误: " + resp.Status)
		return false
	}

	var result OllamaResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		logHelper.Error("JSON解析失败: ", err)
		logHelper.WriteLog("JSON解析失败: " + err.Error())
		return false
	}

	if result.Error != "" {
		logHelper.Error("模型错误: ", result.Error)
		logHelper.WriteLog("模型错误: " + result.Error)
		return false
	}

	logHelper.Info("ollama接口正常")
	return true
}

func ollamaTranslate(prompt string) (string, error) {

	requestBody := OllamaRequest{
		Model:  configs.Setting.LLM.Model, // 替换实际使用的模型
		Prompt: prompt,
		Stream: false,
		Options: Options{
			Temperature: configs.Setting.LLM.Temperature, // 控制生成随机性（0-1）
			MaxTokens:   configs.Setting.LLM.MaxTokens,   // 最大输出长度
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("JSON编码失败: %w", err)
	}

	client := &http.Client{Timeout: time.Second * time.Duration(configs.Setting.LLM.MaxResponseTime)} // 大模型响应较慢
	resp, err := client.Post(
		fmt.Sprintf("%s/api/generate", configs.Setting.LLM.BaseUrl[configs.Setting.LLM.Provider]),
		"application/json",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return "", fmt.Errorf("API请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API返回错误: %s (%d)", string(body), resp.StatusCode)
	}

	var result OllamaResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("JSON解析失败: %w", err)
	}

	if result.Error != "" {
		return "", fmt.Errorf("模型错误: %s", result.Error)
	}

	return result.Response, nil
}
