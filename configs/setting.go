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
		Translate:        "Ctrl+Shift+T",
		Capture:          "Ctrl+Shift+O",
		CaptureTranslate: "Ctrl+Shift+P",
	},
	DefaultTray: false,
	Version:     "1.0.0",
}
