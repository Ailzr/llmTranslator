package ocr

import (
	"encoding/json"
	"io"
	"llmTranslator/configs"
	"llmTranslator/langMap"
	"llmTranslator/logHelper"
	"net/http"
	"os"
	"strings"
)

type DangoOCR struct {
}

type DangoResponse struct {
	Code      int
	Data      []dangoData
	Message   string
	RequestId string
}

type dangoData struct {
	Coordinate struct {
		LowerLeft  []float32
		LowerRight []float32
		UpperLeft  []float32
		UpperRight []float32
	}
	Score float64
	Words string
}

type dangoRequest struct {
	ImagePath string
	Language  string
}

func (d *DangoOCR) TestOCR(testFilePath string) bool {
	if d.GetOCR(testFilePath) == "" {
		return false
	}
	return true
}

func parseDangoResponse(dangoResponse *DangoResponse) string {
	result := ""
	for _, data := range dangoResponse.Data {
		result += data.Words
	}
	return result
}

func (d *DangoOCR) GetOCR(filePath string) string {
	baseUrl := configs.Setting.OCR.BaseUrl[configs.Setting.OCR.Provider]
	lang := langMap.LangMapToDango[configs.Setting.OCR.Lang]
	workDir, err := os.Getwd()
	if err != nil {
		logHelper.Error("获取工作目录失败: %v", err)
		return ""
	}
	filePath = workDir + "/" + filePath

	req := &dangoRequest{
		ImagePath: filePath,
		Language:  lang,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		logHelper.Error("json marshal error: %v", err)
		return ""
	}
	apiURL := baseUrl + "/ocr/api"

	res, err := http.Post(apiURL, "application/json", strings.NewReader(string(reqBody)))
	if err != nil {
		logHelper.Error("ocr api error: %v", err)
		return ""
	}
	defer res.Body.Close()
	dangoResponse := &DangoResponse{}
	body, err := io.ReadAll(res.Body)
	err = json.Unmarshal(body, dangoResponse)
	if err != nil {
		logHelper.Error("json unmarshal error: %v", err)
		return ""
	}
	if dangoResponse.Code != 0 {
		logHelper.Error("ocr api error: %v", dangoResponse.Message)
		return ""
	}

	return parseDangoResponse(dangoResponse)
}
