package ocr

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"llmTranslator/langMap"
	"llmTranslator/logHelper"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type BaiduOCRResponse struct {
	WordsResult    []map[string]string `json:"words_result"`
	WordsResultNum int                 `json:"words_result_num"`
	LogId          int64               `json:"log_id"`
}

func ocrTestByBaidu(testFilePath string) bool {
	if ocrByBaidu(testFilePath) == "" {
		return false
	}
	return true
}

func getAccessToken() error {
	apiKey := viper.GetString("ocr.api_key.baidu")
	apiSecret := viper.GetString("ocr.api_secret.baidu")
	if apiKey == "" || apiSecret == "" {
		return fmt.Errorf("百度OCR API Key或Secret为空")
	}
	accessUrl := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s", apiKey, apiSecret)

	payload := strings.NewReader(``)
	client := &http.Client{}
	req, err := http.NewRequest("POST", accessUrl, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	resp := &map[string]interface{}{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return err
	}
	viper.Set("ocr.baidu.access_token", (*resp)["access_token"].(string))
	viper.Set("ocr.baidu.generate_time", time.Now())
	err = viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func checkAccessToken() bool {
	if viper.GetString("ocr.baidu.access_token") == "" {
		err := getAccessToken()
		if err != nil {
			logHelper.Error("获取百度OCR API Token失败: %v", err)
			logHelper.WriteLog("获取百度OCR API Token失败: %v", err)
			return false
		}
	}
	//百度官方说有效期最长时间为30天，这里设置为时间超过25天判断为过期
	if time.Now().Sub(viper.GetTime("ocr.baidu.generate_time")).Hours() > 24*25 {
		err := getAccessToken()
		if err != nil {
			logHelper.Error("获取百度OCR API Token失败: %v", err)
			logHelper.WriteLog("获取百度OCR API Token失败: %v", err)
			return false
		}
	}
	return true
}

func jointOCRResult(ocrResult *BaiduOCRResponse) string {
	result := ""
	for _, item := range ocrResult.WordsResult {
		result += item["words"] + "\n"
	}
	return result
}

func ocrByBaidu(filePath string) string {
	if !checkAccessToken() {
		return ""
	}

	accessToken := viper.GetString("ocr.baidu.access_token")
	baseUrl := viper.GetString("ocr.base_url.baidu")

	reqUrl := fmt.Sprintf("%s?access_token=%s", baseUrl, accessToken)

	image, err := os.ReadFile(filePath)
	if err != nil {
		logHelper.Error("读取tmp.png失败: %v", err)
		logHelper.WriteLog("读取tmp.png失败: %v", err)
		return ""
	}
	// 将图片转换为base64编码
	base64Image := base64.StdEncoding.EncodeToString(image)

	params := url.Values{}
	params.Add("access_token", accessToken)
	params.Add("image", base64Image)
	params.Add("language_type", langMap.LangMapToBaidu[viper.GetString("ocr.lang")])

	res, err := http.Post(reqUrl, "application/x-www-form-urlencoded", bytes.NewBufferString(params.Encode()))
	if err != nil {
		logHelper.Error("发送请求失败: %v", err)
		logHelper.WriteLog("发送请求失败: %v", err)
		return ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logHelper.Error("读取响应失败: %v", err)
		logHelper.WriteLog("读取响应失败: %v", err)
		return ""
	}

	data := &BaiduOCRResponse{}
	err = json.Unmarshal(body, data)
	if err != nil {
		logHelper.Error("解析响应失败: %v", err)
		logHelper.WriteLog("解析响应失败: %v", err)
		return ""
	}

	return jointOCRResult(data)
}
