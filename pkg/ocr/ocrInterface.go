package ocr

import "github.com/spf13/viper"

func GetOCRResult() string {
	sourceLang := viper.GetString("ocr.lang")
	return ocrByPaddle("tmp_img/tmp.png", sourceLang)
}

func OCRTest() bool {
	provider := viper.GetString("ocr.provider")
	switch provider {
	case "paddle-ocr":
		return ocrTestByPaddle()
	case "baidu-ocr":
		return ocrTestByBaidu()
	default:
		return false
	}
}
