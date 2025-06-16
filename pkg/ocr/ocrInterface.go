package ocr

import (
	"fmt"
	"llmTranslator/configs"
	"llmTranslator/logHelper"
)

type InterfaceOCR interface {
	GetOCR(filePath string) string
	TestOCR(testFilePath string) bool
}

const tmpFilePath = "tmp_img/tmp.png"
const testFilePath = "ocrTest/test_"

var LocalOCR = []string{"paddle", "dango"}

func NewOCRTool(provider string) (InterfaceOCR, error) {
	switch provider {
	case "paddle":
		return &PaddleOCR{}, nil
	case "dango":
		return &DangoOCR{}, nil
	case "baidu":
		return &BaiduOCR{}, nil
	default:
		return nil, fmt.Errorf("不支持的ocr供应商")
	}
}

func GetOCRResult() string {
	ocr, err := NewOCRTool(configs.Setting.OCR.Provider)
	if err != nil {
		logHelper.Error("%v", err)
		return ""
	}
	return ocr.GetOCR(tmpFilePath)
}

func OCRTest() bool {
	ocr, err := NewOCRTool(configs.Setting.OCR.Provider)
	if err != nil {
		logHelper.Error("%v", err)
		return false
	}
	return ocr.TestOCR(testFilePath)
}
