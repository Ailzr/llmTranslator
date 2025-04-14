package ocr

import "github.com/spf13/viper"

const tmpFilePath = "tmp_img/tmp.png"

func GetOCRResult() string {
	provider := viper.GetString("ocr.provider")
	switch provider {
	case "paddle":
		return ocrByPaddle(tmpFilePath)
	case "baidu":
		return ocrByBaidu(tmpFilePath)
	default:
		return ""
	}
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
