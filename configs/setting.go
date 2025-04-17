package configs

var Setting = Config{
	Capture: Capture{
		StartX: 0,
		StartY: 0,
		EndX:   0,
		EndY:   0,
	},
	LLM: LLM{
		BaseUrl: map[string]string{
			"ollama": "http://localhost:11434",
		},
		MaxResponseTime: 60,
		MaxTokens:       2000,
		Model:           "deepseek-r1:8b",
		Provider:        "ollama",
		Support:         []string{"deepseek-r1:8b", "mistral:latest", "EasonONLINE/Sakura-qwen2.5-v1.0:7b"},
		Temperature:     0.5,
		Prompt:          "你是一个翻译助手，将以下文本翻译成%s，保持专业术语准确，保留数字和专有名词，不要回复其他内容，仅回复翻译出来的文本！需要翻译的内容：\n%s",
	},
	OCR: OCR{
		BaseUrl: map[string]string{
			"paddle": "http://localhost:5000",
			"dango":  "http://localhost:6666",
			"baidu":  "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic",
		},
		Provider: "dango",
		Lang:     "japan",
		Baidu: Baidu{
			AccessToken: "",
			APIKey:      "",
			APISecret:   "",
		},
	},
	UI: UI{
		Theme: "dark",
	},
	HotKey: HotKey{
		Translate:          "Ctrl+Shift+T",
		Capture:            "Ctrl+Shift+O",
		CaptureTranslate:   "Ctrl+Shift+P",
		CaptureToClipboard: "Ctrl+Shift+Q",
	},
	AppSetting: AppSetting{
		DefaultTray: false,
		ShowRawText: false,
		TargetLang:  "zh-CN",
	},
	Version: "1.0.0",
}
