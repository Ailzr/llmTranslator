package ocr

import (
	"llmTranslator/configs"
)

const tmpFilePath = "tmp_img/tmp.png"
const testFilePath = "test.png"

var LocalOCR = []string{"paddle", "dango"}

func GetOCRResult() string {
	provider := configs.Setting.OCR.Provider
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
	provider := configs.Setting.OCR.Provider
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
