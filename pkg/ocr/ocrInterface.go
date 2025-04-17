package ocr

import (
	"llmTranslator/configs"
)

const tmpFilePath = "tmp_img/tmp.png"
const testFilePath = "ocrTest/test_"

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
	testPath := testFilePath + configs.Setting.OCR.Lang + ".png"
	switch provider {
	case "paddle":
		return ocrTestByPaddle(testPath)
	case "dango":
		return ocrTestByDango(testPath)
	case "baidu":
		return ocrTestByBaidu(testPath)
	default:
		return false
	}
}
