package ui

import (
	"golang.design/x/hotkey"
	"llmTranslator/configs"
	"llmTranslator/logHelper"
)

const (
	TranslateHotKey             = "translate"
	CaptureHotKey               = "capture"
	CaptureTranslateHotKey      = "translate_capture"
	CaptureImgToClipboardHotKey = "capture_clipboard"
)

var (
	translateHotKey        *hotkey.Hotkey
	captureHotKey          *hotkey.Hotkey
	captureTranslateHotKey *hotkey.Hotkey
	captureImgToClipboard  *hotkey.Hotkey
)

func AddTranslateHotKey() {
	// 注册全局热键
	tHotKey := configs.Setting.HotKey.Translate
	modKeys, key, err := ParseHotKey(tHotKey)
	if err != nil {
		logHelper.Error("解析热键失败: %v", err)
		logHelper.WriteLog("解析热键失败: %v", err)
		return
	}
	translateHotKey = hotkey.New(modKeys, key)
	// 注册热键
	if err = translateHotKey.Register(); err != nil {
		logHelper.Error("注册热键失败: %v", err)
		logHelper.WriteLog("注册热键失败: %v", err)
		return
	} else {
		go func() {
			for range translateHotKey.Keydown() {
				mw.Translate()
			}
		}()
	}
}

func AddCaptureRectangleHotKey() {
	cHotKey := configs.Setting.HotKey.Capture
	modKeys, key, err := ParseHotKey(cHotKey)
	if err != nil {
		logHelper.Error("解析热键失败: %v", err)
		logHelper.WriteLog("解析热键失败: %v", err)
		return
	}
	captureHotKey = hotkey.New(modKeys, key)
	if err = captureHotKey.Register(); err != nil {
		logHelper.Error("注册热键失败: %v", err)
		logHelper.WriteLog("注册热键失败: %v", err)
		return
	} else {
		go func() {
			for range captureHotKey.Keydown() {
				mw.CaptureAndSaveSelection()
			}
		}()
	}
}

func AddCaptureTranslateHotKey() {
	ctHotKey := configs.Setting.HotKey.CaptureTranslate
	modKeys, key, err := ParseHotKey(ctHotKey)
	if err != nil {
		logHelper.Error("解析热键失败: %v", err)
		logHelper.WriteLog("解析热键失败: %v", err)
		return
	}
	captureTranslateHotKey = hotkey.New(modKeys, key)
	if err = captureTranslateHotKey.Register(); err != nil {
		logHelper.Error("注册热键失败: %v", err)
		logHelper.WriteLog("注册热键失败: %v", err)
		return
	} else {
		go func() {
			for range captureTranslateHotKey.Keydown() {
				mw.CaptureAndTranslate()
			}
		}()
	}
}

func AddCaptureImgToClipboradHotKey() {
	citcHotKey := configs.Setting.HotKey.CaptureToClipboard
	modKeys, key, err := ParseHotKey(citcHotKey)
	if err != nil {
		logHelper.Error("解析热键失败: %v", err)
		logHelper.WriteLog("解析热键失败: %v", err)
		return
	}
	captureImgToClipboard = hotkey.New(modKeys, key)
	if err = captureImgToClipboard.Register(); err != nil {
		logHelper.Error("注册热键失败: %v", err)
		logHelper.WriteLog("注册热键失败: %v", err)
		return
	} else {
		go func() {
			for range captureImgToClipboard.Keydown() {
				mw.CaptureToClipboard()
			}
		}()
	}
}

func UnregisterHotKey(closeKey string) {
	switch closeKey {
	case TranslateHotKey:
		err := translateHotKey.Unregister()
		if err != nil {
			logHelper.Error("注销热键失败: %v", err)
			logHelper.WriteLog("注销热键失败: %v", err)
		}
	case CaptureHotKey:
		err := captureHotKey.Unregister()
		if err != nil {
			logHelper.Error("注销热键失败: %v", err)
			logHelper.WriteLog("注销热键失败: %v", err)
		}
	case CaptureTranslateHotKey:
		err := captureTranslateHotKey.Unregister()
		if err != nil {
			logHelper.Error("注销热键失败: %v", err)
			logHelper.WriteLog("注销热键失败: %v", err)
		}
	case CaptureImgToClipboardHotKey:
		err := captureImgToClipboard.Unregister()
		if err != nil {
			logHelper.Error("注销热键失败: %v", err)
			logHelper.WriteLog("注销热键失败: %v", err)
		}
	}
}

func UnregisterAllHotKey() {
	UnregisterHotKey(TranslateHotKey)
	UnregisterHotKey(CaptureHotKey)
	UnregisterHotKey(CaptureTranslateHotKey)
	UnregisterHotKey(CaptureImgToClipboardHotKey)
}

func RegisterAllHotKey() {
	AddTranslateHotKey()
	AddCaptureRectangleHotKey()
	AddCaptureTranslateHotKey()
	AddCaptureImgToClipboradHotKey()
}
