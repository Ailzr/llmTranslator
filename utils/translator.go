package utils

import (
	"llmTranslator/logHelper"
	"llmTranslator/pkg/llm"
	"llmTranslator/pkg/ocr"
)

func GetTranslate() string {
	img, err := CaptureImg()
	SaveImgToPng(img, "tmp")
	if err != nil {
		logHelper.Error(err.Error())
		logHelper.WriteLog(err.Error())
		return "截图失败"
	}

	text := ocr.GetOCRResult()
	if text == "" {
		logHelper.Error("OCR识别失败")
		logHelper.WriteLog("OCR识别失败")
		return "OCR识别失败"
	}

	result := llm.Translate(text, "中文")
	return result
}
