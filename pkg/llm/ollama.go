package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"llmTranslator/logHelper"
	"net/http"
	"time"
)

const ollamaUrl = "http://localhost:11434/api/generate"

// Ollama 请求结构
type OllamaRequest struct {
	Model   string `json:"model"`
	Prompt  string `json:"prompt"`
	Stream  bool   `json:"stream"`
	Options struct {
		Temperature float64 `json:"temperature,omitempty"`
		MaxTokens   int     `json:"num_predict,omitempty"`
	} `json:"options,omitempty"`
}

// Ollama 响应结构
type OllamaResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
	Error    string `json:"error,omitempty"`
}

func ollamaTest() bool {
	baseUrl := viper.GetString("llm.base_url.ollama")
	maxResponseTime := viper.GetInt64("llm.max_response_time")
	client := &http.Client{Timeout: time.Duration(maxResponseTime) * time.Second}
	model := viper.GetString("llm.model")

	jsonBody, err := json.Marshal(OllamaRequest{Model: model})
	resp, err := client.Post(
		fmt.Sprintf("%s/api/generate", baseUrl),
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

func ollamaTranslate(prompt, text, sourceLang, targetLang string) (string, error) {

	model := viper.GetString("llm.model")
	temperature := viper.GetFloat64("llm.temperature")
	maxTokens := viper.GetInt("llm.max_tokens")
	baseUrl := viper.GetString("llm.base_url.ollama")
	maxResponseTime := viper.GetInt64("llm.max_response_time")

	requestBody := OllamaRequest{
		Model:  model, // 替换实际使用的模型
		Prompt: prompt,
		Stream: false,
		Options: struct {
			Temperature float64 `json:"temperature,omitempty"`
			MaxTokens   int     `json:"num_predict,omitempty"`
		}{
			Temperature: temperature, // 控制生成随机性（0-1）
			MaxTokens:   maxTokens,   // 最大输出长度
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("JSON编码失败: %w", err)
	}

	client := &http.Client{Timeout: time.Second * time.Duration(maxResponseTime)} // 大模型响应较慢
	resp, err := client.Post(
		fmt.Sprintf("%s/api/generate", baseUrl),
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
