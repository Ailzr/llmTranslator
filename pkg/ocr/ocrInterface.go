package ocr

import "github.com/spf13/viper"

const tmpFilePath = "tmp_img/tmp.png"

func GetOCRResult() string {
	sourceLang := viper.GetString("ocr.lang")
	return ocrByPaddle(tmpFilePath, sourceLang)
}

func OCRTest() bool {
	provider := viper.GetString("ocr.provider")
	switch provider {
	case "paddle":
		return ocrTestByPaddle()
	case "baidu":
		return ocrTestByBaidu()
	default:
		return false
	}
}
