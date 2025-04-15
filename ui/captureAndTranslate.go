package ui

import "fyne.io/fyne/v2"

func (mw *MainWindow) captureAndTranslate() {
	//TODO 截屏翻译功能
	fyne.Do(func() {
		mw.CaptureWindow.Show()
	})

}
