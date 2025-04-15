package ui

import (
	"fmt"
	"golang.design/x/hotkey"
	"llmTranslator/configs"
	"llmTranslator/logHelper"
)

const (
	TranslateHotKey        = "translate"
	CaptureHotKey          = "capture"
	CaptureTranslateHotKey = "translate_capture"
)

var (
	translateHotKey        *hotkey.Hotkey
	captureHotKey          *hotkey.Hotkey
	captureTranslateHotKey *hotkey.Hotkey
)

func AddTranslateHotKey() error {
	// 注册全局热键
	tHotKey := configs.Setting.HotKey.Translate
	modKeys, key, err := ParseHotKey(tHotKey)
	if err != nil {
		logHelper.Error("解析热键失败: %v", err)
		logHelper.WriteLog("解析热键失败: %v", err)
		return err
	}
	translateHotKey = hotkey.New(modKeys, key)
	// 注册热键
	if err = translateHotKey.Register(); err != nil {
		logHelper.Error("注册热键失败: %v", err)
		logHelper.WriteLog("注册热键失败: %v", err)
		return fmt.Errorf("框选区翻译快捷键已被注册")
	} else {
		go func() {
			for range translateHotKey.Keydown() {
				mw.Translate()
			}
		}()
	}
	return nil
}

func AddCaptureRectangleHotKey() error {
	cHotKey := configs.Setting.HotKey.Capture
	modKeys, key, err := ParseHotKey(cHotKey)
	if err != nil {
		logHelper.Error("解析热键失败: %v", err)
		logHelper.WriteLog("解析热键失败: %v", err)
		return err
	}
	captureHotKey = hotkey.New(modKeys, key)
	if err = captureHotKey.Register(); err != nil {
		logHelper.Error("注册热键失败: %v", err)
		logHelper.WriteLog("注册热键失败: %v", err)
		return fmt.Errorf("截屏快捷键已被注册")
	} else {
		go func() {
			for range captureHotKey.Keydown() {
				mw.CaptureRectangle()
			}
		}()
	}
	return nil
}

func AddCaptureTranslateHotKey() error {
	ctHotKey := configs.Setting.HotKey.CaptureTranslate
	modKeys, key, err := ParseHotKey(ctHotKey)
	if err != nil {
		logHelper.Error("解析热键失败: %v", err)
		logHelper.WriteLog("解析热键失败: %v", err)
		return err
	}
	captureTranslateHotKey = hotkey.New(modKeys, key)
	if err = captureTranslateHotKey.Register(); err != nil {
		logHelper.Error("注册热键失败: %v", err)
		logHelper.WriteLog("注册热键失败: %v", err)
		return fmt.Errorf("截屏快捷键已被注册")
	} else {
		go func() {
			for range captureHotKey.Keydown() {
				//TODO 截屏翻译功能
				mw.CaptureRectangle()
			}
		}()
	}
	return nil
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
	}
}

func UnregisterAllHotKey() {
	UnregisterHotKey(TranslateHotKey)
	UnregisterHotKey(CaptureHotKey)
	UnregisterHotKey(CaptureTranslateHotKey)
}
