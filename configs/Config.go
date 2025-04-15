package configs

type Config struct {
	Capture Capture `json:"capture"`
	LLM     LLM     `json:"llm"`
	OCR     OCR     `json:"ocr"`
	UI      UI      `json:"ui"`
	HotKey  HotKey  `json:"hotkey"`
	Version string  `json:"version"`
}

type Capture struct {
	StartX int `json:"start_x"`
	StartY int `json:"start_y"`
	EndX   int `json:"end_x"`
	EndY   int `json:"end_y"`
}

type LLM struct {
	APPID           map[string]string `json:"appid"`
	APIKey          map[string]string `json:"api_key"`
	APISecret       map[string]string `json:"api_secret"`
	BaseUrl         BaseUrl           `json:"base_url"`
	MaxResponseTime int               `json:"max_response_time"`
	MaxTokens       int               `json:"max_tokens"`
	Model           string            `json:"model"`
	Provider        string            `json:"provider"`
	Support         []string          `json:"support"`
	Temperature     float32           `json:"temperature"`
}

type OCR struct {
	APPID     map[string]string `json:"appid"`
	APIKey    map[string]string `json:"api_key"`
	APISecret map[string]string `json:"api_secret"`
	BaseUrl   BaseUrl           `json:"base_url"`
	Provider  string            `json:"provider"`
	Lang      string            `json:"lang"`
}

type UI struct {
	Theme string `json:"theme"`
}

type BaseUrl map[string]string

type HotKey struct {
	Translate        string `json:"translate"`
	Capture          string `json:"capture"`
	CaptureTranslate string `json:"capture_translate"`
}

func getDefaultConfig() *Config {
	return &Config{
		Capture: Capture{
			StartX: 0,
			StartY: 0,
			EndX:   0,
			EndY:   0,
		},
		LLM: LLM{
			APPID:     map[string]string{"ollama": "", "xunfei": ""},
			APIKey:    map[string]string{"ollama": "", "xunfei": ""},
			APISecret: map[string]string{"ollama": "", "xunfei": ""},
			BaseUrl: BaseUrl{
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
			APPID:     map[string]string{"paddle": "", "dango": "", "baidu": ""},
			APIKey:    map[string]string{"paddle": "", "dango": "", "baidu": ""},
			APISecret: map[string]string{"paddle": "", "dango": "", "baidu": ""},
			BaseUrl: BaseUrl{
				"paddle": "http://localhost:5000",
				"dango":  "http://localhost:6666",
				"baidu":  "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic",
			},
			Provider: "dango",
			Lang:     "japan",
		},
		UI: UI{
			Theme: "dark",
		},
		HotKey: HotKey{
			Translate:        "Ctrl+Shift+T",
			Capture:          "Ctrl+Shift+O",
			CaptureTranslate: "Ctrl+Shift+P",
		},
		Version: "1.0.0",
	}
}
