package configs

import (
	"bytes"
	"encoding/json"
	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
	"llmTranslator/logHelper"
	"os"
	"time"
)

type Config struct {
	Capture    Capture    `json:"capture" mapstructure:"capture"`
	LLM        LLM        `json:"llm" mapstructure:"llm"`
	OCR        OCR        `json:"ocr" mapstructure:"ocr"`
	UI         UI         `json:"ui" mapstructure:"ui"`
	HotKey     HotKey     `json:"hotkey" mapstructure:"hotkey"`
	AppSetting AppSetting `json:"app_setting" mapstructure:"app_setting"`
	Version    string     `json:"version" mapstructure:"version"`
}

type Capture struct {
	StartX int `json:"start_x" mapstructure:"start_x"`
	StartY int `json:"start_y" mapstructure:"start_y"`
	EndX   int `json:"end_x" mapstructure:"end_x"`
	EndY   int `json:"end_y" mapstructure:"end_y"`
}

type LLM struct {
	BaseUrl         map[string]string `json:"base_url" mapstructure:"base_url"`
	MaxResponseTime int               `json:"max_response_time" mapstructure:"max_response_time"`
	MaxTokens       int               `json:"max_tokens" mapstructure:"max_tokens"`
	Model           string            `json:"model" mapstructure:"model"`
	Provider        string            `json:"provider" mapstructure:"provider"`
	Support         []string          `json:"support" mapstructure:"support"`
	Temperature     float32           `json:"temperature" mapstructure:"temperature"`
}

type OCR struct {
	BaseUrl  map[string]string `json:"base_url" mapstructure:"base_url"`
	Provider string            `json:"provider" mapstructure:"provider"`
	Lang     string            `json:"lang" mapstructure:"lang"`
	Baidu    Baidu             `json:"baidu" mapstructure:"baidu"`
}

type UI struct {
	Theme string `json:"theme" mapstructure:"theme"`
}

type HotKey struct {
	Translate          string `json:"translate" mapstructure:"translate"`
	Capture            string `json:"capture" mapstructure:"capture"`
	CaptureTranslate   string `json:"capture_translate" mapstructure:"capture_translate"`
	CaptureToClipboard string `json:"capture_to_clipboard" mapstructure:"capture_to_clipboard"`
}

type Baidu struct {
	AccessToken  string    `json:"access_token" mapstructure:"access_token"`
	GenerateTime time.Time `json:"generate_time" mapstructure:"generate_time"`
	APIKey       string    `json:"api_key" mapstructure:"api_key"`
	APISecret    string    `json:"api_secret" mapstructure:"api_secret"`
}

type AppSetting struct {
	DefaultTray bool   `json:"default_tray" mapstructure:"default_tray"`
	ShowRawText bool   `json:"show_raw_text" mapstructure:"show_raw_text"`
	TargetLang  string `json:"target_lang" mapstructure:"target_lang"`
}

func createDefaultConfig() {
	file, err := os.OpenFile("configs/setting.json", os.O_CREATE, 0666)
	if err != nil {
		logHelper.Error("创建配置文件错误: %v", err)
		logHelper.WriteLog("创建配置文件错误: %v", err)
		return
	}
	err = file.Close()
	if err != nil {
		logHelper.Error("关闭配置文件错误: %v", err)
		logHelper.WriteLog("关闭配置文件错误: %v", err)
		return
	}

	WriteSettingToFile()
}

func LoadSettingByFile() {
	if err := viper.ReadInConfig(); err != nil {
		logHelper.Error("读取配置文件错误: %v", err)
		logHelper.WriteLog("读取配置文件错误: %v", err)
		return
	}
	decodeHook := mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToTimeHookFunc(time.RFC3339),
	)
	if err := viper.Unmarshal(&Setting, viper.DecodeHook(decodeHook)); err != nil {
		logHelper.Error("配置解析失败: %v", err)
		logHelper.WriteLog("配置解析失败: %v", err)
		return
	}
}

func WriteSettingToFile() {
	marshal, err := json.Marshal(Setting)
	if err != nil {
		logHelper.Error("创建默认配置时JSON序列化失败: %v", err)
		logHelper.WriteLog("创建默认配置时JSON序列化失败: %v", err)
		return
	}
	err = viper.ReadConfig(bytes.NewReader(marshal))
	if err != nil {
		logHelper.Error("创建默认配置时读取配置文件错误: %v", err)
		logHelper.WriteLog("创建默认配置时读取配置文件错误: %v", err)
		return
	}

	err = viper.WriteConfig()
	if err != nil {
		logHelper.Error("创建默认配置时写入配置文件错误: %v", err)
		logHelper.WriteLog("创建默认配置时写入配置文件错误: %v", err)
		return
	}
}
