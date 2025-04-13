package configs

import (
	"github.com/spf13/viper"
	"llmTranslator/logHelper"
	"os"
)

type OCRConfig struct {
	Engine string `json:"engine"` // paddleocr / tesseract
	Lang   string `json:"lang"`   // en / japan / etc
	Url    string `json:"url"`    // 请求地址
}

type LLMConfig struct {
	Provider  string `json:"provider"`   // ollama
	ModelName string `json:"model"`      // 模型名称
	BaseURL   string `json:"base_url"`   // API地址
	APIKey    string `json:"api_key"`    // API密钥
	APISecret string `json:"api_secret"` // API密钥
}

type AppConfig struct {
	OCR OCRConfig `json:"ocr"`
	LLM LLMConfig `json:"llm"`
}

var Config = &AppConfig{}

func init() {
	//使用viper从config.yaml中读取配置信息
	//获取文件夹路径
	workDir, _ := os.Getwd()
	//设置配置文件名和路径
	viper.SetConfigName("setting")
	viper.AddConfigPath(workDir + "/configs")
	//设置配置文件类型
	viper.SetConfigType("json")
	//读取配置信息
	err := viper.ReadInConfig()
	//处理错误
	if err != nil {
		logHelper.Debug("config load error: %v", err)
		logHelper.WriteLog("config load error: %v", err)
	}

	LoadConfig()
	//如果无错误，显示配置文件读取成功
	logHelper.Info("config load success")
}

func LoadConfig() {
	Config.OCR = OCRConfig{
		Engine: viper.GetString("ocr.engine"),
		Lang:   viper.GetString("ocr.lang"),
		Url:    viper.GetString("ocr.url"),
	}
	Config.LLM = LLMConfig{
		Provider:  viper.GetString("llm.provider"),
		ModelName: viper.GetString("llm.model"),
		BaseURL:   viper.GetString("llm.url"),
		APIKey:    viper.GetString("llm.api_key"),
		APISecret: viper.GetString("llm.api_secret"),
	}
}

func SaveConfig(cfg *AppConfig) error {
	logHelper.Debug("save config: %v", cfg)
	// 重新读取配置文件
	LoadConfig()
	return nil
}
