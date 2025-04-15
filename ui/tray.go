package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"net"
)

func startTray(mw *MainWindow) {
	// 系统托盘
	if desk, ok := mw.App.(desktop.App); ok {
		m := fyne.NewMenu("llmTranslator",
			fyne.NewMenuItem("显示", func() {
				mw.Window.Show()
				mw.isTray = false
			}))
		desk.SetSystemTrayMenu(m)
		desk.SetSystemTrayIcon(resourceLlmTranslatorPng)
	}

	// 将关闭按钮修改为隐藏
	mw.Window.SetCloseIntercept(func() {
		mw.Window.Hide()

		//将托盘状态设置为true
		mw.isTray = true
	})

	// 将翻译结果展示窗口的关闭按钮也修改为隐藏
	mw.TranslatorWindow.SetCloseIntercept(func() {
		mw.TranslatorWindow.Hide()
	})

}

// 发送激活信号到已有实例
func sendActivateSignal() {
	conn, err := net.Dial("tcp", ipcPort)
	if err != nil {
		return
	}
	defer conn.Close()
	_, _ = conn.Write([]byte("activate"))
}

// 监听激活信号
func listenForActivateSignal(mw *MainWindow) {
	for {
		conn, err := mw.Listener.Accept()
		if err != nil {
			continue
		}
		buf := make([]byte, 8)
		_, err = conn.Read(buf)
		if err != nil {
			return
		}
		if string(buf) == "activate" {
			// 主线程中更新 UI
			fyne.CurrentApp().SendNotification(fyne.NewNotification("提示", "已激活窗口"))
			if mw.Window != nil {
				mw.Window.Show()
			}
		}
		_ = conn.Close()
	}
}
