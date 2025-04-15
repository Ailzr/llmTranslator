package ocr

import "github.com/spf13/viper"

const tmpFilePath = "tmp_img/tmp.png"
const testFilePath = "test.png"

func GetOCRResult() string {
	provider := viper.GetString("ocr.provider")
	switch provider {
	case "paddle":
		return ocrByPaddle(tmpFilePath)
	case "dango":
		return ocrByDango(tmpFilePath)
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
		return ocrTestByPaddle(testFilePath)
	case "dango":
		return ocrTestByDango(testFilePath)
	case "baidu":
		return ocrTestByBaidu(testFilePath)
	default:
		return false
	}
}
