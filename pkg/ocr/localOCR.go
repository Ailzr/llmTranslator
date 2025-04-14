package ocr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"llmTranslator/logHelper"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// API响应结构体
type OCRResponse struct {
	Text string `json:"text"`
}

func ocrTestByPaddle() bool {
	//TODO Paddle-OCR测试
	return true
}

// 通过文件上传调用OCR
func ocrByPaddle(filePath, lang string) string {

	apiURL := viper.GetString("ocr.base_url.paddle")
	file, err := os.Open(filePath)
	if err != nil {
		logHelper.Error("打开文件失败: %v", err)
		logHelper.WriteLog("打开文件失败: %v", err)
		return ""
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加文件字段
	part, err := writer.CreateFormFile("image", filepath.Base(file.Name()))
	if err != nil {
		logHelper.Error("创建表单文件失败: %v", err)
		logHelper.WriteLog("创建表单文件失败: %v", err)
		return ""
	}
	_, err = io.Copy(part, file)
	if err != nil {
		logHelper.Error("写入文件内容失败: %v", err)
		logHelper.WriteLog("写入文件内容失败: %v", err)
		return ""
	}

	// 添加语言参数
	_ = writer.WriteField("lang", lang)

	// 必须显式关闭才能生成正确的multipart内容
	err = writer.Close()
	if err != nil {
		logHelper.Error("关闭multipart写入器失败: %v", err)
		logHelper.WriteLog("关闭multipart写入器失败: %v", err)
		return ""
	}

	req, err := http.NewRequest("POST", apiURL+"/ocr", body)
	if err != nil {
		logHelper.Error("创建请求失败: %v", err)
		logHelper.WriteLog("创建请求失败: %v", err)
		return ""
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := sendRequest(req)
	if err != nil {
		logHelper.Error("OCR请求失败: %v", err)
		logHelper.WriteLog("OCR请求失败: %v", err)
		return ""
	}

	return resp.Text
}

// 公共请求发送方法
func sendRequest(req *http.Request) (*OCRResponse, error) {
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求发送失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API返回错误: %s (%d)", string(errorBody), resp.StatusCode)
	}

	var result OCRResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("JSON解析失败: %w", err)
	}

	return &result, nil
}
