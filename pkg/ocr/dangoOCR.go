package ocr

import (
	"encoding/json"
	"github.com/spf13/viper"
	"io"
	"llmTranslator/langMap"
	"llmTranslator/logHelper"
	"net/http"
	"os"
	"strings"
)

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

func ocrTestByDango(testFilePath string) bool {
	if ocrByDango(testFilePath) == "" {
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

func ocrByDango(filePath string) string {
	baseUrl := viper.GetString("ocr.base_url.dango")
	lang := langMap.LangMapToDango[viper.GetString("ocr.lang")]
	workDir, err := os.Getwd()
	if err != nil {
		logHelper.Error("获取工作目录失败: %v", err)
		logHelper.WriteLog("获取工作目录失败: %v", err)
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
		logHelper.WriteLog("json marshal error: %v", err)
		return ""
	}
	apiURL := baseUrl + "/ocr/api"

	res, err := http.Post(apiURL, "application/json", strings.NewReader(string(reqBody)))
	if err != nil {
		logHelper.Error("ocr api error: %v", err)
		logHelper.WriteLog("ocr api error: %v", err)
		return ""
	}
	defer res.Body.Close()
	dangoResponse := &DangoResponse{}
	body, err := io.ReadAll(res.Body)
	err = json.Unmarshal(body, dangoResponse)
	if err != nil {
		logHelper.Error("json unmarshal error: %v", err)
		logHelper.WriteLog("json unmarshal error: %v", err)
		return ""
	}
	if dangoResponse.Code != 0 {
		logHelper.Error("ocr api error: %v", dangoResponse.Message)
		logHelper.WriteLog("ocr api error: %v", dangoResponse.Message)
		return ""
	}

	return parseDangoResponse(dangoResponse)
}
