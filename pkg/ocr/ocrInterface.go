package ocr

import "github.com/spf13/viper"

func GetOCRResult() string {
	sourceLang := viper.GetString("ocr.lang")
	return ocrByPaddle("tmp_img/tmp.png", sourceLang)
}
