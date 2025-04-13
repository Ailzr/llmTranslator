package hotkey

import (
	"golang.design/x/hotkey"
	"llmTranslator/logHelper"
	"llmTranslator/ui"
)

func AddTranslateHotKey(mw *ui.MainWindow) {
	// 注册全局热键
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyT)
	// 注册热键
	if err := hk.Register(); err != nil {
		logHelper.Error("注册热键失败: %v", err)
		logHelper.WriteLog("注册热键失败: %v", err)
		return
	} else {
		go func() {
			for range hk.Keydown() {
				mw.App.Driver().DoFromGoroutine(func() {
					mw.Translate()
				}, false)
			}
		}()
	}
}

func AddCaptureRectangleHotKey(mw *ui.MainWindow) {
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyO)
	if err := hk.Register(); err != nil {
		logHelper.Error("注册热键失败: %v", err)
		logHelper.WriteLog("注册热键失败: %v", err)
		return
	} else {
		go func() {
			for range hk.Keydown() {
				mw.App.Driver().DoFromGoroutine(func() {
					mw.CaptureRectangle()
				}, false)
			}
		}()
	}
}
